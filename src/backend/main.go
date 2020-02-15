// FlatTrack.io - server.go

//
// Copyright (C) 2018 Caleb Woodbine <@>
//
// This file is part of FlatTrack.io.
//
// FlatTrack is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// FlatTrack.io is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with FlatTrack.io.  If not, see <https://www.gnu.org/licenses/>.
//

package main

import (
	"log"
	"net/http"
	"time"

	"gitlab.com/flattrack/flattrack.io/src/backend/common"
	"gitlab.com/flattrack/flattrack.io/src/backend/routes"
	"github.com/ddo/go-vue-handler"
	"github.com/gorilla/mux"
)

func handleWebserver() {
	// manage starting of webserver
	port := common.GetAppPort()
	router := mux.NewRouter().StrictSlash(true)
	router.Host("flattrack.io")
	router.HandleFunc("/api", routes.APIroot).Methods("GET")
	router.HandleFunc("/api/interested", routes.APIinterested).Methods("POST")
	router.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./dist/robots.txt")
	})
	router.HandleFunc("/api/{.*}", routes.APIUnknownEndpoint)
	router.PathPrefix("/").Handler(vue.Handler("./dist/")).Methods("GET")
	router.Use(common.Logging)
	srv := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening on", port)
	log.Fatal(srv.ListenAndServe())
}

func main() {
	// initialise the app
	common.InitJSONstore("")
	handleWebserver()
}
