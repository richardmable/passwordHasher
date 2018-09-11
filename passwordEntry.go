package main

import (
	"fmt"
	"log"
)

func passwordCLineEntry() []byte {
	fmt.Println("Enter a password to hash:")
	var pwd string
	// storing input
	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Println(err)
	}
	return []byte(pwd)
}
