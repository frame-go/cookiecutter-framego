package cache

import (
	"github.com/frame-go/framego"

	"{{ cookiecutter.go_module }}/internal/db"
)

type Manager struct {
	db *db.Manager
}

func NewManager(app framego.App, dbManager *db.Manager) *Manager {
	return &Manager{
		db: dbManager,
	}
}
