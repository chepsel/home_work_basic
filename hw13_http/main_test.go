package main

import (
	"os"
	"testing"

	"github.com/chepsel/home_work_basic/hw13_http/client"
	"github.com/chepsel/home_work_basic/hw13_http/server"
)

var _ = func() bool {
	testing.Init()
	os.Args = append(os.Args, "-u=http://localhost:8083/v1/restapi/animal")
	go fakeServer()
	return true
}()

func fakeServer() {
	server.Server("localhost", "8083")
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
