package category

import (
	"github.com/at-vudang95/go-food-market-api/shared/usecase"
	"github.com/jinzhu/gorm"
)

// UsecaseInterface interface.
type UsecaseInterface interface {
}

// Usecase struct.
type Usecase struct {
	usecase.BaseUsecase
	db         *gorm.DB
	repository RepositoryInterface
}

// NewUsecase responses new Usecase instance.
func NewUsecase(bu *usecase.BaseUsecase, master *gorm.DB, r RepositoryInterface) *Usecase {
	return &Usecase{BaseUsecase: *bu, db: master, repository: r}
}
