/*
프레임워크 사용으로 인한 영향을 줄이기 위해
순수 Go Http모듈을 활용합니다.
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var host int = 8080

func main() {
	var err error

	/*
		< Request >
		/request?from=1&to=2

		[ Query Parameter ]
		- from: number
		- to: number
	*/
	http.HandleFunc("/request", RestRequestHandler)

	log.Printf("Listening server on port %d\n", host)
	if err = http.ListenAndServe(fmt.Sprintf(":%d", host), nil); err != nil {
		log.Fatalf("Fail to start server: %v\n", err)
	}

}

type RestResponse struct {
	ResponseNumber int `json:"response_number"`
}

func RestRequestHandler(w http.ResponseWriter, r *http.Request) {
	// Error
	var err error
	// Query Params
	var from, to, sumAll int
	// Response struct
	var response *RestResponse = new(RestResponse)
	// JSON byteArray
	var plainJSON []byte

	switch r.Method {
	case http.MethodGet:
		// Parse Query & Validate
		if from, err = strconv.Atoi(r.URL.Query().Get("from")); err != nil {
			http.Error(w, "Validation Fail: Query 'from' should be type 'int'", http.StatusBadRequest)
			return
		}
		if to, err = strconv.Atoi(r.URL.Query().Get("to")); err != nil {
			http.Error(w, "Validation Fail: Query 'to' should be type 'int'", http.StatusBadRequest)
			return
		}
		if from > to {
			http.Error(w, "Validation Fail: Query 'from' should be lower than Query 'to'", http.StatusBadRequest)
			return
		}

		// Add all of the doubled numbers
		for i := from; i < to; i++ {
			sumAll += i * i
		}

		response.ResponseNumber = sumAll
		// Marshal Struct to Plain JSON
		if plainJSON, err = json.Marshal(response); err != nil {
			http.Error(w, "Fail to marshal struct to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(plainJSON)
	// Ignore request with invalid HTTP metho
	default:
		http.Error(w, "Invalid HTTP method request", http.StatusMethodNotAllowed)
	}
}
