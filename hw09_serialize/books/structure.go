package books

import "encoding/json"

func (b *Book) UnmarshalJSON(input []byte) error {
	type TempStructure Book
	var tmpJSON TempStructure
	err := json.Unmarshal(input, &tmpJSON)
	if err != nil {
		return err
	}
	*b = Book(tmpJSON)
	return nil
}

func (b *Book) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(*b)
	return j, err
}

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author,omitempty"`
	Year   uint32  `json:"year"`
	Size   uint32  `json:"size,omitempty"`
	Rate   float32 `json:"rate,omitempty"`
}

func NewBook(i string, t string, a string, y uint32, s uint32, r float32) *Book {
	return &Book{
		ID:     i,
		Title:  t,
		Author: a,
		Year:   y,
		Size:   s,
		Rate:   r,
	}
}
