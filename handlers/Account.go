package handlers

import (
	"github.com/soufianiso/letterboxd/types"
	"github.com/soufianiso/letterboxd/utils"
	"encoding/json"
	"fmt"
	"net/http"

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
	utils.WriteHeaders(w)
	fmt.Fprint(w,"ww")
	return nil
			
}

