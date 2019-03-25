package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	template "text/template"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

type Config struct {
	Xresource Xresource `json:"xresources"`
	I3Config  I3Config  `json:"i3wm"`
}

type Xresource struct {
	Background  string `json:"background"`
	Foreground  string `json:"foreground"`
	CursorColor string `json:"cursorcolor"`
	Color0      string `json:"color0"`
	Color1      string `json:"color1"`
	Color2      string `json:"color2"`
	Color3      string `json:"color3"`
	Color4      string `json:"color4"`
	Color5      string `json:"color5"`
	Color6      string `json:"color6"`
	Color7      string `json:"color7"`
	Color8      string `json:"color8"`
	Color9      string `json:"color9"`
	Color10     string `json:"color10"`
	Color11     string `json:"color11"`
	Color12     string `json:"color12"`
	Color13     string `json:"color13"`
	Color14     string `json:"color14"`
	Color15     string `json:"color15"`
}

type I3Config struct {
	ClientBackground      string `json:"client.background"`
	ClientFocused         string `json:"client.focused"`
	ClientUnfocused       string `json:"client.unfocused"`
	ClientFocusedInactive string `json:"client.focused_inactive"`
	ClientUrgent          string `json:"client.urgent"`
	ClientPlaceholder     string `json:"client.placeholder"`
}

var swapCmd = &cobra.Command{
	Use:   "swap",
	Short: "Swap resource files with chosen theme",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadFile("./themes/" + args[0] + ".json")
		check(err)
		var config Config
		err = json.Unmarshal(data, &config)
		check(err)
		t := template.Must(template.ParseFiles("templates/xresources.template"))
		dst, err := os.Create("tmp/.Xresources")
		check(err)
		defer dst.Close()
		err = t.Execute(dst, config.Xresource)
		check(err)
		err = dst.Sync()
		check(err)
		t = template.Must(template.ParseFiles("templates/i3config.template"))
		dst, err = os.Create("tmp/.i3config")
		check(err)
		defer dst.Close()
		err = t.Execute(dst, config.I3Config)
		check(err)
		fmt.Println("Successfully generated new config files...")
		// remove old, move new
		home, err := homedir.Dir()
		check(err)
		oldLocation1 := home + "/.Xresources"
		oldLocation2 := home + "/.i3/config"
		newLocation1 := "./tmp/old.Xresources"
		newLocation2 := "./tmp/old.i3config"
		newFile1 := "./tmp/.Xresources"
		newFile2 := "./tmp/.i3config"
		fmt.Println("Moving old config files to tmp directory...")
		err = os.Rename(oldLocation1, newLocation1)
		check(err)
		err = os.Rename(oldLocation2, newLocation2)
		check(err)
		fmt.Println("Moving new configs to default location...")
		err = os.Rename(newFile1, oldLocation1)
		check(err)
		err = os.Rename(newFile2, oldLocation2)
		check(err)
		fmt.Println("Success, please run 'xrdb ~/.Xresources', restart i3 and kill your current urxvt")
		// run xrdb
		// maybe prompt to kill current urxvt session
	},
}

func init() {
	RootCmd.AddCommand(swapCmd)
}
