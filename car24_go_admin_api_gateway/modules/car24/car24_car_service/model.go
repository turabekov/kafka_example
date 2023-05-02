package car24_car_service

type Model struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BrandID   string `json:"brand_id"`
	BrandData *Brand `json:"brand_data"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateModel struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	BrandID string `json:"brand_id"`
}

type UpdateModel struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	BrandID string `json:"brand_id"`
}

type DeleteModel struct {
	ID string `json:"id" swaggerignore:"true"`
}

type ModelQueryParamModel struct {
	Search  string `json:"search"`
	Offset  int    `json:"offset" default:"0"`
	Limit   int    `json:"limit" default:"10"`
	BrandId string `json:"brand_id"`
}

type ModelListModel struct {
	Models []*Model `json:"cars"`
	Count  int      `json:"count"`
}
