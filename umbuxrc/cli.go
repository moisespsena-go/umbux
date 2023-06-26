package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/moisespsena-go/umbux"
)

var prettyPrint bool
var lineNumbers bool

func init() {
	flag.BoolVar(&prettyPrint, "prettyprint", true, "Use pretty indentation in output html.")
	flag.BoolVar(&prettyPrint, "pp", true, "Use pretty indentation in output html.")

	flag.BoolVar(&lineNumbers, "linenos", true, "Enable debugging information in output html.")
	flag.BoolVar(&lineNumbers, "ln", true, "Enable debugging information in output html.")

	flag.Parse()
}

func main() {
	// input := flag.Arg(0)

	// if len(input) == 0 {
	//	fmt.Fprintln(os.Stderr, "Please provide an input file. (umbuxc input.umbux)")
	//	os.Exit(1)
	// }

	cmp := umbux.New()
	cmp.PrettyPrint = prettyPrint
	cmp.LineNumbers = lineNumbers

	err := cmp.Parse("p.a")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = cmp.CompileWriter(os.Stdout)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
