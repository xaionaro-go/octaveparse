package oct

import (
	"fmt"
)

type Sqrt struct {
	Base Node
}

var _ Node = (*Sqrt)(nil)

func (s *Sqrt) String() string {
	return fmt.Sprintf("sqrt(%s)", s.Base)
}

func (Sqrt) TypeID() TypeID {
	return TypeIDSqrt
}
