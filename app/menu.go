package app

import (
	"fmt"
	"strings"
	"usmgr/internal/ui"

	"github.com/rivo/tview"
)

func newMenu() *tview.TextView {
	var menuItems = [][]string{
		{"F1", "F1"},
		{"F2", "F2"},
		{"Enter", "commands"},
	}
	menu := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter)
	menu.SetBackgroundColor(ui.Styles.Menu.BgColor)
	var menuList []string
	for i := 0; i < len(menuItems); i++ {
		key, item := genMenuItem(menuItems[i])
		if i == len(menuItems)-1 {
			item = item + " "
		}
		menuList = append(menuList, key+item)
	}
	fmt.Fprintf(menu, "%s", strings.Join(menuList, " "))
	return menu
}

func genMenuItem(items []string) (string, string) {

	key := fmt.Sprintf("[%s:%s:b] <%s>", ui.GetColorName(ui.Styles.Menu.FgColor), ui.GetColorName(ui.Styles.Menu.BgColor), items[0])
	desc := fmt.Sprintf("[%s:%s:b] %s", ui.GetColorName(ui.Styles.Menu.Item.FgColor), ui.GetColorName(ui.Styles.Menu.Item.BgColor), strings.ToUpper(items[1]))

	return key, desc
}
