package api

import "strconv"

func DeleteTodo(id string) error{
	ID, _ := strconv.Atoi(id)
	sqlStatement := `SELECT from todo where id=`
	_, err := DataBasePtr.Query(sqlStatement)
	println(err.Error())
	return err
}
