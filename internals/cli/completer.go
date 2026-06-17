package cli

import "github.com/c-bata/go-prompt"

func GuestCommands() []prompt.Suggest {
	return []prompt.Suggest{
		{Text: "register"},
		{Text: "login"},
		{Text: "help"},
		{Text: "exit"},
	}
}

func UserCommands() []prompt.Suggest {
	return []prompt.Suggest{
		{Text: "whoami"},
		{Text: "enable-2fa"},
		{Text: "disable-2fa"},
		{Text: "logout"},
		{Text: "help"},
		{Text: "exit"},
	}
}
