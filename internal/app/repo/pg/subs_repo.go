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

// NewSubsRepoDB returns new SubsRepoDB instance.
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
func (r *subsRepoPG) Update(subs *entity.SubscriptionUpdate) (*entity.Subscription, error) {
	// update subs
	err := r.dbStorage.Model(&entity.Subscription{}).
		Where("id = ?", subs.ID).
		Updates(subs).Error
	if err != nil {
		return nil, fmt.Errorf("update: %w", err)
	}

	// get updated subs by ID
	subsFromDB, err := r.GetByID(subs.ID)
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

	dbQuery := r.dbStorage.Model(&entity.Subscription{})
	// apply main conditions
	if filter.UserID != "" {
		dbQuery = dbQuery.Where("user_id = ?", filter.UserID)
	}
	if filter.ServiceName != "" {
		dbQuery = dbQuery.Where("service_name = ?", filter.ServiceName)
	}

	dateCond := r.dbStorage.Model(&entity.Subscription{})
	// collect date condition
	if filter.StartDate != nil {
		dateCond = dateCond.Or("start_date <= ?::date AND end_date IS NULL", filter.StartDate)
	}
	if filter.EndDate != nil {
		dateCond = dateCond.Or("end_date = ?::date", filter.EndDate)
	}
	if filter.StartDate != nil && filter.EndDate != nil {
		dateCond = dateCond.Or("start_date >= ?::date AND start_date < ?::date",
			filter.StartDate, filter.EndDate).
			Or("start_date <= ?::date AND end_date >= ?::date", filter.StartDate, filter.EndDate).
			Or("end_date >= ?::date AND end_date <= ?::date", filter.StartDate, filter.EndDate)
	}
	// connect date condition to main conditions
	dbQuery = dbQuery.Where(dateCond)

	// select prices
	err := dbQuery.Pluck("price", &prices).Error
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
