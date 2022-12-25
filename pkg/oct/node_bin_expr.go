package oct

import "fmt"

type BinExpr struct {
	LHS Node
	RHS Node
	Op  BinOp
}

var _ Node = (*BinExpr)(nil)

func (s BinExpr) String() string {
	return fmt.Sprintf("(%s %s %s)", s.LHS, s.Op, s.RHS)
}
