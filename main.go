package main

import (
	"flag"
	"fmt"
	_ "github.com/kr/pretty"
	"os"
)

func main() {
	os.Exit(realMain(os.Args[1:]))
}

func realMain(args []string) int {

	var dry bool
	var test bool
	var withAuth bool
	var mappingFile string
	var eventFile string
	var hostTemplate string

	flags := flag.NewFlagSet("deployer-tools", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Printf(helpText)
		flags.PrintDefaults()
	}

	flags.BoolVar(&withAuth, "with-auth", false, "Uses the --with-registry-auth flag.")
	flags.BoolVar(&dry, "dry", false, "Prefixes all output with '#' characters.")
	flags.BoolVar(&test, "test", false, "Use to test scripts.")
	flags.StringVar(&mappingFile, "mapping", "", "The location of the mapping file.")
	flags.StringVar(&eventFile, "event", "", "The location of the event file.")
	flags.StringVar(&hostTemplate, "host-template", "", "The template used to build the host DNS entry.")

	if err := flags.Parse(args); err != nil {
		flags.Usage()
		return 1
	}

	if len(args) > 0 && args[0] == "help" {
		flags.Usage()
		return 1
	}

	if test == true {
		fmt.Println("echo 'deployer-tools test'")
		return 0
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
		return 1
	}

	event, err := ReadEvent(eventFile)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	command := NewCommand(
		WithMapping(mapping),
		WithDry(dry),
		WithEvent(event),
		WithHostTemplate(hostTemplate),
		WithAuth(withAuth),
	)
	if command == nil {
		return 1
	}

	err = command.Run()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}

const helpText = `Usage: deployer-tools [options]
  Helps deploy things.
Options:
`
