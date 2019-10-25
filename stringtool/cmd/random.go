package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
	"strconv"
)

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "generate random string",
	Long:  "This subcommand will return a random string",
	Run:   runRandom,
}

func runRandom(cmd *cobra.Command, args []string) {
	if len(args) >= 1 {
		var charLen int
		var err error
		if charLen, err = strconv.Atoi(args[0]); err != nil {
			fmt.Println(err.Error())
		}
		str := random(charLen)
		fmt.Println(str)
	} else {
		fmt.Println("String length is missing")
	}


}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func random(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func init() {
	RootCmd.AddCommand(randomCmd)
}
