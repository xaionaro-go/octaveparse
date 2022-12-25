package oct

import (
	"fmt"
	"strings"
)

type Array []Node

var _ Node = (Matrix)(nil)

func (m Array) String() string {
	var items []string
	for _, item := range m {
		items = append(items, item.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(items, ", "))
}
