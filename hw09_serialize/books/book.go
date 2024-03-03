package books

import (
	"encoding/json"
)

type JSONBook struct {
	Books []Book // позволяет хранить объекты как по одному так и множеством
}

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author,omitempty"`
	Year   uint32  `json:"year"`
	Size   uint32  `json:"size,omitempty"`
	Rate   float32 `json:"rate,omitempty"`
}

func (ctr *Book) UnmarshalJSON(data []byte) error {
	type dropDefaultInf Book
	err := json.Unmarshal(data, (*dropDefaultInf)(ctr))
	if err != nil {
		return err
	}
	return nil
}

func UnmarshalJSONSlice(data []byte) ([]*Book, error) {
	var ctr []*Book
	if err := json.Unmarshal(data, &ctr); err != nil {
		return nil, err
	}
	return ctr, nil
}

func MarshalJSONSlice(ctr []*Book) ([]byte, error) {
	result, err := json.Marshal(ctr)
	return result, err
}

func (ctr *Book) MarshalJSON() ([]byte, error) {
	type dropDefaultInf Book
	result, err := json.Marshal((*dropDefaultInf)(ctr))
	return result, err
}
