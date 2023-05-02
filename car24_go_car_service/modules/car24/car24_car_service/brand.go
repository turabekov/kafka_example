package car24_car_service

type Brand struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateBrand struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateBrand struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DeleteBrand struct {
	ID string `json:"id" swaggerignore:"true"`
}

type BrandQueryParamModel struct {
	Search string `json:"search"`
	Offset int    `json:"offset" default:"0"`
	Limit  int    `json:"limit" default:"10"`
}

type BrandListModel struct {
	Brands []Brand `json:"cars"`
	Count  int     `json:"count"`
}
