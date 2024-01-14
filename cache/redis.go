package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewCache(addr, password string) Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
	return Cache{
		Client: rdb,
	}
}

type Cache struct {
	Client *redis.Client
}

func (c *Cache) SetCache(id, volume string) {
	ctx := context.Background()
	id = fmt.Sprintf("guild-id:%s", id)
	c.Client.HSet(ctx, id, "volume", volume).Err()
	c.Client.Expire(ctx, id, 5*time.Minute).Err()

}
func (c *Cache) GetCache(id string) map[string]string {
	ctx := context.Background()
	id = fmt.Sprintf("guild-id:%s", id)
	return c.Client.HGetAll(ctx, id).Val()
}
