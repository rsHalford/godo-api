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
package model

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rsHalford/goapi/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Todo struct {
	gorm.Model
	Title    string `json:"title"`
	Body     string `json:"body"`
	Priority bool   `json:"priority"`
	Status   bool   `json:"status"`
}

func InitDB() {
	dsn := "user=" + config.GetString("DB_USERNAME") + " password=" + config.GetString("DB_PASSWORD") + " dbname=" + config.GetString("DB_NAME") + " port=5432"
	dbLoaded, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to DB")
	}
	dbLoaded.AutoMigrate(&Todo{})
	db = dbLoaded
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	db.Create(&todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todos []Todo
	db.Find(&todos)
	json.NewEncoder(w).Encode(todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	var todo Todo
	db.First(&todo, vars["id"])
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	db.Save(&todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todo Todo
	w.Header().Set("Content-Type", "application/json")
	db.Delete(&todo, vars["id"])
	json.NewEncoder(w).Encode("Delete successfully.")
}
