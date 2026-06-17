package cli

import "fmt"

func ShowHelp(loggedIn bool) {

	if loggedIn {
		fmt.Println("Available Commands:")
		fmt.Println("whoami")
		fmt.Println("enable-2fa")
		fmt.Println("disable-2fa")
		fmt.Println("logout")
		fmt.Println("help")
		fmt.Println("exit")
		return
	}

	fmt.Println("Available Commands:")
	fmt.Println("register")
	fmt.Println("login")
	fmt.Println("help")
	fmt.Println("exit")
}
