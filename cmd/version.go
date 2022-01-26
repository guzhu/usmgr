package cmd

import (
	"fmt"
	"usmgr/internal/color"

	"github.com/spf13/cobra"
)

func versionCmd() *cobra.Command {
	var short bool

	command := cobra.Command{
		Use:   "version",
		Short: "Print version/build info",
		Long:  "Print version/build information",
		Run: func(cmd *cobra.Command, args []string) {
			printVersion(short)
		},
	}

	command.PersistentFlags().BoolVarP(&short, "short", "s", false, "Prints USmgr version info in short format")

	return &command
}

func printVersion(short bool) {
	const fmat = "%-20s %s\n"
	var outputColor color.Paint

	if short {
		outputColor = -1
	} else {
		outputColor = color.Cyan
		printLogo(outputColor)
	}
	printTuple(fmat, "Version", version, outputColor)
	printTuple(fmat, "Commit", commit, outputColor)
	printTuple(fmat, "Date", date, outputColor)
}

func printTuple(fmat, section, value string, outputColor color.Paint) {
	if outputColor != -1 {
		fmt.Printf(fmat, color.Colorize(section+":", outputColor), color.Colorize(value, color.LightGray))
		return
	}
	fmt.Printf(fmat, section, value)
}