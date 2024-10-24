package utils

import (
	"flag"
	"fmt"
)

func PrintHelp() {
	fmt.Println("mdf - markdown format - is a markdown viewer that displays a .md file live in a browser for better formatting.")
	fmt.Println("\nUsage:")
	fmt.Println("  mdf [options]")
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
	fmt.Println("\nExamples:")
	fmt.Println("  mdf --filepath /path/to/file.md")
	fmt.Println("  mdf --filepath /path/to/file.md --theme dracula")
	fmt.Println("  mdf --cheatsheet")
	fmt.Println("  mdf --help")
}
