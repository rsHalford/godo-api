/*
Copyright © 2021 Richard Halford <richard@xhalford.com>
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rsHalford/goapi/model"
)

var (
	username = hasher("admin")
	password = hasher("secret")
	realm    = "Please enter your username and password to gain access to this API"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage")
}

func BasicAuth(handler http.HandlerFunc, username, password []byte, realm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare(hasher(user), username) != 1 || subtle.ConstantTimeCompare(hasher(pass), password) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			http.Error(w, "Unauthorised Access.", http.StatusUnauthorized)
			return
		}
		handler(w, r)
	}
}

func hasher(s string) []byte {
	value := sha256.Sum256([]byte(s))
	return value[:]
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HomePage)
	router.HandleFunc("/api/v1/todo", BasicAuth(model.CreateTodo, username, password, realm)).Methods("POST")
	router.HandleFunc("/api/v1/todo", BasicAuth(model.GetTodos, username, password, realm)).Methods("GET")
	router.HandleFunc("api/v1/todo/{id}", BasicAuth(model.GetTodo, username, password, realm)).Methods("GET")
	router.HandleFunc("api/v1/todo/{id}", BasicAuth(model.UpdateTodo, username, password, realm)).Methods("PUT")
	router.HandleFunc("api/v1/todo/{id}", BasicAuth(model.DeleteTodo, username, password, realm)).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	model.InitDB()
	handleRequests()
}
