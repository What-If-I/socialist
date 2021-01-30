package pages

import (
	"net/http"
	"strconv"
)

func (h *Handlers) FormFill(w http.ResponseWriter, r *http.Request) {
	idParam := h.session.GetUserID(r)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	user, err := h.repo.User.GetProfile(id)
	if err != nil {
		panic(err)
	}

	if err := h.tmpl["user_form_edit"].Execute(w, user); err != nil {
		panic(err)
	}
}
