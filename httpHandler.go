package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func handlerHash(w http.ResponseWriter, r *http.Request) {
	// handle all request methods
	switch r.Method {
	case http.MethodGet:
		http.Error(w, "GET request method not allowed.", 405)
	case http.MethodPost:
		user := &User{}
		// tell client to expect JSON on return
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("Request received...processing...")
		// arbitrary 5 second delay
		time.Sleep(5 * time.Second)
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
	case http.MethodPut:
		http.Error(w, "PUT request method not allowed.", 405)
	case http.MethodDelete:
		http.Error(w, "DELETE request method not allowed.", 405)
	default:
		http.Error(w, "Invalid request method.", 405)
	}

}
