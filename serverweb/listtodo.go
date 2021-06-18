package serverweb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test/api"
)

func ListTodo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("create", len(api.List()))
	json.NewEncoder(w).Encode(api.List())
	fmt.Println( api.List())
}