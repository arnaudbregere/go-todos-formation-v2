package serverweb

import (
	"fmt"
	"io"
	"net/http"
)

/*func Accueil(w http.ResponseWriter, req *http.Request) {
	token, err := req.URL.Query()["token"]
	if err != nil {
		println(err.Error())
	} else {
		fmt.Println("accueil", token)
		io.WriteString(w, "Hello from a " + token[0])
	}

}*/

func Accueil(w http.ResponseWriter, req *http.Request) {
	token := req.URL.Query()["token"]
		fmt.Println("accueil", token)
		io.WriteString(w, "toto " )

}

// Create /toto (hello en param)
// enregistrer des variables globales
// Comment créer une structure en Go
// Nom et premom mettre dans une liste avec un ID


// Check les variables => Liens Pascal
// Prochaine séance sur les variables
// https://devopssec.fr/article/cours-apprendre-langage-programmation-go#begin-article-section
//
