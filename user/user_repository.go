package user

import (
	"github.com/at-vudang95/go-food-market-api/shared/repository"

	"github.com/jinzhu/gorm"
)

// RepositoryInterface interface.
type RepositoryInterface interface {
}

// Repository struct.
type Repository struct {
	repository.BaseRepository
	// connect master database.
	masterDB *gorm.DB
	// connect read replica database.
	readDB *gorm.DB
	// redis connect Redis.
	// redis *redis.Conn
}

// NewRepository responses new Repository instance.
func NewRepository(br *repository.BaseRepository, master *gorm.DB, read *gorm.DB) *Repository {
	return &Repository{BaseRepository: *br, masterDB: master, readDB: read}
}
