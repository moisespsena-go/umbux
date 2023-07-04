package umbux

import (
	"time"

	"github.com/moisespsena-go/umbu/html/template"
)

type (
	Cache interface {
		Get(name string, mod time.Time) *template.Template
		Store(name string, mod time.Time, tmpl *template.Template) error
	}
	noCache struct {
	}

	MemCache map[string]*template.Template
)

var NoCache noCache

func (noCache) Get(string, time.Time) *template.Template {
	return nil
}

func (noCache) Store(string, time.Time, *template.Template) error {
	return nil
}

func (m MemCache) Get(name string, _ time.Time) *template.Template {
	return m[name]
}

func (m MemCache) Store(name string, _ time.Time, tmpl *template.Template) error {
	m[name] = tmpl
	return nil
}
