package compiler

import iofs "io/fs"

type FileFinder interface {
	Find(name string) (f iofs.File, templateName string, err error)
}

type FinderFS struct {
	FS iofs.FS
}

func (ff FinderFS) Find(name string) (f iofs.File, templateName string, err error) {
	if f, err = ff.FS.Open(name); err != nil {
		return
	}
	templateName = name
	return
}

func NewFinderFS(FS iofs.FS) *FinderFS {
	return &FinderFS{FS: FS}
}
