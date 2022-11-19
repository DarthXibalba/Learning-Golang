# Learning-Golang
Repository for Learning Golang

## Useful Starter Links
- [Learn Go](https://go.dev/learn/)
- [Go by Example](https://gobyexample.com/)
- [Gophercises](https://gophercises.com/)

## Project/Directory Structure & Links
- [Greetings Tutorial / Create a Go Module](https://go.dev/doc/tutorial/create-module)
- [Gophercise - Quiz Game](https://courses.calhoun.io/lessons/les_goph_01)  
  - [Github repo](https://github.com/gophercises/quiz)
- [Gophercise - HTML Link Parser](https://courses.calhoun.io/lessons/les_goph_16)
  - [Github repo](https://github.com/gophercises/link)
- [Gophercise - Sitemap Builder](https://courses.calhoun.io/lessons/les_goph_24)
  - [Github repo](https://github.com/gophercises/sitemap)

# Go Basics Tutorial
## Module Paths
A *module path* is the canonical name for a module, declared with the `module` [directive](https://go.dev/ref/mod#go-mod-file-module) in the module's `go.mod` file. A Module's path is the prefix for the package paths within the module.  
  
A module path should describe both what the module does and where to find it. Typically a module path consists of a repository root path, a directory within the repository (usually empty), and a major version suffix (only for major version 2 or higher).  
- The *repository root path* is the portion of the module path that corresponds to the root directory of the version control repository where the module is developed. Most modules are defined in their repository's root directory, so this is usually the entire path.
- For example, `golang.org/x/net` is the repository root path for the module of the same name. See [Finding a repository for a module path](https://go.dev/ref/mod#vcs-find) for information on how the `go` command locates a repository using HTTP requests derived from a module path.

## go.mod files
A module is defined by a UTF-8 encoded text file name `go.mod` in its root directory. The `go.mod` file is line-oriented. Each line holds a single directive, made up of a keyword followed by arguments. For example:
```
module example.com/my/thing

go 1.12

require example.com/other/thing v1.0.2
require example.com/new/thing/v2 v2.3.4
exclude example.com/old/thing v1.2.3
replace example.com/bad/thing v1.4.5 => example.com/good/thing v1.4.5
retract [v1.9.0, v1.9.5]
```
The leading keyword can be factored out of adjacent lines to create a block, like in Go imports.
```
require (
    example.com/new/thing/v2 v2.3.4
    example.com/old/thing v1.2.3
)
```
The `go.mod` file is designed to be human readable and machine writable. The `go` command provides several subcommands that change `go.mod` files.

### [go mod init](https://go.dev/ref/mod#go-mod-init)
Usage:
```
go mod init [module-path]
```
Example:
```
go mod init
go mod init example.com/m
```
The `go mod init` command initializes and writes a new `go.mod` file in the current directory, in effect creating a new module rooted at the current directory. The `go.mod` file must not already exist.  
  
`init` accepts one optional argument, the [module path](https://go.dev/ref/mod#glos-module-path) for the new module. See [Module paths] for instructions on choosing a module path. If the module path argument is omitted, `init` will attempt to infer the module path using import comments in `.go` files, vendoring tool configuration files, and the current directory (if in `GOPATH`).

### [go mod tidy](https://go.dev/ref/mod#go-mod-tidy)
Usage:
```
go mod tidy [-e] [-v] [-go=version] [-compat=version]
```

`go mod tidy` ensures that the `go.mod` file matches the source code in the module. It adds any missing module requirements necessary to build the current module's packages and dependencies, and it removes requirements on modules that don't provide any relevant packages. It also adds any missing entires to `go.sum` and removes unnecessary entries.

# Instructions for starting a new Go Project
### 1. Create a module direcotry for your Go source code
```
mkdir %moduleName%
cd %moduleName%
```
### 2. Enable dependency tracking for your code
When your code imports packages contained in other modules, you manage those dependencies through your code's own module. That module is defined by a go.mod file that tracks the modules taht provide those packages. That go.mod file stays with your code, including in your source code repository.  
  
To enable depnedncy tracking for your code by creating a go.mod file, run the `go mod init` [command](#go-mod-init), giving it the name of the module your code will be in. The name is the module's module path.  
  
In actual development, the module path will typically be the repository location where your source code will be kept. For example, the module path might be `github.com/mymodule`. If you plan to publish your module for others to use, the module path *must* be a location from which Go tools can download your module. For more about naming a module with a module path, see [Managing dependencies](https://go.dev/doc/modules/managing-dependencies#naming_module).  
  
For the purposes of this README, use `example/mymodule`:
```
$ go mod init example/mymodule
go creating new go.mod: module example/mymodule
```

### 3. Create a mymodule.go file
```
package main

import "fmt"

func main() {
  fmt.Println("Hellow, World!")
}
```
In this code, you:
- Declare a `mymodule` package (a package is a way to group functions, and it's made up of all the files in the same directory).
- Import the popular `fmt` package, whcih contains functions for formatting text.
- Implement a `main` function to print a message to the console. A `main` function executes by default when you run the `mymodule` package.

### 4. Run your code
```
$ go run .
Hello, World!
```

### 5. Call code in an external package
In your code, import the `rsc.io/quote` package and add a call to its `Go` function.
```
package mymodule

import "fmt"

import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}
```

### 6. Add/Update new module requirements and sums
Go will add the `quote` module as a requirement, as well as a go.sum file for use in [authenticating the module](https://go.dev/ref/mod#authenticating).
```
$ go mod tidy
go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.2
```

## Package Naming
- short, no camel/snake case
- descriptive
- same as directory name
- small and many packages

### Package Scope
- block
    - for example, a variable can live within a block (such as a function {} block)
- unexported
    - if a variable resides outside of a function block and starts with a lowercase letter, then this variable is only accessible within the same package
- exported
    - if a variable resides outside of a function block and starts with an Uppercase letter, then this variable can be exported and will be accessible outside of the package
  
`Note: Circular References (for packages) are not allowed`  

## Directory Structure
Isolate each module to it's own seperate subdirectory:
```
repo/
--module1/
------module1.go
--module2/
------module2.go
------module3/
----------module3.go
--main.go
--go.mod
--go.sum
```

### Example Directory Structure for a bigger project
```
/app
----/handlers
-------- handlers.go  { ninja handlers, dojo handlers }
----/models
-------- models.go  { ninja struct, dojo struct }
-------- requests.go
-------- responses.go
----/persistence
-------- persistence.go
----/main
--------main.go  { main function, router setup, server kickoff }
```

#### This structure adheres to:
1. Simple and straight-forward naming schemes
2. Main and non-main packages
3. Minimum nesting
4. Packages declared in these files will match with the directory names

## Cmdline Setup
1. Run for each module
```
go mod init github.com/darthxibalba/%repository%/%module%
```
  
2. Replace github.com with a local directory in `go.mod` file
```
replace github.com/darthxibalba/%repository%/%module% => ../../%module%
```
  
3. Run to update module
```
go mod tidy
```