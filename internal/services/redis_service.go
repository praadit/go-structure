package services

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"time"

// 	"github.com/go-redis/redis/v8"
// 	"gitlab.com/m8851/pmo-echo-api/config"
// 	"gitlab.com/m8851/pmo-echo-api/helper"
// )

// type RedisServiceInterface interface {
// 	SetCacheValue(key string, val interface{}, ttl time.Duration)
// 	SetCacheString(key string, val string, ttl time.Duration)
// 	GetCacheValue(key string) string
// 	DeleteCache(key ...string)
// }
// type RedisService struct {
// 	rdb *redis.Client
// }

// func NewRedisService(rdb *redis.Client) RedisServiceInterface {
// 	return &RedisService{
// 		rdb: rdb,
// 	}
// }

// func (service *RedisService) SetCacheValue(key string, val interface{}, ttl time.Duration) {
// 	result := service.rdb.Set(context.Background(), formatKey(key), serializeInterface(val), ttl)
// 	helper.PanicIfError(result.Err(), nil)
// }

// func (service *RedisService) SetCacheString(key string, val string, ttl time.Duration) {
// 	result := service.rdb.Set(context.Background(), formatKey(key), val, ttl)
// 	helper.PanicIfError(result.Err(), nil)
// }

// func (service *RedisService) GetCacheValue(key string) string {
// 	result := service.rdb.Get(context.Background(), formatKey(key))

// 	savedVal := helper.ValidateRedisVal(result.Val(), result.Err())
// 	return savedVal
// }

// func (service *RedisService) DeleteCache(keys ...string) {
// 	var formatedKey []string
// 	for _, key := range keys {
// 		formatedKey = append(formatedKey, formatKey(key))
// 	}
// 	result := service.rdb.Del(context.Background(), formatedKey...)
// 	helper.PanicIfError(result.Err(), nil)
// }

// func formatKey(key string) string {
// 	return fmt.Sprintf(config.AppConfig[config.RedisInitialKey] + ":" + key)
// }

// func serializeInterface(val interface{}) string {
// 	stringified, err := json.Marshal(val)
// 	helper.PanicIfError(err, nil)

// 	return string(stringified)
// }
