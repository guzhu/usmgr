package cmd

import (
	"fmt"
	"usmgr/config"
	"usmgr/internal/color"
	"usmgr/internal/ui"

	"github.com/spf13/cobra"
)

func infoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Print configuration info",
		Long:  "Print configuration information",
		Run: func(cmd *cobra.Command, args []string) {
			printInfo()
		},
	}
}

func printInfo() {
	const fmat = "%-25s %s\n"

	printLogo(color.Cyan)
	printTuple(fmat, "Configuration", config.USMgrConfigFile, color.Cyan)
	printTuple(fmat, "Logs", config.USMgrLogs, color.Cyan)
	printTuple(fmat, "Screen Dumps", config.USMgrDumpDir, color.Cyan)
}

func printLogo(c color.Paint) {
	for _, l := range ui.LogoSmall {
		fmt.Println(color.Colorize(l, c))
	}
	fmt.Println()
}
