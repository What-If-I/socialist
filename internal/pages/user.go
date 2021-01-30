package pages

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h *Handlers) ShowUser(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	profile, err := h.repo.User.GetProfile(id)
	if err != nil {
		panic(err)
	}

	if err := h.tmpl["user"].Execute(w, profile); err != nil {
		panic(err)
	}
}
