package oct

type BinOp string

const (
	BinOpPlus     = BinOp("+")
	BinOpMinus    = BinOp("-")
	BinOpMultiply = BinOp("*")
	BinOpDivide   = BinOp("/")
	BinOpPower    = BinOp("**")
)

var _ Node = BinOp("")

func (s BinOp) String() string {
	return string(s)
}

func (BinOp) TypeID() TypeID {
	return TypeIDBinOp
}
