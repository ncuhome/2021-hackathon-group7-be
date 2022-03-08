package dao

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"reflect"
	"time"
)

type JsonCache struct {
	Data   interface{} //必须传引用
	ID     string
	Prefix string //ID的描述，如果为空则默认为Data的类型名+":"
}

// 需提供ID，Data
func (s JsonCache) SetData(expiration time.Duration) error {
	if s.Prefix == "" {
		s.Prefix = reflect.TypeOf(s.Data).Elem().String()
	}
	key := CacheConfigObj.Prefix + s.Prefix + s.ID
	DataBytes, err := json.Marshal(s.Data)
	if err != nil {
		return err
	}

	return Cache.Set(key, string(DataBytes), expiration).Err()
}

// 需提供ID及data模型，写入Data
func (s JsonCache) GetData() error {
	if s.Prefix == "" {
		s.Prefix = reflect.TypeOf(s.Data).Elem().String()
	}
	key := CacheConfigObj.Prefix + s.Prefix + s.ID

	value, err := Cache.Get(key).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(value), s.Data)
	if err != nil {
		return err
	}

	return nil
}

// 需提供ID及data模型，写入Data，并更新
func (s JsonCache) GetDataEpr(expiration time.Duration) error {
	if s.Prefix == "" {
		s.Prefix = reflect.TypeOf(s.Data).Elem().String()
	}
	key := CacheConfigObj.Prefix + s.Prefix + s.ID

	value, err := Cache.Get(key).Result()
	if err != nil {
		return err
	}

	Cache.PExpire(key, expiration)

	err = json.Unmarshal([]byte(value), s.Data)
	if err != nil {
		return err
	}

	return nil
}

// 需提供ID及data模型，删除Data
func (s JsonCache) DelData() error {
	if s.Prefix == "" {
		s.Prefix = reflect.TypeOf(s.Data).Elem().String()
	}
	key := CacheConfigObj.Prefix + s.Prefix + s.ID

	err := Cache.Del(key).Err()
	if err != nil && err != redis.Nil {
		return err
	}

	return nil
}
