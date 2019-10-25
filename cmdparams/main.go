package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	// cobra RunE functions ((Persistent)(Pre|Post)RunE) is more priority than Run functions ((Persistent)(Pre|Post)Run)
	// run functions executing order, please see cobra.Command comments
	rootCmd := &cobra.Command{
		Use:                    "cmdparams",
		Aliases:                nil,
		SuggestFor:             nil,
		Short:                  "cobra.Command example",
		Long:                   `This is a cobra.Command example ,
try to use all of the params to learn how to use it.
`,
		Example:                `./cmdparams -h(--help)
./cmdparams --version
...
`,
		ValidArgs:              nil,
		Args:                   nil,
		ArgAliases:             nil,
		BashCompletionFunction: "",
		Deprecated:             `because of some reason
`,
		Hidden:                 false,
		Annotations:            nil,
		Version:                "0.0.1-alpha",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("cmdparams is persistent pre running...")
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("cmdparams is persistent pre running E...")
			return nil
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("cmdparams is pre running...")
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("cmdparams is pre running E...")
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("cmdparams is running...")
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("cmdparams is running E...")
			return nil
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("cmdparams is post running...")
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("cmdparams is post running E...")
			return nil
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("cmdparams is persistent post running...")
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("cmdparams is persistent post running E...")
			return errors.New("something went wrong")
		},
		SilenceErrors:              false, // change to true hide error streams
		SilenceUsage:               false, // change to true hide usage statement when error occur.
		DisableFlagParsing:         false,
		DisableAutoGenTag:          false,
		DisableFlagsInUseLine:      false,
		DisableSuggestions:         false,
		SuggestionsMinimumDistance: 0,
		TraverseChildren:           false,
		FParseErrWhitelist:         cobra.FParseErrWhitelist{},
	}
	if err := rootCmd.Execute(); err != nil {
	}
}
