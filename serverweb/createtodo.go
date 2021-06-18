package serverweb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"test/api"
)

func Create(w http.ResponseWriter, req *http.Request) {
	titre := req.URL.Query()["titre"]
	description := req.URL.Query()["description"]
	duedate := req.URL.Query()["duedate"]
	EnableCors(&w)
	if titre != nil && description != nil && duedate != nil {
		strduedate, _ := strconv.Atoi(duedate[0])
		api.Create(titre[0], description[0], strduedate)

		idCreate, err := api.CreateTodoInDb(titre[0], description[0], duedate[0])

		fmt.Println("*****iDCREATE*****", idCreate, err)


		fmt.Println("create", len(api.List()))
		//io.WriteString(w, "Afficher ma Liste" )
		//fmt.Println( api.List())
		json.NewEncoder(w).Encode(api.List())
	} else {
		io.WriteString(w, "Il manque les paramètres" )
	}
}