package services

import (
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"github.com/Webglhost-QA-Backend/backend/internal/app/repositories"
)

type PhoneInfoMap map[string]string

type PhoneService interface {
	AddPhone(phone *models.Phone) error
	UpdatePhone(phone *models.Phone) error
	FindAllPhone() ([]PhoneInfoMap, error)
	FindOnePhone(serial string) (models.Phone, error)
}

type PhoneServiceImpl struct {
	phoneRepo repositories.PhoneRepository
}

func NewPhoneService(repo repositories.PhoneRepository) PhoneServiceImpl {
	return PhoneServiceImpl{
		phoneRepo: repo,
	}
}

func (p PhoneServiceImpl) AddPhone(phone *models.Phone) error {
	return p.phoneRepo.InsertOne(phone)
}

func (p PhoneServiceImpl) UpdatePhone(phone *models.Phone) error {
	return p.phoneRepo.Update(phone)
}

func (p PhoneServiceImpl) FindOnePhone(serial string) (models.Phone, error) {
	phone, err := p.phoneRepo.FindOne(serial)
	return phone, err
}

func (p PhoneServiceImpl) FindAllPhone() ([]PhoneInfoMap, error) {
	phones, err := p.phoneRepo.FindAll()
	if err != nil {
		return nil, err
	}
	phoneMaps := make([]PhoneInfoMap, len(phones))
	for i, phone := range phones {
		phoneMaps[i] = PhoneInfoMap{
			"serial":           phone.Serial,
			"manufacturer":     phone.Manufacturer,
			"model":            phone.Model,
			"androidVersion":   phone.AndroidVersion,
			"cpuabi":           phone.Cpuabi,
			"marketName":       phone.MarketName,
			"marketNameSymbol": phone.MarketNameSymbol,
		}
	}
	return phoneMaps, nil
}
