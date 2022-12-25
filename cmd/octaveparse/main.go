package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/xaionaro-go/octaveparse/pkg/format/fmtc"
	"github.com/xaionaro-go/octaveparse/pkg/format/fmtgo"
	"github.com/xaionaro-go/octaveparse/pkg/oct"
)

type format int

const (
	formatUndefined = format(iota)
	formatUnknown
	formatGo
	formatC
)

func (f format) String() string {
	switch f {
	case formatUndefined:
		return "<format>"
	case formatGo:
		return "go"
	case formatC:
		return "c"
	}
	return fmt.Sprintf("undefined_%d", int(f))
}
func (f *format) Set(in string) error {
	switch strings.ToLower(strings.Trim(in, " \n\t\r")) {
	case "go":
		*f = formatGo
	case "c":
		*f = formatC
	default:
		*f = formatUnknown
		return fmt.Errorf("unknown format '%s'", in)
	}
	return nil
}

func main() {
	outFormat := formatGo
	flag.Var(&outFormat, "format", "available options: 'c', 'go'")
	symsFlag := flag.String("syms", "", "a comma-separated list of symbols to add to the parser")
	flag.Parse()

	if flag.NArg() > 1 {
		log.Fatalf("too many arguments: %d > 1", flag.NArg())
	}

	syms := strings.Split(*symsFlag, ",")

	inputReader := os.Stdin
	if flag.NArg() == 1 {
		inputPath := flag.Arg(0)
		f, err := os.Open(inputPath)
		if err != nil {
			log.Fatalf("cannot open '%s': %v", inputPath, err)
		}
		defer f.Close()
		inputReader = f
	}

	octaveExprBytes, err := io.ReadAll(inputReader)
	if err != nil {
		log.Fatalf("unable to read the input: %v", err)
	}

	parser := oct.NewParser()
	for _, sym := range syms {
		parser.AddSym(strings.Trim(sym, " \t\r\n"))
	}
	octaveExpr, err := parser.Parse(string(octaveExprBytes))
	if err != nil {
		log.Fatalf("unable to parse the octave expression '%s': %v", octaveExprBytes, err)
	}

	outWriter := os.Stdout

	switch outFormat {
	case formatGo:
		err = fmtgo.Format(outWriter, octaveExpr)
	case formatC:
		err = fmtc.Format(outWriter, octaveExpr)
	default:
		log.Fatalf("unexpected -format value: '%s'", outFormat)
	}
	if err != nil {
		log.Fatalf("unable to format with '%s': %v", outFormat, err)
	}
	fmt.Fprintf(outWriter, "\n")
}
