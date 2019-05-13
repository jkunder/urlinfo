/*
 * URL Lookup Service
 * Handles Post Requests : POST {'url': url, 'status': 'ALLOW|BLOCK'}
 * and GET Requests : GET /urlinfo/1/{hostname_and_port}/{original_path_and_query_string}
 * The lookup service keeps the status of the url that is posted to it
 * and responds with the status on receiving a lookup 'GET' request
 */

package main

import (
	"fmt"
	"net/http"
	"strings"
)


func main() {
	// Map to store the url status
	var urlstatus = make(map[string]string)

	http.HandleFunc("/", func (w http.ResponseWriter, req *http.Request) {

		switch req.Method {
		case "GET":
			lookupUrl := strings.Split(req.URL.Path,"/")[1]
			returnStatus, ok := urlstatus[lookupUrl]
			if ok == false {
				returnStatus  = "UNKNOWN"
			}
			fmt.Fprintf(w, "URL : %s Status : %s", req.URL,
				returnStatus)

		case "POST":
			// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
			if err := req.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			fmt.Fprintf(w, "URL status updated r.PostFrom = %v\n",
				req.PostForm)
			url := req.FormValue("url")
			rating := req.FormValue("status")
			urlstatus[url]=rating
		}
	})

	//Start http server listening on port 8080
	http.ListenAndServe(":8080", nil)

}

