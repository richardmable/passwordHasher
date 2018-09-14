package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var stats = &Stats{}

func handlerHash(w http.ResponseWriter, r *http.Request) {
	// run stats generator each time method is called
	defer statsGenerator(time.Now(), stats)
	// check for POST method
	if r.Method != "POST" {
		http.Error(w, "method not allowed.", 405)
	} else {
		user := &User{}
		// tell client to expect JSON on return
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("Request received...processing...")
		// arbitrary 5 second delay
		// generate a random seed to the sequence differs on each start
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		timeDelay := r1.Intn(7)
		time.Sleep(time.Duration(timeDelay) * time.Second)
		r.ParseForm()
		for k, v := range r.Form {
			// convert to string
			vString := strings.Join(v, "")
			// ensure we are not hashing other possible form values
			if k == "password" {
				user.HashPwd = hashPassword([]byte(vString))
			}
		}
		// convert to json
		userJson, err := json.Marshal(user)
		checkError(err)
		// response is OK
		w.WriteHeader(http.StatusOK)
		// write JSON response
		w.Write(userJson)
		fmt.Println("Request processed")
	}

}

func handlerStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed.", 405)
	} else {
		fmt.Println("Statistics requested...")
		statsJson, err := json.Marshal(stats)
		checkError(err)
		w.WriteHeader(http.StatusOK)
		w.Write(statsJson)
		fmt.Println("Statistics delivered!")
	}
}
