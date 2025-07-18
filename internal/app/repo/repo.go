// Package repo contains interfaces of repositories for all entities.
// Its implementations like DB repositories, mocks,
// etc. are in sub-packages with the same names.
package repo

import (
	"SubscriptionAggregator/internal/app/entity"
)

type SubsRepoDB interface {
	Create(subs *entity.Subscription) error
	GetByID(id string) (*entity.Subscription, error)
	Update(subs *entity.Subscription) (*entity.Subscription, error)
	Delete(id string) error
	GetList() (entity.SubscriptionList, error)
	GetSum(filter *entity.SubscriptionSumFilter) (int, error)
}
