package api

func UpdateTodo(id string, titre string, description string, dueDate string) string{
	sqlStatement := `UPDATE todo SET titre = $1, description = $2, date = $3 WHERE id = $4`
	DataBasePtr.QueryRow(sqlStatement, titre, description, dueDate, id)
	println(id)
	return id
}


