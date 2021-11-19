package serverweb

import (
	"encoding/json"
	"io"
	"net/http"
	"test/api"
)

func Delete(w http.ResponseWriter, req *http.Request) {
	EnableCors(&w)
	id := req.URL.Query()["id"]
	if len(id) > 0  {
		api.DeleteTodo(id[0])
		var myReturn api.TodoReturnRequest
		myReturn.CodeError = "OK"
		myReturn.Params = "Delete"

		json.NewEncoder(w).Encode(myReturn)
	} else {
		io.WriteString(w, "Il manque les paramètres" )
	}
}
