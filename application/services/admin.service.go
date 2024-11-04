package services

import (
	"github.com/google/uuid"

	"github.com/joaofilippe/edu-uni-srv/core/entities/admin"
)

type AdminService struct {
}

func (as *AdminService) Create(admin *adminentities.Admin) error {
	return nil
}
func (as *AdminService) FindAll() ([]*adminentities.Admin, error) {
	return []*adminentities.Admin{}, nil
}
func (as *AdminService) FindById(id uuid.UUID) (*adminentities.Admin, error) {
	return &adminentities.Admin{}, nil
}
func (as *AdminService) FindByUserID(userID uuid.UUID) (*adminentities.Admin, error) {
	return &adminentities.Admin{}, nil
}
func (as *AdminService) Update(admin *adminentities.Admin) error {
	return nil
}
func (as *AdminService) Delete(id uuid.UUID) error {
	return nil
}
