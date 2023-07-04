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
	input := flag.Arg(0)

	if len(input) == 0 {
		fmt.Fprintln(os.Stderr, "Please provide an input file. (umbuxc input.pug)")
		os.Exit(1)
	}

	tfs := umbux.NewTemplateFS(umbux.NewFinderFS(os.DirFS(".")), umbux.NoCache, umbux.Options{true, true, nil})
	if t, err := tfs.Open(input); err != nil {
		fmt.Println(err)
	} else {
		t.Option()
		err := t.Execute(os.Stdout, map[string]any{"items": []string{"a", "b"}, "foo": func(args ...interface{}) interface{} {
			return args[0]
		}})
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
	}
}
