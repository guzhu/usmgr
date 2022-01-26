package app

import (
	"fmt"
	"github.com/rivo/tview"
	"usmgr/internal/ui"
)

// InfoBarViewHeight info bar height
const (
	connectionCellRow = 0
	hostnameCellRow   = 1
	osCellRow         = 2
)

// Head implements the info bar primitive
type Head struct {
   *tview.Table
	title  string
}

// NewInfoBar returns info bar view
func newHead() *Head {
	h := &Head{
		Table: tview.NewTable(),
		title: "USmgr",
	}

	emptyCell := func() *tview.TableCell {
		return tview.NewTableCell("")
	}

	headerColor := ui.GetColorName(ui.Styles.InfoBar.ItemFgColor)

	h.SetBorder(true)

	h.SetCell(connectionCellRow, 0, emptyCell())
	h.SetCell(connectionCellRow, 1, tview.NewTableCell(fmt.Sprintf("[%s::]%s", headerColor, "connection:")))
	h.SetCell(connectionCellRow, 2, emptyCell())

	h.SetCell(hostnameCellRow, 0, emptyCell())
	h.SetCell(hostnameCellRow, 1, tview.NewTableCell(fmt.Sprintf("[%s::]%s", headerColor, "Hostname:")))
	h.SetCell(hostnameCellRow, 2, emptyCell())

	h.SetCell(osCellRow, 0, emptyCell())
	h.SetCell(osCellRow, 1, tview.NewTableCell(fmt.Sprintf("[%s::]%s", headerColor, "OS type:")))
	h.SetCell(osCellRow, 2, emptyCell())

	return  h
}

// UpdateBasicInfo updates hostname, kernel and os type values
func (h *Head) UpdateBasicInfo(hostname string, kernel string, ostype string) {
	h.GetCell(hostnameCellRow, 2).SetText(hostname)
	h.GetCell(osCellRow, 2).SetText(ostype)
	h.GetCell(connectionCellRow, 2).SetText(kernel)
}
