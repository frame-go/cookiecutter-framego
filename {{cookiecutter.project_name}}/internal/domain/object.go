package domain

import (
	"context"

	"github.com/frame-go/framego/errors"
	"google.golang.org/grpc/codes"

	"{{ cookiecutter.go_module }}/api/{{ cookiecutter.service_package_name }}"
	"{{ cookiecutter.go_module }}/internal/models"
)

func (s *Service) CreateObject(ctx context.Context, object *{{ cookiecutter.service_package_name }}.Object) (*{{ cookiecutter.service_package_name }}.Object, error) {
	model := &models.Object{
		Name: object.Name,
		Data: object.Data,
	}
	err := s.db.CreateObject(ctx, model)
	if err != nil {
		return nil, errors.Wrap(err, "query_db_error").WithGRPCCode(codes.Internal).With("message", err.Error())
	}
	return s.GetObjectByName(ctx, object.Name)
}

func (s *Service) GetObjectByName(ctx context.Context, name string) (*{{ cookiecutter.service_package_name }}.Object, error) {
	model, err := s.db.GetObjectByName(ctx, name)
	if err != nil {
		return nil, errors.Wrap(err, "object_not_found").WithGRPCCode(codes.NotFound).With("name", name)
	}
	object := &{{ cookiecutter.service_package_name }}.Object{
		Name: model.Name,
		Data: model.Data,
	}
	return object, nil
}

func (s *Service) UpdateObjectDataByName(ctx context.Context, name string, data string) (*{{ cookiecutter.service_package_name }}.Object, error) {
	err := s.db.UpdateObjectDataByName(ctx, name, data)
	if err != nil {
		return nil, errors.Wrap(err, "query_db_error").WithGRPCCode(codes.Internal).With("message", err.Error())
	}
	return s.GetObjectByName(ctx, name)
}

func (s *Service) DeleteObject(ctx context.Context, name string) error {
	err := s.db.DeleteObjectByName(ctx, name)
	if err != nil {
		return errors.Wrap(err, "query_db_error").WithGRPCCode(codes.Internal).With("name", name)
	}
	return nil
}
