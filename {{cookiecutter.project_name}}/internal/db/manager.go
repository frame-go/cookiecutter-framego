package db

import (
	"context"

	"github.com/frame-go/framego"
	"github.com/frame-go/framego/log"
	"github.com/frame-go/framego/uniqueid"
	"gorm.io/gorm"
)

type Manager struct {
	db          *gorm.DB
	idGenerator uniqueid.Generator
}

func NewManager(m framego.DatabaseManager, idGenerator uniqueid.Generator) *Manager {
	db := m.GetDatabaseClient("{{ cookiecutter.service_name }}")
	if db == nil {
		log.Logger.Fatal().Msg("database_client_not_configured")
	}
	return &Manager{
		db:          db,
		idGenerator: idGenerator,
	}
}

func (m *Manager) Ping(ctx context.Context) error {
	sqlDB, err := m.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.PingContext(ctx)
}
