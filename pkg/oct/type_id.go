package oct

import "fmt"

type TypeID int

const (
	TypeIDUndefined = TypeID(iota)
	TypeIDBinExpr
	TypeIDBinOp
	TypeIDConst
	TypeIDI
	TypeIDMatrix
	TypeIDSqrt
	TypeIDSym
	TypeIDParenthesisOpen
	TypeIDParenthesisClose
	TypeIDBracketOpen
	TypeIDBracketClose
	TypeIDComma
	TypeIDSpace
)

func (t TypeID) String() string {
	switch t {
	case TypeIDUndefined:
		return "<undefined>"
	case TypeIDBinExpr:
		return "binary_expression"
	case TypeIDBinOp:
		return "binary_operator"
	case TypeIDConst:
		return "constant"
	case TypeIDI:
		return "i"
	case TypeIDMatrix:
		return "matrix"
	case TypeIDSqrt:
		return "sqrt"
	case TypeIDSym:
		return "sym"
	case TypeIDParenthesisOpen:
		return "("
	case TypeIDParenthesisClose:
		return ")"
	case TypeIDBracketOpen:
		return "["
	case TypeIDBracketClose:
		return "]"
	case TypeIDComma:
		return ","
	case TypeIDSpace:
		return "SPACE"
	}
	return fmt.Sprintf("unknown_%d", t)
}
