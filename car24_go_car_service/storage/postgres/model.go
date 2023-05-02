package postgres

import (
	"database/sql"
	"fmt"

	cs "gitlab.udevs.io/car24/car24_go_car_service/modules/car24/car24_car_service"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/helper"
	"gitlab.udevs.io/car24/car24_go_car_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type modelRepo struct {
	db *sqlx.DB
}

func NewModel(db *sqlx.DB) repo.ModelI {
	return &modelRepo{
		db: db,
	}
}

func (br *modelRepo) Create(car *cs.CreateModel) (string, error) {
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
		model(
			id,
			name,
			brand_id,
			updated_at
		)
	VALUES ($1, $2, $3, now())`

	_, err = tsx.Exec(
		query,
		car.ID,
		car.Name,
		car.BrandID,
	)

	if err != nil {
		return "", err
	}

	return car.ID, nil
}

func (br *modelRepo) Get(id string) (*cs.Model, error) {
	var (
		name      sql.NullString
		brandId   sql.NullString
		updatedAt sql.NullString
		createdAt sql.NullString

		brandName      sql.NullString
		brandUpdatedAt sql.NullString
		brandCreatedAt sql.NullString
	)

	query := `
		SELECT
			m.name,
			m.brand_id,

			b.name,
			b.updated_at,
			b.created_at,

			m.updated_at,
			m.created_at
		FROM model AS m
		LEFT JOIN brand AS b ON m.brand_id = b.id
		WHERE m.id = $1
	`

	row := br.db.QueryRow(query, id)
	err := row.Scan(
		&name,
		&brandId,
		&brandName,
		&brandUpdatedAt,
		&brandCreatedAt,
		&updatedAt,
		&createdAt,
	)
	if err != nil {
		return nil, err
	}

	brand := &cs.Brand{
		ID:        brandId.String,
		Name:      brandName.String,
		UpdatedAt: brandUpdatedAt.String,
		CreatedAt: createdAt.String,
	}

	return &cs.Model{
		ID:        id,
		Name:      name.String,
		BrandID:   brandId.String,
		BrandData: brand,
		UpdatedAt: updatedAt.String,
		CreatedAt: createdAt.String,
	}, nil
}

func (br *modelRepo) GetAll(queryParam cs.ModelQueryParamModel) (res cs.ModelListModel, err error) {

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
			name,
			brand_id,
			updated_at,
			created_at
		FROM
			model
	`

	if len(queryParam.BrandId) > 0 {
		filter += " AND brand_id = '" + queryParam.BrandId + "' "
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
			id        sql.NullString
			name      sql.NullString
			brandId   sql.NullString
			updatedAt sql.NullString
			createdAt sql.NullString
		)

		err := rows.Scan(
			&res.Count,
			&id,
			&name,
			&brandId,
			&updatedAt,
			&createdAt,
		)
		if err != nil {
			return res, err
		}

		res.Models = append(res.Models, cs.Model{
			ID:        id.String,
			Name:      name.String,
			BrandID:   brandId.String,
			UpdatedAt: updatedAt.String,
			CreatedAt: createdAt.String,
		})
	}

	return res, nil
}

func (br *modelRepo) Update(car *cs.UpdateModel) (err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE 
				model
			SET
				name = :name,
				brand_id = :brand_id,
				updated_at = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":       car.ID,
		"name":     car.Name,
		"brand_id": car.BrandID,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	_, err = br.db.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func (br *modelRepo) Delete(id string) error {
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

	query := `UPDATE car SET model_id = NULL WHERE model_id=$1`
	_, err = tsx.Exec(query, id)

	if err != nil {
		return err
	}

	query = `DELETE FROM model WHERE id=$1`
	_, err = tsx.Exec(query, id)

	if err != nil {
		return err
	}

	return err
}
