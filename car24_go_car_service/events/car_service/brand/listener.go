package car

import (
	"context"
	"encoding/json"
	"net/http"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	cs "gitlab.udevs.io/car24/car24_go_car_service/modules/car24/car24_car_service"
	"gitlab.udevs.io/car24/car24_go_car_service/modules/car24/response"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/helper"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/logger"
)

func (c *BrandService) Create(ctx context.Context, event cloudevents.Event) response.Response {
	var (
		brand cs.CreateBrand
	)

	c.log.Info("brand create", logger.Any("event", event))

	err := json.Unmarshal(event.DataEncoded, &brand)

	resp, hasError := helper.HandleBadRequest(c.log, brand.ID, "Error while unmarshalling", err)
	if hasError {
		resp.CorrelationID = brand.ID
		return resp
	}

	_, err = c.storage.Brand().Create(&brand)

	resp, hasError = helper.HandleInternal(c.log, brand.ID, "Error while creating car", err)
	if hasError {
		resp.CorrelationID = brand.ID
		return resp
	}

	c.log.Info("Successfully created Car", logger.String("id", brand.ID))

	err = c.kafka.Push("v1.car_service.brand.created", event)
	resp, hasError = helper.HandleInternal(c.log, brand.ID, "Error while creating car", err)
	if hasError {
		resp.CorrelationID = brand.ID
		return resp
	}

	resp = response.Response{
		ID:         brand.ID,
		StatusCode: http.StatusOK,
		Data:       brand,
		Action:     "create",
	}

	return resp
}

func (c *BrandService) Update(ctx context.Context, event cloudevents.Event) response.Response {
	var (
		brand cs.UpdateBrand
	)

	c.log.Info("brand update", logger.Any("event", event))

	err := json.Unmarshal(event.DataEncoded, &brand)
	resp, hasError := helper.HandleBadRequest(c.log, brand.ID, "Error while marshaling", err)
	if hasError {
		resp.CorrelationID = brand.ID
		return resp
	}

	err = c.storage.Brand().Update(&brand)
	resp, hasError = helper.HandleInternal(c.log, brand.ID, "Error while updating car", err)
	if hasError {
		resp.CorrelationID = brand.ID
		return resp
	}

	c.log.Info("Successfully Updated Car", logger.String("id", brand.ID))

	err = c.kafka.Push("v1.car_service.brand.updated", event)
	resp, hasError = helper.HandleInternal(c.log, brand.ID, "Error while updating car", err)
	if hasError {
		resp.CorrelationID = brand.ID
		return resp
	}

	return response.Response{
		ID:         brand.ID,
		StatusCode: http.StatusOK,
		Action:     "update",
	}
}

func (c *BrandService) Delete(ctx context.Context, event cloudevents.Event) response.Response {
	var (
		car cs.DeleteCarModel
	)

	c.log.Info("Car delete", logger.Any("event", event))

	err := json.Unmarshal(event.DataEncoded, &car)
	resp, hasError := helper.HandleInternal(c.log, car.ID, "Error while unmarshalling car", err)
	if hasError {
		resp.CorrelationID = car.ID
		return resp
	}

	id := car.ID

	err = c.storage.Brand().Delete(string(id))
	resp, hasError = helper.HandleInternal(c.log, car.ID, "Error while deleting car", err)
	if hasError {
		resp.CorrelationID = car.ID
		return resp
	}

	c.log.Info("Successfully Deleted Car", logger.String("id", car.ID))

	err = c.kafka.Push("v1.car_service.brand.deleted", event)
	resp, hasError = helper.HandleInternal(c.log, car.ID, "Error while deleting car", err)
	if hasError {
		resp.CorrelationID = car.ID
		return resp
	}

	return response.Response{
		ID:         car.ID,
		StatusCode: http.StatusOK,
		Action:     "delete",
	}
}
