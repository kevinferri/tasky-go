package store

type Snippet struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func GetSnippetById(id string) (Snippet, error) {
	var snippet Snippet

	query := `select * from snippets WHERE id=$1;`
	row := db.QueryRow(query, id)
	err := row.Scan(&snippet.ID, &snippet.Title)

	return snippet, err
}

func GetAllSnippets() ([]Snippet, error) {
	var snippets []Snippet

	query := `select * from snippets;`
	rows, err := db.Query(query)

	if err != nil {
		return snippets, err
	}

	defer rows.Close()

	for rows.Next() {
		var id, title string

		err := rows.Scan(&id, &title)
		if err != nil {
			return snippets, err
		}

		snippet := Snippet{
			ID:    id,
			Title: title,
		}

		snippets = append(snippets, snippet)
	}

	return snippets, nil
}

func CreateSnippet(s Snippet) error {
	query := `insert into snippets(title) values($1);`
	err := db.QueryRow(query, s.Title).Scan(&s.ID)

	return err
}
