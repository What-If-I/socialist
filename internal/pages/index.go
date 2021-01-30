package pages

import (
	"net/http"
	"strconv"
)

func (h *Handlers) Index(w http.ResponseWriter, r *http.Request) {
	idParam := h.session.GetUserID(r)
	if idParam == "" {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	user, err := h.repo.User.GetProfile(id)
	if err != nil {
		panic(err)
	}

	if user.Name == "" {
		http.Redirect(w, r, "/fill-form", http.StatusTemporaryRedirect)
		return
	}

	if err := h.tmpl["index"].Execute(w, user); err != nil {
		panic(err)
	}
}
