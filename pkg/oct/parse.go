package oct

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
)

type Parser struct {
	Syms []string
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) AddSym(syms ...string) {
	p.Syms = append(p.Syms, syms...)
}

func (p *Parser) Parse(in string) (Node, error) {
	lexer, err := p.compile()
	if err != nil {
		return nil, fmt.Errorf("internal error: unable to prepare a lexer machine: %w", err)
	}

	scanner, err := lexer.Scanner([]byte(in))
	if err != nil {
		return nil, fmt.Errorf("internal error: unable to initialize a lexer scanner: %w", err)
	}

	result, token, err := scanOne(scanner)
	if err != nil {
		return nil, err
	}

	token, _ = skipSpaceIfAny(token, scanner)
	if token != nil {
		return nil, fmt.Errorf("got an extra token in the end: %+v", token)
	}

	return result, nil
}

func scanOne(s *lexmachine.Scanner) (Node, *lexmachine.Token, error) {
	tok, err, eof := s.Next()
	if eof {
		return nil, nil, fmt.Errorf("unexpected end")
	}
	if err != nil {
		return nil, nil, fmt.Errorf("unable to parse: %w", err)
	}
	token := tok.(*lexmachine.Token)

	return parseOne(token, s)
}

func parseOne(token *lexmachine.Token, s *lexmachine.Scanner) (Node, *lexmachine.Token, error) {
	var terms []Node
	for {
		term, err := parseTerm(token, s)
		if err == errNotAnTerm {
			break
		}
		if err == errEmpty {
			break
		}
		if err != nil {
			return nil, nil, fmt.Errorf("unable to parse term: %w", err)
		}
		if term == nil {
			return nil, nil, fmt.Errorf("internal error: term == nil")
		}

		terms = append(terms, term)

		tok, err, eof := s.Next()
		if eof {
			token = nil
			break
		}
		if err != nil {
			return nil, nil, fmt.Errorf("unable to parse: %w", err)
		}
		token = tok.(*lexmachine.Token)
	}

	result, err := compileNode(terms)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to compile the expression using terms: %v (scanner TC: %v): %w", terms, s.TC, err)
	}
	return result, token, nil
}

func compileNode(terms []Node) (Node, error) {
	for _, ops := range [][]BinOp{
		{BinOpPower},
		{BinOpMultiply, BinOpDivide},
		{BinOpPlus, BinOpMinus},
	} {
		dstIdx := 0
		for idx := 0; idx < len(terms); idx++ {
			term := terms[idx]
			op, ok := term.(BinOp)
			if !ok {
				terms[dstIdx] = term
				dstIdx++
				continue
			}
			found := false
			for _, opCmp := range ops {
				if op == opCmp {
					found = true
					break
				}
			}
			if !found {
				terms[dstIdx] = term
				dstIdx++
				continue
			}
			if idx == len(terms)-1 {
				return nil, fmt.Errorf("found '%v' without RHS", term)
			}
			var lhs Node
			if dstIdx == 0 {
				lhs = ConstInt(0) // for example expression "-1" should be interpreted as "0 - 1".
				dstIdx++
			} else {
				lhs = terms[dstIdx-1]
			}
			terms[dstIdx-1] = &BinExpr{
				Op:  op,
				LHS: lhs,
				RHS: terms[idx+1],
			}
			idx++
		}
		terms = terms[:dstIdx]
	}
	if len(terms) != 1 {
		return nil, fmt.Errorf("expected to find exactly one term, but found %d: %v", len(terms), terms)
	}

	return terms[0], nil
}

type errNotAnItemT struct{}

var errNotAnTerm = errNotAnItemT{}

func (err errNotAnItemT) Error() string {
	return "not an item"
}

type errEmptyT struct{}

var errEmpty = errEmptyT{}

func (err errEmptyT) Error() string {
	return "empty"
}

func skipSpaceIfAny(token *lexmachine.Token, s *lexmachine.Scanner) (*lexmachine.Token, error) {
	if TypeID(token.Type) != TypeIDSpace {
		return token, nil
	}
	tok, err, eof := s.Next()
	if eof {
		return nil, errEmpty
	}
	if err != nil {
		return nil, fmt.Errorf("unable to parse: %w", err)
	}

	return tok.(*lexmachine.Token), nil
}

