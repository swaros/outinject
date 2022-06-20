package outinject

import "fmt"

type PlainParse struct {
}

func (c PlainParse) Parse(i ...interface{}) string {
	return fmt.Sprint(i...)
}

func (c PlainParse) Enable(mo *MOut) bool {
	return true
}
