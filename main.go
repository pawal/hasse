package main

import (
	"fmt"
	"os"

	"github.com/pawal/go-hass"
)

func main() {
	h := hass.NewAccess("http://localhost:8123", "")
	err := h.CheckAPI()
	if err != nil {
		panic(err)
	}
	args := os.Args[1:]
	if len(args) == 0 {
		list, err := h.FilterStates("switch", "lock", "light")
		if err != nil {
			panic(err)
		}
		for d := range list {
			fmt.Printf("%s (%s): %s\n", list[d].EntityID,
				list[d].Attributes.FriendlyName,
				list[d].State)
		}
		os.Exit(0)
	}

	// assume to only show status on device
	if len(args) == 1 {
		s, err := h.GetState(args[0])
		if err != nil {
			fmt.Printf("Something went wrong: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s state: %s\n", args[0], s.State)
		os.Exit(0)
	}

	// assume change state on device
	if len(args) == 2 {
		s, err := h.GetState(args[0])
		if err != nil {
			fmt.Printf("Something went wrong: %s\n", err)
			os.Exit(1)
		}
		dev, err := h.GetDevice(s)
		if err != nil {
			fmt.Printf("Something went wrong: %s\n", err)
			os.Exit(1)
		}
		if args[1] == "on" {
			dev.On()
		} else if args[1] == "off" {
			dev.Off()
		} else if args[1] == "toggle" {
			dev.Toggle()
		} else {
			fmt.Printf("Unrecognized command '%s'\n", args[1])
		}
	}
}
