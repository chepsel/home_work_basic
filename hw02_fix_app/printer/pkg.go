package printer

import (
	"fmt"

	"github.com/chepsel/home_work_basic/hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) {
	for i := 0; i < len(staff); i++ {
		fmt.Println(staff[i])
	}
}
