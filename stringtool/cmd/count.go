package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count [str]",
	Short: "return string length",
	Long:  "return the certain string length",
	Run:   runCount,
}

func runCount(cmd *cobra.Command, args []string) {
	logrus.Debug("args", args)
	if len(args) >= 1 {
		str := args[0]
		fmt.Println(len(str))
	} else {
		logrus.Error("String is missing")
	}
}

func init() {
	RootCmd.AddCommand(countCmd)
}
