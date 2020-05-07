/*
  routes
    common
      commonly used functions for endpoints
*/

package routes

import (
	"database/sql"
	"encoding/json"
	"gitlab.com/flattrack/flattrack.io/src/backend/common"
	"gitlab.com/flattrack/flattrack.io/src/backend/types"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"

	"github.com/ddo/go-vue-handler"
	"github.com/gorilla/mux"
)

func Logging(next http.Handler) http.Handler {
	// log all requests
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Printf("%v %v %v %v %v %v", r.Header["User-Agent"], r.Method, r.URL, r.Proto, r.Response, r.RemoteAddr)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// JSONResponse
// sends a JSON response
func JSONResponse(r *http.Request, w http.ResponseWriter, code int, output types.JSONMessageResponse) {
	output.Metadata.URL = r.RequestURI
	output.Metadata.Timestamp = time.Now().Unix()
	output.Metadata.Version = common.GetAppBuildVersion()
	response, _ := json.Marshal(output)
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// HandleWebserver
// manage starting of webserver
func HandleWebserver(db *sql.DB) {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println("Metrics listening on :2112")
		http.ListenAndServe(":2112", nil)
	}()
	port := common.GetAppPort()
	router := mux.NewRouter().StrictSlash(true)
	apiEndpointPrefix := "/api"

	router.HandleFunc(apiEndpointPrefix, GetRoot)
	for _, endpoint := range GetEndpoints(apiEndpointPrefix, db) {
		router.HandleFunc(endpoint.EndpointPath, endpoint.HandlerFunc).Methods(endpoint.HttpMethod, http.MethodOptions)
	}

	router.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./dist/robots.txt")
	})
	router.HandleFunc("/api/{.*}", UnknownEndpoint)
	router.PathPrefix("/").Handler(vue.Handler("./dist/")).Methods("GET")
	router.Use(Logging)
	srv := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening on", port)
	log.Fatal(srv.ListenAndServe())
}
