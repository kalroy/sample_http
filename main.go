package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Request struct {
	Val string `json:"value"`
}

var keyVal map[string]string = map[string]string{"key": "Kallol"}

func respondWithError(response http.ResponseWriter, statusCode int, msg string) {
	respondWithJSON(response, statusCode, map[string]string{"error": msg})
}

func respondWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}

func updateKey(response http.ResponseWriter, request *http.Request) {
	var req Request
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	}
	keyVal["key"] = req.Val
	respondWithJSON(response, http.StatusOK, keyVal)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/key", updateKey).Methods("PUT")

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println(err)
	}
}
