package books

import (
	"encoding/json"
	"fmt"
)

type JSONBook struct {
	Books []Book
}

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author,omitempty"`
	Year   uint32  `json:"year"`
	Size   uint32  `json:"size,omitempty"`
	Rate   float32 `json:"rate,omitempty"`
}

type SentinelError string

func (err SentinelError) Error() string {
	return string(err)
}

const (
	EmptySlice SentinelError = "Empty slice"
)

func (ctr *JSONBook) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return fmt.Errorf("no bytes to unmarshal")
	}
	switch b[0] {
	case '{':
		return ctr.unmarshalSingle(b)
	case '[':
		return ctr.unmarshalMany(b)
	}
	err := ctr.unmarshalMany(b)
	if err != nil {
		return ctr.unmarshalSingle(b)
	}
	return nil
}

func (ctr *JSONBook) unmarshalSingle(b []byte) error {
	var t Book
	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}
	ctr.Books = []Book{t}
	return nil
}

func (ctr *JSONBook) unmarshalMany(b []byte) error {
	var books []Book
	err := json.Unmarshal(b, &books)
	if err != nil {
		return err
	}
	ctr.Books = books
	return nil
}

func (ctr *JSONBook) MarshalJSON() ([]byte, error) {
	switch len(ctr.Books) {
	case 0:
		return nil, EmptySlice
	case 1:
		j, err := json.Marshal(ctr.Books[0])
		return j, err
	default:
		var emptyBook Book
		var resultJSON []byte
		for i, element := range ctr.Books {
			if element != emptyBook {
				result, err := json.Marshal(ctr.Books[i])
				if err == nil {
					if len(resultJSON) > 0 {
						resultJSON = append(resultJSON, []byte(`,`)...)
					}
					resultJSON = append(resultJSON, result...)
				}
			}
		}
		resultJSON = append([]byte(`[`), resultJSON...)
		resultJSON = append(resultJSON, []byte(`]`)...)
		return resultJSON, nil
	}
}
