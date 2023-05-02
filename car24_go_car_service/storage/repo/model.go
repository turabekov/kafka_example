package repo

import (
	"gitlab.udevs.io/car24/car24_go_car_service/modules/car24/car24_car_service"
)

type ModelI interface {
	Create(car *car24_car_service.CreateModel) (string, error)
	Get(id string) (*car24_car_service.Model, error)
	GetAll(car24_car_service.ModelQueryParamModel) (car24_car_service.ModelListModel, error)
	Update(car *car24_car_service.UpdateModel) error
	Delete(id string) error
}
