package pg

import (
	goerrors "errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"SubscriptionAggregator/internal/app/entity"
)

type ServiceRepoPG struct {
	dbStorage *gorm.DB
}

// NewServiceRepoDB returns new ServiceRepoDB instance.
func NewServiceRepoDB(dbStorage *gorm.DB) *ServiceRepoPG {
	return &ServiceRepoPG{
		dbStorage: dbStorage,
	}
}

// GetByNameOrCreate returns service with given service name if it exists.
// Else it creates new service with given ID and returns it.
func (r *ServiceRepoPG) GetByNameOrCreate(service *entity.Service) error {
	err := r.dbStorage.Where("name = ?", service.Name).First(service).Error
	// if record not found
	if goerrors.Is(err, gorm.ErrRecordNotFound) {
		// create new service record
		return r.create(service)
	}
	if err != nil {
		return fmt.Errorf("get by name: %w", err)
	}
	return nil
}

// create creates new service.
// All necessary fields must be presented. ID will be generated.
func (r *ServiceRepoPG) create(service *entity.Service) error {
	service.ID = uuid.NewString()
	if err := r.dbStorage.Create(service).Error; err != nil {
		return fmt.Errorf("create: %w", err)
	}
	return nil
}
