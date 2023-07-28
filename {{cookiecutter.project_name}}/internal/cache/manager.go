package cache

import (
	"github.com/frame-go/framego"

	"{{ cookiecutter.go_module }}/internal/db"
)

type Manager struct {
	db *db.Manager
}

func NewManager(a framego.App, dbManager *db.Manager) *Manager {
	m := &Manager{
		db: dbManager,
	}
	return m
}
