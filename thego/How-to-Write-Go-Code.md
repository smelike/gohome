How to Write Go Code

Introdution

This document demonstrates the development of a simple Go package inside a module and introduces the go tool, the standard way to fetch, build, and install Go modules, packages, and commands.


Code organization

Go programs are organized into packages.

A package is a collection of source files in the same directory that are compiled together.

Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.

A repository contains one ore more modules.
A module is a collection of related Go packages that are released together.
A Go repository typically contains only one module, located at the root of the repository.
A file named go.mod there declares the module path: the import path prefix for all packages within the module. The module contains the packages in the directory containing its go.mod file asa well as subdirectories of that directory, up to the next subdirectory containing another go.mod file (if any).



> go install

// set the default value for an environment variable
> go en -w GOBIN=/somewhere/else/bin

// unset a variable previously set by go env -w, use go env -u:
> go env -u GOBIN

> go list -f '{{.Target}}'


### Importing packages from your module

Let's write a morestrings package and use it from the hello program. First, create a directory for the package named $HOME/hello/morestrings, and then a file named reverse.go in that directory with the following contents:
// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package morestrings

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
Because our ReverseRunes function begins with an upper-case letter, it is exported, and can be used in other packages that import our morestrings package.
Let's test that the package compiles with go build:
$ cd $HOME/hello/morestrings
$ go build
$
This won't produce an output file. Instead it saves the compiled package in the local build cache.

### Importing packages from remote modules
An import path can describe how to obtain the package source code using a revision control system such as Git or Mercurial. The go tool uses this property to automatically fetch packages from remote repositories. For instance, to use github.com/google/go-cmp/cmp in your program:
package main

import (
    "fmt"

    "example/user/hello/morestrings"
    "github.com/google/go-cmp/cmp"
)

func main() {
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
    fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
Now that you have a dependency on an external module, you need to download that module and record its version in your go.mod file. The go mod tidy command adds missing module requirements for imported packages and removes requirements on modules that aren't used anymore.
$ go mod tidy
go: finding module for package github.com/google/go-cmp/cmp
go: found github.com/google/go-cmp/cmp in github.com/google/go-cmp v0.5.4
$ go install example/user/hello
$ hello
Hello, Go!
  string(
-     "Hello World",
+     "Hello Go",
  )
$ cat go.mod
module example/user/hello

go 1.16

require github.com/google/go-cmp v0.5.4
$
Module dependencies are automatically downloaded to the pkg/mod subdirectory of the directory indicated by the GOPATH environment variable. The downloaded contents for a given version of a module are shared among all other modules that require that version, so the go command marks those files and directories as read-only. To remove all downloaded modules, you can pass the -modcache flag to go clean:
$ go clean -modcache

## Testing

> go test