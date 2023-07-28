package domain

import (
	"github.com/frame-go/framego"

	"{{ cookiecutter.go_module }}/internal/cache"
	"{{ cookiecutter.go_module }}/internal/db"
)

type Service struct {
	app   framego.App
	db    *db.Manager
	cache *cache.Manager
}

func NewService(app framego.App) *Service {
	s := &Service{
		app: app,
	}
	s.db = db.NewManager(app)
	s.cache = cache.NewManager(app, s.db)
	return s
}
