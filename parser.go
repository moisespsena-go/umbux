package umbux

import (
	"bytes"
	"io"
	iofs "io/fs"
	"path"
	"time"

	"github.com/moisespsena-go/umbu/html/template"
	"github.com/moisespsena-go/umbux/compiler"
	"github.com/moisespsena-go/umbux/runtime"
)

func Parse(opts *Options, finder compiler.FileFinder, name string, r io.Reader) (t *template.Template, err error) {
	var data []byte
	if data, err = io.ReadAll(r); err != nil {
		return
	}

	indentString := ""

	if opts.PrettyPrint {
		indentString = "  "
	}

	ctx := compiler.NewContext(finder, indentString)

	var root *compiler.Root
	if root, err = ctx.ParseReader(name, bytes.NewBuffer(data)); err != nil {
		return
	}

	var s string
	if s, err = ctx.Compile(root); err != nil {
		return
	}

	if t, err = template.New(name).Parse(s); err != nil {
		return
	}

	t.Funcs(runtime.FuncMap)
	return
}

func ParseFS(options *Options, fs iofs.FS, ext ...string) (templates MemCache, err error) {
	templates = make(MemCache)

	var (
		finder = NewFinderFS(fs)
		exts   = map[string]bool{}
	)

	if len(exts) == 0 {
		exts[".pug"] = true
	} else {
		for _, ext := range ext {
			exts[ext] = true
		}
	}

	err = finder.Walk(func(pth string) bool {
		_, ok := exts[path.Ext(pth)]
		return ok
	}, func(f iofs.File, templateName string) (err error) {
		defer f.Close()
		var t *template.Template
		if t, err = Parse(options, finder, templateName, f); err != nil {
			return
		}
		templates.Store(templateName, time.Time{}, t)
		return nil
	})
	return
}
