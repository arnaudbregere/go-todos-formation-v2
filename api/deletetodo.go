package api

import (
	"fmt"
	"strconv"
)

func DeleteTodo(id string) int{
	index := -1
	fmt.Println("id", id)
	for i, todo := range Todos {

		fmt.Println("todo", todo)
		s := strconv.Itoa(todo.Id)
		if s == id {
			index = i
		}
	}

	if index >= 0 {
		Todos = append(Todos[:index], Todos[(index+1):]...)
	}
	return index
}


