package pages

import (
	"net"
	"net/http"
	"strconv"
	"strings"

	"Socialist/internal/repository"
)

func (h *Handlers) ListUsers(w http.ResponseWriter, r *http.Request) {
	idUser, err := strconv.Atoi(h.session.GetUserID(r))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	var nameQ, surnameQ, userQuery string
	if queryUser := r.URL.Query()["user"]; len(queryUser) != 0 {
		userQuery = queryUser[0]
		words := strings.Split(userQuery, " ")
		nameQ = words[0]
		if len(words) > 1 {
			surnameQ = strings.Join(words[1:], " ")
		}
	}

	var users []repository.ProfileBrief
	if nameQ != "" {
		users, err = h.repo.User.Search(r.Context(), idUser, nameQ, surnameQ)
	} else {
		users, err = h.repo.User.ListBrief(idUser)
		if err != nil {
			panic(err)
		}
	}

	ctx := struct {
		Users     []repository.ProfileBrief
		UserQuery string
	}{
		Users:     users,
		UserQuery: userQuery,
	}

	if err := h.tmpl["users"].Execute(w, ctx); err != nil {
		if _, ok := err.(*net.OpError); ok {
			return
		}
		panic(err)
	}
}
