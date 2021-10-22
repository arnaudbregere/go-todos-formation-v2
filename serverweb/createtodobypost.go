package serverweb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"test/api"
)

func CreateByPost(w http.ResponseWriter, req *http.Request) {
	EnableCors(&w)

	//decoder := json.NewDecoder(req.Body)
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		fmt.Println("ERR***********", err.Error())
		return
	}

	var data api.TodoParamsPost

	errJSON := json.Unmarshal(b, &data)
	if errJSON != nil {
		fmt.Println("errJSON***********", errJSON.Error())
		return
	}

	api.Create(data.Params.Titre, data.Params.Description, data.Params.Duedate)

	idCreate, err := api.CreateTodoInDb(data.Params.Titre, data.Params.Description, data.Params.Duedate)

	fmt.Println("*****iDCREATE*****", idCreate, err)

	ListTodo, _ := api.List()
	fmt.Println("create", len(ListTodo))

	json.NewEncoder(w).Encode(ListTodo)

}



