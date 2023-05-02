package postgres

import (
	"database/sql"
	"fmt"

	cs "gitlab.udevs.io/car24/car24_go_car_service/modules/car24/car24_car_service"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/helper"
	"gitlab.udevs.io/car24/car24_go_car_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type carRepo struct {
	db *sqlx.DB
}

func NewCar(db *sqlx.DB) repo.CarI {
	return &carRepo{
		db: db,
	}
}

func (br *carRepo) Create(car *cs.CreateCarModel) (string, error) {
	tsx, err := br.db.Begin()
	if err != nil {
		return "", err
	}
	defer func() {
		if err != nil {
			_ = tsx.Rollback()
		} else {
			err := tsx.Commit()
			if err != nil {
				fmt.Println("While committing transaction ", err)
			}
		}
	}()

	query := `
	INSERT INTO 
		car(
			id,
			model_id,
			state_number,
			updated_at
		)
	VALUES ($1, $2, $3, now())`

	_, err = tsx.Exec(
		query,
		car.ID,
		car.ModelID,
		car.StateNumber,
	)

	if err != nil {
		return "", err
	}

	return car.ID, nil
}

func (br *carRepo) Get(id string) (*cs.CarModel, error) {
	var (
		modelId     sql.NullString
		stateNumber sql.NullString
		updatedAt   sql.NullString
		createdAt   sql.NullString

		modelName      sql.NullString
		modelBrandId   sql.NullString
		modelUpdatedAt sql.NullString
		modelCreatedAt sql.NullString
	)

	query := `
		SELECT
			c.model_id,

			m.name,
			m.brand_id,
			m.updated_at,
			m.created_at,

			c.state_number,
			c.updated_at,
			c.created_at
		FROM car AS c
		LEFT JOIN model AS m ON c.model_id = m.id
		WHERE c.id = $1
	`

	row := br.db.QueryRow(query, id)
	err := row.Scan(
		&modelId,
		&modelName,
		&modelBrandId,
		&modelUpdatedAt,
		&modelCreatedAt,

		&stateNumber,
		&updatedAt,
		&createdAt,
	)
	if err != nil {
		return nil, err
	}

	m := &cs.Model{
		ID:        modelId.String,
		Name:      modelName.String,
		BrandID:   modelBrandId.String,
		UpdatedAt: modelUpdatedAt.String,
		CreatedAt: modelCreatedAt.String,
	}

	return &cs.CarModel{
		ID:          id,
		ModelID:     modelId.String,
		ModelData:   m,
		StateNumber: stateNumber.String,
		UpdatedAt:   updatedAt.String,
		CreatedAt:   createdAt.String,
	}, nil
}

func (br *carRepo) GetAll(queryParam cs.CarQueryParamModel) (res cs.CarListModel, err error) {

	var (
		filter = "WHERE 1=1"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		params = map[string]interface{}{}
		query  string
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			model_id,
			state_number,
			updated_at,
			created_at
		FROM
			car
	`

	if len(queryParam.ModelId) > 0 {
		filter += " AND model_id = '" + queryParam.ModelId + "' "
	}

	if queryParam.Offset > 0 {
		offset = " OFFSET :offset"
		params["offset"] = queryParam.Offset
	}

	if queryParam.Limit > 0 {
		limit = " LIMIT :limit"
		params["limit"] = queryParam.Limit
	}

	query += filter + offset + limit
	query, args := helper.ReplaceQueryParams(query, params)
	rows, err := br.db.Query(query, args...)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var (
			id          sql.NullString
			modelId     sql.NullString
			stateNumber sql.NullString
			updatedAt   sql.NullString
			createdAt   sql.NullString
		)

		err := rows.Scan(
			&res.Count,
			&id,
			&modelId,
			&stateNumber,
			&updatedAt,
			&createdAt,
		)
		if err != nil {
			return res, err
		}

		res.Cars = append(res.Cars, cs.CarModel{
			ID:          id.String,
			ModelID:     modelId.String,
			StateNumber: stateNumber.String,
			UpdatedAt:   updatedAt.String,
			CreatedAt:   createdAt.String,
		})
	}

	return res, nil
}

func (br *carRepo) Update(car *cs.UpdateCarModel) (err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE 
			    car
			SET
				model_id = :model_id,
				state_number = :state_number,
				updated_at = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":           car.ID,
		"model_id":     car.ModelID,
		"state_number": car.StateNumber,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	_, err = br.db.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func (br *carRepo) Delete(id string) error {
	tsx, err := br.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tsx.Rollback()
		} else {
			err := tsx.Commit()
			if err != nil {
				fmt.Println("While committing transaction ", err)
			}
		}
	}()

	query := `DELETE FROM car WHERE id=$1`
	_, err = tsx.Exec(query, id)

	if err != nil {
		return err
	}

	return err
}
