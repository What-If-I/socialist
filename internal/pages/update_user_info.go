package pages

import (
	"fmt"
	"net/http"
	"strconv"

	"Socialist/internal/repository"
)

type userInfo struct {
	Name      string `schema:"name,required"`
	Surname   string `schema:"surname,required"`
	Age       int    `schema:"age,required"`
	Gender    string `schema:"gender,required"`
	Interests string `schema:"interests,required"`
	City      string `schema:"city,required"`
}

func (h *Handlers) UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	var form userInfo
	if err := decodeForm(&form, r); err != nil {
		fmt.Print(err)
	}

	userID, err := strconv.Atoi(h.session.GetUserID(r))
	if err != nil {
		panic(err)
	}

	err = h.repo.User.CreateUpdateProfile(userID, repository.Profile{
		ProfileBrief: repository.ProfileBrief{
			Name:    form.Name,
			Surname: form.Surname,
		},
		Age:       form.Age,
		Gender:    form.Gender,
		Interests: form.Interests,
		City:      form.City,
	})
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/", 301)
}
