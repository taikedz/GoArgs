package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/taikedz/goargs/goargs"
)

const WAVE_MOJI string = "👋"

const SALUTE_EN string = "Hello"
const AND_EN string = "and"

const SALUTE_FR string = "Salut"
const AND_FR string = "et"

func main() {
	// Declare a new parser, and its basic help string
	parser := goargs.NewParser("salute NAME [--wave N] [--grouped] [--with SALUTATION] [--lang LANGUAGE]")

	// Declare a string flag "--with", its default value "Hello", and its help string
	// The returned value is a pointer to a memory location of type string
	salutation := parser.String("with", "%", "Salutation to use")
	parser.SetShortFlag('w', "with")

	// Again with an int flag
	wave_count := parser.Int("wave", 0, fmt.Sprintf("Add N hand wave emojis (%s)", WAVE_MOJI))
	parser.SetShortFlag('W', "wave")

	// Again with a bool flag
	grouped := parser.Bool("grouped", false, "Group names to a single salutation")
	parser.SetShortFlag('g', "grouped")

	// Mutually exclusive modes
	mode := parser.Mode("lang", "en", map[rune]string{'e': "en", 'f': "fr"}, "Language to use for salutation")

	// Perform the parse. If `--help` is found amongst the flags, prints the help and exits
	if err := parser.ParseCliArgs(); err != nil {
		fmt.Printf("! -> %v\n", err)
		os.Exit(1)
	}

	// The variable is a pointed, remember to dereference!
	mojis := wave(*wave_count)

	var ph_salute string
	var ph_and string

	if *salutation == "%" {
		switch *mode {
		case "en":
			ph_salute = SALUTE_EN
			ph_and = AND_EN
		case "fr":
			ph_salute = SALUTE_FR
			ph_and = AND_FR
		}
	}

	if *grouped {
		names := parser.Args()
		lead_names := strings.Join(names[:len(names)-1], ", ")
		last_name := names[len(names)-1]

		fmt.Printf("%s %s %s %s %s\n", ph_salute, lead_names, ph_and, last_name, mojis)
	} else {
		for _, name := range parser.Args() {
			fmt.Printf("%s, %s %s\n", ph_salute, name, mojis)
		}
	}
}

func wave(times int) string {
	var hands []string

	for i := 0; i < times; i++ {
		hands = append(hands, WAVE_MOJI)
	}

	return strings.Join(hands, "")
}
