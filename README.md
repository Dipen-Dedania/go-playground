# go-playground
Sample programs with GO

Download GO from -> https://golang.org/doc/install?download=go1.10.windows-386.msi

## Hello World

```
go build 01-hello-world.go
```

This will create an .exe file. You can directly run go file by

```
go run 01-hello-world.go
```

## Creating your first API

Make sure that you have GOPATH is set as your environment variable and it's a different than your GO instalation
Now, install a module *mux* to route your request.
```
go get github.com/gorilla/mux
```
this will install a package into virtual environment. Now navigate to first-api.go file and build it. Run exe file and fireup http://localhost:8000/people. That was pretty quick. Now, I don't want that console when I run my application. Yes, you can do it by adding below flag
```
go build -ldflags "-H windowsgui" 03-first-api.go
```