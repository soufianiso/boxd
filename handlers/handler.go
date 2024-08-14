package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/soufianiso/letterboxd/types"
	"github.com/soufianiso/letterboxd/utils"
)


func GetFilms(w http.ResponseWriter, r *http.Request) error{
	if r.Method == "POST"{
		user := types.User{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil{
			return err
		}
		userjson, err := json.Marshal(&user)
		if err != nil{
			return err
		}

		err = utils.WriteHeaders(w)
		if err != nil{
			return err
		}
	

		w.Write(userjson)



	}
	return nil
}

