# GoArgs - a better Go Arguments Parser

Go's default `flag` library is rudimentary; GoArgs aims to provide a simple yet more featureful module for parsing arguments.

Notably:

* Flags can appear intermixed with positional arguments
* Short flags can be combined with single-dash notation
* Long flags are specified with double-dash notation
* Parser provides a function taking a user-defined token list
* Parser can opt to ignore unknown arguments, or error on unknown arguments, as-needed.

This basic feature set allows the user to build up a rich subcommands structure with flexible argument positioning.

