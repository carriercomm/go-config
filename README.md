# go-config

`go-config` is a library written in Golang designed to provide an easy interface
with which to deal with configuration objects. Unlike many other similar
libraries, `go-config` treats configuration objects as if they are members of a
tree; that is to say that they can have parents and children.  If a value is
missing in a child configuration object, it will look to its parent for the
answer recursively.

## Getting Started

### Directory layout

`go-config` makes certain assumptions about your directory layout.  It assumes
you have a single configuration directory (usually called "config") in which
there are several JSON files corresponding to your configuration sets.  An
example directory layout is presented below:

```
config/
├── default.json
├── development.json
└── production.json

0 directories, 3 files
```

Each file is scoped to a single environment, and all files are JSON-encoded.
Valid environment names are listed below:

  - `"production"`
  - `"staging"`
  - `"development"`
  - `"default"`

### Declaring your environment

`go-config` picks up your environment by searching through OS-level environment
variables.  It searches through an exported variable it assumes to be called
`ENV`.  Values contained in this variable are allowed to be any case, so long as
they are one of the four valid environments as provided above.

### Fetching values

An example is provided below for fetching a value out of a configuration object:

```go
package main

func main() {
        cfg, err := config.NewConfiguration("./config")
        if err != nil {
                panic(err)
        }

        fmt.Println(cfg.String("mysql.user"))
}
```

From the above example, there are a few useful things to point out, namely the
lookup syntax.  For mapped values (`map[string]interface{}`), you are allowed to
delimit those lookups by a `.`.  For array lookups, you are allowed to delimit
the indice of the array you are interested in by a `.`.

#### Examples

```go
func main() {
        // ...

        cfg.String("db.mysql.username")
        cfg.String("dbs.1.username")
        cfg.String("db.mysql.users.0.username")
}
```

## License

MIT.
