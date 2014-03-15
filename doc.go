/*

GO Bindings for Csound6

This wrapper is still very experimental. It has been tested only on Linux.
It needs a proper installation of Csound6 with header files in the include path in the csound directory
(e.g. csound/csound.h). libcsound64 and libcsnd6 have to be in the PATH.

You can install this package with go get:

  go get github.com/fggp/go-csnd6

Or you can download a zip archive of the project using the 'Download ZIP' button on the right.
You'll get a zip file named 'go-csnd6-master.zip'. Decompressing it you'll get a directory named 'go-csnd6-master'.
Rename this directory to 'go-csnd6' and move it to '$GOPATH/src/github/fggp'. Enter into
the '$GOPATH/src/github/fggp/go-csnd6' directory. You can eventually adapt the #cgo directives
in csnd.go to your system. Finally install the package with `go install`.

This wrapper is intended to be used with a double build of csound6.

Go 1.2 is needed for the C++ support of go build

Use examples can be seen here: https://github.com/kunstmusik/csoundAPI_examples/tree/master/go
*/
package csnd6
