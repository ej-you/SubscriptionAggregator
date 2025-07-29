package pg

import (
	goerrors "errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"SubscriptionAggregator/internal/app/entity"
	"SubscriptionAggregator/internal/app/repo"
)

var _ repo.ServiceRepoDB = (*serviceRepoPG)(nil)

// ServiceRepoDB implementation.
type serviceRepoPG struct {
	dbStorage *gorm.DB
}

// NewServiceRepoDB returns new ServiceRepoDB instance.
func NewServiceRepoDB(dbStorage *gorm.DB) repo.ServiceRepoDB {
	return &serviceRepoPG{
		dbStorage: dbStorage,
	}
}

// GetByNameOrCreate returns service with given service name if it exists.
// Else it creates new service with given ID and returns it.
func (r *serviceRepoPG) GetByNameOrCreate(service *entity.Service) error {
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
	// err := r.dbStorage.Debug().Where("name = ?", service.Name).
	// 	FirstOrCreate(service).Error
	// if err != nil {
	// 	return fmt.Errorf("get or create service: %w", err)
	// }
	// return nil
}

// create creates new service.
// All necessary fields must be presented. ID will be generated.
func (r *serviceRepoPG) create(service *entity.Service) error {
	service.ID = uuid.NewString()
	if err := r.dbStorage.Create(service).Error; err != nil {
		return fmt.Errorf("create: %w", err)
	}
	return nil
}
