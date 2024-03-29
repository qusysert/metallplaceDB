package repository

import (
	"context"
	"fmt"
	"metallplace/internal/app/model"
	"metallplace/pkg/gopkg-db"
	"time"
)

// AddMaterialValue Adding value to certain property of a product for a certain date
func (r *Repository) AddMaterialValue(ctx context.Context, uid int, propertyName string, valueFloat float64, valueStr string, createdOn time.Time) error {
	propertyId, err := r.GetPropertyId(ctx, propertyName)
	if err != nil {
		return err
	}

	_, err = db.FromContext(ctx).Exec(ctx, `
				INSERT INTO material_value (uid, property_id, value_decimal, value_str, created_on)
					VALUES ($1, $2, $3, $4, $5) ON CONFLICT (uid, property_id, created_on) DO UPDATE SET value_decimal=EXCLUDED.value_decimal RETURNING id`, uid, propertyId, valueFloat, valueStr, createdOn)
	if err != nil {
		return fmt.Errorf("Can't add value %w", err)
	}

	return nil
}

// GetMaterialValueForPeriod Gets all price records ob property of material for given time period
func (r *Repository) GetMaterialValueForPeriod(ctx context.Context, uid, propertyId int, start string, finish string) ([]model.Price, float64, error) {
	var priceFeed []model.Price
	var price model.Price
	var prevPrice float64

	// Getting main price feed
	rows, err := db.FromContext(ctx).Query(ctx, `SELECT created_on, value_decimal
		FROM material_value WHERE uid=$1 AND property_id=$4 AND created_on >= $2 AND
			created_on <= $3 ORDER BY created_on ASC`, uid, start, finish, propertyId)
	if err != nil {
		return nil, 0, fmt.Errorf("Can't get material price %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&price.Date, &price.Value)
		if err != nil {
			return nil, 0, fmt.Errorf("Can't read from rows in get_prices %w", err)
		}
		priceFeed = append(priceFeed, price)
	}

	// Getting price prior to main feed (mainly used for calculating changes for earliest record of main feed)
	row := db.FromContext(ctx).QueryRow(ctx, `SELECT value_decimal
		FROM material_value WHERE uid=$1 AND property_id=$3 AND created_on < $2 ORDER BY created_on DESC LIMIT 1`, uid, start, propertyId)
	err = row.Scan(&prevPrice)
	if err != nil {
		prevPrice = 0
	}

	return priceFeed, prevPrice, nil
}
