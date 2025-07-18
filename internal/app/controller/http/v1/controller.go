// Package http/v1 is a first version of HTTP-controller.
// It provides registers for HTTP-routes and controllers with handlers for them.
package v1

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	"SubscriptionAggregator/internal/app/entity"
	"SubscriptionAggregator/internal/app/errors"
	"SubscriptionAggregator/internal/app/usecase"
	"SubscriptionAggregator/internal/pkg/validator"
)

// SubsController is a HTTP-controller for subs usecase.
type SubsController struct {
	subsUC usecase.SubsUsecase
	valid  validator.Validator
}

// NewSubsController returns new SubsController.
func NewSubsController(subsUC usecase.SubsUsecase, valid validator.Validator) *SubsController {
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
// @param			Sub	body		inSubs	true	"Информация о подписке"
// @success		201	{object}	entity.Subscription
// @failure		400	"Невалидное тело запроса"
func (c *SubsController) Create(ctx *fiber.Ctx) error {
	bodyData := &inSubs{}
	// parse body
	if err := ctx.BodyParser(bodyData); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}
	// validate parsed data
	if err := c.valid.Validate(bodyData); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}

	subs := entity.Subscription{
		ServiceName: bodyData.ServiceName,
		Price:       bodyData.Price,
		UserID:      bodyData.UserID,
		StartDate:   bodyData.StartDate,
		EndDate:     bodyData.EndDate,
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
		return fmt.Errorf("parse body: %w", err)
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
// @param			id	path		string	true	"UUID подписки"
// @param			Sub	body		inSubs	true	"Информация о подписке"
// @success		200	{object}	entity.Subscription
// @failure		400	"Невалидный параметр или тело запроса"
// @failure		404	"Подписка не найдена"
func (c *SubsController) Update(ctx *fiber.Ctx) error {
	pathData := &inPathUUID{}
	// parse path-params
	if err := ctx.ParamsParser(pathData); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}
	// validate parsed data
	if err := c.valid.Validate(pathData); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}
	bodyData := &inSubs{}
	// parse body
	if err := ctx.BodyParser(bodyData); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}
	// validate parsed data
	if err := c.valid.Validate(bodyData); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}

	subs := entity.Subscription{
		ID:          pathData.ID,
		ServiceName: bodyData.ServiceName,
		Price:       bodyData.Price,
		UserID:      bodyData.UserID,
		StartDate:   bodyData.StartDate,
		EndDate:     bodyData.EndDate,
	}

	// update subs
	if err := c.subsUC.Update(&subs); err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(subs)
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
		return fmt.Errorf("parse body: %w", err)
	}
	// validate parsed data
	if err := c.valid.Validate(pathData); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}

	// get subs
	if err := c.subsUC.Delete(pathData.ID); err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

// @summary		Получить все записи подписок
// @description	Получение всех записей подписок.
// @router			/subs [get]
// @id				get-all-subs
// @tags			subs-crudl
// @success		200	{object}	entity.SubscriptionList
func (c *SubsController) GetAll(ctx *fiber.Ctx) error {
	// get all subs
	subsList, err := c.subsUC.GetAll()
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(subsList)
}

// @summary		Получить суммарную стоимость подписок
// @description	Получение суммарной стоимости всех подписок за выбранный период с фильтрацией по id пользователя и названию подписки.
// @router			/subs/sum [get]
// @id				get-subs-sum
// @tags			subs-advanced
// @param			Filter	query		inSubSumFilter	false	"Поля для фильтрации"
// @success		200		{object}	entity.SubscriptionSum
// @failure		400		"Невалидный(ые) параметр(ы) запроса"
func (c *SubsController) GetSum(ctx *fiber.Ctx) error {
	queryData := &inSubSumFilter{}
	// parse path-params
	if err := ctx.QueryParser(queryData); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}
	// validate parsed data
	if err := c.valid.Validate(queryData); err != nil {
		return fmt.Errorf("%w: %s", errors.ErrValidateData, err.Error())
	}

	subSumFilter := entity.SubscriptionSumFilter{
		ServiceName: queryData.ServiceName,
		UserID:      queryData.UserID,
		StartDate:   queryData.StartDate,
		EndDate:     queryData.EndDate,
	}
	// get subs
	subsSum, err := c.subsUC.GetSum(&subSumFilter)
	if err != nil {
		return err
	}
	totalData := entity.SubscriptionSum{Sum: subsSum}
	return ctx.Status(fiber.StatusOK).JSON(totalData)
}
