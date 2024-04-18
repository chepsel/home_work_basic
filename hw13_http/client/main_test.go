package main

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var _ = func() bool {
	testing.Init()
	os.Args = append(os.Args, "-u=http://localhost:8082/v1/restapi/animal")
	go fakeServer()
	return true
}()

func fakeServer() {
	server := gin.Default()
	v1 := server.Group("/v1")
	restAPI := v1.Group("/restapi")
	{
		restAPI.GET("/animal", func(c *gin.Context) {
			c.JSONP(200, nil)
		})
		restAPI.POST("/animal", func(c *gin.Context) {
			c.JSONP(200, nil)
		})
	}
	server.Run(":8082")

}

func TestMarshalJSON(t *testing.T) {
	testCases := []struct {
		want1     []byte
		want2     []byte
		desc      string
		input1    *Animal
		testError bool
	}{
		{
			desc:  "check valid",
			want1: []byte(`{"id":"Vitaly","name":"Kapibara","age":12,"weight":33,"hight":44}`),
			input1: &Animal{
				ID:     "Vitaly",
				Name:   "Kapibara",
				Age:    12,
				Weight: 33,
				Hight:  44,
			},
			testError: false,
		},
		{
			desc:  "check one",
			want1: []byte(`{"id":"Bober","name":"Kurva","age":12,"weight":33,"hight":44}`),
			input1: &Animal{
				ID:     "Bober",
				Name:   "Kurva",
				Age:    12,
				Weight: 33,
				Hight:  44,
			},
			testError: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := tC.input1.MarshalJSON()
			if tC.testError {
				if err == nil {
					t.Errorf("missing error")
				}
			} else {
				tC.want1 = append(tC.want1, tC.want2...)
				assert.Equal(t, string(tC.want1), string(got))
			}
		})
	}
}

func TestNewAnimal(t *testing.T) {
	testCases := []struct {
		want   *Animal
		desc   string
		input1 *Animal
	}{
		{
			desc:   "NewAnimal Check 1",
			input1: NewAnimal("Vitaly", "Kapibara", 12, 33, 44),
			want: &Animal{
				ID:     "Vitaly",
				Name:   "Kapibara",
				Age:    12,
				Weight: 33,
				Hight:  44,
			},
		},
		{
			desc:   "NewAnimal Check 2",
			input1: NewAnimal("Vitaly1", "Kapibara2", 12, 33, 44),
			want: &Animal{
				ID:     "Vitaly1",
				Name:   "Kapibara2",
				Age:    12,
				Weight: 33,
				Hight:  44,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.want, tC.input1)
		})
	}
}

func TestGet(t *testing.T) {
	testCases := []struct {
		desc   string
		input1 *Animal
	}{
		{
			desc:   "NewAnimal Check 1",
			input1: NewAnimal("Vitaly", "Kapibara", 12, 33, 44),
		},

		{
			desc:   "NewAnimal Check 2",
			input1: NewAnimal("Bober", "Bober", 12, 33, 44),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := tC.input1.Get()
			if err != nil {
				t.Errorf(err.Error())
			}
		})
	}
}

func TestPost(t *testing.T) {
	testCases := []struct {
		desc   string
		input1 *Animal
	}{
		{
			desc:   "NewAnimal Check 1",
			input1: NewAnimal("Vitaly", "Kapibara", 12, 33, 44),
		},

		{
			desc:   "NewAnimal Check 2",
			input1: NewAnimal("Bober", "Bober", 12, 33, 44),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := tC.input1.Post()
			if err != nil {
				t.Errorf(err.Error())
			}
		})
	}
}
