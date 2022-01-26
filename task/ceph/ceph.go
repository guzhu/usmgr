package ceph

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
	"strings"
	"usmgr/internal/ui"
)

type CephTask struct {
	*tview.Table
	title string
}

func NewCeph() *CephTask  {
	c := &CephTask{
		Table: tview.NewTable(),
		title: "ceph",
	}

	c.SetSelectedFunc(c.selectable)

	headers := []string{"name", "", "", ""}
	expansions := []int{4, 4, 4, 8}
	alignment := []int{ui.L, ui.L, ui.L, ui.L}

	fgColor := ui.Styles.PageTable.HeaderRow.FgColor
	bgColor := ui.Styles.PageTable.HeaderRow.BgColor

	for i := 0; i < len(headers); i++ {
		c.SetCell(0, i,
			tview.NewTableCell(fmt.Sprintf("[black::b]%s", strings.ToUpper(headers[i]))).
				SetExpansion(1).
				SetBackgroundColor(bgColor).
				SetTextColor(fgColor).
				SetAlign(tview.AlignLeft).
				SetSelectable(false))
	}

	osd := []string{"osd", "", "", ""}
	pool := []string{"pool", "", "", ""}
	vol := []string{"vol", "", "", ""}
	mds := []string{"mds", "", "", ""}

	ui.AddTableData(c.Table, 1, [][]string{osd}, alignment, expansions, tcell.ColorYellow, true)
	ui.AddTableData(c.Table, 2, [][]string{pool}, alignment, expansions, tcell.ColorYellow, true)
	ui.AddTableData(c.Table, 3, [][]string{vol}, alignment, expansions, tcell.ColorYellow, true)
	ui.AddTableData(c.Table, 4, [][]string{mds}, alignment, expansions, tcell.ColorYellow, true)

	c.SetFixed(1, 1)
	c.SetSelectable(true, false)

	return c
}

func (c *CephTask) GetTitle() string {
	return c.title
}

func (c *CephTask) selectable(row int, column int) {
	log.Debug().Msgf("row: %v column: %v",row, column)
	// s.GetCell(row, column).SetTextColor(tcell.ColorRed)
	c.SetSelectable(true, false)
}