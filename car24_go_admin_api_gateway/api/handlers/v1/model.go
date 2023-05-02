package v1

import (
	"net/http"

	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/modules/car24/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/modules/car24/car24_car_service"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

// @Security ApiKeyAuth
// @Router /v1/model [post]
// @Summary Create Model
// @Description API for creating model
// @Tags model
// @Accept json
// @Produce json
// @Param Platform-Id header string true "platform-id" default(7d4a4c38-dd84-4902-b744-0488b80a4c01)
// @Param Model body car24_car_service.CreateModel true "model"
// @Success 201 {object} response.ResponseOK
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
func (h *handlerV1) CreateModel(c *gin.Context) {
	var model car24_car_service.CreateModel

	err := c.ShouldBindJSON(&model)
	if HandleError(c, h.log, err, "error while binding body to model") {
		return
	}

	correlationID, err := uuid.NewRandom()
	if HandleError(c, h.log, err, "Error while generating new uuid") {
		return
	}

	id, _ := uuid.NewRandom()
	model.ID = id.String()
	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetSource(c.FullPath())
	e.SetType("v1.car_service.model.create")
	err = e.SetData(cloudevents.ApplicationJSON, model)
	if HandleError(c, h.log, err, "error while setting event data") {
		return
	}

	err = h.kafka.Push("v1.car_service.model.create", e)
	if HandleError(c, h.log, err, "error while publishing event") {
		return
	}

	c.JSON(http.StatusOK, response.ResponseOK{
		Message: correlationID.String(),
	})
}

// @Security ApiKeyAuth
// @Router /v1/model/{id} [get]
// @Summary Get Model
// @Description API for getting model
// @Tags model
// @Accept json
// @Produce json
// @Param Platform-Id header string true "platform-id" default(7d4a4c38-dd84-4902-b744-0488b80a4c01)
// @Param id path string true "id"
// @Success 200 {object} car24_car_service.Model
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
func (h *handlerV1) GetModel(c *gin.Context) {
	_ = h.MakeProxy(c, h.cfg.CarServiceURL, c.Request.URL.Path)
}

// @Security ApiKeyAuth
// @Router /v1/model [get]
// @Summary Get models
// @Description API for getting all models
// @Tags model
// @Accept json
// @Produce json
// @Param find query car24_car_service.ModelQueryParamModel false "filters"
// @Success 200 {object} car24_car_service.ModelListModel
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
func (h *handlerV1) GetAllModels(c *gin.Context) {

	_ = h.MakeProxy(c, h.cfg.CarServiceURL, c.Request.URL.Path)
}

// @Security ApiKeyAuth
// @Router /v1/model/{id} [put]
// @Summary Update Model
// @Description API for updating model
// @Tags model
// @Accept json
// @Produce json
// @Param Platform-Id header string true "platform-id" default(7d4a4c38-dd84-4902-b744-0488b80a4c01)
// @Param id path string true "id"
// @Param Model body car24_car_service.UpdateModel true "car"
// @Success 201 {object} response.ResponseOK
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
func (h *handlerV1) UpdateModel(c *gin.Context) {
	var (
		model car24_car_service.UpdateModel
	)

	err := c.ShouldBindJSON(&model)
	if HandleError(c, h.log, err, "error while binding body to model") {
		return
	}

	correlationID, err := uuid.NewRandom()
	if HandleError(c, h.log, err, "Error while generating new uuid") {
		return
	}

	model.ID = c.Param("id")

	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetSource(c.FullPath())
	e.SetType("v1.car_service.model.update")
	err = e.SetData(cloudevents.ApplicationJSON, model)
	if HandleError(c, h.log, err, "error while setting event data") {
		return
	}

	err = h.kafka.Push("v1.car_service.model.update", e)
	if HandleError(c, h.log, err, "error while publishing event") {
		return
	}

	c.JSON(http.StatusOK, response.ResponseOK{
		Message: correlationID.String(),
	})
}

// @Security ApiKeyAuth
// @Router /v1/model/{id} [delete]
// @Summary Delete Model
// @Description API for deleting model
// @Tags model
// @Accept json
// @Produce json
// @Param Platform-Id header string true "platform-id" default(7d4a4c38-dd84-4902-b744-0488b80a4c01)
// @Param id path string true "id"
// @Success 201 {object} response.ResponseOK
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
func (h *handlerV1) DeleteModel(c *gin.Context) {
	var (
		request car24_car_service.DeleteModel
	)

	correlationID, err := uuid.NewRandom()
	if HandleError(c, h.log, err, "Error while generating new uuid") {
		return
	}

	request.ID = c.Param("id")

	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetSource(c.FullPath())
	e.SetType("v1.car_service.model.delete")
	err = e.SetData(cloudevents.ApplicationJSON, request)
	if HandleError(c, h.log, err, "error while setting event data") {
		return
	}

	err = h.kafka.Push("v1.car_service.model.delete", e)
	if HandleError(c, h.log, err, "error while publishing event") {
		return
	}

	c.JSON(http.StatusOK, response.ResponseOK{
		Message: correlationID.String(),
	})
}
