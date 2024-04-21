package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type SentinelError string

func (err SentinelError) Error() string {
	return string(err)
}

const (
	NotFound       SentinelError = "Not Found"
	WrongParams    SentinelError = "Wrong request params"
	WrongStructure SentinelError = "Wrong structure"
	RequestError   SentinelError = "Request failed"
	MissingFlag    SentinelError = "URL flag is missing"
)

const (
	DefaultURL string = "default"
)

var url string

type Animal struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Age    int8   `json:"age,omitempty"`
	Weight int8   `json:"weight,omitempty"`
	Hight  int8   `json:"hight,omitempty"`
}

func (ctr *Animal) MarshalJSON() ([]byte, error) {
	type dropDefaultInf Animal
	result, err := json.Marshal((*dropDefaultInf)(ctr))
	return result, err
}

func NewAnimal(id string, name string, age int8, weight int8, hight int8) *Animal {
	return &Animal{
		ID:     id,
		Name:   name,
		Age:    age,
		Weight: weight,
		Hight:  hight,
	}
}

func (ctr *Animal) Post() error {
	body, err := ctr.MarshalJSON()
	if err != nil {
		return WrongStructure
	}
	log.Println(" POST - Resuest string:", string(body))
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return WrongParams
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Println(" POST - response status:", resp.Status)
	if resp.StatusCode != http.StatusOK {
		return RequestError
	}
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println(" POST - response Body:", string(result))
	return nil
}

func (ctr *Animal) Get() error {
	var address strings.Builder
	address.WriteString(url)
	address.WriteString("?id=")
	address.WriteString(ctr.ID)
	log.Println(" GET - Request string:", address.String())
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, address.String(), nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Println(" GET - response status:", resp.Status)
	if resp.StatusCode != http.StatusOK {
		return RequestError
	}
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println(" GET - response Body:", string(result))
	return nil
}

func Client(input string) error {
	if input == DefaultURL {
		log.Printf("- wrong start params \n- url:\"%s\";\n", input)
		return MissingFlag
	}
	url = input
	data := NewAnimal("Vitaly", "Kapibara", 12, 33, 44)
	err := data.Post()
	if err != nil {
		return err
	}
	data.Get()
	if err != nil {
		return err
	}
	log.Println("- End")
	return nil
}
