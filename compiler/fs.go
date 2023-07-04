package compiler

import iofs "io/fs"

type FileFindWalker interface {
	Walk(accept func(path string) bool, do func(f iofs.File, templateName string) error) error
}

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

func (ff FinderFS) Walk(accept func(path string) bool, do func(f iofs.File, templateName string) error) error {
	return iofs.WalkDir(ff.FS, ".", func(path string, d iofs.DirEntry, ierr error) (err error) {
		if ierr != nil {
			return ierr
		}

		if accept(path) {
			var f iofs.File
			if f, err = ff.FS.Open(path); err != nil {
				return
			}
			err = do(f, path)
		}
		return
	})
}

func NewFinderFS(FS iofs.FS) *FinderFS {
	return &FinderFS{FS: FS}
}
