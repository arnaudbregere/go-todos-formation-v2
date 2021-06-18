package api

func Create(titre string, description string, dueDate int) (error, int) {
	var myTodo Todo

	myTodo.Titre = titre
	myTodo.Description = description
	myTodo.DueDate = dueDate
	myTodo.Id = ClePrimaire
	ClePrimaire++

	Todos = append(Todos, myTodo)

	return nil, ClePrimaire
}


