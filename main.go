package main

import (
	"fmt"
	authController "go-jwt-web/controllers/authController"
	productcontroller "go-jwt-web/controllers/productController"
	"go-jwt-web/middlewares"
	"go-jwt-web/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/", authController.Home).Methods("GET")
	r.HandleFunc("/login", authController.Login).Methods("POST")
	r.HandleFunc("/register", authController.Register).Methods("POST")
	r.HandleFunc("/logout", authController.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productcontroller.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleware)

	port := "8088"
	server := http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	fmt.Println("menjalankan server....")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("server berjalan pada port")
	}

}
