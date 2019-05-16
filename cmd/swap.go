package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	template "text/template"

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
	themeChoice    string
	outLocation    string
	templateChoice string
)

var swapCmd = &cobra.Command{
	Use:   "swap",
	Short: "Swap resource file with chosen theme",
	Long: `
	Swap command takes in a theme json file from the
	theme directory and merges it into the theme template provided
	in the template directory. The new file is then saved to the out
	location provided and attempts to restart the application's session`,
	Run: func(cmd *cobra.Command, args []string) {

		if themeChoice == "" {
			log.Fatal(errors.New("no theme was given"))
		}

		if templateChoice == "" {
			log.Fatal(errors.New("no template was given"))
		}

		if outLocation == "" {
			log.Fatal(errors.New("no out location was given"))
		}

		var config Config
		err := readThemeConfig(themeChoice, &config)
		if err != nil {
			log.Fatal(err)
		}

		err = parseAndExecute(templateChoice, outLocation, config)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Please restart your current session")

		if strings.Contains(templateChoice, "Xresources") {
			cmd := exec.Command("/usr/bin/xrdb", "-merge", outLocation)
			err = cmd.Run()
			if err != nil {
				log.Fatalf("failed to merge changes into current xrdb session: %v", err)
			}
		}
	},
}

func readThemeConfig(tc string, config *Config) error {
	data, err := ioutil.ReadFile(tc)
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

func parseAndExecute(tmpl, outPath string, config Config) error {
	t := template.Must(template.ParseFiles(tmpl))
	dst, err := os.Create(outPath)
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
	fmt.Println("Successfully generated new theme")
	return nil
}

func init() {
	swapCmd.Flags().StringVarP(&themeChoice, "theme", "t", "", "path of theme to parse")
	swapCmd.Flags().StringVarP(&templateChoice, "template", "c", "", "application config template to use")
	swapCmd.Flags().StringVarP(&outLocation, "out", "o", "", "location to write the config to")
	RootCmd.AddCommand(swapCmd)
}
