package cmd

import (
	"io"
	"log"
	"os"
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
		err = copyi3Config(home, d)
		if err != nil {
			log.Fatal(err)
		}
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

// consolidate these copy funcs to one
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

func copyi3Config(path, date string) error {
	log.Println("Attempting to copy ~/.i3/config into tmp/old...")
	if _, err := os.Stat(path + "/.i3/config"); os.IsNotExist(err) {
		return err
	}
	log.Println(".i3/config was found...")
	src, err := os.Open(path + "/.i3/config")
	if err != nil {
		return err
	}
	defer src.Close()
	log.Println("Found file, attempting to copy...")
	dst, err := os.Create("tmp/old/" + time.Now().String() + ".i3config")
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
	log.Println("Successfully copied .i3/config to the temporary directory")
	return nil
}

func copyi3StatusConf(path, date string) error {
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
