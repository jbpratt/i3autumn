package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	template "text/template"

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
	ClientUrget           string `json:"client.urgent"`
	ClientPlaceholder     string `json:"client.placeholder"`
}

var swapCmd = &cobra.Command{
	Use:   "swap",
	Short: "Swap resource files with chosen theme",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadFile("./themes/" + args[0] + ".json")
		if err != nil {
			log.Fatal(err)
		}
		var config Config
		err = json.Unmarshal(data, &config)
		if err != nil {
			log.Fatal(err)
		}
		t := template.Must(template.ParseFiles("templates/xresources.template"))
		err = t.Execute(os.Stdout, config.Xresource)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(swapCmd)
}
