package pages

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h *Handlers) FriendAdd(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	idFriend, err := strconv.Atoi(idParam)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	idUser, err := strconv.Atoi(h.session.GetUserID(r))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if err := h.repo.User.AddFriend(idUser, idFriend); err != nil {
		panic(err)
	}

	friend, err := h.repo.User.GetProfile(idFriend)
	if err != nil {
		panic(err)
	}

	if err := h.tmpl["friend_add"].Execute(w, friend); err != nil {
		panic(err)
	}
}
