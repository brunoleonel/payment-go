package repositories

import (
	"github.com/jinzhu/gorm"
)

//OperationRepository handles the operations related to payments on the database
type OperationRepository interface {
}

type operationRepository struct {
	db *gorm.DB
}

//NewOperationRepository returns a new OperationRepository instance
func NewOperationRepository(db *gorm.DB) OperationRepository {
	return &operationRepository{
		db: db,
	}
}
