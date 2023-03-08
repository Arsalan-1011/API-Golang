package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func handler() {
	muxRoutes := mux.NewRouter().StrictSlash(true)
	muxRoutes.HandleFunc("/admin/{ID}", getAdmin).Methods("GET")
	muxRoutes.HandleFunc("/admin/", getAllAdmins).Methods("GET")
	muxRoutes.HandleFunc("/admin", createAdmin).Methods("POST")
	muxRoutes.HandleFunc("/admin/{ID}", updateAdmin).Methods("PUT")
	muxRoutes.HandleFunc("/admin/{ID}", deleteAdmin).Methods("DELETE")

	// Creative
	muxRoutes.HandleFunc("/creative/", AddCreative).Methods("POST")
	muxRoutes.HandleFunc("/creative/", getCrearives).Methods("GET")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"DELETE", "GET", "POST", "PUT"},
	})

	handler := c.Handler(muxRoutes)
	log.Println("server started on 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
