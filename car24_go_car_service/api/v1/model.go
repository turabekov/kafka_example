package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.udevs.io/car24/car24_go_car_service/modules/car24/car24_car_service"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/logger"
)

// @Router /v1/model/{id} [get]
// @Summary Get Model
// @Description API for getting model
// @Tags model
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.CarModel
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
func (h *handlerV1) GetModel(c *gin.Context) {
	id := c.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		h.log.Error("Error whiling getting model", logger.Any("error", err))
		h.HandleBadRequest(c, err, "Id format should be uuid")
		return
	}

	resp, err := h.storage.Model().Get(id)
	if err != nil {
		h.log.Error("Error whiling getting model", logger.Any("error", err))
		h.HandleError(c, err, "Model not found")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router /v1/model [get]
// @Summary Get models
// @Description API for get models
// @Tags model
// @Accept json
// @Produce json
// @Param find query response.GetAllModelsRequest false "filters"
// @Success 201 {object} response.GetAllModelsResponse
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
func (h *handlerV1) GetAllModels(c *gin.Context) {
	var (
		queryParam car24_car_service.ModelQueryParamModel
		err        error
	)

	queryParam.Offset, err = ParseQueryParam(c, "offset", "0")
	if err != nil {
		h.HandleError(c, err, "wrong offset input")
		return
	}
	queryParam.Limit, err = ParseQueryParam(c, "limit", "10")
	if err != nil {
		h.HandleError(c, err, "wrong limit input")
		return
	}

	queryParam.BrandId = c.Query("brand_id")

	h.log.Info("----GETALL_MODELS--->", logger.Any("Request", queryParam))

	resp, err := h.storage.Model().GetAll(queryParam)
	if err != nil {
		h.log.Error("Error whiling getting model", logger.Any("error", err))
		h.HandleInternalServerError(c, err, "Something went wrong")
		return
	}

	h.log.Info("----GETALL_MODELS--->", logger.Any("Response", resp))

	h.handleSuccessResponse(c, 200, "ok", resp)
}
