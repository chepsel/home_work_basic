package fixapp

import (
	"testing"

	"github.com/chepsel/home_work_basic/hw06_testing/fixapp/types"
	"github.com/stretchr/testify/assert"
)

func TestEmployee(t *testing.T) {
	want := "User ID: 10; Age: 25; Name: Rob; Department ID: 3; "
	structure := types.Employee{UserID: 10, Age: 25, Name: "Rob", DepartmentID: 3}
	got := structure.String()
	assert.Equal(t, want, got)
}

func TestFixApp(t *testing.T) { // tdt - шаблон готовый для использования(просто напиши tdt и жми enter)
	testCases := []struct {
		desc  string
		want  []types.Employee
		input string
	}{
		{
			desc:  "Data 3",
			input: "./data3.json",
			want: []types.Employee{
				{
					UserID:       4,
					Age:          64,
					Name:         "Petr",
					DepartmentID: 3,
				}, {
					UserID:       17,
					Age:          89,
					Name:         "Oleg",
					DepartmentID: 0,
				},
			},
		},
		{
			desc:  "Data 2",
			input: "./data2.json",
			want: []types.Employee{
				{
					UserID:       13,
					Age:          33,
					Name:         "Ignat",
					DepartmentID: 5,
				},
				{
					UserID:       9,
					Age:          0,
					Name:         "Vasiliy",
					DepartmentID: 1,
				},
			},
		},
		{
			desc:  "Data 1",
			input: "./data.json",
			want: []types.Employee{
				{
					UserID:       10,
					Age:          25,
					Name:         "Rob",
					DepartmentID: 3,
				}, {
					UserID:       11,
					Age:          30,
					Name:         "George",
					DepartmentID: 2,
				},
			},
		},
		{
			desc:  "Data Null",
			input: "",
			want: []types.Employee{
				{
					UserID:       10,
					Age:          25,
					Name:         "Rob",
					DepartmentID: 3,
				}, {
					UserID:       11,
					Age:          30,
					Name:         "George",
					DepartmentID: 2,
				},
			},
		},
		{
			desc:  "Data Null",
			input: ".",
			want:  nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.input != "." {
				got, _ := FixApp(tC.input)
				assert.Equal(t, tC.want, got)
			} else {
				_, err := FixApp(tC.input)
				if err == nil {
					t.Errorf("missing error")
				}
			}
		})
	}
}
