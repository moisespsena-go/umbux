package main

import (
	"fmt"
	"os"

	"github.com/moisespsena-go/umbux"
)

func main() {
	var (
		options = umbux.Options{false, false, nil}
	)
	templates, err := umbux.ParseFS(&options, os.DirFS("examples"))

	if err != nil {
		panic(err)
	}

	t := templates["mixin.pug"]
	err = t.Execute(os.Stdout, map[string]any{"items": []string{"a", "b"}, "foo": func(args ...interface{}) interface{} {
		return args[0]
	}})
	fmt.Println(err)
}
