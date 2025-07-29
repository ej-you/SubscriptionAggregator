// Package http/v1 is a first version of HTTP-controller.
// It provides registers for HTTP-routes and controllers with handlers for them.
package v1

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	"SubscriptionAggregator/internal/app/entity"
	"SubscriptionAggregator/internal/app/errors"
	"SubscriptionAggregator/internal/pkg/validator"
)

type SubsUsecase interface {
	Create(subs *entity.Subscription) error
	GetByID(id string) (*entity.Subscription, error)
	Update(subs *entity.SubscriptionUpdate) (*entity.Subscription, error)
	Delete(id string) error
	GetList(subsList *entity.SubscriptionList) error
	GetSum(filter *entity.SubscriptionSumFilter) (int, error)
}

// SubsController is a HTTP-controller for subs usecase.
type SubsController struct {
	subsUC SubsUsecase
	valid  validator.Validator
}

// NewSubsController returns new subs controller.
func NewSubsController(subsUC SubsUsecase, valid validator.Validator) *SubsController {
	return &SubsController{
		subsUC: subsUC,
		valid:  valid,
	}
}

// @summary		Создать запись подписки
// @description	Создание новой записи подписки.
// @router			/subs [post]
// @id				create-sub
// @tags			subs-crudl
// @param			Sub	body		inSubsCreate	true	"Информация о подписке"
// @success		201	{object}	entity.Subscription
// @failure		400	"Невалидное тело запроса"
func (c *SubsController) Create(ctx *fiber.Ctx) error {
	bodyData := &inSubsCreate{}
	// parse body
	if err := ctx.BodyParser(bodyData); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}
	// validate parsed data
	if err := c.valid.Validate(bodyData); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}
	// parse dates
	if err := bodyData.ParseDates(); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}

	subs := entity.Subscription{
		ServiceName: bodyData.ServiceName,
		Price:       bodyData.Price,
		UserID:      bodyData.UserID,
		StartDate:   bodyData.StartDateParsed,
		EndDate:     bodyData.EndDateParsed,
	}
	// create subs
	if err := c.subsUC.Create(&subs); err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(subs)
}

// @summary		Получить запись подписки
// @description	Получение записи подписки по её ID.
// @router			/subs/{id} [get]
// @id				get-sub
// @tags			subs-crudl
// @param			id	path		string	true	"UUID подписки"
// @success		200	{object}	entity.Subscription
// @failure		400	"Невалидный параметр запроса"
// @failure		404	"Подписка не найдена"
func (c *SubsController) GetByID(ctx *fiber.Ctx) error {
	pathData := &inPathUUID{}
	// parse path-params
	if err := ctx.ParamsParser(pathData); err != nil {
		return fmt.Errorf("parse path: %w", err)
	}
	// validate parsed data
	if err := c.valid.Validate(pathData); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}

	// get subs
	subs, err := c.subsUC.GetByID(pathData.ID)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(subs)
}

// @summary		Обновить запись подписки
// @description	Обновление записи подписки по её ID.
// @router			/subs/{id} [patch]
// @id				update-sub
// @tags			subs-crudl
// @param			id	path		string			true	"UUID подписки"
// @param			Sub	body		inSubsUpdate	true	"Информация о подписке"
// @success		200	{object}	entity.Subscription
// @failure		400	"Невалидный параметр или тело запроса"
// @failure		404	"Подписка не найдена"
func (c *SubsController) Update(ctx *fiber.Ctx) error {
	pathData := &inPathUUID{}
	// parse path-params
	if err := ctx.ParamsParser(pathData); err != nil {
		return fmt.Errorf("parse path: %w", err)
	}
	// validate parsed data
	if err := c.valid.Validate(pathData); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}
	bodyData := &inSubsUpdate{}
	// parse body
	if err := ctx.BodyParser(bodyData); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}
	// validate parsed data
	if err := c.valid.Validate(bodyData); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}
	// parse dates
	if err := bodyData.ParseDates(); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}

	subs := entity.SubscriptionUpdate{
		ID:          pathData.ID,
		ServiceName: bodyData.ServiceName,
		Price:       bodyData.Price,
		UserID:      bodyData.UserID,
		StartDate:   bodyData.StartDateParsed,
		EndDate:     bodyData.EndDateParsed,
	}
	// update subs
	updatedSubs, err := c.subsUC.Update(&subs)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(updatedSubs)
}

