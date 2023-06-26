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
)

var NoCache noCache

func (noCache) Get(name string, mod time.Time) *template.Template {
	return nil
}

func (noCache) Store(name string, mod time.Time, tmpl *template.Template) error {
	return nil
}
