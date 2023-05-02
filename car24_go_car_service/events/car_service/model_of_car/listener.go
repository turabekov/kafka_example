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

func (c *ModelService) Create(ctx context.Context, event cloudevents.Event) response.Response {
	var (
		model cs.CreateModel
	)

	c.log.Info("Model model", logger.Any("event", event))

	err := json.Unmarshal(event.DataEncoded, &model)

	resp, hasError := helper.HandleBadRequest(c.log, model.ID, "Error while unmarshalling", err)
	if hasError {
		resp.CorrelationID = model.ID
		return resp
	}

	_, err = c.storage.Model().Create(&model)

	resp, hasError = helper.HandleInternal(c.log, model.ID, "Error while creating model", err)
	if hasError {
		resp.CorrelationID = model.ID
		return resp
	}

	c.log.Info("Successfully created Model", logger.String("id", model.ID))

	err = c.kafka.Push("v1.car_service.model.created", event)
	resp, hasError = helper.HandleInternal(c.log, model.ID, "Error while creating model", err)
	if hasError {
		resp.CorrelationID = model.ID
		return resp
	}

	resp = response.Response{
		ID:         model.ID,
		StatusCode: http.StatusOK,
		Data:       model,
		Action:     "create",
	}

	return resp
}

func (c *ModelService) Update(ctx context.Context, event cloudevents.Event) response.Response {
	var (
		model cs.UpdateModel
	)

	c.log.Info("model update", logger.Any("event", event))

	err := json.Unmarshal(event.DataEncoded, &model)
	resp, hasError := helper.HandleBadRequest(c.log, model.ID, "Error while marshaling", err)
	if hasError {
		resp.CorrelationID = model.ID
		return resp
	}

	err = c.storage.Model().Update(&model)
	resp, hasError = helper.HandleInternal(c.log, model.ID, "Error while updating model", err)
	if hasError {
		resp.CorrelationID = model.ID
		return resp
	}

	c.log.Info("Successfully Updated model", logger.String("id", model.ID))

	err = c.kafka.Push("v1.car_service.model.updated", event)
	resp, hasError = helper.HandleInternal(c.log, model.ID, "Error while updating model", err)
	if hasError {
		resp.CorrelationID = model.ID
		return resp
	}

	return response.Response{
		ID:         model.ID,
		StatusCode: http.StatusOK,
		Action:     "update",
	}
}

func (c *ModelService) Delete(ctx context.Context, event cloudevents.Event) response.Response {
	var (
		model cs.DeleteModel
	)

	c.log.Info("Model delete", logger.Any("event", event))

	err := json.Unmarshal(event.DataEncoded, &model)
	resp, hasError := helper.HandleInternal(c.log, model.ID, "Error while unmarshalling model", err)
	if hasError {
		resp.CorrelationID = model.ID
		return resp
	}

	id := model.ID

	err = c.storage.Model().Delete(string(id))
	resp, hasError = helper.HandleInternal(c.log, model.ID, "Error while deleting model", err)
	if hasError {
		resp.CorrelationID = model.ID
		return resp
	}

	c.log.Info("Successfully Deleted model", logger.String("id", model.ID))

	err = c.kafka.Push("v1.car_service.model.deleted", event)
	resp, hasError = helper.HandleInternal(c.log, model.ID, "Error while deleting model", err)
	if hasError {
		resp.CorrelationID = model.ID
		return resp
	}

	return response.Response{
		ID:         model.ID,
		StatusCode: http.StatusOK,
		Action:     "delete",
	}
}
