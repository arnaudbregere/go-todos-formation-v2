package serverweb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test/api"
)

func ListTodo(w http.ResponseWriter, req *http.Request) {
	EnableCors(&w)
	ListTodo, _ := api.List()
	fmt.Println("create", len(ListTodo))
	json.NewEncoder(w).Encode(ListTodo)
	fmt.Println( api.List())
}
