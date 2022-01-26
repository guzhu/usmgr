package ui

import (
	"github.com/gdamore/tcell/v2"
)

const (
	// IDLength max ID length to display
	IDLength = 12
)

// GetColorName returns convert tcell color to its name
func GetColorName(color tcell.Color) string {
	for name, c := range tcell.ColorNames {
		if c == color {
			return name
		}
	}
	return ""
}

// AlignStringListWidth returns max string len in the list.
func AlignStringListWidth(list []string) ([]string, int) {
	var (
		max         = 0
		alignedList []string
	)
	for _, item := range list {
		if len(item) > max {
			max = len(item)
		}
	}
	for _, item := range list {
		if len(item) < max {
			need := max - len(item)
			for i := 0; i < need; i++ {
				item = item + " "
			}
		}
		alignedList = append(alignedList, item)
	}
	return alignedList, max
}


type theme struct {
	InfoBar            infoBar
	Menu               menu
	PageTable          pageTable
	CommandDialog      commandDialog
	ConfirmDialog      confirmDialog
	ImageSearchDialog  imageSearchDialog
	ImageHistoryDialog imageHistoryDialog
}

type infoBar struct {
	ItemFgColor  tcell.Color
	ValueFgColor tcell.Color
	ProgressBar  progressBar
}
type menu struct {
	FgColor tcell.Color
	BgColor tcell.Color
	Item    menuItem
}
type menuItem struct {
	FgColor tcell.Color
	BgColor tcell.Color
}

type progressBar struct {
	FgColor       tcell.Color
	BarOKColor    tcell.Color
	BarWarnColor  tcell.Color
	BarCritColor  tcell.Color
	BarEmptyColor tcell.Color
}

type pageTable struct {
	FgColor   tcell.Color
	BgColor   tcell.Color
	HeaderRow headerRow
}

type headerRow struct {
	FgColor tcell.Color
	BgColor tcell.Color
}

type commandDialog struct {
	BgColor   tcell.Color
	FgColor   tcell.Color
	HeaderRow headerRow
}

type confirmDialog struct {
	BgColor tcell.Color
	FgColor tcell.Color
}

type imageSearchDialog struct {
	BgColor                tcell.Color
	FgColor                tcell.Color
	ResultHeaderRow        headerRow
	ResultTableBgColor     tcell.Color
	ResultTableBorderColor tcell.Color
}

type imageHistoryDialog struct {
	BgColor   tcell.Color
	FgColor   tcell.Color
	HeaderRow headerRow
}
