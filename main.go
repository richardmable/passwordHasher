package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("Program started...")
	var programOption int64
	fmt.Println("Select a program option by entering a number:\n 1: Command line input to return SHA512 Base64 encoded hash\n 2: Hash and encode passwords over HTTP")
	_, err := fmt.Scan(&programOption)
	if err != nil {
		log.Println(err)
	}
	// command line mode to take a user inputted string and return a base64 SHA512 encoded string
	if programOption == 1 {
		fmt.Println("Program 1, command line input started")
		// infinite loop to take inifinite entries, don't have to restart each time
		for {
			pwd := hashPassword(passwordCLineEntry())
			fmt.Println("Your base64 encoded password:")
			fmt.Println(pwd)
		}
		// does the same as program 1, but over http, and can handle multiple connections
	} else if programOption == 2 {
		fmt.Println("Program 2, http mode started")
		// same encoding as first program
		// delay 5 seconds
		// use a goroutine
		// handle multiple connections
		// how to test multiple connections
		service := ":8080"
		tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
		checkError(err)
		listener, err := net.ListenTCP("tcp", tcpAddr)
		checkError(err)
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println(err)
				continue
			}
			go handleClient(conn)
		}
		// http.HandleFunc("/", handler)
		// log.Fatal(http.ListenAndServe(":8080", nil))

	}
}

func handleClient(conn net.Conn) {
	// be sure connections close
	defer conn.Close()
	time.Sleep(5 * time.Second)
	daytime := time.Now().String()
	fmt.Printf("Connected at: ", daytime)
	conn.Write([]byte(daytime))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// Change your program so that when launched your code starts and listens for HTTP requests on a provided port. Accept ​POST​ requests on the ​/hash​ endpoint with a form field named password ​to provide the value to hash. The response should be the base64 encoded string of the SHA512 hash of the provided password. The server should not respond immediately, it should leave the socket open for 5 seconds before responding.
