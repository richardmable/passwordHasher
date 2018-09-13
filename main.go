package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Program started...")
	var programOption int64
	fmt.Println("Select a program option by entering a number:\n 1: Command line input to return SHA512 Base64 encoded hash\n 2: Hash and encode passwords over HTTP")
	_, err := fmt.Scan(&programOption)
	checkError(err)
	// command line mode to take a user inputted string and return a base64 SHA512 encoded string
	if programOption == 1 {
		fmt.Println("Program 1, command line input started")
		// infinite loop to take inifinite entries, don't have to restart each time
		for {
			pwd := hashPassword(passwordCLineEntry())
			fmt.Println("Your base64 encoded password:")
			fmt.Println(pwd)
		}
	} else if programOption == 2 {
		// does the same as program 1, but over http, and can handle multiple connections
		fmt.Println("Program 2, http mode started")
		http.HandleFunc("/hash", handlerHash)
		log.Fatal(http.ListenAndServe(":8080", nil))

	}
}
