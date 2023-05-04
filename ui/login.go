package ui

import (
	"fmt"

	"github.com/andlabs/ui"
)

func login_window() error {
	err := ui.Main(func() {
		window := ui.NewWindow("Login", 400, 200, false)

		// TODO: Add username and password fields and a login button

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		window.Show()

	})

	if err != nil {
		return fmt.Errorf("error opening login window: %v", err)
	}

	username_text_field := ui.NewEntry()
	password_text_field := ui.NewPasswordEntry()
	login_button := ui.NewButton("Login")

	form := ui.NewForm()
	form.SetPadded(true)
	form.Append("Username", username_text_field, false)
	form.Append("Password", password_text_field, false)
	form.Append("", login_button, false)

	window.SetChild(form)

	login_button.OnClicked(func(*ui.Button) {
		username := username_text_field.Text()
		password := password_text_field.Text()

		// TODO: pass the credentials to the login function

		window.Destroy()

	return nil
}
