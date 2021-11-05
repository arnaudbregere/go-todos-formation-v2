package api

func DeleteTodo(id string) error{
	sqlStatement := `DELETE FROM todo WHERE id = $1`
	DataBasePtr.QueryRow(sqlStatement, id)
	println("DELETE id", id)
	return nil
}
