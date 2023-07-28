package db

import (
	"context"
	"time"

	"{{ cookiecutter.go_module }}/internal/models"
)

func (m *Manager) CreateObject(ctx context.Context, object *models.Object) error {
	object.CreateTime = uint32(time.Now().Unix())
	result := m.{{ cookiecutter.__service_name_camel }}DB.WithContext(ctx).Create(object)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *Manager) GetObject(ctx context.Context, id uint64) (*models.Object, error) {
	object := &models.Object{}
	result := m.{{ cookiecutter.__service_name_camel }}DB.WithContext(ctx).First(object, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return object, nil
}

func (m *Manager) GetObjectByName(ctx context.Context, name string) (*models.Object, error) {
	object := &models.Object{}
	result := m.{{ cookiecutter.__service_name_camel }}DB.WithContext(ctx).Where("`name` = ?", name).First(object)
	if result.Error != nil {
		return nil, result.Error
	}
	return object, nil
}

func (m *Manager) UpdateObject(ctx context.Context, object *models.Object) error {
	object.UpdateTime = uint32(time.Now().Unix())
	result := m.{{ cookiecutter.__service_name_camel }}DB.WithContext(ctx).Save(object)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *Manager) UpdateObjectDataByName(ctx context.Context, name string, data interface{}) error {
	updateFields := map[string]interface{}{
		"data":        data,
		"update_time": uint32(time.Now().Unix()),
	}
	result := m.{{ cookiecutter.__service_name_camel }}DB.WithContext(ctx).Model(&models.Object{}).Where("`name` = ?", name).Updates(updateFields)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *Manager) DeleteObject(ctx context.Context, id uint64) error {
	result := m.{{ cookiecutter.__service_name_camel }}DB.WithContext(ctx).Delete(&models.Object{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *Manager) DeleteObjectByName(ctx context.Context, name string) error {
	result := m.{{ cookiecutter.__service_name_camel }}DB.WithContext(ctx).Where("`name` = ?", name).Delete(&models.Object{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
