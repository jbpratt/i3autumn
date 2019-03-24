package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		dt := time.Now()
		d := dt.Format("01-02-2006")
		fmt.Println("Making tmp directory...")
		err := os.MkdirAll("./tmp/", os.ModePerm)
		check(err)
		home, err := homedir.Dir()
		check(err)
		copyXResources(home, d)
		copyi3Config(home, d)
	},
}

func copyXResources(path, date string) {
	src, err := os.Open(filepath.Join(path, "/.Xresources"))
	check(err)
	fmt.Println("Found file " + src.Name() + " ...")
	defer src.Close()
	dst, err := os.Create("tmp/" + date + ".Xresources")
	fmt.Println("Copying " + src.Name() + " to " + dst.Name())
	check(err)
	defer dst.Close()
	_, err = io.Copy(dst, src)
	check(err)
	fmt.Println("Successfully copied files..")
	err = dst.Sync()
	check(err)
}

func copyi3Config(path, date string) {
	src, err := os.Open(filepath.Join(path, "/.i3/config"))
	check(err)
	fmt.Println("Found file " + src.Name() + " ...")
	defer src.Close()
	dst, err := os.Create("tmp/" + date + ".i3.config")
	fmt.Println("Copying " + src.Name() + " to " + dst.Name())
	check(err)
	defer dst.Close()
	_, err = io.Copy(dst, src)
	check(err)
	fmt.Println("Successfully copied files..")
	err = dst.Sync()
	check(err)
}
func init() {
	RootCmd.AddCommand(backupCmd)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
