package main

import (
	"fmt"
	"os"
)

var Version string

func main() {
	os.Exit(realMain())
}

func realMain() int {
  if len(os.Args) == 1 {
    printUsage()
    return 1
  }
  switch os.Args[1] {
  case "version":
    printVersion();
    return 0;
	default:
    printUsage()
		fmt.Printf("\nError: %q is not valid command.\n", os.Args[1])
		return 1;
	}
}

func printVersion() {
	fmt.Fprintf(os.Stderr, "%s\n", Version)
}

func printUsage() {
	fmt.Printf(helpText)
}

const helpText = `Usage: configr [command] [options]
  Configr configures things.
Commands:
  version
`
