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
)

// GetRoot
// root endpoints of the API
func GetRoot(w http.ResponseWriter, r *http.Request) {
	// root of api
	JSONResponse(r, w, 200, types.JSONMessageResponse{
		Metadata: types.JSONResponseMetadata{
			Response: "Welcome to FlatTrack.io",
		},
	})
}

// PostInterested
// submits an email for alerting when FlatTrack is ready later on
func PostInterested(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseCode := 400
		responseMessage := "Failed to submit email address as interested"

		var interestedSpec types.InterestedSpec
		body, err := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &interestedSpec)

		responseMessage, err = interested.AddEmailToInterested(db, interestedSpec)
		if err == nil {
			responseCode = 200
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

func UnknownEndpoint(w http.ResponseWriter, r *http.Request) {
	// unknown endpoint
	JSONResponse(r, w, 404, types.JSONMessageResponse{
		Metadata: types.JSONResponseMetadata{
			Response: "This endpoint doesn't seem to exist.",
		},
	})
}
