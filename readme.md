# Password Hasher

### A password hasher in Go

* Can return a base64 representation of a SHA512 encoded string, provided over HTTP or the command line
* Can be run as a command line program or as a server
* Supports graceful shutdown to process all open requests before quitting
* Can provide statistics on average time of all requests, and total requests made to the server

### Instructions:

1. You'll need to have Go (1.8 or later) installed and setup properly
    * Best to see here: [Golang Install](https://golang.org/doc/install)
2. After that is setup, download or clone repo into proper directory within your Go code folder.
3. Start the program by entering `go run *.go` which will compile, build, and run the program.
4. From here are presented with 3 options:
    * 1 is to run the program as a simple command line interface
    * 2 is to run the program as a simple server that accepts a `password` as a form field to /hash and returns the hash over HTTP
    * 3 is the same as 2, but supports graceful shutdown and a /stats endpoint
5. Go forth and enjoy

