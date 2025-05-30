package services

import (
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"github.com/Webglhost-QA-Backend/backend/internal/app/repositories"
	"github.com/Webglhost-QA-Backend/backend/pkg/cache_client"
)

type WatcherService interface {
	AddWatcher(watcher *models.Watcher) error
	DeleteWatcher(watcher *models.Watcher) error
	UpdateWatcher(watcher *models.Watcher) error
	FindAllWatchers() ([]*models.Watcher, error)
	FindOneWatcher(resourceId string) (*models.Watcher, error)
	FindByIdWatcher(watcher *models.Watcher) (*models.Watcher, error)
	RefreshCache(env, runtime string) error
}

type WatcherServiceImpl struct {
	watcherRepo repositories.WatcherRepository
	redis       *cache_client.Redis
}

func NewWatcherService(repo repositories.WatcherRepository, cache *cache_client.Redis) *WatcherServiceImpl {
	return &WatcherServiceImpl{
		watcherRepo: repo,
		redis:       cache,
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

func (s *WatcherServiceImpl) RefreshCache(env, runtime string) error {
	watchers, err := s.FindAllWatchers()
	if err != nil {
		return err
	}
	space := "watcher"
	brands := cache_client.Get_all_brand_from_watcher(watchers)
	s.redis.ClearKey(space, env, runtime)
	for _, brand := range brands {
		permissionFilter := map[string]string{
			"brand": brand,
			"tag":   "PERMISSION",
		}
		permissionWatchers, err2 := s.watcherRepo.FindByCategory(permissionFilter)
		if err2 != nil {
			return err2
		}
		permissionDatas := cache_client.Convert(permissionWatchers, brand, permissionFilter["tag"])
		for _, oneData := range permissionDatas {
			s.redis.SetKey(space, env, runtime, cache_client.WatcherCache(oneData))
		}

		installFilter := map[string]string{
			"brand": brand,
			"tag":   "PERMISSION",
		}
		installWatchers, err3 := s.watcherRepo.FindByCategory(installFilter)
		if err3 != nil {
			return err2
		}
		installDatas := cache_client.Convert(installWatchers, brand, installFilter["tag"])
		for _, oneData := range installDatas {
			s.redis.SetKey(space, env, runtime, cache_client.WatcherCache(oneData))
		}
	}

	for _, brand := range brands {
		permissionFilter := map[string]string{
			"brand": brand,
			"tag":   "INSTALL",
		}
		permissionWatchers, err2 := s.watcherRepo.FindByCategory(permissionFilter)
		if err2 != nil {
			return err2
		}
		permissionDatas := cache_client.Convert(permissionWatchers, brand, permissionFilter["tag"])
		for _, oneData := range permissionDatas {
			s.redis.SetKey(space, env, runtime, cache_client.WatcherCache(oneData))
		}

		installFilter := map[string]string{
			"brand": brand,
			"tag":   "INSTALL",
		}
		installWatchers, err3 := s.watcherRepo.FindByCategory(installFilter)
		if err3 != nil {
			return err2
		}
		installDatas := cache_client.Convert(installWatchers, brand, installFilter["tag"])
		for _, oneData := range installDatas {
			s.redis.SetKey(space, env, runtime, cache_client.WatcherCache(oneData))
		}
	}
	return nil
}
