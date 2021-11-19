package api

func DeleteTodo(id string) error{
	sqlStatement := `DELETE FROM todo WHERE id = $1`
	DataBasePtr.QueryRow(sqlStatement, id)
	return nil
}
