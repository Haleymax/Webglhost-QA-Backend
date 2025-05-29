package services

import (
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"github.com/Webglhost-QA-Backend/backend/internal/app/repositories"
)

type WatcherService interface {
	AddWatcher(watcher *models.Watcher) error
	DeleteWatcher(watcher *models.Watcher) error
	UpdateWatcher(watcher *models.Watcher) error
	FindAllWatchers() ([]*models.Watcher, error)
	FindOneWatcher(resourceId string) (*models.Watcher, error)
	FindByIdWatcher(watcher *models.Watcher) (*models.Watcher, error)
}

type WatcherServiceImpl struct {
	watcherRepo repositories.WatcherRepository
}

func NewWatcherService(repo repositories.WatcherRepository) *WatcherServiceImpl {
	return &WatcherServiceImpl{
		watcherRepo: repo,
	}
}

func (s *WatcherServiceImpl) AddWatcher(watcher *models.Watcher) error {
	return s.watcherRepo.Insert(watcher)
}

func (s *WatcherServiceImpl) DeleteWatcher(watcher *models.Watcher) error {
	return s.watcherRepo.Delete(watcher)
}

func (s *WatcherServiceImpl) UpdateWatcher(watcher *models.Watcher) error {
	return s.watcherRepo.Update(watcher)
}

func (s *WatcherServiceImpl) FindAllWatchers() ([]*models.Watcher, error) {
	return s.watcherRepo.FindAll()
}

func (s *WatcherServiceImpl) FindOneWatcher(resourceId string) (*models.Watcher, error) {
	filter := map[string]string{
		"resource": resourceId,
	}
	return s.watcherRepo.FindOne(filter)
}

func (s *WatcherServiceImpl) FindByIdWatcher(watcher *models.Watcher) (*models.Watcher, error) {
	return s.watcherRepo.FindByID(watcher)
}
