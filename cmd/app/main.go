package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"CLI-login-system/internals/auth"
	"CLI-login-system/internals/cli"
	"CLI-login-system/internals/database"
	"CLI-login-system/internals/mfa"
	"CLI-login-system/internals/session"

	prompt "github.com/c-bata/go-prompt"
)

func main() {

	db, err := database.Open("data/app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = database.InitSchema(
		db,
		"migrations/001_initial_schema.sql",
	)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := &database.UserRepository{
		DB: db,
	}

	sessionRepo := &database.SessionRepository{
		DB: db,
	}

	authService := &auth.Service{
		Repo: userRepo,
	}

	sessionService := &session.Service{
		Repo: sessionRepo,
	}

	var currentUserID int64
	var currentUserUsername string
	var currentToken string
	var currentExpiry time.Time

	completer := func(d prompt.Document) []prompt.Suggest {
		if currentUserID == 0 {
			return prompt.FilterHasPrefix(
				cli.GuestCommands(),
				d.GetWordBeforeCursor(),
				true,
			)
		}

		return prompt.FilterHasPrefix(
			cli.UserCommands(),
			d.GetWordBeforeCursor(),
			true,
		)
	}

	executor := func(command string) {
		args := strings.Fields(command)
		if len(args) == 0 {
			return
		}

		baseCommand := args[0]

		switch baseCommand {
		case "register":
			if len(args) < 3 {
				fmt.Println("Usage: register <username> <password>")
				return
			}
			username := args[1]
			password := args[2]

			err := authService.Register(username, password)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Registration successful")

		case "login":
			if currentUserID != 0 {
				fmt.Println("Already logged in. Logout first.")
				return
			}
			if len(args) < 3 {
				fmt.Println("Usage: login <username> <password> [totp_code]")
				return
			}
			username := args[1]
			password := args[2]

			user, err := authService.Login(username, password)
			if err != nil {
				fmt.Println("Login error:", err)
				return
			}

			if user.MFAEnabled {
				if len(args) < 4 {
					fmt.Println("MFA is enabled for this account.")
					fmt.Println("Please run: login <username> <password> <totp_code>")
					return
				}
				code := args[3]
				if !mfa.Verify(code, user.TOTPSecret.String) {
					fmt.Println("Invalid TOTP code")
					return
				}
			}

			token, expiry, err := sessionService.Create(user.ID)
			if err != nil {
				fmt.Println("Session error:", err)
				return
			}

			err = userRepo.UpdateLastLogin(user.ID)
			if err != nil {
				fmt.Println("Session error:", err)
				return
			}

			currentUserID = user.ID
			currentUserUsername = user.Username
			currentToken = token
			currentExpiry = expiry

			fmt.Println("Login successful")
			fmt.Println("Username:", user.Username)
			fmt.Println("Registration Date:", formatTime(user.CreatedAt))
			if user.MFAEnabled {
				fmt.Println("MFA Status: Enabled")
			} else {
				fmt.Println("MFA Status: Disabled")
			}
			if !user.LastLogin.Valid {
				fmt.Println("Last Login: Never")
			} else {
				fmt.Println("Last Login:", formatTime(user.LastLogin.String))
			}
			fmt.Println(
				"Session Expiration:",
				currentExpiry.Format("2006-01-02 15:04:05"),
			)

		case "whoami":
			if currentUserUsername == "" {
				fmt.Println("Please login first")
				return
			}
			if time.Now().After(currentExpiry) {
				currentUserID = 0
				currentUserUsername = ""
				currentToken = ""
				fmt.Println("Session expired. Please login again.")
				return
			}
			user, err := userRepo.GetUserByID(currentUserID)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("\nCurrent User")
			fmt.Println("------------")
			fmt.Println("Username:", user.Username)
			fmt.Println("Registration Date:", formatTime(user.CreatedAt))
			if user.MFAEnabled {
				fmt.Println("MFA Status: Enabled")
			} else {
				fmt.Println("MFA Status: Disabled")
			}
			fmt.Println(
				"Session Expiration:",
				currentExpiry.Format("2006-01-02 15:04:05"),
			)
			if !user.LastLogin.Valid {
				fmt.Println("Last Login: Never")
			} else {
				fmt.Println("Last Login:", formatTime(user.LastLogin.String))
			}

		case "logout":
			if currentToken == "" {
				fmt.Println("No active session")
				return
			}
			err := sessionRepo.DeactivateSession(currentToken)
			if err != nil {
				fmt.Println(err)
				return
			}
			currentUserID = 0
			currentUserUsername = ""
			currentToken = ""
			currentExpiry = time.Time{}
			fmt.Println("Logged out successfully")

		case "enable-2fa":
			if currentUserID == 0 {
				fmt.Println("Please login first")
				return
			}
			setup, err := mfa.Generate(
				currentUserUsername,
			)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = userRepo.EnableMFA(
				currentUserID,
				setup.Secret,
			)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("MFA Enabled")
			fmt.Println("Secret:")
			fmt.Println(setup.Secret)
			fmt.Println("Provisioning URL:")
			fmt.Println(setup.URL)

		case "disable-2fa":
			if currentUserID == 0 {
				fmt.Println("Please login first")
				return
			}
			err := userRepo.DisableMFA(
				currentUserID,
			)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("MFA disabled")

		case "help":
			if currentUserID == 0 {
				fmt.Println("Available Commands:")
				fmt.Println("register <username> <password>")
				fmt.Println("login <username> <password> [totp_code]")
				fmt.Println("help")
				fmt.Println("exit")
				return
			}
			if time.Now().After(currentExpiry) {
				currentUserID = 0
				currentUserUsername = ""
				currentToken = ""
				fmt.Println("Session expired. Please login again.")
				return
			}
			fmt.Println("Available Commands:")
			fmt.Println("whoami")
			fmt.Println("enable-2fa")
			fmt.Println("disable-2fa")
			fmt.Println("logout")
			fmt.Println("help")
			fmt.Println("exit")

		case "exit":
			fmt.Println("Goodbye")
			os.Exit(0)

		default:
			fmt.Println("Unknown command. Type 'help'.")
		}
	}

	fmt.Println("CLI Login System")
	fmt.Println("Type 'help' to see available commands.")

	p := prompt.New(
		executor,
		completer,
	)
	p.Run()
}

func formatTime(value string) string {
	t, err := time.Parse(
		time.RFC3339,
		value,
	)

	if err != nil {
		return value
	}

	return t.Format(
		"2006-01-02 15:04:05",
	)
}
