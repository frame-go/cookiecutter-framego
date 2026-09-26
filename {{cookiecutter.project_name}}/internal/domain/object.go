package domain

import (
	"context"
	"time"

	"github.com/frame-go/framego/errors"
	"github.com/frame-go/framego/uniqueid"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"

	"{{ cookiecutter.go_module }}/api/{{ cookiecutter.service_package_name }}"
	"{{ cookiecutter.go_module }}/internal/models"
)

const defaultObjectPageSize = 20

func (s *Service) GetObject(ctx context.Context, req *{{ cookiecutter.service_package_name }}.GetObjectRequest) (*{{ cookiecutter.service_package_name }}.Object, error) {
	id, err := uniqueid.ParseID(req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "invalid_object_id").WithGRPCCode(codes.InvalidArgument)
	}
	object, err := s.db.GetObject(ctx, id)
	if err != nil {
		return nil, wrapDBNotFoundError(err, "object_not_found")
	}
	return objectToProto(object), nil
}

func (s *Service) ListObjects(ctx context.Context, req *{{ cookiecutter.service_package_name }}.ListObjectsRequest) ([]*{{ cookiecutter.service_package_name }}.Object, int32, error) {
	limit := int(req.Limit)
	if limit <= 0 {
		limit = defaultObjectPageSize
	}
	rows, total, err := s.db.ListObjects(ctx, int(req.Offset), limit)
	if err != nil {
		return nil, 0, errors.Wrap(err, "list_objects_error").WithGRPCCode(codes.Internal)
	}
	objects := make([]*{{ cookiecutter.service_package_name }}.Object, 0, len(rows))
	for _, row := range rows {
		objects = append(objects, objectToProto(row))
	}
	return objects, int32(total), nil
}

func (s *Service) CreateObject(ctx context.Context, req *{{ cookiecutter.service_package_name }}.CreateObjectRequest) (*{{ cookiecutter.service_package_name }}.Object, error) {
	now := time.Now()
	object := &models.Object{
		Id:          s.idGenerator.NewID(),
		Name:        req.Name,
		Description: req.Description,
		Status:      int16({{ cookiecutter.service_package_name }}.ObjectStatus_OBJECT_STATUS_ACTIVE),
		CreateTime:  now,
		UpdateTime:  now,
	}
	if err := s.db.CreateObject(ctx, object); err != nil {
		return nil, errors.Wrap(err, "create_object_error").WithGRPCCode(codes.Internal)
	}
	return objectToProto(object), nil
}

func (s *Service) UpdateObject(ctx context.Context, req *{{ cookiecutter.service_package_name }}.UpdateObjectRequest) (*{{ cookiecutter.service_package_name }}.Object, error) {
	id, err := uniqueid.ParseID(req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "invalid_object_id").WithGRPCCode(codes.InvalidArgument)
	}
	object, err := s.db.GetObject(ctx, id)
	if err != nil {
		return nil, wrapDBNotFoundError(err, "object_not_found")
	}
	object.Description = req.Description
	object.UpdateTime = time.Now()
	if err := s.db.UpdateObject(ctx, object); err != nil {
		return nil, errors.Wrap(err, "update_object_error").WithGRPCCode(codes.Internal)
	}
	return objectToProto(object), nil
}

func (s *Service) DeleteObject(ctx context.Context, req *{{ cookiecutter.service_package_name }}.DeleteObjectRequest) error {
	id, err := uniqueid.ParseID(req.Id)
	if err != nil {
		return errors.Wrap(err, "invalid_object_id").WithGRPCCode(codes.InvalidArgument)
	}
	if _, err := s.db.GetObject(ctx, id); err != nil {
		return wrapDBNotFoundError(err, "object_not_found")
	}
	if err := s.db.UpdateObjectStatus(ctx, id, int16({{ cookiecutter.service_package_name }}.ObjectStatus_OBJECT_STATUS_DELETED)); err != nil {
		return errors.Wrap(err, "delete_object_error").WithGRPCCode(codes.Internal)
	}
	return nil
}

func objectToProto(m *models.Object) *{{ cookiecutter.service_package_name }}.Object {
	return &{{ cookiecutter.service_package_name }}.Object{
		Id:          m.Id.String(),
		Name:        m.Name,
		Description: m.Description,
		Status:      int32(m.Status),
		CreateTime:  timestamppb.New(m.CreateTime),
		UpdateTime:  timestamppb.New(m.UpdateTime),
	}
}
