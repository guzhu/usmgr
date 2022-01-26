package ssh

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/kevinburke/ssh_config"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strings"
	"usmgr/internal/ui"
)

type SSHTask struct {
	*tview.Table
	title string
	data [][]string
	pages *tview.Pages
}

func NewSSH() *SSHTask  {
	s := &SSHTask{
		Table: tview.NewTable(),
		title: "ssh",
	}

	s.SetSelectedFunc(s.selectable)

	headers := []string{"Name", "User", "Host", "Port"}
	expansions := []int{2, 2, 3, 3 }
	alignment := []int{ui.L, ui.L, ui.L, ui.L}

	fgColor := ui.Styles.PageTable.FgColor
	bgColor := ui.Styles.PageTable.BgColor

	for i := 0; i < len(headers); i++ {
		s.SetCell(0, i,
			tview.NewTableCell(fmt.Sprintf("[black::b]%s", strings.ToUpper(headers[i]))).
				SetExpansion(1).
				SetBackgroundColor(bgColor).
				SetTextColor(fgColor).
				SetAlign(tview.AlignLeft).
				SetSelectable(false))
	}
	hosts := s.getHosts()
	log.Debug().Msgf("%v",hosts)
	if len(hosts) > 0 {
		for index, host :=  range hosts {
			ui.AddTableData(s.Table, 1+index, [][]string{host}, alignment, expansions, tcell.ColorYellow, true)
		}
	}

	s.SetFixed(1, 1)
	s.SetSelectable(true, false)

	return s
}

func (s *SSHTask) GetTitle() string {
	return s.title
}

func (s *SSHTask) selectable(row int, column int) {
	log.Debug().Msgf("row: %v column: %v",row, column)
	// s.GetCell(row, column).SetTextColor(tcell.ColorRed)
	s.SetSelectable(true, false)

	s.pages.SwitchToPage(s.GetCell(row,0).Text)

	log.Debug().Msgf("switch to %s",s.GetCell(row,0).Text)
}

func (s *SSHTask) getHosts() (data [][]string) {
	f, _ := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "config"))

	cfg, _ := ssh_config.Decode(f)


	expansions := []int{4, 4}
	alignment := []int{ui.L, ui.L}

	s.pages = tview.NewPages()

	for _, host := range cfg.Hosts {
		var r []string

		hostinfo := s.hostInfo()

		for _, p := range host.Patterns {
			if p.String() == "*"{
				break
			}

			hostinfo.SetTitle(p.String())

			r = append(r, p.String())
			t,_ := cfg.Get( p.String(),"User")
			r = append(r,t)
			t,_ = cfg.Get( p.String(),"HostName")
			r = append(r,t)
			t,_ = cfg.Get( p.String(),"Port")

			if len(t) == 0 {
				r = append(r,"22")
			} else {
				r = append(r,t)
			}

			break
		}


		if len(r) != 0 {
			data = append(data, r)
		}

		// var m [][]string
		for index, node := range host.Nodes{
			var n []string
			n = append(n, node.String())
			ui.AddTableData(hostinfo, 1+index, [][]string{n}, alignment, expansions, tcell.ColorYellow,true)

		}
		// m = append(m, n)
		// ui.AddTableData(hostinfo, 1+index, data, alignment, expansions, tcell.ColorYellow,true)

		s.pages.AddPage(hostinfo.GetTitle(),hostinfo,true,false)
	}
	s.data = data
	return data
}

func (s *SSHTask) hostInfo() *tview.Table {
	h := tview.NewTable()

	fgColor := ui.Styles.PageTable.FgColor
	bgColor := ui.Styles.PageTable.BgColor

	headers := []string{"name", "value"}
	for i := 0; i < len(headers); i++ {
		h.SetCell(0, i,
			tview.NewTableCell(fmt.Sprintf("[black::b]%s", strings.ToUpper(headers[i]))).
				SetExpansion(1).
				SetBackgroundColor(bgColor).
				SetTextColor(fgColor).
				SetAlign(tview.AlignLeft).
				SetSelectable(false))
	}

	return h
}
