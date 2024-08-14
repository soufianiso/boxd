package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "github.com/gorilla/mux"
)

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func GetFilms(w http.ResponseWriter, r *http.Request)  error {
	fruit := response2{
		Page: 1,
		Fruits: []int{1,2},	
	}
	json_data, err := json.Marshal(fruit)

	if err != nil{
		return err
	}
	fmt.Fprint(w,string(json_data))
	return nil
}

