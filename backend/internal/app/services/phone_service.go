package services

import (
	"errors"
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"github.com/Webglhost-QA-Backend/backend/internal/app/repositories"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type PhoneInfoMap map[string]string

type PhoneService interface {
	AddPhone(phone models.Phone) (bool, string)
	UpdatePhone(phone *models.Phone) error
	FindAllPhone() ([]PhoneInfoMap, error)
	FindOnePhone(serial string) (models.Phone, error)
	DeletePhone(serial string) error
}

type PhoneServiceImpl struct {
	phoneRepo repositories.PhoneRepository
}

func NewPhoneService(repo repositories.PhoneRepository) PhoneServiceImpl {
	return PhoneServiceImpl{
		phoneRepo: repo,
	}
}

func (p PhoneServiceImpl) AddPhone(phone models.Phone) (bool, string) {
	exis, err := p.FindOnePhone(phone.Serial)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("phone not exists")
		} else {
			log.Println(err)
			return false, err.Error()
		}
	}

	var empty models.Phone
	if exis != empty {
		return false, "phone is exist"
	}

	phone.CreateTime = time.Now()
	phone.UpdateTime = time.Now()
	err = p.phoneRepo.InsertOne(phone)
	if err != nil {
		log.Println(err)
		return false, err.Error()
	}
	return true, "success add phone"
}

func (p PhoneServiceImpl) UpdatePhone(phone *models.Phone) error {
	phone.UpdateTime = time.Now()
	return p.phoneRepo.Update(phone)
}

func (p PhoneServiceImpl) FindOnePhone(serial string) (models.Phone, error) {
	phone, err := p.phoneRepo.FindOne(serial)
	return phone, err
}

func (p PhoneServiceImpl) DeletePhone(serial string) error {
	phone, err := p.phoneRepo.FindOne(serial)
	if err != nil {
		log.Printf("delete phone err: %v", err)
		return err
	}
	if err = p.phoneRepo.Delete(phone); err != nil {
		log.Printf("delete phone err: %v", err)
		return err
	}
	return nil
}

func (p PhoneServiceImpl) FindAllPhone() ([]PhoneInfoMap, error) {
	phones, err := p.phoneRepo.FindAll()
	if err != nil {
		return nil, err
	}
	phoneMaps := make([]PhoneInfoMap, len(phones))
	for i, phone := range phones {
		phoneMaps[i] = PhoneInfoMap{
			"id":               phone.ID,
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
