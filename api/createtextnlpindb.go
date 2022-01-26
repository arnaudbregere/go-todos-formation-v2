package api

func CreateTextNLPInDb(textNLP string, textFingerPrint string) (int, error) {
	sqlStatement := `INSERT INTO "text_nlp" ("text_brut", "text_fingerprint") VALUES ($1, $2) RETURNING id`
	id := 0
	err := DataBasePtr.QueryRow(sqlStatement, textNLP, textFingerPrint).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
