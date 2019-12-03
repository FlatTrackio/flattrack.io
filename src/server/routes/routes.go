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
	"net/http"
	"regexp"
        "flattrack.io/src/server/types"
        "flattrack.io/src/server/common"
)

func APIroot (w http.ResponseWriter, r *http.Request) {
        packageJSON := common.LoadPackageJSON()
        common.JSONResponse(w, 200, types.JSONMessageResponse{Message: "Welcome to FlatTrack.io", Version: packageJSON["version"].(string)})
}

func APIinterested (w http.ResponseWriter, r *http.Request) {
	var emailInForm string
	r.ParseForm()
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
        if r.Form["email"] == nil {
                common.JSONResponse(w, 400, types.JSONMessageResponse{Message: "Email address field not found."})
                return
        }
	emailInForm = r.Form["email"][0]

        if emailInForm == "" {
                common.JSONResponse(w, 400, types.JSONMessageResponse{Message: "Email address field is empty."})
                return
        }
        if len(emailInForm) > 70 {
                common.JSONResponse(w, 400, types.JSONMessageResponse{Message: "Email address is longer than 70 characters."})
                return
        }

	if matches := emailRegex.MatchString(emailInForm); !matches {
		common.JSONResponse(w, 400, types.JSONMessageResponse{Message: "Content in email form does not match the specifications for an email."})
                return
	}

        if ! common.JSONstoreExists() {
                common.InitJSONstore()
        }

	emailStore := common.ReadJSONstore("")
        if _, found := common.Find(emailStore.Emails, emailInForm); found {
		common.JSONResponse(w, 200, types.JSONMessageResponse{Message: "Added to notify list successfully."})
                return
        }

	emailStore.Emails = append(emailStore.Emails, emailInForm)

	successfullyWritten := common.WriteJSONstore(emailStore)
	if successfullyWritten {
		common.JSONResponse(w, 200, types.JSONMessageResponse{Message: "Added to notify list successfully."})
                return
	} else {
		common.JSONResponse(w, 400, types.JSONMessageResponse{Message: "An error occured."})
                return
	}
}

func UnknownPage (w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/#/unknown-page", 302)
}