// Package usecase contains usecase implementations for entities.
package usecase

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"SubscriptionAggregator/internal/app/entity"
)

type ServiceRepoDB interface {
	GetByNameOrCreate(service *entity.Service) error
}

type SubsRepoDB interface {
	Create(subs *entity.Subscription) error
	GetByID(id string) (*entity.Subscription, error)
	Update(subs *entity.SubscriptionUpdate) (*entity.Subscription, error)
	Delete(id string) error
	GetList() (entity.SubscriptionList, error)
	GetSum(filter *entity.SubscriptionSumFilter) (int, error)
}

type SubsUsecase struct {
	serviceRepoDB ServiceRepoDB
	subsRepoDB    SubsRepoDB
}

// NewSubsUsecase returns new subs usecase instance.
func NewSubsUsecase(serviceRepoDB ServiceRepoDB, subsRepoDB SubsRepoDB) *SubsUsecase {
	return &SubsUsecase{
		subsRepoDB:    subsRepoDB,
		serviceRepoDB: serviceRepoDB,
	}
}

// Create creates new subs.
// All required fields must be presented. ID is auto-generated.
func (u *SubsUsecase) Create(subs *entity.Subscription) error {
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
func (u *SubsUsecase) GetByID(id string) (*entity.Subscription, error) {
	subs, err := u.subsRepoDB.GetByID(id)
	return subs, errors.Wrap(err, "get subs by id")
}

// Update updates all subs fields with given data by giving book ID.
// ID and all required fields must be presented.
func (u *SubsUsecase) Update(subs *entity.SubscriptionUpdate) (*entity.Subscription, error) {
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
func (u *SubsUsecase) Delete(id string) error {
	err := u.subsRepoDB.Delete(id)
	return errors.Wrap(err, "delete subs")
}

// GetAll gets all subs.
// TODO: add pagination
func (u *SubsUsecase) GetAll() (entity.SubscriptionList, error) {
	subsList, err := u.subsRepoDB.GetList()
	return subsList, errors.Wrap(err, "get all subs")
}

// GetSum returns sum of subs prices filtered by filter.
func (u *SubsUsecase) GetSum(filter *entity.SubscriptionSumFilter) (int, error) {
	totalPrice, err := u.subsRepoDB.GetSum(filter)
	return totalPrice, errors.Wrap(err, "get subs prices sum")
}
