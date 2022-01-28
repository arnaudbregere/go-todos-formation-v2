package api

func ListTextNlp(start string, limit string) ([]TextNLPInDb, error) {
	var listNlp []TextNLPInDb
	sqlStatement := `SELECT * FROM "text_nlp" LIMIT $2 OFFSET $1`
	rows, err := DataBasePtr.Query(sqlStatement, start, limit)
	if err != nil {
		return listNlp, err
	}
	for rows.Next() {
		var textNlp TextNLPInDb
		if err := rows.Scan(&textNlp.Id, &textNlp.TextBrut, &textNlp.TextFingerPrint, &textNlp.DateUpdate); err != nil {
			return listNlp, err
		}
		listNlp = append(listNlp, textNlp)
	}
	return listNlp, nil
}