func parseTerm(token *lexmachine.Token, s *lexmachine.Scanner) (Node, error) {
	token, err := skipSpaceIfAny(token, s)
	if err != nil {
		return nil, err
	}

	var result Node
	switch TypeID(token.Type) {
	case TypeIDMatrix:
		result, err = scanMatrix(s)
	case TypeIDSqrt:
		result, err = scanFuncArg(s)
		if result != nil {
			result = &Sqrt{Base: result}
		}
	case TypeIDParenthesisOpen:
		result, err = parseParenthesisInternals(
			s,
			TypeIDParenthesisOpen, TypeIDParenthesisClose,
			func(s *lexmachine.Scanner) (Node, *lexmachine.Token, error) {
				return scanOne(s)
			},
		)
	case TypeIDBracketOpen:
		var nodes []Node
		nodes, token, err = scanList(s, TypeIDBracketClose)
		if err != nil {
			return nil, fmt.Errorf("unable to parse %+v: %w", token, err)
		}
		if nodes != nil {
			result = Array(nodes)
		}
		if TypeID(token.Type) != TypeIDBracketClose {
			return nil, fmt.Errorf("expected ']', but got: %+v", token)
		}
	case TypeIDConst:
		result, err = parseConst(token)
	case TypeIDBinOp:
		return parseBinOp(token)
	case TypeIDSym:
		return Sym(token.Lexeme), nil
	case TypeIDI:
		return I, nil
	default:
		return nil, errNotAnTerm
	}

	if err != nil {
		return nil, fmt.Errorf("unable to parse data associated with token %v: %w", token, err)
	}

	return result, nil
}

func parseBinOp(token *lexmachine.Token) (Node, error) {
	op := BinOp(token.Lexeme)
	switch op {
	case BinOpPower, BinOpMultiply, BinOpDivide, BinOpPlus, BinOpMinus:
		return op, nil
	}
	return nil, fmt.Errorf("internal error: invalid binary operator: '%s'", op)
}

func parseConst(token *lexmachine.Token) (Node, error) {
	in := string(token.Lexeme)
	switch {
	case strings.Contains(in, "."):
		f, err := strconv.ParseFloat(in, 64)
		if err != nil {
			return nil, fmt.Errorf("unable to parse a float number '%s': %w", in, err)
		}
		return ConstFloat(f), nil
	default:
		i, err := strconv.ParseInt(in, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("unable to parse an integer '%s': %w", in, err)
		}
		return ConstInt(i), nil
	}
}

func scanList(
	s *lexmachine.Scanner,
	listEndTypeID TypeID,
) ([]Node, *lexmachine.Token, error) {
	var result []Node
	for {
		tok, err, eof := s.Next()
		if eof {
			return nil, nil, fmt.Errorf("unexpected end")
		}
		if err != nil {
			return nil, nil, fmt.Errorf("unable to parse: %w", err)
		}
		token := tok.(*lexmachine.Token)

		node, token, err := parseOne(token, s)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to parse item %d: %w", len(result), err)
		}
		result = append(result, node)

		if token == nil {
			return result, nil, nil
		}
		switch TypeID(token.Type) {
		case TypeIDComma:
		case listEndTypeID:
			return result, token, nil
		default:
			return nil, nil, fmt.Errorf("expected a comma or '%v', but received: %+v", token, listEndTypeID)
		}
	}
}

func scanMatrix(s *lexmachine.Scanner) (Matrix, error) {
	arg, err := scanFuncArg(s)
	if err != nil {
		return nil, fmt.Errorf("invalid matrix-argument: %w", err)
	}

	rows, ok := arg.(Array)
	if !ok {
		return nil, fmt.Errorf("expected Martix (an Array of an Array), but got %T instead", arg)
	}

	var result Matrix
	for _, row := range rows {
		cols, ok := row.(Array)
		if !ok {
			return nil, fmt.Errorf("expected Martix (an Array of an Array), but got an Array with %T inside instead", arg)
		}
		result = append(result, cols)
	}

	return result, nil
}

func scanFuncArg(s *lexmachine.Scanner) (Node, error) {
	args, err := scanFuncArgs(s)
	if err != nil {
		return nil, fmt.Errorf("unable to scan function arguments: %w", err)
	}
	if len(args) != 1 {
		return nil, fmt.Errorf("expected 1 argument, but got %d", len(args))
	}
	return args[0], nil
}

