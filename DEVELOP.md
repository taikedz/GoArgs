# Development notes for this branch

Rename to `tk-opts` and place arg parsing and options loading in separate submodules.

```go
import (
    "github.com/taikedz/tk-opts/tkargs"
    "github.com/taikedz/tk-opts/tkconf"
)

file_sources := []string{
    "%/config/myapp.json", // "%" is "path to current binary
    "/etc/myapp/myapp.json",
    "~/.config/myapp/myapp.json",
    "./myapp.json",
//    "other-myapp.json" // INAVLID - needs a path type (includes "/" like in "%/", "/", "~/", "./")
}
config := tkconf.NewJsonLoader(file_sources)
// or goconfig.NewEnvParser(file_sources)
// or goconfig.NewIniParser(file_sources)
// or goconfig.NewYamlParser(file_sources) // optional, since Yaml is likely an external dependency again

// A variable to populate - we will use its reference
var name string

/* How to extract a Json field
 * Example data:
 *     {"person": {"name":"Alix"} }
 * providing a hard-coded default
 */
name_p := config.String("/person/name", "Jason", "MYAPP_USERNAME") // includes an environment variable override. Use "" for none?
//config.StringVar(&name, "/person/name", "Jason", "MYAPP_USERNAME") // auto-create var pointer

// For env, just a key: config.String("name", "Jason", "MYAPP_USERNAME")
// For ini, section and key: config.String("User:name", "Jason", "MYAPP_USERNAME")
// For ini, default section and key: config.String(":name", "Jason", "MYAPP_USERNAME")

config.Parse() // actively loads file data, and sets values/defaults
// _After_ which, we can get the CLI options and override as needed

parser := tkargs.NewParser("People")
config.StringVar(&name, "name", name, "Name of user")
//config.StringVar(name_p, "name", name, "Name of user") // re-use auto-created var pointer
```

Order of resolution of `tkconf` is:

* files, in declaration order, later overriding previous
* environment variables, overriding files

This allows the package to take care of options configurations.
