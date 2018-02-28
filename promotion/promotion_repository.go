package promotion

import (
	"github.com/at-vudang95/go-food-market-api/shared/repository"
	"github.com/jinzhu/gorm"
)

// RepositoryInterface interface
type RepositoryInterface interface {
}

// Repository struct.
type Repository struct {
	repository.BaseRepository
	// connect master database.
	MasterDB *gorm.DB
}

// NewRepository responses new Repository instance.
func NewRepository(br *repository.BaseRepository, master *gorm.DB) *Repository {
	return &Repository{BaseRepository: *br, MasterDB: master}
}
