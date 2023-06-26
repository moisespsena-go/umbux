package umbux

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestCompiler_CompileMixin(t *testing.T) {
	compileExpect(t, `
mixin repeat($value, $count)
	+repeat($value, $level+1)
+repeat(1, 0)
`, `{{define "repeat" $value $count}}{{template "repeat" . $value ($level + 1)}}{{end}}{{template "repeat" . 1 0}}`)
}

func TestCompiler_CompileConcat(t *testing.T) {
	err := compileW(os.Stdout, `
#{-1}
`)
	fmt.Println(err)
}

func compileExpect(t *testing.T, tpl, expected string) {
	var w bytes.Buffer
	if err := compileW(&w, tpl); err != nil {
		t.Fatal(err)
	}
	expect(strings.TrimSpace(w.String()), strings.TrimSpace(expected), t)
}

func compile(tpl string) (result string, err error) {
	result, err = CompileToString(tpl, Options{false, false, nil, nil})
	return
}

func compileW(w io.Writer, tpl string) (err error) {
	return CompileToWriter(w, tpl, Options{false, false, nil, nil})
}
