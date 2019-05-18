package util

import "fmt"

func ExampleParseSetting() {
	fmt.Println(ParseSetting("*color0:                          #1e1e20"))
	// Output:
	// {color0 #1e1e20}
}

func ExampleReadConfig() {
	config, _ := ReadConfig("./config-test.txt")

	for _, s := range config {
		fmt.Println(s)
	}

	// Output:
	// {color0 #1e1e20}
	// {color8 #e6a57a}
	// {color1 #e6a57a}
	// {color9 #e6a57a}
	// {color2 #e39866}
	// {color10 #e39866}
	// {color3 #df8b54}
	// {color11 #df8b54}
	// {color4 #dc7f41}
	// {color12 #dc7f41}
	// {color5 #85678f}
	// {color13 #c6723a}
	// {color6 #b06534}
	// {color14 #b06534}
	// {color7 #f1cbb3}
	// {color15 #f1cbb3}

}
