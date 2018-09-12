package main

import (
	"fmt"
	"net"
	"time"
)

func handleClient(conn net.Conn) {
	// be sure connections close
	defer conn.Close()
	time.Sleep(5 * time.Second)
	daytime := time.Now().String()
	fmt.Printf("Connected at: %s\n", daytime)
	conn.Write([]byte(daytime))
}
