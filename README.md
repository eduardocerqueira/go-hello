# Go basic

playing with Go and VCode ...

pre-requisites:
* [Go installed](https://golang.org/dl/)
* [Vcode-go extension](https://github.com/golang/vscode-go)

```
# first time, installing and running
go version
go install .
export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
hello

# after changing any code
go build
go install .
hello
```

running tests
```
cd morestrings
go test
```

if needed, helpful commands:

```
go get github.com/google/go-cmp/cmp
go clean -modcache
```

## links
* https://golang.org/doc/code
* https://github.com/golang/vscode-go
* https://code.visualstudio.com/docs/languages/go
* https://golang.org/doc/tutorial/create-module
* https://gobyexample.com
* https://www.digitalocean.com/community/tutorials/debugging-go-code-with-visual-studio-code
