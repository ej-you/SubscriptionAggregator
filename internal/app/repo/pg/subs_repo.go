// Package pg contains PostgreSQL DB repos implementations for entities.
package pg

import (
	goerrors "errors"
	"fmt"

	"gorm.io/gorm"

	"SubscriptionAggregator/internal/app/entity"
	"SubscriptionAggregator/internal/app/errors"
	"SubscriptionAggregator/internal/app/repo"
)

var _ repo.SubsRepoDB = (*subsRepoPG)(nil)

// SubsRepoDB implementation.
type subsRepoPG struct {
	dbStorage *gorm.DB
}

func NewSubsRepoDB(dbStorage *gorm.DB) repo.SubsRepoDB {
	return &subsRepoPG{
		dbStorage: dbStorage,
	}
}

// Create creates new subscription.
// All necessary fields must be presented.
func (r *subsRepoPG) Create(subs *entity.Subscription) error {
	if err := r.dbStorage.Create(subs).Error; err != nil {
		return fmt.Errorf("create: %w", err)
	}
	return nil
}

// GetByID gets subscription by given ID and returns it.
func (r *subsRepoPG) GetByID(id string) (*entity.Subscription, error) {
	subs := &entity.Subscription{}

	err := r.dbStorage.Where("id = ?", id).First(subs).Error
	// if record not found
	if goerrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("get by id: %w", err)
	}
	return subs, nil
}

// Update updates all subscription fields with given data by giving subs id.
func (r *subsRepoPG) Update(subs *entity.Subscription) error {
	// check that given subs exists
	if _, err := r.GetByID(subs.ID); err != nil {
		return fmt.Errorf("update: %w", err)
	}
	// update subs
	if err := r.dbStorage.Save(subs).Error; err != nil {
		return fmt.Errorf("update: %w", err)
	}
	return nil
}

// Delete deletes subscription by its ID.
func (r *subsRepoPG) Delete(id string) error {
	if err := r.dbStorage.Delete(&entity.Subscription{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return nil
}

// GetList gets all subscriptions and returns it.
func (r *subsRepoPG) GetList() (entity.SubscriptionList, error) {
	var subsList entity.SubscriptionList

	if err := r.dbStorage.Find(&subsList).Error; err != nil {
		return nil, fmt.Errorf("get list: %w", err)
	}
	return subsList, nil
}

// GetSum returns sum of subs prices filtered by given filter.
func (r *subsRepoPG) GetSum(filter *entity.SubscriptionSumFilter) (int, error) {
	var prices []int

	dbSuery := r.dbStorage.Model(&entity.Subscription{})
	if filter.UserID != "" {
		dbSuery = dbSuery.Where("user_id = ?", filter.UserID)
	}
	if filter.ServiceName != "" {
		dbSuery = dbSuery.Where("service_name = ?", filter.ServiceName)
	}
	if filter.StartDate != nil {
		dbSuery = dbSuery.Where("start_date >= ?::date", filter.StartDate)
	}
	if filter.EndDate != nil {
		dbSuery = dbSuery.Where("end_date <= ?::date", filter.EndDate)
	}
	// select prices
	err := dbSuery.Pluck("price", &prices).Error
	if err != nil {
		return 0, fmt.Errorf("get sum: %w", err)
	}

	// sum gotten prices
	var totalPrice int
	for _, price := range prices {
		totalPrice += price
	}
	return totalPrice, nil
}
