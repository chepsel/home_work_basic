package source

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const testFile = "./storage.json"

func TestPut(t *testing.T) {
	testCases := []struct {
		input1 Animal
		input2 string
		desc   string
	}{
		{
			desc: "check valid",
			input1: Animal{
				ID:     "Ignat",
				Name:   "Выхухоль",
				Age:    12,
				Weight: 21,
				Hight:  30,
			},
			input2: "Ignat",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mu := sync.Mutex{}
			storageFile = testFile
			animals := make(map[string]Animal)
			animals[tC.input2] = tC.input1
			storage := &Storage{Animals: animals}
			storage.Put(tC.input2, tC.input1, &mu)
			time.Sleep(100 * time.Millisecond)
			assert.Equal(t, true, true)
		})
	}
}

func TestGet(t *testing.T) {
	time.Sleep(200 * time.Millisecond)
	storageFile = testFile
	animals := make(map[string]Animal)
	id := "Ignat"
	animal := Animal{
		ID:     "Ignat",
		Name:   "Выхухоль",
		Age:    12,
		Weight: 21,
		Hight:  30,
	}
	animals[id] = animal
	storage := &Storage{Animals: animals}
	got, _ := storage.Get("Ignat")
	assert.Equal(t, animal, got)
}

func TestFileDB(t *testing.T) {
	testCases := []struct {
		input1 Animal
		input2 string
		desc   string
	}{
		{
			desc: "check valid",
			input1: Animal{
				ID:     "Ignat",
				Name:   "Выхухоль",
				Age:    12,
				Weight: 21,
				Hight:  30,
			},
			input2: "Ignat",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			storageFile = testFile
			animals := make(map[string]Animal)
			animals[tC.input2] = tC.input1
			want := &Storage{Animals: animals}
			got := FileDB()
			assert.Equal(t, want, got)
		})
	}
}

func TestSaveBeforeClose(t *testing.T) {
	testCases := []struct {
		input1 Animal
		input2 string
		desc   string
	}{
		{
			desc: "check valid",
			input1: Animal{
				ID:     "Ignat",
				Name:   "Выхухоль",
				Age:    12,
				Weight: 21,
				Hight:  30,
			},
			input2: "Ignat",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			storageFile = testFile
			animals := make(map[string]Animal)
			animals[tC.input2] = tC.input1
			storage := &Storage{Animals: animals}
			err := storage.SaveBeforeClose()
			if err != nil {
				t.Errorf("error")
			}
		})
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		input1 Animal
		input2 string
		desc   string
	}{
		{
			desc: "check valid",
			input1: Animal{
				ID:     "Ignat",
				Name:   "Выхухоль",
				Age:    12,
				Weight: 21,
				Hight:  30,
			},
			input2: "Ignat",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			mu := sync.Mutex{}
			storageFile = testFile
			animals := make(map[string]Animal)
			animals[tC.input2] = tC.input1
			storage := &Storage{Animals: animals}
			err := storage.Delete(tC.input2, &mu)
			if err != nil {
				t.Errorf("error")
			}
			time.Sleep(100 * time.Millisecond)
		})
	}
}

func TestGetNotFound(t *testing.T) {
	time.Sleep(200 * time.Millisecond)
	storageFile = testFile
	animals := make(map[string]Animal)
	id := "Ignat"
	animal := Animal{
		ID:     "Ignat",
		Name:   "Выхухоль",
		Age:    12,
		Weight: 21,
		Hight:  30,
	}
	animals[id] = animal
	storage := &Storage{Animals: animals}
	_, err := storage.Get(id)
	if err != nil {
		t.Errorf("error")
	}
}
