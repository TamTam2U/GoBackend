package agencyRepository

import (
	"github.com/dev-newus/GoAlinBackend/src/models/entities"
	"gorm.io/gorm"
)

type AgencyRepo struct {
	db *gorm.DB
}

func NewAgencyRepo(db *gorm.DB) *AgencyRepo {
	return &AgencyRepo{
		db: db,
	}
}

func (repo *AgencyRepo) Create(agency entities.Agency) (entities.Agency, error) {
	result := repo.db.Create(&agency)
	if result.Error != nil {
		return agency, result.Error
	}

	return agency, nil
}
