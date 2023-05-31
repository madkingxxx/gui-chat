package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type gui struct {
	app fyne.App
	mW  fyne.Window
}

func New() *gui {
	a := app.New()
	mainWindow := a.NewWindow("Go Chat")
	configureWindow(mainWindow)
	return &gui{
		app: a,
		mW:  mainWindow,
	}
}

func configureWindow(w fyne.Window) {
	w.Resize(fyne.NewSize(800, 600))
	// w.SetFixedSize(true)
	w.CenterOnScreen()
	w.SetPadded(true)
}

// LoginWindows creates a login window.
func (g *gui) LoginWindows() {
	login := widget.NewEntry()
	password := widget.NewPasswordEntry()

	forum := widget.NewForm(
		widget.NewFormItem("Login", login),
		widget.NewFormItem("Password", password),
	)

	forum.OnSubmit = func() {
		g.ChatWindows()
	}

	g.mW.SetContent(forum)
	g.mW.Show()
}

func resizeWidget(w fyne.Widget, w1, h1 int) fyne.Widget {
	w.Resize(fyne.NewSize(float32(w1), float32(h1)))
	return w
}

func (g *gui) ChatWindows() {
	// text list
	// dynamical list of users in the chat
	userList := getUserList()
	// main layout
	layout := container.New(layout.NewGridLayoutWithColumns(2))

	// adding users and chat text
	myIcon := widget.NewIcon(theme.AccountIcon())
	topPanel := container.NewBorder(nil, nil, widget.NewLabel("Online users"), myIcon)

	layout.Add(topPanel)
	layout.Add(resizeWidget(widget.NewLabel("Chat"), 50, 50))

	// user list box
	userBox := container.NewVBox()
	for _, user := range userList {
		userText := widget.NewLabel(user.Name)
		userBox.Add(userText)
	}
	layout.Add(userBox)

	// chat box
	g.mW.SetContent(layout)

	g.mW.Show()
}

// Run runs the application.
func (g *gui) Run() {
	g.ChatWindows()
	g.app.Run()
}
