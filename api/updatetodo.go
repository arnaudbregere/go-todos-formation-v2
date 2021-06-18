package api

import (
	"fmt"
	"strconv"
)

func UpdateTodo(id string, titre string, description string, dueDate string) int{
	index := -1
	fmt.Println("id", id)
	for i, todo := range Todos {
		fmt.Println("todo", todo)
		s := strconv.Itoa(todo.Id)
		d, _ := strconv.Atoi(dueDate)
		if s == id {
			index = i
			todo.Description = description
			todo.Titre = titre
			todo.DueDate = d
			Todos = append(Todos[:index], Todos[(index+1):]...)
			Todos = append(Todos, todo)
		}
	}
	return index
}


