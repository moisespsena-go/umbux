package umbux

import (
	"bytes"
	"fmt"
	"io"
	iofs "io/fs"
	"os"

	"github.com/moisespsena-go/umbu/html/template"
	"github.com/moisespsena-go/umbux/compiler"
)

type FileFinder = compiler.FileFinder

type FinderFS = compiler.FinderFS

func NewFinderFS(FS iofs.FS) *FinderFS {
	return &FinderFS{FS: FS}
}

type TemplateFS struct {
	Finder  FileFinder
	Cache   Cache
	Options Options
}

func NewTemplateFS(finder FileFinder, cache Cache, options Options) *TemplateFS {
	return &TemplateFS{Finder: finder, Cache: cache, Options: options}
}

func (fs *TemplateFS) Open(name string) (t *template.Template, err error) {
	var (
		f            iofs.File
		templateName string
	)
	if f, templateName, err = fs.Finder.Find(name); err != nil {
		return
	}

	defer f.Close()

	var fi os.FileInfo
	if fi, err = f.Stat(); err != nil {
		return
	}

	t = fs.Cache.Get(templateName, fi.ModTime())
	if t == nil {
		var data []byte
		if data, err = io.ReadAll(f); err != nil {
			return
		}

		indentString := ""

		if fs.Options.PrettyPrint {
			indentString = "  "
		}

		ctx := compiler.NewContext(fs.Finder, indentString)

		var root *compiler.Root
		if root, err = ctx.ParseReader(name, bytes.NewBuffer(data)); err != nil {
			return
		}

		var s string
		if s, err = ctx.Compile(root); err != nil {
			return
		}

		fmt.Println(s)

		if t, err = template.New(templateName).Parse(s); err != nil {
			return
		}
		if err = fs.Cache.Store(templateName, fi.ModTime(), t); err != nil {
			return
		}
	}

	return
}
