package ui

import (
	"fmt"
	"strings"
	"usmgr/config"

	"github.com/rivo/tview"
)

// LogoSmall USmgr small log.
var LogoSmall = []string{
	` ____         ____ `,
	`|    |       |    |`,
	`|    |       |    |`,
	`|    |       |    |`,
	`|    |_______|    |`,
	`|                 |`,
	`\\_______________// `,
}

// LogoBig USmgr big logo for splash page.
var LogoBig = []string{
	` ____         ____ `,
	`|    |       |    |`,
	`|    |       |    |`,
	`|    |       |    |`,
	`|    |_______|    |`,
	`|                 |`,
	`\\_______________// `,
}

// Splash represents a splash screen.
type Splash struct {
	*tview.Flex
}

// NewSplash instantiates a new splash screen with product and company info.
func NewSplash(styles *config.Styles, version string) *Splash {
	s := Splash{Flex: tview.NewFlex()}
	s.SetBackgroundColor(styles.BgColor())

	logo := tview.NewTextView()
	logo.SetDynamicColors(true)
	logo.SetTextAlign(tview.AlignCenter)
	s.layoutLogo(logo, styles)

	vers := tview.NewTextView()
	vers.SetDynamicColors(true)
	vers.SetTextAlign(tview.AlignCenter)
	s.layoutRev(vers, version, styles)

	s.SetDirection(tview.FlexRow)
	s.AddItem(logo, 10, 1, false)
	s.AddItem(vers, 1, 1, false)

	return &s
}

func (s *Splash) layoutLogo(t *tview.TextView, styles *config.Styles) {
	logo := strings.Join(LogoBig, fmt.Sprintf("\n[%s::b]", styles.Body().LogoColor))
	fmt.Fprintf(t, "%s[%s::b]%s\n",
		strings.Repeat("\n", 2),
		styles.Body().LogoColor,
		logo)
}

func (s *Splash) layoutRev(t *tview.TextView, rev string, styles *config.Styles) {
	fmt.Fprintf(t, "[%s::b]Revision [red::b]%s", styles.Body().FgColor, rev)
}
