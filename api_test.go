package main

import (
	"bytes"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var newUser = "jsmith"

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}

func TestDeleteUser(t *testing.T) {
	payload := []byte(`{"username": "` + newUser + `"}`)
	req, err := http.NewRequest("DELETE", "/deleteUser", bytes.NewBuffer(payload))

	checkError(err, t)
	q := req.URL.Query()
	q.Add("username", "jsmith")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()

	r := mux.NewRouter()

	r.HandleFunc("/deleteUser", deleteUser).Methods("DELETE")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}

}

func TestCreateUser(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	payload := []byte(`{"firstName":"John", "lastName": "Smith" ,  "username": "` + newUser + `" ,   "darkMode": true }`)
	req, err := http.NewRequest("POST", "/createUser", bytes.NewBuffer(payload))

	checkError(err, t)

	rr := httptest.NewRecorder()

	r := mux.NewRouter()

	r.HandleFunc("/createUser", createUser).Methods("POST")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}

	expected := string(`{"message":"New user created"}`)

	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")
}

func TestDuplicateUsername(t *testing.T) {
	payload := []byte(`{"firstName":"John", "lastName": "Smith" ,  "username": "` + newUser + `" ,   "darkMode": true }`)
	req, err := http.NewRequest("POST", "/createUser", bytes.NewBuffer(payload))

	checkError(err, t)

	rr := httptest.NewRecorder()

	r := mux.NewRouter()

	r.HandleFunc("/createUser", createUser).Methods("POST")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}

	expected := string(`{"message":"Username already used"}`)

	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")
}

func TestSearch(t *testing.T) {

	payload := []byte(`{"searchString": "` + newUser + `"}`)
	req, err := http.NewRequest("GET", "/search", bytes.NewBuffer(payload))

	checkError(err, t)

	rr := httptest.NewRecorder()

	r := mux.NewRouter()

	r.HandleFunc("/search", search).Methods("GET")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}

}

func TestListUsers(t *testing.T) {

	req, err := http.NewRequest("GET", "/listUsers", nil)

	checkError(err, t)

	rr := httptest.NewRecorder()

	r := mux.NewRouter()

	r.HandleFunc("/listUsers", listUsers).Methods("GET")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}

}

func TestUpdateName(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	payload := []byte(`{"firstName":"` + randomString(13) + `", "lastName": "` + randomString(13) + `", "username": "` + newUser + `" }`)
	req, err := http.NewRequest("PUT", "/updateName", bytes.NewBuffer(payload))

	checkError(err, t)

	rr := httptest.NewRecorder()

	r := mux.NewRouter()

	r.HandleFunc("/updateName", updateName).Methods("PUT")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}

	expected := string(`{"message":"User updated"}`)

	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")
}

func TestToggleDarkMode(t *testing.T) {
	payload := []byte(`{"username": "` + newUser + `"}`)
	req, err := http.NewRequest("PUT", "/toggleDarkMode", bytes.NewBuffer(payload))

	checkError(err, t)

	rr := httptest.NewRecorder()

	r := mux.NewRouter()

	r.HandleFunc("/toggleDarkMode", toggleDarkMode).Methods("PUT")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}

	expected := string(`{"message":"User updated"}`)

	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")
}
