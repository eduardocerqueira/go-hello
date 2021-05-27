# Go basic

play with Go and VCode

pre-requisites:
* [Go installed](https://golang.org/dl/)
* [Vcode-go extension](https://github.com/golang/vscode-go)

```
go version
go install .
export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
hello
```

## links
* https://golang.org/doc/code
* https://github.com/golang/vscode-go
* https://code.visualstudio.com/docs/languages/go
* https://www.digitalocean.com/community/tutorials/debugging-go-code-with-visual-studio-code
