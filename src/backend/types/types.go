// FlatTrack.io - types.go

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

package types

import (
	"net/http"
)

// JSONResponseMetadata ...
// metadata field in response
type JSONResponseMetadata struct {
	URL       string `json:"selfLink"`
	Version   string `json:"version"`
	Timestamp int64  `json:"timestamp"`
	Response  string `json:"response"`
}

// JSONMessageResponse ...
// generic structure for responses
type JSONMessageResponse struct {
	Metadata JSONResponseMetadata `json:"metadata"`
	Spec     interface{}          `json:"spec"`
	Data     interface{}          `json:"data,omitempty"`
}

// InterestedSpec ...
// fields for email notification form
type InterestedSpec struct {
	Email string `json:"email"`
}

// Endpoints ...
// all API endpoints stored in an array
type Endpoints []struct {
	EndpointPath string
	HandlerFunc  http.HandlerFunc
	HTTPMethod   string
}
