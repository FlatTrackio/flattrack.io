// FlatTrack.io - routes.go

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

package routes

import (
	"gitlab.com/flattrack/flattrack.io/src/backend/common"
	"gitlab.com/flattrack/flattrack.io/src/backend/types"
	"net/http"
	"regexp"
)

func APIroot(w http.ResponseWriter, r *http.Request) {
	// root of api
	common.JSONResponse(r, w, 200, types.JSONMessageResponse{
		Metadata: types.JSONResponseMetadata{
			Response: "Welcome to FlatTrack.io",
		},
	})
}

func APIinterested(w http.ResponseWriter, r *http.Request) {
	// interested endpoint of api
	var emailInForm string
	r.ParseForm()
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if r.Form["email"] == nil {
		common.JSONResponse(r, w, 400, types.JSONMessageResponse{
			Metadata: types.JSONResponseMetadata{
				Response: "Email address field not found.",
			},
		})
		return
	}
	emailInForm = r.Form["email"][0]

	if emailInForm == "" {
		common.JSONResponse(r, w, 400, types.JSONMessageResponse{
			Metadata: types.JSONResponseMetadata{
				Response: "Email address field is empty.",
			},
		})
		return
	}
	if len(emailInForm) > 70 {
		common.JSONResponse(r, w, 400, types.JSONMessageResponse{
			Metadata: types.JSONResponseMetadata{
				Response: "Email address is longer than 70 characters.",
			},
		})
		return
	}

	if matches := emailRegex.MatchString(emailInForm); !matches {
		common.JSONResponse(r, w, 400, types.JSONMessageResponse{
			Metadata: types.JSONResponseMetadata{
				Response: "Content in email form does not match the specifications for an email.",
			},
		})
		return
	}

	if !common.JSONstoreExists() {
		common.InitJSONstore("")
	}

	emailStore := common.ReadJSONstore("")
	if _, found := common.Find(emailStore.Emails, emailInForm); found {
		common.JSONResponse(r, w, 200, types.JSONMessageResponse{
			Metadata: types.JSONResponseMetadata{
				Response: "Added to notify list successfully.",
			},
		})
		return
	}

	emailStore.Emails = append(emailStore.Emails, emailInForm)

	successfullyWritten := common.WriteJSONstore(emailStore)
	if successfullyWritten {
		common.JSONResponse(r, w, 200, types.JSONMessageResponse{
			Metadata: types.JSONResponseMetadata{
				Response: "Added to notify list successfully.",
			},
		})
		return
	} else {
		common.JSONResponse(r, w, 400, types.JSONMessageResponse{
			Metadata: types.JSONResponseMetadata{
				Response: "An error occured.",
			},
		})
		return
	}
}

func APIUnknownEndpoint(w http.ResponseWriter, r *http.Request) {
	// unknown endpoint
	common.JSONResponse(r, w, 404, types.JSONMessageResponse{
		Metadata: types.JSONResponseMetadata{
			Response: "This endpoint doesn't seem to exist.",
		},
	})
}