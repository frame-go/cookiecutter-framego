package handlers

import (
	"context"

	"github.com/frame-go/framego/errors"
	"github.com/frame-go/framego/log"

	"{{ cookiecutter.go_module }}/api/{{ cookiecutter.service_package_name }}"
)

func (s *server) GetObject(ctx context.Context, req *{{ cookiecutter.service_package_name }}.GetObjectRequest) (*{{ cookiecutter.service_package_name }}.GetObjectResponse, error) {
	object, err := s.service.GetObject(ctx, req)
	if err != nil {
		errors.LogError(log.FromContext(ctx).Warn(), err).Interface("request", req).Msg("get_object_error")
		return nil, err
	}
	log.FromContext(ctx).Info().Str("object_id", object.Id).Msg("get_object")
	return &{{ cookiecutter.service_package_name }}.GetObjectResponse{Object: object}, nil
}

func (s *server) ListObjects(ctx context.Context, req *{{ cookiecutter.service_package_name }}.ListObjectsRequest) (*{{ cookiecutter.service_package_name }}.ListObjectsResponse, error) {
	objects, total, err := s.service.ListObjects(ctx, req)
	if err != nil {
		errors.LogError(log.FromContext(ctx).Warn(), err).Interface("request", req).Msg("list_objects_error")
		return nil, err
	}
	log.FromContext(ctx).Info().Int("count", len(objects)).Int32("total", total).Msg("list_objects")
	return &{{ cookiecutter.service_package_name }}.ListObjectsResponse{Objects: objects, Total: total}, nil
}

func (s *server) CreateObject(ctx context.Context, req *{{ cookiecutter.service_package_name }}.CreateObjectRequest) (*{{ cookiecutter.service_package_name }}.CreateObjectResponse, error) {
	ctx = context.WithoutCancel(ctx)
	object, err := s.service.CreateObject(ctx, req)
	if err != nil {
		errors.LogError(log.FromContext(ctx).Warn(), err).Interface("request", req).Msg("create_object_error")
		return nil, err
	}
	log.FromContext(ctx).Info().Str("object_id", object.Id).Msg("create_object")
	return &{{ cookiecutter.service_package_name }}.CreateObjectResponse{Object: object}, nil
}

func (s *server) UpdateObject(ctx context.Context, req *{{ cookiecutter.service_package_name }}.UpdateObjectRequest) (*{{ cookiecutter.service_package_name }}.UpdateObjectResponse, error) {
	ctx = context.WithoutCancel(ctx)
	object, err := s.service.UpdateObject(ctx, req)
	if err != nil {
		errors.LogError(log.FromContext(ctx).Warn(), err).Interface("request", req).Msg("update_object_error")
		return nil, err
	}
	log.FromContext(ctx).Info().Str("object_id", object.Id).Msg("update_object")
	return &{{ cookiecutter.service_package_name }}.UpdateObjectResponse{Object: object}, nil
}

func (s *server) DeleteObject(ctx context.Context, req *{{ cookiecutter.service_package_name }}.DeleteObjectRequest) (*{{ cookiecutter.service_package_name }}.DeleteObjectResponse, error) {
	ctx = context.WithoutCancel(ctx)
	err := s.service.DeleteObject(ctx, req)
	if err != nil {
		errors.LogError(log.FromContext(ctx).Warn(), err).Interface("request", req).Msg("delete_object_error")
		return nil, err
	}
	log.FromContext(ctx).Info().Str("object_id", req.Id).Msg("delete_object")
	return &{{ cookiecutter.service_package_name }}.DeleteObjectResponse{}, nil
}
