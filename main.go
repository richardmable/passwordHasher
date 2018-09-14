package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Program started...")
	var programOption int64
	fmt.Println(
		`Select a program option by entering a number:
        1: Command line input to return SHA512 Base64 encoded hash
        2: Hash and encode passwords over HTTP 
        3: Same as 2, but with the ability to send 
        a GET request to /shutdown to shutdown 
        the server once work is completed, and a /stats endpoint`)
	_, err := fmt.Scan(&programOption)
	checkError(err)
	// command line mode to take a user inputted
	// string and return a base64 SHA512 encoded string
	if programOption == 1 {
		fmt.Println("Program 1, command line input started")
		// infinite loop to take infinite entries, don't have to restart each time
		for {
			pwd := hashPassword(passwordCLineEntry())
			fmt.Println("Your base64 encoded password:")
			fmt.Println(pwd)
		}
		// ideally would set the port in the .env or similar
	} else if programOption == 2 {
		// does the same as program 1, but over http,
		// and can handle multiple connections
		fmt.Println("Program 2, http mode started")
		// set some timeouts
		s := &http.Server{
			Addr:           ":8080",
			Handler:        nil,
			ReadTimeout:    15 * time.Second,
			WriteTimeout:   15 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		http.HandleFunc("/hash", handlerHash)
		log.Fatal(s.ListenAndServe())

	} else if programOption == 3 {
		// does the same as program 2 but provides a
		// /shutdown endpoint to shutdown gracefully
		// and provides a /stats endpoint
		fmt.Println("Program 3, http mode started w/shutdown and stats enabled")
		// set some timeouts
		svr := &http.Server{
			Addr:           ":8080",
			Handler:        nil,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		// handle the hashing
		http.HandleFunc("/hash", handlerHash)
		http.HandleFunc("/stats", handlerStats)
		// don't want the server to block, as we need to check for shutdown signals
		go func() {
			if err := svr.ListenAndServe(); err != http.ErrServerClosed {
				checkError(err)
			}
		}()
		// channels to send signals that shutdown is ok, and can commence
		idleConnsClosed := make(chan struct{})
		sigStop := make(chan bool, 1)
		// run the server shutdown as a goroutine that blocks until shutdown signal is sent
		go gracefulShutdown(svr, idleConnsClosed, sigStop)
		// handle the shutdown request by sending a signal ok to shutdown
		http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				http.Error(w, "method not allowed.", 405)
			} else {
				sigStop <- true
			}

		})
		// block program exit until until all idle connections are closed
		<-idleConnsClosed
	}
}
