package serverweb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"test/api"
)

func SaveTextNlp(w http.ResponseWriter, req *http.Request) {
	EnableCors(&w)
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		fmt.Println("ERR***********", err.Error())
		return
	}

	var data api.TextNLP

	errJSON := json.Unmarshal(b, &data)
	if errJSON != nil {
		fmt.Println("errJSON***********", errJSON.Error())
		return
	}

	idCreate, err := api.CreateTextNLPInDb(data.TextBase, data.TextNLPFingerPrint)

	fmt.Println("*****iDCREATE*****", idCreate, err)

	var myReturn api.TodoReturnRequest
	myReturn.CodeError = "OK"
	myReturn.Params = ""

	json.NewEncoder(w).Encode(myReturn)

}
