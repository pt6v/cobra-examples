package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "return string length",
	Long:  "return the certain string length",
	Run:   runCount,
}

func runCount(cmd *cobra.Command, args []string) {
	if len(args) >= 1 {
		str := args[0]
		fmt.Println(len(str))
	} else {
		fmt.Println("String is missing")
	}
}

func init() {
	RootCmd.AddCommand(countCmd)
}
