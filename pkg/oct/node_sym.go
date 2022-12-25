package oct

type Sym string

var _ Node = Sym("")

func (s Sym) String() string {
	return string(s)
}

func (Sym) TypeID() TypeID {
	return TypeIDSym
}
