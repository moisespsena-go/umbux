package amber

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestCompiler_CompileMixin(t *testing.T) {
	err := compileW(os.Stdout, `
mixin repeat($value, $count)
	+repeat($value, $level+1)

+repeat(1, 0)
`)
	fmt.Println(err)
}

func TestCompiler_CompileConcat(t *testing.T) {
	err := compileW(os.Stdout, `
#{-1}
`)
	fmt.Println(err)
}

func compile(tpl string) (result string, err error) {
	result, err = CompileToString(tpl, Options{false, false, nil, nil})
	return
}

func compileW(w io.Writer, tpl string) (err error) {
	return CompileToWriter(w, tpl, Options{false, false, nil, nil})
}
