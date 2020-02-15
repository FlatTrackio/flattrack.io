// FlatTrack.io - common.go

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

package common

import (
	"encoding/json"
	"gitlab.com/flattrack/flattrack.io/src/backend/types"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"path/filepath"
)

var (
	currentWorkingDirectory, _ = os.Getwd()
	deploymentInterestedJSON = fmt.Sprintf("%v/%v/%v", currentWorkingDirectory, "deployment", "interested.json")
	packageJSON = LoadPackageJSON()
	appVersion = packageJSON.Version
)

func GetAppPort() (output string) {
	// determine the port for the app to run on
	output = os.Getenv("APP_PORT")
	if output == "" {
		output = ":8080"
	}
	return output
}

func Find(slice []string, val string) (int, bool) {
	// determine if string is in string array
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func LoadPackageJSON() (output types.PackageJSON) {
	// return contents of package.json
	packageJSON, err := os.Open("package.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer packageJSON.Close()
	packageJSONFileContents, _ := ioutil.ReadAll(packageJSON)
	json.Unmarshal([]byte(string(packageJSONFileContents)), &output)
	return output
}

func JSONstoreExists() (exists bool) {
	// determine if deployment/interested.json exists / the app has been initialised
	if _, err := os.Stat(deploymentInterestedJSON); err == nil {
		exists = true
	} else {
		exists = false
	}
	return exists
}

func InitJSONstore(altConfig string) (completed bool) {
	// create deployment/interested.json
	emptyJSONstore := types.EmailStore{}
	filePath := altConfig
	if altConfig == "" {
		filePath = deploymentInterestedJSON
	}
	if !JSONstoreExists() {
		log.Println("Creating json store")
		err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
		if err == nil {
			completed = WriteJSONstore(emptyJSONstore)
		} else {
			fmt.Println("Initialisation error:", err)
			os.Exit(1)
		}
	}
	return completed
}

func DeinitJSONstore(altConfig string) (completed bool) {
	// remove deployment/interested.json (used for tests)
	filePath := altConfig
	if altConfig == "" {
		filePath = deploymentInterestedJSON
	}
	err := os.Remove(filePath)
	completed = true
	if err != nil {
		completed = false
	}
	return completed
}

func ReadJSONstore(altConfig string) (output types.EmailStore) {
	// read contents of deployment/interested.json
	var filePath string
	filePath = altConfig
	if altConfig == "" {
		filePath = deploymentInterestedJSON
	}
	emailStore, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer emailStore.Close()
	emailStoreFileContents, _ := ioutil.ReadAll(emailStore)
	json.Unmarshal([]byte(string(emailStoreFileContents)), &output)
	return output
}

func WriteJSONstore(content types.EmailStore) (completed bool) {
	// write data to deployment/interested.json
	contentAsJSONString, _ := json.Marshal(content)
	err := ioutil.WriteFile(deploymentInterestedJSON, []byte(contentAsJSONString), 0644)
	if err == nil {
		completed = true
	}
	return completed
}

func JSONResponse(r *http.Request, w http.ResponseWriter, code int, output types.JSONMessageResponse) {
	// send a JSON response
	output.Metadata.URL = r.RequestURI
	output.Metadata.URL = r.RequestURI
	output.Metadata.Timestamp = time.Now().Unix()
	output.Metadata.Version = appVersion
	response, _ := json.Marshal(output)
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func Logging(next http.Handler) http.Handler {
	// log all requests
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Printf("%v %v %v %v %v %v", r.Header["User-Agent"], r.Method, r.URL, r.Proto, r.Response, r.RemoteAddr)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
