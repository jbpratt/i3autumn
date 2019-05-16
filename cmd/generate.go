package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a template from a config file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO")
	},
}

func init() {
	RootCmd.AddCommand(generateCmd)
}
