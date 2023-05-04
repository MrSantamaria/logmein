package ui

import (
	"fmt"

	"github.com/MrSantamaria/logmein/utils"
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
			err := utils.Retry(3, func() error {
				_, err := getUserInput(window, usernameTextField, "Username")
				return err
			})
			if err != nil {
				fmt.Errorf("error getting user input for username: %v", err)
			}
			getUserInput(window, passwordTextField, "Password")
			window.Destroy()
		})

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		window.SetMargined(true)

		// TODO: Center window on screen

		window.Show()
	})

	if err != nil {
		return "", "", fmt.Errorf("error opening login window: %v", err)
	}

	return username, password, nil
}

func getUserInput(window *ui.Window, field *ui.Entry, fieldName string) (string, error) {
	for {
		if field.Text() == "" {
			ui.MsgBoxError(window, "Error", fmt.Sprintf("%s is required.", fieldName))
			break
		}
		return field.Text(), nil
	}

	return "", fmt.Errorf("error getting user input for %s", fieldName)
}
