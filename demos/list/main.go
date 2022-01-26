// Demo code for the List primitive.
package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type T struct {
	*tview.Box
	list *tview.List
}

func (t *T) Draw(screen tcell.Screen) {
	t.Box.DrawForSubclass(screen, t)
	t.Box.SetBorder(false)
	x, y, width, height := t.GetInnerRect()
	t.list.SetRect(x, y, width, height)
	t.list.SetBorder(true)
	t.list.Draw(screen)
}

func main() {
	app := tview.NewApplication()
	list := tview.NewList().ShowSecondaryText(false).
		AddItem("List item 1", "Some explanatory text", 'a', nil).
		AddItem("List item 2", "Some explanatory text", 'b', nil).
		AddItem("List item 3", "Some explanatory text", 'c', nil).
		AddItem("List item 4", "Some explanatory text", 'd', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})


	flex := tview.NewFlex().
		AddItem(list, 0, 1, true).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Top"), 0, 1, false)

	t := &T{
		Box:  tview.NewBox(),
	}
	t.list = list
	t.list.SetHighlightFullLine(true)

	// var pages  *tview.Pages
	// pages = tview.NewPages().
	// 	AddPage("main", flex, true, true)

	app.SetRoot(flex, true)

	if err := app.EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
