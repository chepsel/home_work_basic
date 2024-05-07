package main

import (
	"net/http"
	"os"
	"testing"

	"github.com/chepsel/home_work_basic/hw13_http/client"
	"github.com/gin-gonic/gin"
)

var _ = func() bool {
	testing.Init()
	os.Args = append(os.Args, "-u=http://localhost:8083/v1/restapi/animal")
	go fakeServer()
	return true
}()

func fakeServer() {
	server := gin.Default()
	v1 := server.Group("/v1")
	restAPI := v1.Group("/restapi")
	{
		restAPI.GET("/animal", func(c *gin.Context) {
			id := c.Query("id")
			if len(id) == 0 {
				c.JSONP(http.StatusBadRequest, "Wrong request params")
			} else {
				c.JSONP(200, `{"id": "Vitaly","name": "Kapibara","age": 12,"weight": 33,"hight": 44}`)
			}
		})
		restAPI.POST("/animal", func(c *gin.Context) {
			var animal client.Animal
			err := c.ShouldBindJSON(&animal)
			switch {
			case err != nil:
				c.JSONP(http.StatusBadRequest, gin.H{"err": err.Error()})
			case len(animal.ID) > 0 && len(animal.Name) > 0:
				c.JSONP(200, "success")
			default:
				c.JSONP(http.StatusBadRequest, gin.H{"err": "missing id or name key"})
			}
		})
	}
	server.Run(":8083")
}

func TestClient(t *testing.T) {
	testCases := []struct {
		desc      string
		input1    string
		errorTest bool
	}{
		{
			desc:   "NewAnimal Check 1",
			input1: conf.url,
		},
		{
			desc:      "Error check",
			input1:    "http://localhost:8083/v1/restapi/animal3",
			errorTest: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := client.Client(tC.input1)
			if err != nil && tC.errorTest == false {
				t.Errorf(err.Error())
			} else if tC.errorTest && err == nil {
				t.Errorf("lost error")
			}
		})
	}
}
