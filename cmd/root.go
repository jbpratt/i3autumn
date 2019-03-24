package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "i3-autumn",
	Short: "i3-autumn is a CLI theme manager for Xresources and i3",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// cmd here
	// },
}
