package db

import (
	"context"
	"time"

	"github.com/frame-go/framego/uniqueid"

	"{{ cookiecutter.go_module }}/api/{{ cookiecutter.service_package_name }}"
	"{{ cookiecutter.go_module }}/internal/models"
)

func (m *Manager) GetObject(ctx context.Context, id uniqueid.ID) (*models.Object, error) {
	object := &models.Object{}
	err := m.db.WithContext(ctx).
		Where("id = ? AND status <> ?", id, int16({{ cookiecutter.service_package_name }}.ObjectStatus_OBJECT_STATUS_DELETED)).
		First(object).Error
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (m *Manager) ListObjects(ctx context.Context, offset int, limit int) ([]*models.Object, int64, error) {
	query := m.db.WithContext(ctx).Model(&models.Object{}).
		Where("status <> ?", int16({{ cookiecutter.service_package_name }}.ObjectStatus_OBJECT_STATUS_DELETED))
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	query = query.Order("create_time DESC")
	if limit > 0 {
		query = query.Offset(offset).Limit(limit)
	}
	var objects []*models.Object
	if err := query.Find(&objects).Error; err != nil {
		return nil, 0, err
	}
	return objects, total, nil
}

func (m *Manager) CreateObject(ctx context.Context, object *models.Object) error {
	return m.db.WithContext(ctx).Create(object).Error
}

func (m *Manager) UpdateObject(ctx context.Context, object *models.Object) error {
	return m.db.WithContext(ctx).Save(object).Error
}

func (m *Manager) UpdateObjectStatus(ctx context.Context, id uniqueid.ID, status int16) error {
	return m.db.WithContext(ctx).Model(&models.Object{}).
		Where("id = ?", id).
		Updates(map[string]any{"status": status, "update_time": time.Now()}).Error
}
