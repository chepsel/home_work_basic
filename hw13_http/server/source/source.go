package source

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
)

type SentinelError string

func (err SentinelError) Error() string {
	return string(err)
}

const (
	NotFound   SentinelError = "NotFound"
	MissingKey SentinelError = "missing id or name key"
	MissingID  SentinelError = "id is missing"
)

var storageFile = "./server/storage.json"

type Animal struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Age    int8   `json:"age,omitempty"`
	Weight int8   `json:"weight,omitempty"`
	Hight  int8   `json:"hight,omitempty"`
}

type Storage struct {
	Animals map[string]Animal
}

type JSONBook struct {
	Books []Animal // позволяет хранить объекты как по одному так и множеством
}

func FileDB() *Storage {
	data, err := ReadStorage()
	if err != nil {
		log.Fatal(err)
	}
	return &data
}

func (storage *Storage) Get(id string) (Animal, error) {
	var empty Animal
	if val := storage.Animals[id]; val != empty {
		return storage.Animals[id], nil
	}
	return empty, NotFound
}

func (storage *Storage) Put(id string, animals Animal, mu *sync.Mutex) error {
	mu.Lock()
	storage.Animals[id] = animals
	mu.Unlock()
	err := SaveToStorage(storage)
	if err != nil {
		return err
	}
	return nil
}

func (storage *Storage) Delete(id string, mu *sync.Mutex) error {
	_, ok := storage.Animals[id]
	if !ok {
		return MissingID
	}
	mu.Lock()
	delete(storage.Animals, id)
	mu.Unlock()
	err := SaveToStorage(storage)
	if err != nil {
		return err
	}
	return nil
}

func (storage *Storage) SaveBeforeClose() error {
	return SaveToStorage(storage)
}

func SaveToStorage(storage *Storage) error {
	stored := storage.formatToJSON()
	byteArray, err := MarshalJSONSlice(stored)
	if err != nil {
		log.Printf("failed to encode structure: %v", err)
		return err
	}
	file, _ := os.Create(storageFile)
	defer file.Close()
	_, err = file.Write(byteArray)
	if err != nil {
		log.Printf("failed to write: %v", err)
		return err
	}
	return nil
}

func (storage *Storage) formatToJSON() []Animal {
	arrayLengt := len(storage.Animals)
	animals := make([]Animal, arrayLengt)
	var i int
	for _, v := range storage.Animals {
		animals[i] = v
		i++
	}
	return animals
}

func (storage *Storage) formatToMap(input []Animal) {
	tmpMap := make(map[string]Animal)
	for _, v := range input {
		tmpMap[v.ID] = v
	}
	storage.Animals = tmpMap
}

func ReadStorage() (Storage, error) {
	file, err := os.Open(storageFile)
	data := Storage{}
	if err != nil {
		log.Println(err)
		return data, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
		return data, err
	}
	arrayData, err := UnmarshalJSONSlice(byteValue)
	if err != nil {
		log.Println(err)
		return data, err
	}
	data.formatToMap(arrayData)
	return data, nil
}

func UnmarshalJSONSlice(data []byte) ([]Animal, error) {
	var ctr []Animal
	if err := json.Unmarshal(data, &ctr); err != nil {
		return nil, err
	}
	return ctr, nil
}

func MarshalJSONSlice(ctr []Animal) ([]byte, error) {
	result, err := json.Marshal(ctr)
	return result, err
}
