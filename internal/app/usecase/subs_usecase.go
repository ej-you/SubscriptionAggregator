package usecase

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"SubscriptionAggregator/internal/app/entity"
	"SubscriptionAggregator/internal/app/repo"
)

var _ SubsUsecase = (*subsUsecase)(nil)

// SubsUsecase implementation.
type subsUsecase struct {
	subsRepoDB repo.SubsRepoDB
}

func NewSubsUsecase(subsRepoDB repo.SubsRepoDB) SubsUsecase {
	return &subsUsecase{
		subsRepoDB: subsRepoDB,
	}
}

// Create creates new subs.
// All required fields must be presented. ID is auto-generated.
func (u *subsUsecase) Create(subs *entity.Subscription) error {
	subs.ID = uuid.NewString()
	err := u.subsRepoDB.Create(subs)
	return errors.Wrap(err, "create subs")
}

// Get gets one subs by given ID.
func (u *subsUsecase) GetByID(id string) (*entity.Subscription, error) {
	subs, err := u.subsRepoDB.GetByID(id)
	return subs, errors.Wrap(err, "get subs by id")
}

// Update updates all subs fields with given data by giving book ID.
// ID and all required fields must be presented.
func (u *subsUsecase) Update(subs *entity.Subscription) error {
	err := u.subsRepoDB.Update(subs)
	return errors.Wrap(err, "update subs")
}

// Delete deletes subs by its ID.
func (u *subsUsecase) Delete(id string) error {
	err := u.subsRepoDB.Delete(id)
	return errors.Wrap(err, "delete subs")
}

// GetAll gets all subs.
func (u *subsUsecase) GetAll() (entity.SubscriptionList, error) {
	subsList, err := u.subsRepoDB.GetList()
	return subsList, errors.Wrap(err, "get all subs")
}

// GetSum returns sum of subs prices filtered by filter.
func (u *subsUsecase) GetSum(filter *entity.SubscriptionSumFilter) (int, error) {
	totalPrice, err := u.subsRepoDB.GetSum(filter)
	return totalPrice, errors.Wrap(err, "get subs prices sum")
}
