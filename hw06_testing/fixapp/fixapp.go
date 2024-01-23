package fixapp

import (
	"github.com/chepsel/home_work_basic/hw06_testing/fixapp/reader"
	"github.com/chepsel/home_work_basic/hw06_testing/fixapp/types"
)

func FixApp(path string) ([]types.Employee, error) {
	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "./data.json"
	}

	staff, err = reader.ReadJSON(path)
	if err != nil {
		return nil, err
	}
	return staff, nil
}
