package server

import (
	httpv1 "SubscriptionAggregator/internal/app/controller/http/v1"
	repopg "SubscriptionAggregator/internal/app/repo/pg"
	"SubscriptionAggregator/internal/app/usecase"
)

var (
	// check DB repos implementations
	_ usecase.ServiceRepoDB = (*repopg.ServiceRepoPG)(nil)
	_ usecase.SubsRepoDB    = (*repopg.SubsRepoPG)(nil)
	// check usecases implementations
	_ httpv1.SubsUsecase = (*usecase.SubsUsecase)(nil)
)

// registerEndpointsV1 registera all endpoints for 1 version of API.
func (s *httpServer) registerEndpointsV1() {
	// create repos
	subsRepoDB := repopg.NewSubsRepoDB(s.db)
	serviceRepoDB := repopg.NewServiceRepoDB(s.db)
	// create usecases
	subsUsecase := usecase.NewSubsUsecase(serviceRepoDB, subsRepoDB)
	// create controllers
	subsController := httpv1.NewSubsController(subsUsecase, s.valid)
	// register endpoints
	apiV1 := s.fiberApp.Group("/api/v1")
	httpv1.RegisterSubsEndpoints(apiV1, subsController)
}
