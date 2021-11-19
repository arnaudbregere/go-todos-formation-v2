package serverweb

import (
	"encoding/json"
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
		api.UpdateTodo(id[0], titre[0], description[0], dueDate[0])
		var myReturn api.TodoReturnRequest
		myReturn.CodeError = "OK"
		myReturn.Params = ""

		json.NewEncoder(w).Encode(myReturn)
	} else {
		var myReturn api.TodoReturnRequest
		myReturn.CodeError = "ERREUR"
		myReturn.Params = "Liste vide"

		json.NewEncoder(w).Encode(myReturn)
	}
}
