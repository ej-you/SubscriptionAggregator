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

// Update updates subscription.
// It selects subs by given ID and replace all old values (from DB) to new (given).
// It returns full filled updated subs.
func (r *subsRepoPG) Update(subs *entity.Subscription) (*entity.Subscription, error) {
	// get subs by given ID
	subsFromDB, err := r.GetByID(subs.ID)
	if err != nil {
		return nil, fmt.Errorf("update: %w", err)
	}
	// map with fields to update
	updates := make(map[string]any)
	// append non-nil new field values to map
	if len(subs.ServiceName) != 0 {
		updates["service_name"] = subs.ServiceName
		subsFromDB.ServiceName = subs.ServiceName
	}
	if subs.Price != 0 {
		updates["price"] = subs.Price
		subsFromDB.Price = subs.Price
	}
	if len(subs.UserID) != 0 {
		updates["user_id"] = subs.UserID
		subsFromDB.UserID = subs.UserID
	}
	if subs.StartDate != nil {
		updates["start_date"] = subs.StartDate
		subsFromDB.StartDate = subs.StartDate
	}
	if subs.EndDate != nil {
		updates["end_date"] = subs.EndDate
		subsFromDB.EndDate = subs.EndDate
	}

	// if nothing to update
	if len(updates) == 0 {
		return subsFromDB, nil
	}
	// update subs
	err = r.dbStorage.Model(&entity.Subscription{}).
		Where("id = ?", subs.ID).
		Updates(updates).Error
	if err != nil {
		return nil, fmt.Errorf("update: %w", err)
	}
	return subsFromDB, nil
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
