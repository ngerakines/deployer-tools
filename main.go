package main

import (
	"fmt"
	"os"
	_ "github.com/kr/pretty"
	"flag"
)

func main() {
	os.Exit(realMain(os.Args[1:]))
}

func realMain(args []string) int {

	var silent bool
	var mappingFile string
	var eventFile string
	var hostTemplate string

	flags := flag.NewFlagSet("deployer-tools", flag.ExitOnError)
  flags.Usage = func() {
        fmt.Printf(helpText)
        flags.PrintDefaults()
    }

	flags.BoolVar(&silent, "silent", true, "Operate silently with not output. Defaults to true.")
	flags.StringVar(&mappingFile, "mapping", "", "The location of the mapping file.")
	flags.StringVar(&eventFile, "event", "", "The location of the event file.")
	flags.StringVar(&hostTemplate, "host-template", "", "The template used to build the host DNS entry.")

	if err := flags.Parse(args); err != nil {
    if silent == false {
      flags.Usage()
    }
    return 1
  }

  if len(args) > 0 && args[0] == "help" {
    flags.Usage()
    return 1
  }

	if mappingFile == "" || eventFile == "" {
    flags.Usage()
    return 1
  }

	if hostTemplate == "" {
    hostTemplate = "manager01.{{.Cluster}}.internal:2375"
  }

	mapping, err := ReadMapping(mappingFile)
	if err != nil {
		fmt.Println(err)
		return 1;
	}

	event, err := ReadEvent(eventFile)
	if err != nil {
		fmt.Println(err)
		return 1;
	}

	// fmt.Printf("%# v\n", pretty.Formatter(mapping))
	// fmt.Printf("%# v\n", pretty.Formatter(event))

	switch event.Type {
  	case "service.update":
			// fmt.Println("Updating services.")
			if err = ServiceUpdate(mapping, event, hostTemplate); err != nil {
				fmt.Println(err)
				return 1;
			}
		default:
			fmt.Printf("Unknown event type '%s'.", event.Type)
		}

	return 0
}

const helpText = `Usage: deployer-tools [options]
  Helps deploy things.
Options:
`
