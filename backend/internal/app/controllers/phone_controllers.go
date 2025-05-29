package controllers

import "github.com/Webglhost-QA-Backend/backend/internal/app/services"

type PhoneController struct {
	PhoneService  services.PhoneService
	remoteService services.RemoteService
}
