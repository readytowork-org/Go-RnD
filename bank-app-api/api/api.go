package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"sumit.com/go-bank-backend/helpers"
	"sumit.com/go-bank-backend/users"
)

type Login struct {
	Username string
	Password string
}

type ErrResponse struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request) {

	//Ready body
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr()

	// Handle Login
	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr()
	login := users.Login(formattedBody.Username, formattedBody.Password)

	//Prepare response
	if login["message"] == "all is fine" {
		resp := login
		json.NewEncoder(w).Encode(resp)

	} else {
		// Handle Error
		resp := ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func StartApi() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	fmt.Println("App is working on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
