package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strings"
)

var upperCmd = &cobra.Command{
	Use:   "upper [str]",
	Short: "return string uppercase",
	Long:  "return the certain string uppercase",
	Run:   runUpper,
}

func runUpper(cmd *cobra.Command, args []string) {
	logrus.Debug("args", args)
	if len(args) >= 1 {
		str := args[0]
		fmt.Println(strings.ToUpper(str))
	} else {
		logrus.Error("String is missing")
	}
}

func init() {
	RootCmd.AddCommand(upperCmd)
}
