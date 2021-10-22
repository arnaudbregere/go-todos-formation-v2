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
		ListTodo := api.List
		json.NewEncoder(w).Encode(ListTodo)
	} else {
		io.WriteString(w, "Il manque les paramètres" )
	}
}
