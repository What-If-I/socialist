package pages

import (
	"net/http"
)

func (h *Handlers) Logout(w http.ResponseWriter, r *http.Request) {
	h.session.Flush(w)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}
