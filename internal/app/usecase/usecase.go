// Package usecase contains interfaces of usecases
// and its implementations for all entities.
package usecase

import (
	"SubscriptionAggregator/internal/app/entity"
)

type SubsUsecase interface {
	Create(subs *entity.Subscription) error
	GetByID(id string) (*entity.Subscription, error)
	Update(subs *entity.Subscription) error
	Delete(id string) error
	GetAll() (entity.SubscriptionList, error)
	GetSum(filter *entity.SubscriptionSumFilter) (int, error)
}
