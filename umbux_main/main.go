package main

import (
	"fmt"
	"os"

	"github.com/moisespsena-go/umbux"
)

func main() {
	tfs := umbux.NewTemplateFS(umbux.NewFinderFS(os.DirFS(".")), umbux.NoCache, umbux.Options{false, false, nil})
	if t, err := tfs.Open("examples/mixin.pug"); err != nil {
		fmt.Println(err)
	} else {
		t.Option()
		err := t.Execute(os.Stdout, map[string]any{"items": []string{"a", "b"}, "foo": func(args ...interface{}) interface{} {
			return args[0]
		}})
		fmt.Println(err)
	}
}
