package main

import (
	"fmt"
	"os"
	"strings"
)

func ParseArguments(args []string) {
	if len(args) != 2 {
		fmt.Println("Too many or insufficient arguments, please run `repooster owner/name`")
		PrintHelp()
		os.Exit(1)
	}

	if args[1] == "version" || args[1] == "-v" || args[1] == "--version" {
		version := GetCLIVersion()
		fmt.Printf("repooster version %s\n", version)
		os.Exit(0)
	}

	if args[1] == "help" || args[1] == "-h" || args[1] == "--help" {
		fmt.Println("repooster kickstarts the initial configurations of scaffolded GitHub repository")
		PrintHelp()
		os.Exit(0)
	}

}

func ValidateRepoString(r string) {
	if len(strings.Split(r, "/")) != 2 {
		fmt.Printf("Accepted: %s\n", r)
		fmt.Println("Invalid arguments, argument should be `owner/name` format.")
		PrintHelp()
		os.Exit(1)
	}
}

func PrintHelp() {
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  repooster [args|subcommands]")
	fmt.Println("")
	fmt.Println("Arguments:")
	fmt.Println("  owner/name	Strings of GitHub repository")
	fmt.Println("")
	fmt.Println("Subcommands:")
	fmt.Println("  version	Prints the repooster CLI version")
	fmt.Println("  help		Print this help")
}
