package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type gui struct {
	app fyne.App
}

func New() *gui {
	a := app.New()
	w := a.NewWindow("ChatterByte")
	w.Resize(fyne.NewSize(800, 600))
	w.SetFixedSize(true)
	w.CenterOnScreen()
	w.SetPadded(true)

	return &gui{
		app: a,
	}
}

func (g *gui) Login() {

}

func Render() {

	login := widget.NewEntry()
	password := widget.NewPasswordEntry()
	forum := widget.NewForm(
		widget.NewFormItem("Login", login),
		widget.NewFormItem("Password", password),
	)

	forum.OnSubmit = func() {
		println("Login: ", login.Text)
		println("Password: ", password.Text)
		w.Close()
	}
	// window with login and password form
	w.SetContent(forum)
	w.Show()
	a.Run()
}
