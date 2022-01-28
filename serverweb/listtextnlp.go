package serverweb

import (
	"encoding/json"
	"net/http"
	"test/api"
)

func ListTextNlp(w http.ResponseWriter, req *http.Request) {
	EnableCors(&w)
	start := req.URL.Query()["start"]
	limit := req.URL.Query()["limit"]

	myReturn, errList := api.ListTextNlp(start[0], limit[0])
	if errList != nil {
		json.NewEncoder(w).Encode(errList.Error())
	} else {
		json.NewEncoder(w).Encode(myReturn)
	}
}