// @summary		Удалить запись подписки
// @description	Удаление записи подписки по её ID.
// @router			/subs/{id} [delete]
// @id				delete-sub
// @tags			subs-crudl
// @param			id	path	string	true	"UUID подписки"
// @success		204	"Успешное удаление"
// @failure		400	"Невалидный параметр запроса"
func (c *SubsController) Delete(ctx *fiber.Ctx) error {
	pathData := &inPathUUID{}
	// parse path-params
	if err := ctx.ParamsParser(pathData); err != nil {
		return fmt.Errorf("parse path: %w", err)
	}
	// validate parsed data
	if err := c.valid.Validate(pathData); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
		// err = fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
		// return ctx.Status(400).JSON(err.Error())
	}

	// get subs
	if err := c.subsUC.Delete(pathData.ID); err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

// @summary		Получить все записи подписок
// @description	Получение всех записей подписок с пагинацией.
// @router			/subs [get]
// @id				get-all-subs
// @tags			subs-crudl
// @param			limit	query		int	false	"Кол-во результатов на страницу (по умолчанию: 10)"	example:"5"
// @param			page	query		int	false	"Номер страницы (по умолчанию: 1)"					example:"2"
// @success		200		{object}	entity.SubscriptionList
func (c *SubsController) GetList(ctx *fiber.Ctx) error {
	subsList := &entity.SubscriptionList{
		Data:       make([]entity.Subscription, 0),
		Pagination: &entity.SubscriptionPagination{},
	}
	// parse query-params
	if err := ctx.QueryParser(subsList.Pagination); err != nil {
		return fmt.Errorf("parse query: %w", err)
	}
	// validate parsed data
	if err := c.valid.Validate(subsList.Pagination); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}

	// get all subs
	err := c.subsUC.GetList(subsList)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(subsList)
}

// @summary		Получить суммарную стоимость подписок
// @description	Получение суммарной стоимости всех подписок за выбранный период с фильтрацией по id пользователя и названию подписки.
// @router			/subs-sum [get]
// @id				get-subs-sum
// @tags			subs-advanced
// @param			user_id			query		string	false	"UUID пользователя"	example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"
// @param			service_name	query		string	false	"Название сервиса"	example:"Yandex Plus"
// @param			start_date		query		string	false	"Дата начала"		example:"07-2025"
// @param			end_date		query		string	false	"Дата окончания"	example:"08-2025"
// @success		200				{object}	entity.SubscriptionSum
// @failure		400				"Невалидный(ые) параметр(ы) запроса"
func (c *SubsController) GetSum(ctx *fiber.Ctx) error {
	queryData := &inSubSumFilter{}
	// parse path-params
	if err := ctx.QueryParser(queryData); err != nil {
		return fmt.Errorf("parse query: %w", err)
	}
	// validate parsed data
	if err := c.valid.Validate(queryData); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}
	// parse dates
	if err := queryData.ParseDates(); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}

	subSumFilter := entity.SubscriptionSumFilter{
		ServiceName: queryData.ServiceName,
		UserID:      queryData.UserID,
		StartDate:   queryData.StartDateParsed,
		EndDate:     queryData.EndDateParsed,
	}
	// get subs
	subsSum, err := c.subsUC.GetSum(&subSumFilter)
	if err != nil {
		return err
	}
	totalData := entity.SubscriptionSum{Sum: subsSum}
	return ctx.Status(fiber.StatusOK).JSON(totalData)
}
