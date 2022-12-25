package fmtgo

import (
	"fmt"
	"io"
	"strings"

	"github.com/xaionaro-go/octaveparse/pkg/oct"
)

func Format(w io.Writer, expr oct.Node) error {
	return formatOne(w, expr)
}

func formatOne(w io.Writer, node oct.Node) error {
	var (
		format string
		args   []any
	)
	switch node {
	case oct.I:
		format = "complex(0, 1)"
	default:
		switch node := node.(type) {
		case *oct.BinExpr:
			lhs, err := formatOneString(node.LHS)
			if err != nil {
				return fmt.Errorf("unable to format LHS: %w", err)
			}
			rhs, err := formatOneString(node.RHS)
			if err != nil {
				return fmt.Errorf("unable to format RHS: %w", err)
			}
			switch node.Op {
			case oct.BinOpPower:
				format, args = "math.Pow(%s, %s)", []any{lhs, rhs}
			case oct.BinOpPlus, oct.BinOpMinus, oct.BinOpMultiply, oct.BinOpDivide:
				format, args = "(%s %s %s)", []any{lhs, node.Op, rhs}
			default:
				return fmt.Errorf("unknown binary operator: %s", node.Op)
			}
		case oct.ConstInt:
			format, args = "%d", []any{int64(node)}
		case oct.Array:
			return formatArray(w, node)
		case oct.Matrix:
			return formatMatrix(w, node)
		case *oct.Sqrt:
			base, err := formatOneString(node.Base)
			if err != nil {
				return fmt.Errorf("unable to format Base: %w", err)
			}
			format, args = "math.Sqrt(%s)", []any{base}
		case oct.Sym:
			format, args = "%s", []any{string(node)}
		default:
			return fmt.Errorf("unknown node type '%T'", node)
		}
	}
	_, err := fmt.Fprintf(w, format, args...)
	return err
}

func formatOneString(node oct.Node) (string, error) {
	var buf strings.Builder
	if err := formatOne(&buf, node); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func formatArray(w io.Writer, node oct.Array) error {
	if _, err := fmt.Fprintf(w, "[]any{"); err != nil {
		return fmt.Errorf("unable to write array header: %w", err)
	}

	for idx, item := range node {
		if err := formatOne(w, item); err != nil {
			return fmt.Errorf("unable to format item %d: %w", idx, err)
		}
		if idx != len(node)-1 {
			fmt.Fprintf(w, ",")
		}
	}

	if _, err := fmt.Fprintf(w, "}"); err != nil {
		return fmt.Errorf("unable to write array footer: %w", err)
	}

	return nil
}

func formatMatrix(w io.Writer, node oct.Matrix) error {
	if _, err := fmt.Fprintf(w, "[][]any{"); err != nil {
		return fmt.Errorf("unable to write array header: %w", err)
	}

	for rowIdx, row := range node {
		if err := formatArray(w, oct.Array(row)); err != nil {
			return fmt.Errorf("unable to row %d: %w", rowIdx, err)
		}
		if rowIdx != len(node)-1 {
			fmt.Fprintf(w, ",")
		}
	}

	if _, err := fmt.Fprintf(w, "}"); err != nil {
		return fmt.Errorf("unable to write array footer: %w", err)
	}

	return nil
}
