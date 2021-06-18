package api

func CreateTodoInDb(title string, description string, dueDate string) (int, error) {
	sqlStatement := `INSERT INTO "todo" ("titre", "description", "date") VALUES ($1, $2, $3) RETURNING id`
	id := 0
	err := DataBasePtr.QueryRow(sqlStatement, title, description, dueDate).Scan(&id)
	println("id :", id)
	// _, err := myDb.Exec(sqlStatement)
	if err != nil {
		return -1, err
	}
	return id, nil
}
