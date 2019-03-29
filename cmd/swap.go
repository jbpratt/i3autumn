package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	template "text/template"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

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

type Config struct {
	Xresources Xresources `json:"xresources"`
	I3Config   I3Config   `json:"i3wm"`
}

var (
	XresourcesFileName       string = "/.Xresources"
	I3configFileName         string = "/.i3config"
	I3configOriginalFileName string = "/.i3/config"
	I3statusFileName         string = "/.i3status.conf"
	XresourcesTmpl           string = "./templates/Xresources.tmpl"
	I3configTmpl             string = "./templates/i3config.tmpl"
	I3statusTmpl             string = "./templates/i3statusconf.tmpl"
	ThemeExt                 string = ".json"
	ThemeDir                 string = "./themes/"
	PackageDir               string = "/.i3/autumn"
	OldTempDir               string = "/old"
	NewTempDir               string = "/new"
)

var swapCmd = &cobra.Command{
	Use:   "swap",
	Short: "Swap resource files with chosen theme",
	Run: func(cmd *cobra.Command, args []string) {
		// make a backup
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}
		// read theme file
		// unmarshal theme into config struct
		var config Config
		err = readThemeConfig(args[0], &config)
		if err != nil {
			log.Fatal(err)
		}
		err = parseAndExecute(XresourcesTmpl, home, XresourcesFileName, config)
		if err != nil {
			log.Fatal(err)
		}
		err = parseAndExecute(I3configTmpl, home, I3configFileName, config)
		if err != nil {
			log.Fatal(err)
		}
		err = parseAndExecute(I3statusTmpl, home, I3statusFileName, config)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Successfully generated new config files...")
		err = os.Rename(home+XresourcesFileName, home+PackageDir+OldTempDir+XresourcesFileName)
		if err != nil {
			log.Fatal(err)
		}
		err = os.Rename(home+I3configOriginalFileName, home+PackageDir+OldTempDir+I3configFileName)
		if err != nil {
			log.Fatal(err)
		}
		err = os.Rename(home+I3statusFileName, home+PackageDir+OldTempDir+I3statusFileName)
		if err != nil {
			log.Fatal(err)
		}
		err = os.Rename(home+PackageDir+NewTempDir+XresourcesFileName, home+XresourcesFileName)
		if err != nil {
			log.Fatal(err)
		}
		err = os.Rename(home+PackageDir+NewTempDir+I3configFileName, home+I3configOriginalFileName)
		if err != nil {
			log.Fatal(err)
		}
		err = os.Rename(home+PackageDir+NewTempDir+I3statusFileName, home+I3statusFileName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Files have been moved...")
		c := exec.Command("/bin/xrdb", home+XresourcesFileName)
		err = c.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Please restart your current session")
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
	fmt.Println("Successfully parsed theme")
	return nil
}

// parse tmpl
func parseAndExecute(tmpl, home, name string, config Config) error {
	t := template.Must(template.ParseFiles(tmpl))
	dst, err := os.Create(home + PackageDir + NewTempDir + name)
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
