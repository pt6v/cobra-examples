package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var upperCmd = &cobra.Command{
	Use:   "upper",
	Short: "return string uppercase",
	Long:  "return the certain string uppercase",

	Run:   runUpper,
}

func runUpper(cmd *cobra.Command, args []string) {
	if len(args) >= 1 {
		str := args[0]
		fmt.Println(strings.ToUpper(str))
	} else {
		fmt.Println("String is missing")
	}
}

func init() {
	RootCmd.AddCommand(upperCmd)
}
