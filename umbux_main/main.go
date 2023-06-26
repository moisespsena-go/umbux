package main

import (
	"os"
)

func main() {
	umbux.CompileToWriter(os.Stdout, `#{1 ? 2 : 3}`, umbux.Options{false, false, nil, nil})
}
