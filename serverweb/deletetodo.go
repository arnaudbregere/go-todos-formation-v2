package serverweb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"test/api"
)

func Delete(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query()["id"]
	if len(id) > 0  {
		fmt.Println("delete", len(api.List()))
		//s := strconv.Itoa(id)
		api.DeleteTodo(id[0])
		//io.WriteString(w, "Afficher ma Liste" )
		//fmt.Println( api.List())
		json.NewEncoder(w).Encode(api.List())
	} else {
		io.WriteString(w, "Il manque les paramètres" )
	}
}