package app

import (
	"github.com/rivo/tview"
	"time"
)

const (
	splashDelay      = 1 * time.Second
	clusterRefresh   = 15 * time.Second
	clusterInfoWidth = 50
	clusterInfoPad   = 15
	refreshInterval = 1000 * time.Millisecond
)

// App represents an application view.
type App struct {
	version string
	*tview.Application
	head     *Head
	menu        *tview.TextView
	pages       *Pages
	task        *Task
}

// NewApp returns a new app.
func NewApp() *App {
	a := App{
		Application:  tview.NewApplication(),
	}
	a.head = newHead()
	a.menu = newMenu()
	a.task = newTask(&a)
	a.pages = newPages(&a)

	return &a
}

func (a *App) Run() error {
	middle := tview.NewFlex().SetDirection(tview.FlexColumn)
	middle.AddItem(a.task,0,1,true)
	middle.AddItem(a.pages, 0,10,true)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(a.head, 6, 1, false).
		AddItem(middle, 0, 1, true).
		AddItem(a.menu, 1, 1, false)

	a.initUI()

	go a.refresh()

	if err := a.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		return err
	}

	return nil
}