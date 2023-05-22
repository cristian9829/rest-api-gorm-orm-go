package main

import (
	"net/http"

	"github.com/cristian9829/go-gorm-restapi/db"
	"github.com/cristian9829/go-gorm-restapi/models"
	"github.com/cristian9829/go-gorm-restapi/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DetUserHandler).Methods("DELETE")

	//task router
	r.HandleFunc("/task", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/task/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/task", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/task/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
