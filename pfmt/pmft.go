package pfmt

import (
	"fmt"
)

// Time converts integer into minunte:second format (2m3s)
func Time(s int) string {
	m := int(s / 60)
	s = int(s % 60)

	var str string
	if m > 0 {
		str += fmt.Sprintf("%dm", m)
	}
	str += fmt.Sprintf("%ds", s)
	return str
}
