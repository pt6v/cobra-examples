package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	rootCmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Simple command executed.")
		},
	}
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

