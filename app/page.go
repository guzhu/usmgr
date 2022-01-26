package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
	"usmgr/internal/ui"
	"usmgr/task/ceph"
	"usmgr/task/service"
	ssh2 "usmgr/task/ssh"
)

type Pages struct {
	*tview.Pages
	title  string
	actions ui.KeyActions
	app *App
}

func newPages(a *App) *Pages {
	p := &Pages{
		Pages:  tview.NewPages(),
		title: "Pages",
		actions: make(ui.KeyActions),
		app: a,
	}
	p.SetBorder(true)
	p.SetInputCapture(p.keyboard)
	p.bindKeys()

	ceph := ceph.NewCeph()
	srv := service.NewService()
	ssh := ssh2.NewSSH()

	p.AddPage(ceph.GetTitle(), ceph, true, false)
	p.AddPage(srv.GetTitle(), srv, true, false)
	p.AddPage(ssh.GetTitle(), ssh, true, true)

	return p
}


func (p *Pages) keyboard(evt *tcell.EventKey) *tcell.EventKey {
	a, ok := p.actions[ui.AsKey(evt)]
	if  ok {
		return a.Action(evt)
	}
	return evt
}

func (p *Pages) bindKeys() {
	p.actions.Set(ui.KeyActions{
		tcell.KeyTAB: ui.NewKeyAction("table",p.switchFocus ,false),
	})
}

func (p *Pages) switchFocus(evt *tcell.EventKey) *tcell.EventKey {
	log.Debug().Msgf("switch focus to task box")
	p.app.SetFocus(p.app.task)
	return nil
}


