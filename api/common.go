package api

type Todo struct {
	Id          int
	Titre       string
	Description string
	DueDate     string
}

type TodoPost struct {
	Titre       string `json:"titre"`
	Description string `json:"description"`
	Duedate     string `json:"duedate"`
}

type TodoParamsPost struct {
	Params TodoPost `json:"params"`
}

type TodoReturnRequest struct {
	CodeError string `json:"code_error"`
	Params    string `json:"params"`
}

type TextNLP struct {
	TextBase           string `json:"text_base"`
	TextNLPFingerPrint string `json:"text_nlp_finger_print"`
}

type TextNLPInDb struct {
	Id              int    `json:"id"`
	TextBrut        string `json:"text_brut"`
	TextFingerPrint string `json:"text_fingerprint"`
	DateUpdate      string `json:"date_update"`
}
