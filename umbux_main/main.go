package main

import (
	"fmt"
	"os"

	"github.com/moisespsena-go/umbux"
)

func main() {
	tfs := umbux.NewTemplateFS(umbux.NewFinderFS(os.DirFS(".")), umbux.NoCache, umbux.Options{false, false, nil})
	if t, err := tfs.Open("teste.umbux"); err != nil {
		fmt.Println(err)
	} else {
		t.Option()
		t.Execute(os.Stdout, map[string]any{"items": []string{"a", "b"}})
	}
}
