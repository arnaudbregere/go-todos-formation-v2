package serverweb

import (
	"fmt"
	"io"
	"net/http"
	"test/api"
)

func GetTodo(w http.ResponseWriter, req *http.Request) {
	Id := req.URL.Query()["id"]
	if Id != nil {
		for i, s := range api.Todos {
			fmt.Println(i, s)
		}
		io.WriteString(w, "OK pour l'ID" )
	} else {
		io.WriteString(w, "Il manque le paramètre ID" )
	}
}