func scanFuncArgs(s *lexmachine.Scanner) ([]Node, error) {
	result, err := scanParenthesis(s, TypeIDParenthesisOpen, TypeIDParenthesisClose, func(s *lexmachine.Scanner) (Node, *lexmachine.Token, error) {
		nodes, token, err := scanList(s, TypeIDParenthesisClose)
		if err != nil {
			return nil, nil, err
		}
		return Array(nodes), token, nil
	})
	if err != nil {
		return nil, err
	}
	args, ok := result.(Array)
	if !ok {
		return nil, fmt.Errorf("internal error: expected Array, but got %T", result)
	}
	return args, nil
}

func scanParenthesis(
	s *lexmachine.Scanner,
	openTypeID, closeTypeID TypeID,
	scanInternalsFn func(s *lexmachine.Scanner) (Node, *lexmachine.Token, error),
) (Node, error) {
	tok, err, eof := s.Next()
	if eof {
		return nil, fmt.Errorf("unexpected end: expected '%v'", openTypeID)
	}
	if err != nil {
		return nil, fmt.Errorf("unable to parse '%v': %w", openTypeID, err)
	}
	token := tok.(*lexmachine.Token)

	if TypeID(token.Type) != openTypeID {
		return nil, fmt.Errorf("unexpected token %+v: expected '%v'", token, openTypeID)
	}

	return parseParenthesisInternals(s, openTypeID, closeTypeID, scanInternalsFn)
}

func parseParenthesisInternals(
	s *lexmachine.Scanner,
	openTypeID, closeTypeID TypeID,
	scanInternalsFn func(s *lexmachine.Scanner) (Node, *lexmachine.Token, error),
) (Node, error) {
	result, token, err := scanInternalsFn(s)
	if err != nil {
		return nil, fmt.Errorf("unable to scan parenthesis/bracket '%v...%v' internals: %w", openTypeID, closeTypeID, err)
	}

	if token == nil {
		tok, err, eof := s.Next()
		if eof {
			return nil, fmt.Errorf("unexpected end: parenthesis/bracket '%v' was never closed", openTypeID)
		}
		if err != nil {
			return nil, fmt.Errorf("unable to parse token '%v': %w", closeTypeID, err)
		}
		token = tok.(*lexmachine.Token)
	}
	if TypeID(token.Type) != closeTypeID {
		return nil, fmt.Errorf("unexpected token %+v: expected token '%v'", token, closeTypeID)
	}

	return result, nil
}

func token(typeID TypeID) lexmachine.Action {
	return func(s *lexmachine.Scanner, m *machines.Match) (any, error) {
		return s.Token(int(typeID), m.Bytes, m), nil
	}
}

func (p *Parser) compile() (*lexmachine.Lexer, error) {
	binaryOps := []BinOp{
		BinOpPower,
		BinOpMultiply,
		BinOpDivide,
		BinOpPlus,
		BinOpMinus,
	}

	others := map[TypeID]string{
		TypeIDI:                "I",
		TypeIDParenthesisOpen:  "(",
		TypeIDParenthesisClose: ")",
		TypeIDBracketOpen:      "[",
		TypeIDBracketClose:     "]",
		TypeIDComma:            ",",
	}

	lexer := lexmachine.NewLexer()
	lexer.Add([]byte("Matrix"), token(TypeIDMatrix))
	lexer.Add([]byte("sqrt"), token(TypeIDSqrt))
	for _, binOp := range binaryOps {
		lexer.Add([]byte(regexp.QuoteMeta(string(binOp))), token(TypeIDBinOp))
	}
	for typeID, other := range others {
		lexer.Add([]byte(regexp.QuoteMeta(other)), token(typeID))
	}
	for _, sym := range p.Syms {
		lexer.Add([]byte(regexp.QuoteMeta(sym)), token(TypeIDSym))
	}
	lexer.Add([]byte(`[0-9]+(\.[0-9]+)?`), token(TypeIDConst))
	lexer.Add([]byte(`[0-9]+(\.[0-9]+)?`), token(TypeIDConst))
	lexer.Add([]byte(`[ \t\n]+`), token(TypeIDSpace))

	err := lexer.Compile()
	if err != nil {
		return nil, fmt.Errorf("unable to compile the lexer machine: %w", err)
	}

	return lexer, nil
}
