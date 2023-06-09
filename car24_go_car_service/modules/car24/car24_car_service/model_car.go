package car24_car_service

type CarModel struct {
	ID          string `json:"id"`
	ModelID     string `json:"model_id"`
	ModelData   *Model `json:"model_data"`
	StateNumber string `json:"state_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateCarModel struct {
	ID          string `json:"id"`
	ModelID     string `json:"model_id"`
	StateNumber string `json:"state_number"`
}

type UpdateCarModel struct {
	ID          string `json:"id"`
	ModelID     string `json:"model_id"`
	StateNumber string `json:"state_number"`
}

type DeleteCarModel struct {
	ID string `json:"id" swaggerignore:"true"`
}

type CarQueryParamModel struct {
	Search  string `json:"search"`
	Offset  int    `json:"offset" default:"0"`
	Limit   int    `json:"limit" default:"10"`
	ModelId string `json:"model_id"`
}

type CarListModel struct {
	Cars  []CarModel `json:"cars"`
	Count int        `json:"count"`
}
