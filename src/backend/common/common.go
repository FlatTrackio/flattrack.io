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
	"fmt"
	"os"
	"regexp"
)

var (
	APP_BUILD_VERSION      = "0.0.0"
	APP_BUILD_HASH         = "???"
	APP_BUILD_DATE         = "???"
	APP_BUILD_MODE         = "development"
	APP_DB_MIGRATIONS_PATH = "/app/migrations"
)

// GetEnvOrDefault
// given the value of an environment variable, return it's data or if not available a default value
func GetEnvOrDefault(envName string, defaultValue string) (output string) {
	output = os.Getenv(envName)
	if output == "" {
		output = defaultValue
	}
	return output
}

// GetDBdatabase
// return the database's database to use
func GetDBdatabase() (output string) {
	return GetEnvOrDefault("APP_DB_DATABASE", "flattrackio")
}

// GetDBusername
// return the database user to use
func GetDBusername() (output string) {
	return GetEnvOrDefault("APP_DB_USERNAME", "postgres")
}

// GetDBhost
// return the database host to use
func GetDBhost() (output string) {
	return GetEnvOrDefault("APP_DB_HOST", "localhost")
}

// GetDBpassword
// return the database password to use
func GetDBpassword() (output string) {
	return GetEnvOrDefault("APP_DB_PASSWORD", "postgres")
}

// GetMigrationsPath
// return the path of the database migrations to use
func GetMigrationsPath() (output string) {
	envSet := GetEnvOrDefault("APP_DB_MIGRATIONS_PATH", "")
	if envSet != "" {
		return envSet
	}
	if APP_BUILD_MODE == "production" {
		return "/app/migrations"
	}
	pwd, _ := os.Getwd()
	return fmt.Sprintf("%v/migrations", pwd)
}

// GetAppPort
// return the port which the app should serve HTTP on
func GetAppPort() (output string) {
	return GetEnvOrDefault("APP_PORT", ":8080")
}

// GetAppBuildVersion
// return the version of the current FlatTrack instance
func GetAppBuildVersion() string {
	return APP_BUILD_VERSION
}

// GetAppBuildHash
// return the commit which the current FlatTrack binary was built from
func GetAppBuildHash() string {
	return APP_BUILD_HASH
}

// GetAppBuildDate
// return the build date of FlatTrack
func GetAppBuildDate() string {
	return APP_BUILD_DATE
}

// GetAppBuildMode
// return the mode that the app is built in
func GetAppBuildMode() string {
	return APP_BUILD_MODE
}

// FindStringInArray
// determine if string is in string array
func FindStringInArray(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// RegexMatchEmail
// regex check for valid email address string
// must also be <= 70
func RegexMatchEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email) && len(email) <= 70 && email != ""
}

// SetFirstOrSecond
// given first, return it, else return second
func SetFirstOrSecond(first string, second string) string {
	if first != "" {
		return first
	}
	return second
}
