package handlers

import (
	"context"

	"github.com/frame-go/framego"
	"github.com/frame-go/framego/errors"
	"github.com/frame-go/framego/health"
	"github.com/frame-go/framego/log"

	"{{ cookiecutter.go_module }}/api/{{ cookiecutter.service_package_name }}"
	"{{ cookiecutter.go_module }}/internal/domain"
)

type server struct {
	{{ cookiecutter.service_package_name }}.Unimplemented{{ cookiecutter.__service_name_title }}Server
	health.Checker

	app     framego.App
	service *domain.Service
}

func (s *server) HealthCheck(ctx context.Context) error {
	return nil
}

func (s *server) CreateObject(ctx context.Context, req *{{ cookiecutter.service_package_name }}.CreateObjectRequest) (*{{ cookiecutter.service_package_name }}.CreateObjectResponse, error) {
	object, err := s.service.CreateObject(ctx, req.Object)
	if err != nil {
		errors.LogError(log.FromContext(ctx).Warn(), err).Interface("object", req.Object).Msg("create_object_error")
		return nil, err
	}
	resp := &{{ cookiecutter.service_package_name }}.CreateObjectResponse{
		Object: object,
	}
	log.FromContext(ctx).Trace().Interface("object", object).Msg("create_object")
	return resp, nil
}

func (s *server) GetObject(ctx context.Context, req *{{ cookiecutter.service_package_name }}.GetObjectRequest) (*{{ cookiecutter.service_package_name }}.GetObjectResponse, error) {
	object, err := s.service.GetObjectByName(ctx, req.Name)
	if err != nil {
		errors.LogError(log.FromContext(ctx).Warn(), err).Str("name", req.Name).Msg("get_object_error")
		return nil, err
	}
	resp := &{{ cookiecutter.service_package_name }}.GetObjectResponse{
		Object: object,
	}
	return resp, nil
}

func (s *server) UpdateObject(ctx context.Context, req *{{ cookiecutter.service_package_name }}.UpdateObjectRequest) (*{{ cookiecutter.service_package_name }}.UpdateObjectResponse, error) {
	object, err := s.service.UpdateObjectDataByName(ctx, req.Object.Name, req.Object.Data)
	if err != nil {
		errors.LogError(log.FromContext(ctx).Warn(), err).Interface("object", req.Object).Msg("update_object_error")
		return nil, err
	}
	resp := &{{ cookiecutter.service_package_name }}.UpdateObjectResponse{
		Object: object,
	}
	log.FromContext(ctx).Trace().Interface("object", req.Object).Msg("update_object")
	return resp, nil
}

func (s *server) DeleteObject(ctx context.Context, req *{{ cookiecutter.service_package_name }}.DeleteObjectRequest) (*{{ cookiecutter.service_package_name }}.DeleteObjectResponse, error) {
	err := s.service.DeleteObject(ctx, req.Name)
	if err != nil {
		errors.LogError(log.FromContext(ctx).Warn(), err).Str("name", req.Name).Msg("delete_object_error")
		return nil, err
	}
	resp := &{{ cookiecutter.service_package_name }}.DeleteObjectResponse{}
	log.FromContext(ctx).Trace().Interface("name", req.Name).Msg("delete_object")
	return resp, nil
}

func NewServer(app framego.App) {{ cookiecutter.service_package_name }}.{{ cookiecutter.__service_name_title }}Server {
	return &server{
		app:     app,
		service: domain.NewService(app),
	}
}
