package utils

import (
	"context"
	"encoding/json"
	"time"
	"venturan/global"
)

func SetRedisCache(key string, value map[string]interface{}, expiration time.Duration) error {
	val, err := json.Marshal(value)
	if err != nil {
		return err
	}
	errors := global.App.Redis.SetNX(context.Background(), key, val, expiration).Err()
	if errors != nil {
		return errors
	}
	return nil
}
