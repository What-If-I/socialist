package pages

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type loginForm struct {
	Name     string `schema:"username,required"`
	Password string `schema:"password,required"`
}

type UserInfo struct {
	Name      string
	Surname   string
	Age       int
	Gender    string
	Interests string
	City      string
}

func (h *Handlers) RegisterOrLoginPage(w http.ResponseWriter, r *http.Request) {
	if err := h.tmpl["register_login"].Execute(w, nil); err != nil {
		panic(err)
	}
}

func (h *Handlers) RegisterOrLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	var form loginForm
	if err := decodeForm(&form, r); err != nil {
		fmt.Print(err)
	}

	usr, err := h.repo.User.LoginUser(form.Name)
	if err != nil {
		panic(err)
	}

	if usr.ID != "" {
		if ok := comparePasswords(usr.PasswordHash, form.Password); !ok {
			http.Error(w, "Wrong password", http.StatusForbidden)
			return
		}
	} else {
		pwdHash := hashAndSalt(form.Password)
		id, err := h.repo.User.CreateUser(form.Name, pwdHash)
		if err != nil {
			panic(err)
		}
		usr.ID = id
	}

	if err := h.session.SetUserID(r, w, usr.ID); err != nil {
		panic(err)
	}

	if err := h.session.Save(r, w); err != nil {
		panic(err)
	}

	if !usr.FilledForm {
		http.Redirect(w, r, "/fill-form", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func hashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		return false
	}

	return true
}
