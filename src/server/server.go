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
	"net/http"
	"log"
	"github.com/gorilla/mux"
        "flattrack.io/src/server/common"
        "flattrack.io/src/server/routes"
)

func handleWebserver () {
        // manage starting of webserver
        port := common.GetAppPort()
	router := mux.NewRouter().StrictSlash(true)
        router.Host("flattrack.io")
	router.HandleFunc("/api", routes.APIroot).Methods("GET")
	router.HandleFunc("/api/interested", routes.APIinterested).Methods("POST")
	router.HandleFunc("/{.*}", routes.UnknownPage).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/"))).Methods("GET")
        router.Use(common.Logging)
	log.Println("Listening on", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func main () {
        // initialise the app
        common.InitJSONstore()
        handleWebserver()
}
