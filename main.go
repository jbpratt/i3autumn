package main

import (
	"i3-autumn/cmd"
	"log"
)

func main() {
	check(cmd.RootCmd.Execute())
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
