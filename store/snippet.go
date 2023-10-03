package store

import (
	"log"
)

type Snippet struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (s *Snippet) GetAll() ([]Snippet, error) {
	var snippets []Snippet

	query := `select * from snippets;`

	rows, err := db.Query(query)
	if err != nil {
		return snippets, err
	}

	defer rows.Close()

	for rows.Next() {
		var id string
		var title string

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

func (s *Snippet) Create() error {
	query := `insert into snippets(title) values($1);`

	_, err := db.Exec(query, s.Title)

	if err != nil {
		log.Fatal("Could not create snippet:", err.Error())
		return err
	}

	return nil
}
