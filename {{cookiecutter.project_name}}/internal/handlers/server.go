package handlers

import (
	"context"

	"github.com/frame-go/framego"
	"github.com/frame-go/framego/health"

	"{{ cookiecutter.go_module }}/api/{{ cookiecutter.service_package_name }}"
	"{{ cookiecutter.go_module }}/internal/domain"
)

type Server interface {
	{{ cookiecutter.service_package_name }}.{{ cookiecutter.__service_name_title }}Server
	health.Checker
}

type server struct {
	{{ cookiecutter.service_package_name }}.Unimplemented{{ cookiecutter.__service_name_title }}Server

	app     framego.App
	service *domain.Service
}

func NewServer(app framego.App) Server {
	return &server{
		app:     app,
		service: domain.NewService(app),
	}
}

func (s *server) HealthCheck(ctx context.Context) error {
	return s.service.HealthCheck(ctx)
}
