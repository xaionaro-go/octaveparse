package oct

type i struct{}

var I = i{}

var _ Node = I

func (i) String() string {
	return "I"
}

func (i) Value() any {
	return complex(0, 1)
}

func (i) TypeID() TypeID {
	return TypeIDI
}
