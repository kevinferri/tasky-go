package types

import "encoding/json"

type Note struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (n *Note) ToJson() []byte {
	json, _ := json.Marshal(n)
	return json
}
