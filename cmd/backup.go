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
		home, err := homedir.Dir()
		check(err)
		makePackageDirectories(home)
		err = copyXresources(home, d)
		if err != nil {
			log.Fatal(err)
		}
		//copyi3Config(home, d)
	},
}

func makePackageDirectories(home string) error {
	log.Println("Checking if temp directories already exist")
	if _, err := os.Stat("tmp"); os.IsNotExist(err) {
		log.Println("Making new temporary directory...")
		err := os.MkdirAll("tmp/old/", os.ModePerm)
		if err != nil {
			return err
		}
		err = os.Mkdir("tmp/new/", os.ModePerm)
		if err != nil {
			return err
		}
		log.Println("tmp/{new,old} was made...")
	}
	log.Println("Temporary directory already exists")
	return nil
}

func copyXresources(path, date string) error {
	log.Println("Attempting to copy ~/.Xresources into tmp/old...")
	if _, err := os.Stat(path + "/.Xresources"); os.IsNotExist(err) {
		return err
	}
	log.Println(".Xresources was found...")
	src, err := os.Open(path + "/.Xresources")
	if err != nil {
		return err
	}
	defer src.Close()
	log.Println("Found file, attempting to copy...")
	dst, err := os.Create("tmp/old/" + time.Now().String() + ".Xresources")
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	err = dst.Sync()
	if err != nil {
		return err
	}
	log.Println("Successfully copied .Xresources to the temporary directory")
	return nil
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
