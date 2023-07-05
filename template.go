package umbux

import (
	iofs "io/fs"

	"github.com/moisespsena-go/umbu/html/template"
	"github.com/moisespsena-go/umbux/runtime"
)

func NewTemplate(name, compiled string) (t *template.Template, err error) {
	if t, err = template.New(name).Parse(compiled); err != nil {
		return
	}
	t.Funcs(runtime.FuncMap)
	return
}

func WalkTemplates(options *Options, fs iofs.FS, cb func(t *template.Template) error, ext ...string) error {
	return CompileFS(options, fs, func(name string, compiled *string) (err error) {
		var t *template.Template
		if t, err = NewTemplate(name, *compiled); err != nil {
			return
		}
		return cb(t)
	}, ext...)
}
