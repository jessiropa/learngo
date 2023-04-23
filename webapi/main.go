package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	ID    string
	Name  string
	Grade int
}

var data = []student{
	student{"E001", "Jessi", 23},
	student{"E002", "Rian", 21},
	student{"E003", "Dino", 27},
	student{"E004", "Dian", 34},
}

func users(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "Application/json")

	if r.Method == "GET"{
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return 
	}

	http.Error(w, "", http.StatusBadRequest)
}


func user(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET"{
		var id = r.FormValue("id")
		var result [] byte
		var err error

		for _, each := range data {
			for each.ID == id {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}

		http.Error(w,"User not found", http.StatusNotFound)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
func main() {
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("starting web server at http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}