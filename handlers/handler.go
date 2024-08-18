package handlers

import (
	"errors"
	"net/http"

)



func HandleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST"{
		return RegisterAccount(w,r)
	}

	if r.Method == "GET"{
		return GetAccount(w,r)
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	return errors.New("Bad method")
}



