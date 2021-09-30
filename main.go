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
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rsHalford/goapi/model"
)

var (
	username = hasher(os.Getenv("GOAPI_USER"))
	password = hasher(os.Getenv("GOAPI_PASSWORD"))
	realm    = "Please enter your username and password to gain access to this API"
)

type Server struct {
	router *mux.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	// Lets Gorilla work
	s.router.ServeHTTP(w, r)
}

func basicAuth(handler http.HandlerFunc, username, password []byte, realm string) http.HandlerFunc {
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

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/favicon-32x32.png")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/", http.FileServer(http.Dir("./static")))
	router.HandleFunc("/favicon-32x32.png", faviconHandler)
	router.HandleFunc("/api/v1/todo", basicAuth(model.CreateTodo, username, password, realm)).Methods("POST")
	router.HandleFunc("/api/v1/todo", basicAuth(model.GetTodos, username, password, realm)).Methods("GET")
	router.HandleFunc("/api/v1/todo/{id}", basicAuth(model.GetTodo, username, password, realm)).Methods("GET")
	router.HandleFunc("/api/v1/todo/{id}", basicAuth(model.UpdateTodo, username, password, realm)).Methods("PUT")
	router.HandleFunc("/api/v1/todo/{id}", basicAuth(model.DeleteTodo, username, password, realm)).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", &Server{router}))
}

func main() {
	model.InitDB()
	handleRequests()
}
