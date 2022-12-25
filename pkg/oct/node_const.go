package oct

import "fmt"

type ConstFloat float64

var _ Node = (ConstFloat)(0)

func (f ConstFloat) String() string {
	return fmt.Sprintf("%e", float64(f))
}

func (f ConstFloat) Value() any {
	return float64(f)
}

func (ConstFloat) TypeID() TypeID {
	return TypeIDConst
}

type ConstInt int64

var _ Node = (ConstInt)(0)

func (i ConstInt) String() string {
	return fmt.Sprintf("%d", int64(i))
}

func (f ConstInt) Value() any {
	return int64(f)
}

func (ConstInt) TypeID() TypeID {
	return TypeIDConst
}
