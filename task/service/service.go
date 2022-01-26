package service

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
	"strings"
	"usmgr/internal/ui"
)

type ServiceTask struct {
	*tview.Table
	title string
}

func NewService() *ServiceTask  {
	s := &ServiceTask{
		Table: tview.NewTable(),
		title: "service",
	}

	s.SetSelectedFunc(s.selectable)

	headers := []string{"name", "status", "uptime", "desc"}
	expansions := []int{4, 4, 4, 8}
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

	ui.AddTableData(s.Table, 1, [][]string{headers}, alignment, expansions, tcell.ColorYellow, true)
	ui.AddTableData(s.Table, 2, [][]string{headers}, alignment, expansions, tcell.ColorYellow, true)
	ui.AddTableData(s.Table, 3, [][]string{headers}, alignment, expansions, tcell.ColorYellow, true)
	ui.AddTableData(s.Table, 4, [][]string{headers}, alignment, expansions, tcell.ColorYellow, true)

	s.SetFixed(1, 1)
	s.SetSelectable(true, false)

	return s
}

func (s *ServiceTask) GetTitle() string {
	return s.title
}

func (s *ServiceTask) selectable(row int, column int) {
	log.Debug().Msgf("row: %v column: %v",row, column)
	// s.GetCell(row, column).SetTextColor(tcell.ColorRed)
	s.SetSelectable(true, false)
}