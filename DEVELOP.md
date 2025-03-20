# Development notes for this branch

Rename package to `goptions` - option resolution for Go

Provide an options resolution mechanism:

```go
// New parser
parser := goptions.NewParser("People")

// A flag
name := parser.String("name", "unknown", "Their name")

/* How to map the Json fields to options
 * Example data:
 *     {"person": {"name":"Alix"} }
 * Name is from flag name to Json path
 */
config_map := map[string]string := {"name": "person/name"}

// Config files: resolved from earliest to latest
//   NewJsonConfig(config_map map[string]string) allows mapping
//   config paths to flag names
config := goptions.NewJsonConfig(config_map)

// Use NewEnvConfig() for key-value pairs that populate values if not specified on CLI

// Add a series of locations in which to find config files
config.AddFileSources("config.json", []string{
    "%/defaultConfig", // path relative to the current binary ("%/" notation)
    "/etc/myapp", //  a global config
    "~/.local/etc/myapp" // a user home config ("~/" notation)
})

// Add a single config - resolves after all the above
config.AddFile("./myapp-config.json") // simply "this" is implicitly "./this"

parser.AddConfig(config)

// Add an environment variable that resolves if flag is not specified on CLI
// Overrides file specs
parser.AddEnv("MYAPP_USERNAME", "name")

parser.Parse()

age := parser.Config().getInt("age") // get arbitrary values from config
```

Order of resolution is:

* files, in declaration order
* environment variables
* CLI flags

This allows the package to take care of options configurations.

