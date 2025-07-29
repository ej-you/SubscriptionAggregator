package usecase

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"SubscriptionAggregator/internal/app/entity"
	"SubscriptionAggregator/internal/app/repo"
)

var _ SubsUsecase = (*subsUsecase)(nil)

// SubsUsecase implementation.
type subsUsecase struct {
	subsRepoDB    repo.SubsRepoDB
	serviceRepoDB repo.ServiceRepoDB
}

// NewSubsUsecase returns new SubsUsecase instance.
func NewSubsUsecase(subsRepoDB repo.SubsRepoDB, serviceRepoDB repo.ServiceRepoDB) SubsUsecase {
	return &subsUsecase{
		subsRepoDB:    subsRepoDB,
		serviceRepoDB: serviceRepoDB,
	}
}

// Create creates new subs.
// All required fields must be presented. ID is auto-generated.
func (u *subsUsecase) Create(subs *entity.Subscription) error {
	logrus.Infof("Create subs: %+v", subs)

	service := &entity.Service{Name: subs.ServiceName}
	// get or create service
	if err := u.serviceRepoDB.GetByNameOrCreate(service); err != nil {
		return fmt.Errorf("get or create service: %w", err)
	}
	subs.ServiceID = service.ID
	// create subs
	if err := u.subsRepoDB.Create(subs); err != nil {
		return fmt.Errorf("create subs: %w", err)
	}
	return nil
}

// Get gets one subs by given ID.
func (u *subsUsecase) GetByID(id string) (*entity.Subscription, error) {
	subs, err := u.subsRepoDB.GetByID(id)
	return subs, errors.Wrap(err, "get subs by id")
}

// Update updates all subs fields with given data by giving book ID.
// ID and all required fields must be presented.
// TODO: check service update
func (u *subsUsecase) Update(subs *entity.SubscriptionUpdate) (*entity.Subscription, error) {
	// if service name is presented
	if subs.ServiceName != nil {
		service := &entity.Service{Name: *subs.ServiceName}
		// get or create service
		if err := u.serviceRepoDB.GetByNameOrCreate(service); err != nil {
			return nil, fmt.Errorf("get or create service: %w", err)
		}
		subs.ServiceID = &service.ID
	}

	updatedSubs, err := u.subsRepoDB.Update(subs)
	return updatedSubs, errors.Wrap(err, "update subs")
}

// Delete deletes subs by its ID.
func (u *subsUsecase) Delete(id string) error {
	err := u.subsRepoDB.Delete(id)
	return errors.Wrap(err, "delete subs")
}

// GetAll gets all subs.
// TODO: add pagination
func (u *subsUsecase) GetAll() (entity.SubscriptionList, error) {
	subsList, err := u.subsRepoDB.GetList()
	return subsList, errors.Wrap(err, "get all subs")
}

// GetSum returns sum of subs prices filtered by filter.
func (u *subsUsecase) GetSum(filter *entity.SubscriptionSumFilter) (int, error) {
	totalPrice, err := u.subsRepoDB.GetSum(filter)
	return totalPrice, errors.Wrap(err, "get subs prices sum")
}
