package util

import (
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"github.com/Webglhost-QA-Backend/backend/pkg/cache"
	"log"
)

func Get_all_brand_from_watcher(watchers []*models.Watcher) []string {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("error：%v\n", r)
		}
	}()
	var brands []string
	var brands_map = make(map[string]bool)
	for _, watcher := range watchers {
		for _, bds := range watcher.Brand {
			brands_map[bds] = true
		}
	}
	for key := range brands_map {
		brands = append(brands, key)
	}

	return brands
}

func Convert(watchers []models.Watcher, brand, tag string) []cache.WatcherCache {
	/*
		将数据库中的转换为方便插入到redis中的格式
	*/
	result := make([]cache.WatcherCache, len(watchers))
	for _, watcher := range watchers {
		data := new(cache.WatcherCache)
		data.Brand = brand
		data.Name = watcher.Name
		data.Resource = watcher.Resource
		data.Event = watcher.Click
		data.Tag = tag
		result = append(result, *data)
	}
	return result
}
