package main

import (
	"fmt"
)

func main() {
	fmt.Println("Program started...")
	// infinite loop to take inifinite entries, don't have to restart each time
	for {
		pwd := hashPassword(passwordCLineEntry())
		fmt.Println("Your base64 encoding password:")
		fmt.Println(pwd)
	}
}
