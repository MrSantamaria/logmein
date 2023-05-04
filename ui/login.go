package ui

import (
	"fmt"

	"github.com/andlabs/ui"
)

// LoginWindow creates a pop-up window that prompts the user to enter their credentials.
// It returns the entered username and password.
func LoginWindow() (string, string, error) {
	var username, password string

	err := ui.Main(func() {
		window := ui.NewWindow("Login", 400, 200, false)

		usernameTextField := ui.NewEntry()
		passwordTextField := ui.NewPasswordEntry()
		loginButton := ui.NewButton("Login")

		form := ui.NewForm()
		form.SetPadded(true)
		form.Append("Username", usernameTextField, false)
		form.Append("Password", passwordTextField, false)
		form.Append("", loginButton, false)

		window.SetChild(form)

		loginButton.OnClicked(func(*ui.Button) {
			username = usernameTextField.Text()
			password = passwordTextField.Text()
			window.Destroy()
		})

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		window.Show()

	})

	if err != nil {
		return "", "", fmt.Errorf("error opening login window: %v", err)
	}

	return username, password, nil
}
