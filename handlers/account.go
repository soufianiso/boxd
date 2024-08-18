package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/soufianiso/letterboxd/storage"
	"github.com/soufianiso/letterboxd/types"
	"github.com/soufianiso/letterboxd/utils"
)

func RegisterAccount(w http.ResponseWriter, r *http.Request) error{
	user := types.User{}

	if err := json.NewDecoder(r.Body).Decode(&user) ; err != nil{
		return err
	}
	userjson, err := json.Marshal(&user)
	if err != nil{
		return err
	}

	utils.WriteHeaders(w)
	w.Write(userjson)

	return nil
}



func GetAccount(w http.ResponseWriter, r *http.Request) error{
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ss, err := storage.Store.FetchAccount(id)
	fmt.Println(ss)
	// utils.WriteHeaders(w)
	return nil
			
}

