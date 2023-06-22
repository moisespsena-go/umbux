package main

import (
	"os"

	"github.com/eknkc/amber"
)

func main() {
	amber.CompileToWriter(os.Stdout, `#{1 ? 2 : 3}`, amber.Options{false, false, nil, nil})
}
