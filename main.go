package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Program started...")
	var programOption int64
	fmt.Println("Select a program option by entering a number:\n 1: Command line input to return SHA512 Base64 encoded hash\n 2: Hash and encode passwords over HTTP")
	_, err := fmt.Scan(&programOption)
	if err != nil {
		log.Println(err)
	}
	if programOption == 1 {
		fmt.Println("Program 1, command line input started")
		// infinite loop to take inifinite entries, don't have to restart each time
		for {
			pwd := hashPassword(passwordCLineEntry())
			fmt.Println("Your base64 encoded password:")
			fmt.Println(pwd)
		}
	} else if programOption == 2 {
		fmt.Println("Program 2, http mode started")
		//http requests
	}
}
