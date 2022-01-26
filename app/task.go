package app

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
	"strings"
	"usmgr/internal/ui"
	"usmgr/task/ceph"
	"usmgr/task/service"
	ssh2 "usmgr/task/ssh"
)

type Task struct {
	*tview.List
	title  string
	actions ui.KeyActions
	app *App
}

func newTask(a *App) *Task {
	t := &Task{
		List:  tview.NewList(),
		title: "Task",
		actions: make(ui.KeyActions),
		app: a,
	}
	t.SetBorder(true)
	t.SetTitle(t.title)
	t.ShowSecondaryText(false)
	t.SetInputCapture(t.keyboard)
	t.bindKeys()

	ceph := ceph.NewCeph()
	srv := service.NewService()
	ssh := ssh2.NewSSH()

	f := t.swithPage
	t.AddItem(ceph.GetTitle(),"",'c',f)
	t.AddItem(srv.GetTitle(),"",'s',f)
	t.AddItem(ssh.GetTitle(),"",'h',f)

	t.AddItem("Quit", "quit ", 'q', func() {
		t.app.Stop()
	})
	t.SetHighlightFullLine(true)

	return t
}

func (t *Task)swithPage()  {
	mainText,_ := t.GetItemText(t.GetCurrentItem())
	log.Debug().Msgf("switch to pages %s",mainText)
	t.app.pages.SetTitle(fmt.Sprintf("[::b]%s", strings.ToUpper(mainText)))
	t.app.pages.SwitchToPage(mainText)
	// t.app.SetFocus(t.app.pages)
}

func (t *Task) keyboard(evt *tcell.EventKey) *tcell.EventKey {
	a, ok := t.actions[ui.AsKey(evt)]
	if  ok {
		return a.Action(evt)
	}
	return evt
}

func (t *Task) bindKeys() {
	t.actions.Set(ui.KeyActions{
		tcell.KeyUp:  ui.NewKeyAction("up",t.moveup, false),
		tcell.KeyDown: ui.NewKeyAction("down", t.movedown, false),
		ui.KeyJ:  ui.NewKeyAction("down",t.movedown, false),
		ui.KeyK: ui.NewKeyAction("up", t.moveup, false),
		tcell.KeyTAB: ui.NewKeyAction("table",t.switchFocus ,false),
	})
}

func (t *Task) moveup(evt *tcell.EventKey) *tcell.EventKey {
	previousItem := t.GetCurrentItem()
	t.SetCurrentItem(previousItem-1)
	return nil
}

func (t *Task) movedown(evt *tcell.EventKey) *tcell.EventKey {
	previousItem := t.GetCurrentItem()
	t.SetCurrentItem(previousItem+1)
	return nil
}

func (t *Task) switchFocus(evt *tcell.EventKey) *tcell.EventKey {
	log.Debug().Msgf("switch focus to pages box")
	name,_ := t.app.pages.GetFrontPage()
	t.app.pages.SetTitle(fmt.Sprintf("[::b]%s", strings.ToUpper(name)))
	t.app.SetFocus(t.app.pages)
	return nil
}


