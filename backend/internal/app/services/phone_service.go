package services

import (
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"github.com/Webglhost-QA-Backend/backend/internal/app/repositories"
)

type PhoneService interface {
	AddPhone(phone *models.Phone) error
	UpdatePhone(phone *models.Phone) error
	FindAllPhone() ([]models.Phone, error)
	FindOnePhone(serial string) (*models.Phone, error)
}

type PhoneServiceImpl struct {
	phoneRepo repositories.PhoneRepository
}

func NewPhoneService(repo repositories.PhoneRepository) *PhoneServiceImpl {
	return &PhoneServiceImpl{
		phoneRepo: repo,
	}
}

func (p PhoneServiceImpl) AddPhone(phone *models.Phone) error {
	return p.phoneRepo.InsertOne(phone)
}

func (p PhoneServiceImpl) UpdatePhone(phone *models.Phone) error {
	return p.phoneRepo.Update(phone)
}

func (p PhoneServiceImpl) FindAllPhone() ([]models.Phone, error) {
	return p.phoneRepo.FindAll()
}

func (p PhoneServiceImpl) FindOnePhone(serial string) (models.Phone, error) {
	phone, err := p.phoneRepo.FindOne(serial)
	return phone, err
}
