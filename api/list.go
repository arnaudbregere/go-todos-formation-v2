package api

import "sort"

type ById []Todo

func (a ById) Len() int{
	return len(a)
}
func (a ById) Less(i, j int) bool {
	return a[i].Id < a[j].Id
}
func (a ById) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func List() ([]Todo, error) {
	var todos []Todo
	sqlStatement := `SELECT * from todo`
	rows, err := DataBasePtr.Query(sqlStatement)
	if err != nil {
		return todos, err
	}
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.Id, &todo.Titre, &todo.Description, &todo.DueDate);
		err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}
	sort.Sort(ById(todos))
	return todos, nil
}
