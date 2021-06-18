package serverweb

import (
	"io"
	"net/http"
)

func Health(w http.ResponseWriter, req *http.Request) {
	EnableCors(&w)
	io.WriteString(w, "ON" )

}
