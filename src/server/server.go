package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
        "flattrack.io/src/server/common"
        "flattrack.io/src/server/routes"
)

func handleWebserver () {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api", routes.APIroot).Methods("GET")
	router.HandleFunc("/api/interested", routes.APIinterested).Methods("POST")
	router.HandleFunc("/{.*}", routes.UnknownPage).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/"))).Methods("GET")
        router.Use(common.Logging)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main () {
	fmt.Println("Listening on :8080")
        common.InitJSONstore()
	handleWebserver()
}
