package serverweb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"test/api"
)

func Create(w http.ResponseWriter, req *http.Request) {
	titre := req.URL.Query()["titre"]
	description := req.URL.Query()["description"]
	duedate := req.URL.Query()["duedate"]
	EnableCors(&w)
	if titre != nil && description != nil && duedate != nil {
		api.Create(titre[0], description[0], duedate[0])

		idCreate, err := api.CreateTodoInDb(titre[0], description[0], duedate[0])

		fmt.Println("*****iDCREATE*****", idCreate, err)

		ListTodo, _ := api.List()
		fmt.Println("create", len(ListTodo))
		//io.WriteString(w, "Afficher ma Liste" )
		//fmt.Println( api.List())
		json.NewEncoder(w).Encode(ListTodo)
	} else {
		io.WriteString(w, "Il manque les paramètres" )
	}
}
