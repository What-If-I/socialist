package pages

import (
	"net/http"
	"strconv"
)

func (h *Handlers) ListUsers(w http.ResponseWriter, r *http.Request) {
	idUser, err := strconv.Atoi(h.session.GetUserID(r))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	users, err := h.repo.User.ListBrief(idUser)
	if err != nil {
		panic(err)
	}

	if err := h.tmpl["users"].Execute(w, users); err != nil {
		panic(err)
	}
}
