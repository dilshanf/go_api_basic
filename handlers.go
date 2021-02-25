package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

type User struct {
	ID          int
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Username    string `json:"username"`
	DarkMode    bool   `json:"darkMode"`
	DateCreated string
}

type Search struct {
	SearchString string `json:"searchString"`
}

func createUser(w http.ResponseWriter, r *http.Request) {

	var p User

	err := decodeJSONBody(w, r, &p)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			fmt.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	db := createConnection()

	stmt, e := db.Prepare("insert into users(first_name, last_name, username, dark_mode, date_created) values (?, ?, ?, ?, NOW())")
	internalError(e)

	res, e := stmt.Exec(p.FirstName, p.LastName, p.Username, p.DarkMode)
	if e != nil {

		me, ok := e.(*mysql.MySQLError)
		if !ok {
			response(w, "Request Error", http.StatusNotAcceptable)
		}
		if me.Number == 1062 {
			response(w, "Username already used", http.StatusNotAcceptable)
		}
		return
	}

	id, e := res.LastInsertId()

	fmt.Println("New user created : ", id)
	response(w, "New user created", http.StatusOK)
}

func updateName(w http.ResponseWriter, r *http.Request) {

	var p User

	err := decodeJSONBody(w, r, &p)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			fmt.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	db := createConnection()
	stmt, e := db.Prepare("update users set first_name = ?, last_name = ? where username = ?")
	internalError(e)

	res, e := stmt.Exec(p.FirstName, p.LastName, p.Username)
	if e != nil {
		response(w, "Update Error", http.StatusNotAcceptable)
		return
	}

	a, e := res.RowsAffected()

	if a == 1 {
		response(w, "User updated", http.StatusOK)
	} else {
		response(w, "User not updated", http.StatusExpectationFailed)
	}

}

func toggleDarkMode(w http.ResponseWriter, r *http.Request) {

	var user User

	err := decodeJSONBody(w, r, &user)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			fmt.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	db := createConnection()

	sqlStatement := `select dark_mode from users where username = ?`
	rows := db.QueryRow(sqlStatement, user.Username)

	err = rows.Scan(&user.DarkMode)

	switch err {
	case sql.ErrNoRows:
		response(w, "User not found", http.StatusOK)
		return
	case nil:
		uDarkMode := true
		if user.DarkMode == true {
			uDarkMode = false
		}
		stmt, e := db.Prepare("update users set dark_mode = ? where username = ?")
		internalError(e)

		res, e := stmt.Exec(uDarkMode, user.Username)
		if e != nil {
			response(w, "Update Error", http.StatusNotAcceptable)
			return
		}

		a, e := res.RowsAffected()

		if a == 1 {
			response(w, "User updated", http.StatusOK)
		} else {
			response(w, "User not updated", http.StatusOK)
		}
		return

	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	var p User

	err := decodeJSONBody(w, r, &p)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			fmt.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	db := createConnection()
	stmt, e := db.Prepare("delete from users where username=?")
	internalError(e)

	res, e := stmt.Exec(p.Username)
	if e != nil {
		response(w, "Delete Error", http.StatusNotAcceptable)
		return
	}

	a, e := res.RowsAffected()

	if a == 1 {
		response(w, "User deleted", http.StatusOK)
	} else {
		response(w, "User not deleted", http.StatusOK)
	}

}

func listUsers(w http.ResponseWriter, r *http.Request) {

	getUsers(w, r, "")

}

func search(w http.ResponseWriter, r *http.Request) {

	var p Search

	err := decodeJSONBody(w, r, &p)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			fmt.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	getUsers(w, r, p.SearchString)
}

func getUsers(w http.ResponseWriter, r *http.Request, searchString string) {

	db := createConnection()
	rows, e := db.Query("select first_name, last_name, username, date_created, dark_mode from users WHERE (first_name LIKE '%" + searchString + "%') OR (last_name LIKE '%" + searchString + "%') OR (username LIKE '%" + searchString + "%')")
	internalError(e)

	var jsonUsers []*User
	for rows.Next() {
		users := new(User)
		e = rows.Scan(&users.FirstName, &users.LastName, &users.Username, &users.DateCreated, &users.DarkMode)
		internalError(e)
		jsonUsers = append(jsonUsers, users)
	}

	if err := json.NewEncoder(w).Encode(jsonUsers); err != nil {
		fmt.Println(err)
	}
}
