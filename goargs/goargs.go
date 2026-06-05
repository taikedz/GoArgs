package goargs

/* This file describes the usage intent for the new goargs parser

It specifically DOES NOT try to be a drop-in replacement for `flags`

Minimal requirements

- positional and flag options can appear in any order
- short flags available for options
- support for specifying mutually exclusive options
- support for basic types int, uint, string, float, bool
- support for adding a parser of custom type
*/

import "github.com/taikedz/goargs/goargs"

func Parse(tokens []string) {

	//* Intended usage:

	// Create a new parser with a title
	main_parser := goargs.NewParser("A utility to create a user")

	var username string
	var uid int
	create_home := false
	var mode string

	// Subcommand
	var subc string
	subc_action := main_parser.Subcommand([]string{"create", "delete"}, "Create or delete a user")
	create_parser := subc_action.ParserFor("create")

	// Positional arguments have names not starting with "-"
	create_parser.String(&username, "name", "Name of user account")
	create_parser.Int(&uid, "uid", "UID numeral")

	// Optional
	create_parser.String(&create_home, "--create-home", "Whether to create a home dir for the user").Short("-H")

	mode_param := create_parser.Mode(&mode, []string{"user", "admin", "service"}, "What type of user to create")
	mode_param.Short("-m")
	mode_param.ModeShorts([]rune{'U', 'A', 'S'}) // 1:1 corresponding to choices

	main_parser.Parse(
		tokens,
		true, // require flags recognised
	)
	// */
}
