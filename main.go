package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var apiKey = goDotEnvVariable("apiKey")
var dbPath = goDotEnvVariable("dbPath")
var dbName = goDotEnvVariable("dbName")
var dbUser = goDotEnvVariable("dbUser")
var dbPass = goDotEnvVariable("dbPass")
var port = goDotEnvVariable("port")

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/createUser", checkApiKey(createUser)).Methods("POST")
	r.HandleFunc("/updateName", checkApiKey(updateName)).Methods("PUT")
	r.HandleFunc("/toggleDarkMode", checkApiKey(toggleDarkMode)).Methods("PUT")
	r.HandleFunc("/deleteUser", checkApiKey(deleteUser)).Methods("DELETE")
	r.HandleFunc("/listUsers", checkApiKey(listUsers)).Methods("GET")
	r.HandleFunc("/search", checkApiKey(search)).Methods("GET")

	fmt.Println("server running on " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func checkApiKey(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hApiKey := r.Header.Get("APIKey")

		if hApiKey != apiKey {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("401 - Unauthorized"))
			return
		}

		next(w, r)
	}
}
