package umbux

import (
	"bytes"
	"io"
	iofs "io/fs"
	"path"

	"github.com/moisespsena-go/umbux/compiler"
)

func CompileFile(options *Options, finder FileFinder, name string) (s, templateName string, err error) {
	var f iofs.File
	if f, templateName, err = finder.Find(name); err != nil {
		return
	}
	s, err = CompileReader(options, finder, templateName, f)
	return
}

func CompileReader(options *Options, finder FileFinder, name string, f iofs.File) (s string, err error) {
	var data []byte
	if data, err = io.ReadAll(f); err != nil {
		return
	}

	defer f.Close()
	return CompileBytes(options, finder, name, data)
}

func CompileBytes(options *Options, finder FileFinder, name string, data []byte) (s string, err error) {
	if options.PrettyPrint && options.IndentString == "" {
		options.IndentString = "  "
	}

	ctx := compiler.NewContext(finder, options.IndentString)

	var root *compiler.Root
	if root, err = ctx.ParseReader(name, bytes.NewBuffer(data)); err != nil {
		return
	}

	return ctx.Compile(root)
}

func CompileFS(options *Options, fs iofs.FS, cb func(name string, compiled *string) error, ext ...string) (err error) {
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
		var s string
		if s, err = CompileReader(options, finder, templateName, f); err == nil {
			err = cb(templateName, &s)
		}
		return nil
	})
	return
}
