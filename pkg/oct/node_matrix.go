package oct

import (
	"fmt"
	"strings"
)

type Matrix [][]Node

var _ Node = (Matrix)(nil)

func (m Matrix) String() string {
	var rows []string
	for _, row := range m {
		var cols []string
		for _, col := range row {
			cols = append(cols, col.String())
		}
		rows = append(rows, fmt.Sprintf("[%s]", strings.Join(cols, ", ")))
	}
	return fmt.Sprintf("Matrix([%s])", strings.Join(rows, ", "))
}

func (Matrix) TypeID() TypeID {
	return TypeIDMatrix
}
