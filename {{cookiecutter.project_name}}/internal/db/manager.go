package db

import (
	"github.com/frame-go/framego"
	"gorm.io/gorm"
)

type Manager struct {
	{{ cookiecutter.__service_name_camel }}DB *gorm.DB
}

func NewManager(m framego.DatabaseManager) *Manager {
	return &Manager{
		{{ cookiecutter.__service_name_camel }}DB: m.GetDatabaseClient("{{ cookiecutter.service_name }}"),
	}
}
