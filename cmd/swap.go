package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	template "text/template"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

type Config struct {
	Xresources Xresources `json:"xresources"`
	I3Config   I3Config   `json:"i3wm"`
}

type Xresources struct {
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

var (
	XResources         string = ".Xresources"
	I3config           string = ".i3config"
	XresourcesOriginal string = "/.Xresources" // this is used to find original configs
	I3configOriginal   string = "/.i3/config"
	XresourcesTmpl     string = "templates/Xresources.tmpl"
	I3configTmpl       string = "templates/i3config.tmpl"
	ThemeExt           string = ".json"
	ThemeDir           string = "themes/"
	TempDir            string = "tmp/"
	OldTempDir         string = "old/"
	NewTempDir         string = "new/"
)

var swapCmd = &cobra.Command{
	Use:   "swap",
	Short: "Swap resource files with chosen theme",
	Run: func(cmd *cobra.Command, args []string) {
		// make a backup

		// read theme file
		// unmarshal theme into config struct
		var config Config
		err := readThemeConfig(args[0], &config)
		if err != nil {
			log.Fatal(err)
		}
		// parse Xresource
		//t := template.Must(template.ParseFiles(XresourcesTmpl))
		//st, err := os.Create(TempDir + XResources)
		//check(err)
		//defer dst.Close()
		// apply parsed tmpl to data object and writes output to dst
		//err = t.Execute(dst, config.Xresources)
		//check(err)
		//err = dst.Sync()
		//check(err)
		err = parseAndExecute(XresourcesTmpl, config.Xresources)
		if err != nil {
			log.Fatal(err)
		}
		err = parseAndExecute(I3configTmpl, config.I3Config)
		if err != nil {
			log.Fatal(err)
		}
		//t = template.Must(template.ParseFiles(I3configTmpl))
		//dst, err = os.Create("tmp/.i3config")
		//check(err)
		//defer dst.Close()
		//err = t.Execute(dst, config.I3Config)
		//check(err)
		fmt.Println("Successfully generated new config files...")
		// remove old, move new
		home, err := homedir.Dir()
		check(err)
		newLocation1 := "./tmp/old.Xresources"
		newLocation2 := "./tmp/old.i3config"
		fmt.Println("Moving old config files to tmp directory...")
		err = os.Rename(home+XresourcesOriginal, newLocation1)
		check(err)
		err = os.Rename(home+I3configOriginal, newLocation2)
		check(err)
		fmt.Println("Moving new configs to default location...")
		err = os.Rename(TempDir+XResources, home+XresourcesOriginal)
		check(err)
		err = os.Rename(TempDir+I3config, home+I3configOriginal)
		check(err)
		fmt.Println("Success, please run 'xrdb ~/.Xresources', restart i3 and kill your current urxvt")
		// run xrdb
		// maybe prompt to kill current urxvt session
	},
}

// read config
func readThemeConfig(theme string, config *Config) error {
	data, err := ioutil.ReadFile(ThemeDir + theme + ThemeExt)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	return nil
}

// parse tmpl
func parseAndExecute(tmpl string, config interface{}) error {
	t := template.Must(template.ParseFiles(tmpl))
	dst, err := os.Create(TempDir + XResources)
	if err != nil {
		return err
	}
	defer dst.Close()
	err = t.Execute(dst, config)
	if err != nil {
		return err
	}
	err = dst.Sync()
	if err != nil {
		return err
	}
	return nil
}

// move files

func init() {
	RootCmd.AddCommand(swapCmd)
}
