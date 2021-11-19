package api

type Todo struct {
	Id int
	Titre string
	Description string
	DueDate string
}

type TodoPost struct {
	Titre string `json:"titre"`
	Description string `json:"description"`
	Duedate string `json:"duedate"`
}

type TodoParamsPost struct {
	Params TodoPost `json:"params"`
}

type TodoReturnRequest struct {
	CodeError string `json:"code_error"`
	Params string `json:"params"`
}


