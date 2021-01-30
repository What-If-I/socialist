package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

const (
	sessionName = "social_session_id"
	userID      = "user_id"
)

type Storage struct {
	s *sessions.CookieStore
}

func NewStorage(secret string) *Storage {
	s := sessions.NewCookieStore([]byte(secret))
	return &Storage{s: s}
}

func (s Storage) GetUserID(r *http.Request) string {
	session, _ := s.s.Get(r, sessionName)
	if session == nil {
		return ""
	}

	v := session.Values[userID]
	if v == nil {
		return ""
	}
	return v.(string)
}

func (s Storage) SetUserID(r *http.Request, w http.ResponseWriter, userId string) error {
	session, _ := s.s.Get(r, sessionName)
	session.Values[userID] = userId
	return session.Save(r, w)
}

func (s Storage) IsNew(r *http.Request) bool {
	session, _ := s.s.Get(r, sessionName)
	if session == nil {
		return true
	}
	return session.IsNew
}

func (s Storage) Save(r *http.Request, w http.ResponseWriter) error {
	session, _ := s.s.Get(r, sessionName)
	return s.s.Save(r, w, session)
}

func (s Storage) Flush(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     sessionName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
}
