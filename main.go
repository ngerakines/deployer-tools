package main

import (
	"fmt"
	"os"
	"github.com/kr/pretty"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	if len(os.Args) != 3 {
		printUsage()
		return 1
	}

	mapping, err := ReadMapping(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return 1;
	}

	event, err := ReadEvent(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return 1;
	}

	fmt.Printf("%# v\n", pretty.Formatter(mapping))
	fmt.Printf("%# v\n", pretty.Formatter(event))

	switch event.Type {
  	case "service.update":
			fmt.Println("Updating services.")
			if err = ServiceUpdate(mapping, event); err != nil {
				fmt.Println(err)
				return 1;
			}
		default:
			fmt.Printf("Unknown event type '%s'.", event.Type)
		}

	return 0
}

func printUsage() {
	fmt.Printf(helpText)
}

const helpText = `Usage: deployer-tools [mapping] [event]
  Helps deploy things.
`
