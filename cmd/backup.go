package cmd

import (
	"io"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		home, err := homedir.Dir()
		check(err)
		makePackageDirectories(home)
		err = copyFile(home+"/.Xresources", home+"/.i3/autumn/old/.Xresources")
		if err != nil {
			log.Fatal(err)
		}
		err = copyFile(home+"/.i3/config", home+"/.i3/autumn/old/.i3config")
		if err != nil {
			log.Fatal(err)
		}
		err = copyFile(home+"/.i3/status", home+"/.i3/autumn/old/.i3status")
		if err != nil {
			log.Fatal(err)
		}
	},
}

func makePackageDirectories(home string) error {
	log.Println("Checking if temp directories already exist")
	if _, err := os.Stat(home + "/.i3/autumn"); os.IsNotExist(err) {
		log.Println("Making new directory...")
		err := os.MkdirAll(home+"/.i3/autumn/new", os.ModePerm)
		if err != nil {
			return err
		}
		err = os.Mkdir(home+"/.i3/autumn/old", os.ModePerm)
		if err != nil {
			return err
		}
		log.Println("~/.i3/autumn/{new,old} was made...")
		return nil
	}
	// check if {new,old} already exist, make per date dir in old
	log.Println("~/.i3/autumn directory already exists")
	return nil
}

// takes in path (/home/x/.i3/config) copys to newLocation (/home/x/.i3/autumn/old/.i3config)
func copyFile(path, newLocation string) error {
	log.Println("Attempting to copy " + path + " to " + newLocation)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}
	log.Println(path + " was found...")
	src, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer src.Close()
	dst, err := os.Create(newLocation)
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
	log.Println("Success...")
	return nil
}

func init() {
	RootCmd.AddCommand(backupCmd)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
