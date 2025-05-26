package repositories

import "github.com/Webglhost-QA-Backend/backend/internal/app/models"

type watcherRepository interface {
	Insert(watcher *models.Watcher) error
	Update(watcher *models.Watcher) error
}
