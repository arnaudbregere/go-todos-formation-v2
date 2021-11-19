package serverweb

import (
	"encoding/json"
	"net/http"
	"test/api"
)

func ListTodo(w http.ResponseWriter, req *http.Request) {
	EnableCors(&w)
	ListTodo, _ := api.List()
	json.NewEncoder(w).Encode(ListTodo)
}
