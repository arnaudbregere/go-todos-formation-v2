package serverweb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"test/api"
)

func Update(w http.ResponseWriter, req *http.Request) {
	EnableCors(&w)
	id := req.URL.Query()["id"]
	titre := req.URL.Query()["titre"]
	description := req.URL.Query()["description"]
	dueDate := req.URL.Query()["duedate"]

	if len(id) > 0  {
		ListTodo, _ := api.List()
		fmt.Println("update", len(ListTodo))
		//s := strconv.Itoa(id)
		api.UpdateTodo(id[0], titre[0], description[0], dueDate[0])
		//io.WriteString(w, "Mettre à jour ma Liste" )
		//fmt.Println( api.List())
		json.NewEncoder(w).Encode(ListTodo)
	} else {
		io.WriteString(w, "L'API n'a pas été mise à jour" )
	}
}
