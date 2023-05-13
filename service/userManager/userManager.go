package usermanager

import (
	"github.com/Zhima-Mochi/go-user-service/external/cache"
	"github.com/Zhima-Mochi/go-user-service/external/storage"
)

type userManager struct {
	cache   cache.Cache
	storage storage.Storage
}

func NewUserManager(cache cache.Cache, storage storage.Storage) *userManager {
	return &userManager{
		cache:   cache,
		storage: storage,
	}
}
