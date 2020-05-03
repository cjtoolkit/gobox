# Gobox

Tool manager for use with go generate.

## Installation

```sh
$ go get github.com/cjtoolkit/gobox
```

# Using gobox

Create `gobox.toml` in the root of the project, with the example below.

```toml
[[local]]
binPath = "tools"
install = "."

[[module]]
repo = "github.com/cjtoolkit/embedder"
tag = "v1.0.0"
binPath = "tools"
installs = [
	"."
]
```

On top of the go source file add on top. This is just an example

```go
//go:generate gobox tools/embedder internal generated_const.go resources/*
```

If the tools are not installed it, will install automatically.

## Note

It was designed for use in a development environment and is not intended to 
be used in production or deployment environment.  It's meant to install tools for
the developers to use with `go generate`.