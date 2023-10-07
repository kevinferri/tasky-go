package store

type Snippet struct {
	ID    string `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}

func GetSnippetById(id string) (Snippet, error) {
	s := Snippet{}
	err := db.Get(&s, "select id, title from snippets where id=$1", id)

	return s, err
}

func GetAllSnippets() ([]Snippet, error) {
	ss := []Snippet{}
	err := db.Select(&ss, "select id, title from snippets")

	return ss, err
}

func CreateSnippet(s *Snippet) (Snippet, error) {
	err := db.Get(s, "insert into snippets(title) values($1) returning id, title", s.Title)

	return *s, err
}
