package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"Socialist/internal/config"
	myMiddleware "Socialist/internal/middleware"
	"Socialist/internal/mysql"
	"Socialist/internal/pages"
	"Socialist/internal/repository"
	"Socialist/internal/session"
)

type conf struct {
	DbDsn      string `env:"DB_DSN"`
	Secret     string `env:"SECRET"`
	ServerAddr string `env:"SERVER_ADDR" default:":80"`
}

func main() {
	var cfg conf
	if err := config.Load(&cfg); err != nil {
		panic(err)
	}

	db, err := mysql.NewDB(cfg.DbDsn)
	if err != nil {
		panic(err)
	}

	sessionStorage := session.NewStorage(cfg.Secret)

	h := pages.NewHandlers(
		loadTemplates("./templates/"),
		repository.NewRepository(db),
		sessionStorage,
	)

	loginGuard := myMiddleware.LoginOnly(sessionStorage)

	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.CleanPath)

	r.Get("/", h.Index)
	r.Post("/login", h.RegisterOrLogin)
	r.Get("/login", h.RegisterOrLoginPage)
	r.Get("/logout", h.Logout)

	r.With(loginGuard).Get("/fill-form", h.FormFill)
	r.With(loginGuard).Get("/users", h.ListUsers)
	r.With(loginGuard).Get("/user/{id}", h.ShowUser)
	r.With(loginGuard).Get("/friends", h.ListFriends)

	r.With(loginGuard).Post("/update-user-info", h.UpdateUserInfo)
	r.With(loginGuard).Post("/friend-add/{id}", h.FriendAdd)
	r.With(loginGuard).Post("/friend-remove/{id}", h.FriendRemove)

	r.Get("/static/*", http.FileServer(http.Dir("./resources/")).ServeHTTP)

	fmt.Printf("listening on: http://%s\n", cfg.ServerAddr)
	if err := http.ListenAndServe(cfg.ServerAddr, r); err != nil {
		panic(err)
	}
}
func loadTemplates(templatesDir string) map[string]*template.Template {
	layouts, err := filepath.Glob(templatesDir + "layouts/*.gohtml")
	if err != nil {
		panic(err)
	}

	pageTmpls, err := filepath.Glob(templatesDir + "*.gohtml")
	if err != nil {
		panic(err)
	}

	templates := make(map[string]*template.Template, len(pageTmpls))
	for _, page := range pageTmpls {
		files := []string{page}
		files = append(files, layouts...)
		tmplName := strings.TrimSuffix(filepath.Base(page), ".gohtml")
		templates[tmplName] = template.Must(template.ParseFiles(files...))
	}

	return templates
}
