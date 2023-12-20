package main

import (
	"fmt"

	"github.com/chepsel/home_work_basic/hw02_fix_app/printer"
	"github.com/chepsel/home_work_basic/hw02_fix_app/reader"
	"github.com/chepsel/home_work_basic/hw02_fix_app/types"
)

func main() {
	var path string

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)
	if err != nil {
		fmt.Print(err)
	}

	printer.PrintStaff(staff)
}
