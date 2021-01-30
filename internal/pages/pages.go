package pages

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/schema"

	"Socialist/internal/repository"
	"Socialist/internal/session"
)

type Handlers struct {
	tmpl    map[string]*template.Template
	repo    *repository.Repository
	session *session.Storage
}

func NewHandlers(templates map[string]*template.Template, repo *repository.Repository, storage *session.Storage) *Handlers {
	return &Handlers{
		tmpl:    templates,
		repo:    repo,
		session: storage,
	}
}

var decoder = schema.NewDecoder()

func decodeForm(dst interface{}, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	if err := decoder.Decode(dst, r.PostForm); err != nil {
		return err
	}

	return nil
}

func (h *Handlers) PlainPage(tmplName string) func(w http.ResponseWriter, r *http.Request) {
	if _, ok := h.tmpl[tmplName]; !ok {
		panic(fmt.Sprintf("template '%s' doesn't exist", tmplName))
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if err := h.tmpl[tmplName].Execute(w, nil); err != nil {
			panic(err)
		}
	}

}
