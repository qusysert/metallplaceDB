package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"metallplace/internal/app/model"
	"metallplace/pkg/gopkg-db"
)

// AddMaterialSource Adding material - source - market - unit combo
func (r *Repository) AddMaterialSource(ctx context.Context, materialName, sourceName, market, unit string) (int, error) {
	materialId, err := r.GetMaterialId(ctx, materialName)
	var id int
	if err != nil {
		return 0, fmt.Errorf("Can't get material id %w", err)
	}

	sourceId, err := r.GetSourceId(ctx, sourceName)
	if err != nil {
		return 0, fmt.Errorf("Can't get source id %w", err)
	}

	id, err = r.GetMaterialSourceId(ctx, materialName, sourceName, market, unit)
	if err != nil {
		return 0, fmt.Errorf("cant get material-source id %w", err)
	}

	if id != 0 {
		return id, nil
	}

	row := db.FromContext(ctx).QueryRow(
		ctx, `INSERT INTO material_source (material_id, source_id, target_market, unit) 
		VALUES ($1, $2, $3, $4) 
		ON CONFLICT DO NOTHING RETURNING id`,
		materialId, sourceId, market, unit)

	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("cant add material-source: %w", err)
	}

	return id, nil
}

// GetMaterialSourceId Get unique material-source combo by material and source name
func (r *Repository) GetMaterialSourceId(ctx context.Context, materialName, sourceName, market, unit string) (int, error) {
	var id int
	materialId, err := r.GetMaterialId(ctx, materialName)
	if err != nil {
		return 0, fmt.Errorf("Can't get material id %w", err)
	}

	sourceId, err := r.GetSourceId(ctx, sourceName)
	if err != nil {
		return 0, fmt.Errorf("Can't get source id %w", err)
	}

	row := db.FromContext(ctx).QueryRow(ctx, `SELECT id FROM material_source WHERE material_id=$1 AND
		source_id=$2 AND target_market=$3 AND unit=$4`, materialId, sourceId, market, unit)

	err = row.Scan(&id)
	if err == pgx.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, fmt.Errorf("Can't get material-source pair id %w", err)
	}

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) GetMaterialSource(ctx context.Context, id int) (model.MaterialShortInfo, error) {
	var materialId int
	var sourceId int
	var market string
	var unit string

	row := db.FromContext(ctx).QueryRow(ctx, `SELECT material_id, source_id, target_market, unit 
		FROM material_source WHERE id=$1`, id)

	err := row.Scan(&materialId, &sourceId, &market, &unit)
	if err != nil {
		return model.MaterialShortInfo{}, fmt.Errorf("can't get scan row %w", err)
	}

	materialName, err := r.GetMaterialName(ctx, materialId)
	if err != nil {
		return model.MaterialShortInfo{}, fmt.Errorf("can't get material name %w", err)
	}

	sourceName, err := r.GetSourceName(ctx, sourceId)
	if err != nil {
		return model.MaterialShortInfo{}, fmt.Errorf("can't get source name %w", err)
	}

	return model.MaterialShortInfo{Id: id, Name: materialName, Source: sourceName, Market: market, Unit: unit}, nil
}
