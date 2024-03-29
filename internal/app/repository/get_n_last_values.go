package repository

import (
	"context"
	"fmt"
	"metallplace/internal/app/model"
	"metallplace/pkg/gopkg-db"
)

// GetNLastValues Get n values prior to some date
func (r *Repository) GetNLastValues(ctx context.Context, uid, propertyId int, nValues int, finish string) ([]model.Price, error) {
	var priceFeed []model.Price
	var price model.Price

	rows, err := db.FromContext(ctx).Query(ctx,
		"SELECT * FROM "+
			"("+
			"SELECT created_on, value_decimal "+
			"FROM material_value "+
			"WHERE uid=$1 AND property_id=$2 AND created_on <= $4"+
			"ORDER BY created_on "+
			"DESC LIMIT $3 "+
			") AS last_n_values "+
			"ORDER BY created_on ASC", uid, propertyId, nValues, finish)

	if err != nil {
		return nil, fmt.Errorf("Can't get n last material prices %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&price.Date, &price.Value)
		if err != nil {
			return nil, fmt.Errorf("Can't read from rows in get_prices %w", err)
		}
		priceFeed = append(priceFeed, price)
	}

	return priceFeed, nil
}
