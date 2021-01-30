package pages

import (
	"net/http"
	"strconv"
)

func (h *Handlers) ListFriends(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(h.session.GetUserID(r))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	friends, err := h.repo.User.ListFriends(id)
	if err != nil {
		panic(err)
	}

	if err := h.tmpl["friends"].Execute(w, friends); err != nil {
		panic(err)
	}
}
