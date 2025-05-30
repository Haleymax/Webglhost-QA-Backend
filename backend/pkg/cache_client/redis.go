package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Webglhost-QA-Backend/backend/config"
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type WatcherCache struct {
	Brand    string
	Name     string
	Resource string
	Event    string
	Tag      string
}

type Redis struct {
	Cfg    config.RedisConfig
	Client *redis.Client
}

func (r *Redis) Connect() {
	addr := fmt.Sprintf("%s:%d", r.Cfg.HOST, r.Cfg.PORT)
	redisClient := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   r.Cfg.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := redisClient.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}
	r.Client = redisClient

}

func (r *Redis) ClearKey(space, env, runtime string) error {
	if r.Client == nil {
		log.Println("redis dont connect")
		return fmt.Errorf("redis dont connect")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	pattern := fmt.Sprintf("%s:%s:%s:*", env, runtime, space)
	iter := r.Client.Scan(ctx, 0, pattern, 100).Iterator()

	var keys []string
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}

	if err := iter.Err(); err != nil {
		return err
	}

	if len(keys) == 0 {
		return nil
	}

	for i := 0; i < len(keys); i++ {
		_, err := r.Client.Del(ctx, keys[i]).Result()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Redis) SetKey(space, env, runtime string, watcher WatcherCache) error {
	if r.Client == nil {
		log.Println("redis dont connect")
		return fmt.Errorf("redis dont connect")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	pattern := fmt.Sprintf("%s:%s:%s:%s", env, runtime, space, watcher.Brand, watcher.Tag)
	value := map[string]string{
		"name":     watcher.Name,
		"resource": watcher.Resource,
		"event":    watcher.Event,
	}
	value_json_str, _ := json.Marshal(value)
	_, err := r.Client.Set(ctx, pattern, value_json_str, 0).Result()
	return err
}

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

func Convert(watchers []models.Watcher, brand, tag string) []WatcherCache {
	/*
		将数据库中的转换为方便插入到redis中的格式
	*/
	result := make([]WatcherCache, len(watchers))
	for _, watcher := range watchers {
		data := new(WatcherCache)
		data.Brand = brand
		data.Name = watcher.Name
		data.Resource = watcher.Resource
		data.Event = watcher.Click
		data.Tag = tag
		result = append(result, *data)
	}
	return result
}
