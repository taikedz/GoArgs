// splat.go
package main

import (
	"fmt"
	"os"

	"github.com/taikedz/goargs/goargs"
)

func main() {
	// Example of extra args section
	// If `--help` appears before the `--` then help is triggered
	//   else it is a literal argument token as part of the data lines
	parser := goargs.NewParser("write {a|w} FILES -- DATALINES")

	if err := parser.ParseCliArgs(); err != nil {
		parser.PrintHelp()
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	// Get the tokens from _after_ the separator. All tokens, as raw string values.
	data_lines := parser.ExtraArgs()

	// Declare a variable, and use the overall package's (not the parser instance's) `.Unpack()` function
	// to split off the number of required tokens (1 here), and return the remainder
	// Splitting off also performs type coercion to some basic types (string, int, float, bool)
	var mode string
	files_list, err := parser.UnpackArgs(0, &mode)
	if err != nil {
		parser.PrintHelp()
		fmt.Printf("ERROR: %v", err)
		os.Exit(1)
	}

	switch mode {
	case "a":
		fmt.Printf("Appending %d lines to files: %v\n", len(data_lines), files_list)
	case "w":
		fmt.Printf("Erasing files %v and setting %d lines of content\n", files_list, len(data_lines))
	default:
		fmt.Println("Invalid mode - choose 'a' or 'w'")
		os.Exit(2)
	}
	fmt.Println("Content:")

	for _, line := range data_lines {
		fmt.Printf("> %s\n", line)
	}
}
