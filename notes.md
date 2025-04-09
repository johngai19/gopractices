# Golang notes

## Features

- GO doesn't use inheritance

## go cli commands

- `go get` is used to get a package from a URL with specified version, download and compile it. The package will be added in go.mod file.
- `go mod tidy` is used to update the go.mod file to reflect changes.

## Packages and modules

- Go examines the first letter of the names given to the features in a code file. If the first letter is lowercase then the feature can be used only within the package.
- Using a package alias to avoid name conflit `currencyFmt "packages/fmt"`
- Using a dot import allows user to access the function without using a prefix. `. "packages/store"`
- Using the blank identifier as an alias for an imported package if one just need the effect of the initialization function, `_ "packages/data"`

## Golang resources

### go packages

- <https://pkg.go.dev>
- <https://github.com/golang/go/wiki/projects>
