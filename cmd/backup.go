package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var backupPath string

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "backs up the given file to a tmp directory",
	Run: func(cmd *cobra.Command, args []string) {
		if backupPath == "" {
			log.Fatal("provide a file path to backup")
		}
		err := backupFile(backupPath)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func backupFile(path string) error {
	// copy file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// write to tmp dir
	file, err := ioutil.TempFile(os.TempDir(),
		fmt.Sprintf("i3-autumn-*%s", filepath.Base(path)))
	if err != nil {
		return err
	}

	if _, err = file.Write(data); err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	fmt.Println("file made at", file.Name())

	return nil
}

func init() {
	backupCmd.Flags().StringVarP(&backupPath, "file", "f", "", "file to backup")
	RootCmd.AddCommand(backupCmd)
}
