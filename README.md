# grpc-go-redact
## Why?

GRPC structs generated in Golang automaticly implement the `.String()` func printing out all of the fields. This does blocks writing custom overrides to redact secrets directly in the method ensuring secrets are never printed. 

This package runs after GRPC generation to automaticly overrite the `.String()` method with redaction support.

## Install

*  `go get github.com/Azure/grpc-go-redact` or download the
  binaries from releases page.

* Building from source requires go 1.16+

## Usage

Example:

Run `grpc-go-redact` with generated file `test.pb.go`.

```
grpc-go-redact -input=./test.pb.go
```
The custom .String() method will updated directly in the `test.pb.go` file.

## Things to Note

* Unexported fields will be removed when printing the string.
