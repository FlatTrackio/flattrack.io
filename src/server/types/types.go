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

type JSONResponseMetadata struct {
	URL string `json:"url"`
}

// format of a stand JSON response
type JSONMessageResponse struct {
	Message  string               `json:"message"`
	Version  string               `json:"version"`
	Metadata JSONResponseMetadata `json:"metadata"`
}

// how deployment/config.json should be formatted
type EmailStore struct {
	Emails []string `json:"emails"`
}
