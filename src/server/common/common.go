package common

import (
        "fmt"
        "os"
        "io/ioutil"
        "net/http"
        "log"
        "encoding/json"
        "flattrack.io/src/server/types"
)

func Find (slice []string, val string) (int, bool) {
        for i, item := range slice {
                if item == val {
                        return i, true
                }
        }
        return -1, false
}

func LoadPackageJSON () (output map[string]interface{}) {
	packageJSON, err := os.Open("package.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer packageJSON.Close()
	packageJSONFileContents, _ := ioutil.ReadAll(packageJSON)
	json.Unmarshal([]byte(string(packageJSONFileContents)), &output)
	return output
}

func JSONstoreExists () (exists bool) {
	currentWorkingDirectory, _ := os.Getwd()
	if _, err := os.Stat(fmt.Sprintf("%v/%v/%v", currentWorkingDirectory, "deployment", "config.json")); err == nil {
		exists = true
	} else {
		exists = false
	}
	return exists
}

func InitJSONstore () (completed bool) {
	currentWorkingDirectory, _ := os.Getwd()
	emptyJSONstore := types.EmailStore{}
	if ! JSONstoreExists() {
		err := os.MkdirAll(fmt.Sprintf("%v/%v", currentWorkingDirectory, "deployment"), os.ModePerm)
		if err == nil {
			completed = WriteJSONstore(emptyJSONstore)
		} else {
			os.Exit(1)
		}
	}
	return completed
}

func DeinitJSONstore () (completed bool) {
	currentWorkingDirectory, _ := os.Getwd()
	err := os.Remove(fmt.Sprintf("%v/%v/%v", currentWorkingDirectory, "deployment", "config.json"))
	completed = true
	if err != nil {
		completed = false
	}
	return completed
}

func ReadJSONstore () (output types.EmailStore) {
	currentWorkingDirectory, _ := os.Getwd()
	emailStore, err := os.Open(fmt.Sprintf("%v/%v/%v", currentWorkingDirectory, "deployment", "config.json"))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer emailStore.Close()
	emailStoreFileContents, _ := ioutil.ReadAll(emailStore)
	json.Unmarshal([]byte(string(emailStoreFileContents)), &output)
	return output
}

func WriteJSONstore (content types.EmailStore) (completed bool) {
	currentWorkingDirectory, _ := os.Getwd()
	contentAsJSONString, _ := json.Marshal(content)
	err := ioutil.WriteFile(fmt.Sprintf("%v/%v/%v", currentWorkingDirectory, "deployment", "config.json"), []byte(contentAsJSONString), 0644)
	if err == nil {
		completed = true
	}
	return completed
}

func JSONResponse (w http.ResponseWriter, code int, output interface{}) {
	response, _ := json.Marshal(output)
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func Logging(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                // Do stuff here
                log.Printf("%v %v %v %v %v", r.Method, r.URL, r.Proto, r.Response, r.RemoteAddr)
                // Call the next handler, which can be another middleware in the chain, or the final handler.
                next.ServeHTTP(w, r)
        })
}
