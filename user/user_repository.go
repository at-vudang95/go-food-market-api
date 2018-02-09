package user

import (
	"github.com/at-vudang95/go-food-market-api/shared/repository"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// RepositoryInterface interface.
type RepositoryInterface interface {
	FindOrCreate(string) (User, error)
}

// NewRepository responses new Repository instance.
func NewRepository(br *repository.BaseRepository, master *gorm.DB, read *gorm.DB, redis *redis.Conn) *Repository {
	return &Repository{BaseRepository: *br, masterDB: master, readDB: read, redis: redis}
}
