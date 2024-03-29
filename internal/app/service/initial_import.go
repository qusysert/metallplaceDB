package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/xuri/excelize/v2"
	"math"
	"metallplace/internal/app/model"
	"metallplace/internal/pkg/utils"
	"metallplace/pkg/chartclient"
	db "metallplace/pkg/gopkg-db"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (s *Service) InitialImportDaily(ctx context.Context) error {
	dateLayout := "2-Jan-06"
	book, err := excelize.OpenFile("var/analytics.xlsx")
	if err != nil {
		return fmt.Errorf("cannot open exel file %w", err)
	}

	err = db.ExecTx(ctx, func(ctx context.Context) error {
		if err := s.InitImportDailyMaterials(ctx, book, dateLayout); err != nil {
			return fmt.Errorf("error initializing daily import: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("cant exec init import tx: %w", err)
	}

	fmt.Print("Import finished!")
	return nil
}

// InitialImport Importing data from book, using layout written by hand
func (s *Service) InitialImport(ctx context.Context) error {
	dateLayout := "2-Jan-06"
	book, err := excelize.OpenFile("var/analytics.xlsx")
	if err != nil {
		return fmt.Errorf("cannot open exel file %w", err)
	}

	err = db.ExecTx(ctx, func(ctx context.Context) error {
		if err := s.InitImportMaterialsVertical(ctx, book, dateLayout); err != nil {
			return fmt.Errorf("error initializing vertical import: %w", err)
		}
		if err := s.InitImportMaterialsHorizontalWeekly(ctx, book); err != nil {
			return fmt.Errorf("error initializing weekly horizontal import: %w", err)
		}
		// Currently disabled
		//if err := s.ImportRosStat(ctx); err != nil {
		//	return fmt.Errorf("can't import ros stat: %w", err)
		//}

		return nil
	})
	if err != nil {
		return fmt.Errorf("cant exec init import tx: %w", err)
	}

	fmt.Print("Import finished!")
	return nil

}

// ParseXlsxForChart Used for short reports
func (s *Service) ParseXlsxForChart(byte []byte) (chartclient.Request, error) {
	reader := bytes.NewReader(byte)
	book, err := excelize.OpenReader(reader)
	if err != nil {
		return chartclient.Request{}, fmt.Errorf("cannot open exel file %w", err)
	}
	labelColumn := "A"
	materialStartColumn := "B"
	startRow := 2
	startSheet := "Лист1"
	var req chartclient.Request

	for curCol, err := utils.AlphabetToInt(materialStartColumn); true; curCol++ {
		if err != nil {
			return chartclient.Request{}, fmt.Errorf("cant parse xlsx for chart: %w", err)
		}
		var valueFloat float64
		curRow := startRow
		var curDate string
		var materialAndPrices chartclient.YDataSet
		// Reading material name
		col, err := utils.IntToAlphabet(int32(curCol))
		if err != nil {
			return chartclient.Request{}, fmt.Errorf("cant parse excel format columb number: %w", err)
		}
		value, err := book.GetCellValue(startSheet, col+strconv.Itoa(startRow))
		if err != nil {
			return chartclient.Request{}, fmt.Errorf("cant get cell value: %w", err)
		}
		value = strings.TrimSpace(value)
		if value == "" {
			break
		}
		materialAndPrices.Label = value
		wereAnyValues := false

		for row := curRow + 1; true; row++ {
			// Reading prices
			col, err := utils.IntToAlphabet(int32(curCol))
			if err != nil {
				return chartclient.Request{}, fmt.Errorf("cant parse excel format columb number: %w", err)
			}
			value, err = book.GetCellValue(startSheet, col+strconv.Itoa(row))
			if err != nil {
				return chartclient.Request{}, fmt.Errorf("cant get cell value col %d, row %d: %w", row, curCol, err)
			}
			if value == "" {
				isBreak, err := areNextCellsEmpty(book, startSheet, curCol, row, 10)
				if err != nil {
					return chartclient.Request{}, fmt.Errorf("cant check next n values: %w", err)
				}
				if isBreak {
					break
				}
				if !wereAnyValues {
					valueFloat = -1
				}
			} else {
				wereAnyValues = true
				valueFloat, err = strconv.ParseFloat(strings.Join(strings.Fields(value), ""), 64)
				if err != nil {
					return chartclient.Request{}, fmt.Errorf("cant convert string to float: %w", err)
				}
			}
			materialAndPrices.Data = append(materialAndPrices.Data, math.Round(valueFloat*100)/100)

			// Reading labels
			colNumber, err := utils.AlphabetToInt(materialStartColumn)
			if err != nil {
				return chartclient.Request{}, fmt.Errorf("cant get col nuber: %w", err)
			}
			if curCol == colNumber {
				value, err = book.GetCellValue(startSheet, labelColumn+strconv.Itoa(row))
				if err != nil {
					return chartclient.Request{}, fmt.Errorf("cant get cell value: %w", err)
				}
				if value != "" {
					curDate = value
				}
				curLabel := formatMonth(curDate)
				req.XLabelSet = append(req.XLabelSet, curLabel)
			}
		}
		req.YDataSet = append(req.YDataSet, materialAndPrices)
	}
	req.XLabelSet = removeDuplicateLabels(req.XLabelSet)
	title, err := book.GetCellValue(startSheet, "A1")
	if err != nil {
		return chartclient.Request{}, fmt.Errorf("cant get chart title: %w", err)
	}
	req.Options.Title = title
	return req, nil
}

//func (s *Service) ImportRosStat(ctx context.Context) error {
//	fmt.Println("importing rosstat")
//	directory := "var/ros_stat_books"
//	absDir, err := filepath.Abs(directory)
//	if err != nil {
//		return fmt.Errorf("cant get absolute path: %w", err)
//	}
//	err = filepath.Walk(absDir, func(path string, info os.FileInfo, err error) error {
//		if err != nil {
//			return err
//		}
//		pathArr := strings.Split(path, string(os.PathSeparator))
//		if pathArr[len(pathArr)-1] == ".gitkeep" {
//			return nil // skip .gitkeep files
//		}
//		if !info.IsDir() {
//			// Read file to byte array
//			data, err := ioutil.ReadFile(path)
//			if err != nil {
//				return err
//			}
//			// Pass byte array to function
//			err = s.ParseRosStatBook(ctx, data)
//			if err != nil {
//				return fmt.Errorf("error in importing ros stat file `%s` %w", path, err)
//			}
//		}
//		return nil
//	})
//	if err != nil {
//		return fmt.Errorf("error going through directory: %w", err)
//	}
//	return nil
//}

//func (s *Service) ParseRosStatBook(ctx context.Context, byte []byte) error {
//	reader := bytes.NewReader(byte)
//	book, err := excelize.OpenReader(reader)
//	if err != nil {
//		return fmt.Errorf("cannot open exel file %w", err)
//	}
//
//	// Fetching month and year of report from the title
//	title, err := book.GetCellValue("Содержание", "B2")
//	if err != nil {
//		return fmt.Errorf("cannot get title: %w", err)
//	}
//	titleArr := strings.Split(title, " ")
//	if len(titleArr) < 7 {
//		title, err := book.GetCellValue("Содержание", "B3")
//		if err != nil {
//			return fmt.Errorf("cannot get title: %w", err)
//		}
//		titleArr = strings.Split(title, " ")
//		if len(titleArr) < 7 {
//			return fmt.Errorf("cant get title while parsing rosstat")
//		}
//	}
//	year, err := strconv.Atoi(titleArr[len(titleArr)-2])
//	if err != nil {
//		return fmt.Errorf("cannot convert year: %w", err)
//	}
//	month, err := monthStrToNumber(titleArr[len(titleArr)-3])
//	if err != nil {
//		return fmt.Errorf("cannot convert month: %w", err)
//	}
//	date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
//
//	// Getting id of volume property
//
//	volumePropertyId, err := s.AddPropertyIfNotExists(ctx, model.PropertyShortInfo{Name: "Запас", Kind: "decimal"})
//	if err != nil {
//		return fmt.Errorf("cannot get volume id: %w", err)
//	}
//
//	for _, coord := range model.RosStatCoordinates {
//		name, row, err := findInRowRange(book, coord.Sheet, "A", coord.Row, 5, coord.Material)
//		if err != nil {
//			if date.Year() == 2021 && coord.Material == "Чугун литейный" {
//				continue
//			}
//			return fmt.Errorf("cannot get name: %w", err)
//		}
//
//		code, err := book.GetCellValue(coord.Sheet, "B"+strconv.Itoa(row))
//		if err != nil {
//			return fmt.Errorf("cannot get code: %w", err)
//		}
//
//		location, err := book.GetCellValue(coord.Sheet, "A"+strconv.Itoa(row+2))
//		if err != nil {
//			return fmt.Errorf("cannot get location: %w", err)
//		}
//
//		unitCode, err := book.GetCellValue(coord.Sheet, "B"+strconv.Itoa(row+1))
//		if err != nil {
//			return fmt.Errorf("cannot get unit: %w", err)
//		}
//
//		unit, err := utils.OkpdUnitClassifier(unitCode)
//		if err != nil {
//			return fmt.Errorf("cannot convert unit ОКПД id to int: %w", err)
//		}
//
//		fmt.Println(name + ", " + code)
//		err := s.AddUniqueMaterial(ctx, name+", "+code, "", "rosstat.gov.ru", location, unit.Name, "")
//		if err != nil {
//			return fmt.Errorf("cannot add material %s: %w", name, err)
//		}
//
//		err = s.repo.AddMaterialProperty(ctx, materialSourceId, volumePropertyId)
//		if err != nil {
//			return fmt.Errorf("failed to add property: %w", err)
//		}
//
//		volume, err := book.GetCellValue(coord.Sheet, "C"+strconv.Itoa(row+2))
//		if err != nil {
//			return fmt.Errorf("cannot get volume: %w", err)
//		}
//
//		volumeFloat, err := strconv.ParseFloat(volume, 64)
//		if err != nil {
//			return fmt.Errorf("cannot convert volume: %w", err)
//		}
//
//		propertyName, err := s.GetPropertyName(ctx, volumePropertyId)
//		if err != nil {
//			return fmt.Errorf("cannot get property name: %w", err)
//		}
//
//		err = s.repo.AddMaterialValue(ctx, materialSourceId, propertyName, volumeFloat*unit.ValueMultiplication, "", date)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

func (s *Service) InitImportDailyMaterials(ctx context.Context, book *excelize.File, dateLayout string) error {
	for _, material := range model.InitDaily {
		id, err := s.AddUniqueMaterial(ctx, material.UId, material.Name, material.Group, material.Source, material.Market, material.Unit, material.DeliveryType)
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

			err = s.repo.AddMaterialProperty(ctx, id, propertyId)
			if err != nil {
				return fmt.Errorf("failed to add property %s: %w", property.Name, err)
			}
		}

		// Going through material's properties, and reading property values
		for _, property := range material.Properties {
			fmt.Println(property.Name)
			row := property.Row
			for {
				var value string
				valueCellValue, err := book.GetCellValue(property.Sheet, property.Column+strconv.Itoa(row))
				if err != nil {
					return fmt.Errorf("cant parce with GetCellValue %s:%s%d: %v", property.Sheet, property.Column, row, err)
				}
				valueCalc, err := book.CalcCellValue(property.Sheet, property.Column+strconv.Itoa(row))

				if err != nil {
					if err.Error() == "AVERAGE divide by zero" {
						break
					}
					return fmt.Errorf("cant parce with CalcCellValue %s:%s%d: %v", property.Sheet, property.Column, row, err)
				}
				valueCellValue = strings.TrimSpace(valueCellValue)
				valueCalc = strings.TrimSpace(valueCalc)

				if valueCellValue == "" && valueCalc == "" {
					break
				} else if valueCellValue != "" {
					value = valueCellValue
				} else if valueCalc != "" {
					value = valueCalc
				}

				// Calculating date cell, and formatting it
				var dateCell string
				if property.DateColumn != "" {
					dateCell = property.DateColumn + strconv.Itoa(row)
				} else {
					dateCell = material.DateColumn + strconv.Itoa(row)
				}
				style, _ := book.NewStyle(`{"number_format":15}`)
				err = book.SetCellStyle(property.Sheet, dateCell, dateCell, style)
				if err != nil {
					return fmt.Errorf("cant set cell style: %w", err)
				}

				dateStr, err := book.GetCellValue(property.Sheet, dateCell)
				if err != nil {
					return fmt.Errorf("cant parce dateStr %s:%s%d: %v", property.Sheet, property.Column, row, err)
				}
				dateType, err := book.GetCellType(property.Sheet, dateCell)
				if err != nil {
					return err
				}

				// Parsing date
				createdOn, err := time.Parse(dateLayout, dateStr)
				if err != nil {
					return fmt.Errorf("cant parce date [%v,%v] '%v' (%v): %w", property.Sheet, dateCell, dateStr, dateType, err)
				}

				// Checking type of value: string or decimal
				var valueStr string
				var valueDecimal float64
				if property.Kind == "decimal" {
					valueDecimal, err = strconv.ParseFloat(strings.Join(strings.Fields(value), ""), 64)
					if err != nil {
						return fmt.Errorf("cant parce float price at %s %s%d: %v", property.Sheet, property.Column, row, err)
					}
					if material.ConvSettings.Need {
						valueDecimal = material.ConvSettings.Func(valueDecimal, material.ConvSettings.Rate)
					}
				} else {
					valueStr = value
				}

				err = s.repo.AddMaterialValue(ctx, id, property.Name, valueDecimal, valueStr, createdOn)
				if err != nil {
					return fmt.Errorf("failed to add material value: %w", err)
				}

				if material.NeedSplit {
					formula, err := book.GetCellFormula(property.Sheet, property.Column+strconv.Itoa(row))
					if err != nil {
						return fmt.Errorf("failed to get formyla in splitting for daily init import: %v", err)
					}

					regex := regexp.MustCompile(`\((\d+)\+(\d+)\)/2`)
					if regex.MatchString(formula) {
						matches := regex.FindStringSubmatch(formula)
						minPrice, err := strconv.ParseFloat(strings.Join(strings.Fields(matches[1]), ""), 64)
						if err != nil {
							return fmt.Errorf("failed to parce from formula: %v (uid; %d, property: %s, row: %d)", err, material.UId, property.Name, row)
						}
						maxPrice, err := strconv.ParseFloat(strings.Join(strings.Fields(matches[2]), ""), 64)
						if err != nil {
							return fmt.Errorf("failed to parce from formula: %v (uid; %d, property: %s, row: %d)", err, id, property.Name, row)
						}
						if material.ConvSettings.Need {
							minPrice = material.ConvSettings.Func(minPrice, material.ConvSettings.Rate)
							maxPrice = material.ConvSettings.Func(maxPrice, material.ConvSettings.Rate)
						}

						propertyId, err := s.repo.AddPropertyIfNotExists(ctx, model.PropertyShortInfo{Name: "Мин цена", Kind: "decimal"})
						if err != nil {
							return err
						}
						err = s.repo.AddMaterialProperty(ctx, id, propertyId)
						if err != nil {
							return fmt.Errorf("failed to add property %s: %w", property.Name, err)
						}
						propertyId, err = s.repo.AddPropertyIfNotExists(ctx, model.PropertyShortInfo{Name: "Макс цена", Kind: "decimal"})
						if err != nil {
							return err
						}
						err = s.repo.AddMaterialProperty(ctx, id, propertyId)
						if err != nil {
							return fmt.Errorf("failed to add property %s: %w", property.Name, err)
						}

						err = s.repo.AddMaterialValue(ctx, id, "Мин цена", minPrice, valueStr, createdOn)
						if err != nil {
							return fmt.Errorf("failed to add material value: %w", err)
						}
						err = s.repo.AddMaterialValue(ctx, id, "Макс цена", maxPrice, valueStr, createdOn)
						if err != nil {
							return fmt.Errorf("failed to add material value: %w", err)
						}
					}
				}

				row++
				if row%100 == 0 {
					fmt.Print("#")
				}
			}
			fmt.Println("")
		}
	}
	return nil
}

func (s *Service) InitImportMaterialsVertical(ctx context.Context, book *excelize.File, dateLayout string) error {
	for _, material := range model.InitMaterialsVertical {
		id, err := s.AddUniqueMaterial(ctx, material.UId, material.Name, material.Group, material.Source, material.Market, material.Unit, material.DeliveryType)
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

			err = s.repo.AddMaterialProperty(ctx, id, propertyId)
			if err != nil {
				return fmt.Errorf("failed to add property %s: %w", property.Name, err)
			}
		}

		// Going through material's properties, and reading property values
		for _, property := range material.Properties {
			fmt.Println(property.Name)
			row := property.Row
			for {
				var value string
				valueCellValue, err := book.GetCellValue(property.Sheet, property.Column+strconv.Itoa(row))
				if err != nil {
					return err
				}
				valueCalc, err := book.CalcCellValue(property.Sheet, property.Column+strconv.Itoa(row))
				if err != nil {
					return err
				}

				valueCellValue = strings.TrimSpace(valueCellValue)
				valueCalc = strings.TrimSpace(valueCalc)

				if valueCellValue == "" && valueCalc == "" {
					break
				} else if valueCalc != "" {
					value = valueCalc
				} else if valueCellValue != "" {
					value = valueCellValue
				}

				// Calculating date cell, and formatting it
				var dateCell string
				if property.DateColumn != "" {
					dateCell = property.DateColumn + strconv.Itoa(row)
				} else {
					dateCell = material.DateColumn + strconv.Itoa(row)
				}
				style, err := book.NewStyle(`{"number_format":15}`)
				if err != nil {
					return fmt.Errorf("cant create date style: %w", err)
				}
				err = book.SetCellStyle(property.Sheet, dateCell, dateCell, style)
				if err != nil {
					return fmt.Errorf("cant set cell style: %w", err)
				}

				dateStr, err := book.GetCellValue(property.Sheet, dateCell)
				if err != nil {
					return err
				}
				dateType, err := book.GetCellType(property.Sheet, dateCell)
				if err != nil {
					return err
				}

				// Parsing date
				createdOn, errFull := time.Parse(dateLayout, dateStr)
				var errMonth error
				var errWeek error
				if errFull != nil {
					createdOn, errMonth = getMonthDateForPredict(dateStr)
					if errMonth != nil {
						createdOn, errWeek = stringToDate(dateStr, "weekNum")
						if errWeek != nil {
							return fmt.Errorf("can't parce date [%v,%v] '%v' (%v): %w-%w-%w", property.Sheet, dateCell, dateStr, dateType, errFull, errMonth, errWeek)
						}
					}
				}

				// Checking type of value: string or decimal
				var valueStr string
				var valueDecimal float64
				if property.Kind == "decimal" {
					valueDecimal, err = strconv.ParseFloat(strings.Join(strings.Fields(value), ""), 64)
					if err != nil {
						return err
					}
				} else {
					valueStr = value
				}

				err = s.repo.AddMaterialValue(ctx, id, property.Name, valueDecimal, valueStr, createdOn)
				if err != nil {
					return fmt.Errorf("failed to add material value: %w", err)
				}

				row++
				if row%100 == 0 {
					fmt.Print("#")
				}
			}
			fmt.Println("")
		}
	}
	return nil
}

func (s *Service) InitImportMaterialsHorizontalWeekly(ctx context.Context, book *excelize.File) error {
	for _, material := range model.InitMaterialsHorizontalWeekly {
		id, err := s.AddUniqueMaterial(ctx, material.UId, material.Name, material.Group, material.Source, material.Market, material.Unit, material.DeliveryType)
		if err != nil {
			return fmt.Errorf("cant add unique material %v: %w", material.Name, err)
		}

		fmt.Println("Adding material " + material.Name)

		// Adding and tying properties
		for _, property := range material.Properties {
			propertyId, err := s.repo.AddPropertyIfNotExists(ctx, model.PropertyShortInfo{Name: property.Name, Kind: property.Kind})
			if err != nil {
				return fmt.Errorf("cant add/get property %v: %w", property.Name, err)
			}

			err = s.repo.AddMaterialProperty(ctx, id, propertyId)
			if err != nil {
				return fmt.Errorf("cant add material_property %v-%v: %w", material.Name, property.Name, err)
			}
		}
		for _, property := range material.Properties {
			fmt.Println(property.Name)
			col, err := utils.AlphabetToInt(property.Column)
			if err != nil {
				return fmt.Errorf("cant get int from letter: %w", err)
			}
			for {
				c, err := utils.IntToAlphabet(int32(col))
				if err != nil {
					return fmt.Errorf("cant parse excel format columb number: %w", err)
				}
				value, err := book.CalcCellValue(property.Sheet, c+strconv.Itoa(property.Row))
				value = strings.TrimSpace(value)
				if err != nil {
					return fmt.Errorf("cant calc cell %s %s%d value: %w", property.Sheet, property.Column, property.Row, err)
				}

				if value == "" {
					c, err := utils.IntToAlphabet(int32(col))
					if err != nil {
						return fmt.Errorf("cant parse excel format columb number: %w", err)
					}
					value, err = book.GetCellValue(property.Sheet, c+strconv.Itoa(property.Row))
					if err != nil {
						return fmt.Errorf("cant get cell value: %w", err)
					}
					if value == "" {
						break
					}
				}
				c, err = utils.IntToAlphabet(int32(col))
				if err != nil {
					return fmt.Errorf("cant parse excel format columb number: %w", err)
				}
				dateCell := c + material.DateRow
				dateStr, err := book.GetCellValue(property.Sheet, dateCell)
				if err != nil {
					return fmt.Errorf("cant get cell value: %w", err)
				}
				createdOn, err := stringToDate(dateStr, "week")
				if err != nil {
					return fmt.Errorf("Can't parce date [%v,%v]: %w", property.Sheet, dateCell, err)
				}

				// Checking type of value: string or decimal
				var valueStr string
				var valueDecimal float64
				if property.Kind == "decimal" {
					valueDecimal, err = strconv.ParseFloat(strings.ReplaceAll(value, " ", ""), 64)
					if err != nil {
						return err
					}
				} else {
					valueStr = value
				}

				err = s.repo.AddMaterialValue(ctx, id, property.Name, valueDecimal, valueStr, createdOn)
				if err != nil {
					return err
				}

				// After some time interval changed
				colNum, err := utils.AlphabetToInt("GX")
				if err != nil {
					return fmt.Errorf("cant get int from letter: %w", err)
				}
				if col >= colNum {
					col += 5
				} else {
					col += 4
				}

				if col%100 == 0 {
					fmt.Print("#")
				}
			}
		}
	}
	return nil
}
