package main

import (
	"fmt"
)

func passwordCLineEntry() []byte {
	fmt.Println("Enter a password to hash:")
	var pwd string
	// storing input
	_, err := fmt.Scan(&pwd)
	checkError(err)
	return []byte(pwd)
}
