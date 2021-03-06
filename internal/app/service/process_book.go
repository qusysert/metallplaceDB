package service

import (
	"context"
	"fmt"
	"github.com/xuri/excelize/v2"
	"metallplace/internal/app/model"
	"strconv"
	"time"
)

// InitialImport Importing data from book, using layout written by hand
func (s *Service) InitialImport(ctx context.Context) error {
	dateLayout := "2-Jan-06"

	book, err := excelize.OpenFile("var/analytics.xlsx")
	//book, err := excelize.OpenFile("var/testEx.xlsx")
	if err != nil {
		return fmt.Errorf("cannot open exel file %w", err)
	}

	materials := model.InitMaterials

	//err = db.ExecTx(ctx, func(ctx context.Context) error {

	// Going through init_materials list layout
	for _, material := range materials {
		materialId, err := s.AddUniqueMaterial(ctx, material.Name, material.Source, material.Market, material.Unit)
		if err != nil {
			return err
		}

		fmt.Println("Adding material " + material.Name)

		// Adding and tying properties
		for _, property := range material.Properties {
			propertyId, err := s.repo.AddPropertyIfNotExists(ctx, model.PropertyShortInfo{Name: property.Name, Kind: property.Kind})
			if err != nil {
				return err
			}

			err = s.repo.AddMaterialProperty(ctx, materialId, propertyId)
		}

		// Going through material's properties, and reading property values
		for _, property := range material.Properties {
			fmt.Println(property.Name)
			row := property.Row
			for {
				value, err := book.GetCellValue(material.Sheet, property.Column+strconv.Itoa(row))
				if err != nil {
					return err
				}

				if value == "" {
					break
				}

				// Calculating date cell, and formatting it
				dateCell := material.DateColumn + strconv.Itoa(row)
				style, _ := book.NewStyle(`{"number_format":15}`)
				book.SetCellStyle(material.Sheet, dateCell, dateCell, style)

				dateStr, err := book.GetCellValue(material.Sheet, dateCell)
				if err != nil {
					return err
				}
				dateType, err := book.GetCellType(material.Sheet, dateCell)
				if err != nil {
					return err
				}

				// Parsing date
				createdOn, err := time.Parse(dateLayout, dateStr)
				if err != nil {
					return fmt.Errorf("Can't parce date [%v,%v] '%v' (%v): %w", material.Sheet, dateCell, dateStr, dateType, err)
				}

				// Checking type of value: string or decimal
				var valueStr string
				var valueDecimal float64
				if property.Kind == "decimal" {
					valueDecimal, err = strconv.ParseFloat(value, 64)
					if err != nil {
						return err
					}
				} else {
					valueStr = value
				}

				err = s.repo.AddMaterialSource(ctx, material.Name, material.Source, material.Market, material.Unit)
				if err != nil {
					return fmt.Errorf("cann not add material source id: %v", err)
				}

				materialSourceId, err := s.repo.GetMaterialSourceId(ctx, material.Name, material.Source, material.Market, material.Unit)
				if err != nil {
					return fmt.Errorf("cann not get material source id: %v", err)
				}

				err = s.repo.AddMaterialValue(ctx, materialSourceId, property.Name, valueDecimal, valueStr, createdOn)
				if err != nil {
					return err
				}

				row++
				if row%100 == 0 {
					fmt.Print("#")
				}
			}
			fmt.Println("")
		}
	}

	//	return nil
	//})

	if err != nil {
		fmt.Println("cant exec init import tx: %v", err)
	}

	fmt.Print("Import finished!")
	return nil
}
