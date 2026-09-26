package domain

import (
	"context"

	"github.com/frame-go/framego"
	"github.com/frame-go/framego/errors"
	"github.com/frame-go/framego/log"
	"github.com/frame-go/framego/uniqueid"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"

	"{{ cookiecutter.go_module }}/internal/cache"
	"{{ cookiecutter.go_module }}/internal/db"
)

type Service struct {
	app         framego.App
	idGenerator uniqueid.Generator
	db          *db.Manager
	cache       *cache.Manager
}

func NewService(app framego.App) *Service {
	s := &Service{
		app: app,
	}
	s.idGenerator = app.GetIDGenerator()
	if s.idGenerator == nil {
		log.Logger.Fatal().Msg("id_generator_not_configured")
	}
	s.db = db.NewManager(app, s.idGenerator)
	s.cache = cache.NewManager(app, s.db)
	return s
}

func (s *Service) HealthCheck(ctx context.Context) error {
	return s.db.Ping(ctx)
}

func wrapDBNotFoundError(err error, notFoundCode string) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.Wrap(err, notFoundCode).WithGRPCCode(codes.NotFound)
	}
	return errors.Wrap(err, "db_error").WithGRPCCode(codes.Internal)
}
