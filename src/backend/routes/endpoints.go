/*
  routes
    endpoints
      store all the endpoints under API
*/

package routes

import (
	"net/http"

	"database/sql"
	"gitlab.com/flattrack/flattrack.io/src/backend/types"
)

// GetEndpoints
// group all API endpoints
func GetEndpoints(endpointPrefix string, db *sql.DB) types.Endpoints {
	return types.Endpoints{
		{
			EndpointPath: endpointPrefix + "/interested",
			HandlerFunc:  PostInterested(db),
			HttpMethod:   http.MethodPost,
		},
	}
}
