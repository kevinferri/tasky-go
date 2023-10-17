package store

import "github.com/jmoiron/sqlx"

type Snippet struct {
	ID    string `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}

type SnippetStore struct {
	Store
}

func NewSnippetStore(db *sqlx.DB) *SnippetStore {
	return &SnippetStore{
		Store{db: db},
	}
}

func (store *SnippetStore) ById(id string) (Snippet, error) {
	s := Snippet{}
	err := store.db.Get(&s, "select id, title from snippets where id=$1", id)

	return s, err
}

func (store *SnippetStore) All() ([]Snippet, error) {
	ss := []Snippet{}
	err := store.db.Select(&ss, "select id, title from snippets")

	return ss, err
}

func (store *SnippetStore) Create(s *Snippet) (Snippet, error) {
	err := store.db.Get(s, "insert into snippets(title) values($1) returning id, title", s.Title)

	return *s, err
}
