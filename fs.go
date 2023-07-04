package umbux

import (
	iofs "io/fs"
	"os"

	"github.com/moisespsena-go/umbu/html/template"
	"github.com/moisespsena-go/umbux/compiler"
	"github.com/moisespsena-go/umbux/runtime"
)

type (
	FileFindWalker = compiler.FileFindWalker
	FileFinder     = compiler.FileFinder
	FinderFS       = compiler.FinderFS
)

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
		if t, err = Parse(&fs.Options, fs.Finder, templateName, f); err != nil {
			return
		}

		if err = fs.Cache.Store(templateName, fi.ModTime(), t); err != nil {
			return
		}

		t.Funcs(runtime.FuncMap)
	}

	return
}

func (fs *TemplateFS) ExecutorOf(name string) (exc *template.Executor, err error) {
	var t *template.Template
	if t, err = fs.Open(name); err != nil {
		return
	}
	exc = t.CreateExecutor()
	exc.StateOptions.DotOverrideDisabled = true
	return
}
