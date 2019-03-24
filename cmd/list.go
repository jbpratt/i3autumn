package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available themes",
	Run: func(cmd *cobra.Command, args []string) {
		files, err := ioutil.ReadDir("./themes")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Available themes:\n=============================")
		for _, file := range files {
			fmt.Println(strings.TrimSuffix(file.Name(), ".json"))
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
