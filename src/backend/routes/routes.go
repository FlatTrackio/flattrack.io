/*
  routes
    routes
      endpoint definitions
*/

package routes

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"gitlab.com/flattrack/flattrack.io/src/backend/interested"
	"gitlab.com/flattrack/flattrack.io/src/backend/types"
	"gitlab.com/flattrack/flattrack.io/src/backend/health"
	"gitlab.com/flattrack/flattrack.io/src/backend/feed"
)

// GetRoot ...
// root endpoints of the API
func GetRoot(w http.ResponseWriter, r *http.Request) {
	// root of api
	JSONResponse(r, w, http.StatusOK, types.JSONMessageResponse{
		Metadata: types.JSONResponseMetadata{
			Response: "Welcome to FlatTrack.io",
		},
	})
}

// PostInterested ...
// submits an email for alerting when FlatTrack is ready later on
func PostInterested(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseCode := http.StatusInternalServerError
		responseMessage := "Failed to submit email address as interested"

		var interestedSpec types.InterestedSpec
		body, err := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &interestedSpec)

		responseMessage, err = interested.AddEmailToInterested(db, interestedSpec)
		if err == nil {
			responseCode = http.StatusOK
		} else {
			log.Println(err)
		}

		JSONResponse(r, w, responseCode, types.JSONMessageResponse{
			Metadata: types.JSONResponseMetadata{
				Response: responseMessage,
			},
		})
	}
}

// GetLatestRSSPost ...
// responds with the latest blog post,
func GetLatestRSSPost(w http.ResponseWriter, r *http.Request) {
	response := "Failed to find latest post"
	responseCode := http.StatusInternalServerError

	post, err := feed.GetLatestRSSPost()
	if err == nil {
		response = "Successfully found the latest post"
		responseCode = http.StatusOK
	}

	JSONResponse(r, w, responseCode, types.JSONMessageResponse{
		Metadata: types.JSONResponseMetadata{
			Response: response,
		},
		Spec: post,
	})
}

// UnknownEndpoint ...
// handle wildcard endpoints
func UnknownEndpoint(w http.ResponseWriter, r *http.Request) {
	// unknown endpoint
	JSONResponse(r, w, http.StatusNotFound, types.JSONMessageResponse{
		Metadata: types.JSONResponseMetadata{
			Response: "This endpoint doesn't seem to exist.",
		},
	})
}

// Healthz ...
// HTTP handler for health checks
func Healthz(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := "App unhealthy"
		code := http.StatusInternalServerError

		err := health.Healthy(db)
		if err == nil {
			response = "App healthy"
			code = http.StatusOK
		}
		JSONresp := types.JSONMessageResponse{
			Metadata: types.JSONResponseMetadata{
				Response: response,
			},
			Data: err == nil,
		}
		JSONResponse(r, w, code, JSONresp)
	}
}
