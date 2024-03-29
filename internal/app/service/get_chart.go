package service

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"metallplace/internal/app/model"
	"metallplace/pkg/chartclient"
	"os"
	"time"
)

func (s *Service) GetChart(ctx context.Context, chartPack model.ChartPack) ([]byte, error) {
	var req chartclient.Request
	var isFirst = true

	for _, id := range chartPack.MaterialIdList {
		material, err := s.repo.GetMaterialSource(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("cant get material_source: %w", err)
		}

		dataset := chartclient.YDataSet{Label: material.Name, Data: []float64{}}
		start := chartPack.Start.Format("2006-01-02")
		finish := chartPack.Finish.Format("2006-01-02")
		var feed []model.Price

		switch chartPack.Scale {
		case "day":
			feed, _, err = s.repo.GetMaterialValueForPeriod(ctx, id, chartPack.PropertyId, start, finish)
			if err != nil {
				return nil, fmt.Errorf("cant get material_value: %w", err)
			}
		case "month":
			feed, _, err = s.GetMonthlyAvgFeed(ctx, id, chartPack.PropertyId, start, finish)
			if err != nil {
				return nil, fmt.Errorf("cant get material_value: %w", err)
			}
			// Appending 3 predict values to the feed
			if chartPack.Predict {
				// Calculating predict values dates
				predictStartTime := chartPack.Finish.AddDate(0, 1, 0)
				predictStart := time.Date(predictStartTime.Year(), predictStartTime.Month(), 1, 0, 0, 0, 0, predictStartTime.Location()).Format("2006-01-02")
				predictFinish := chartPack.Finish.AddDate(0, 3, 0).Format("2006-01-02")
				predictPropertyId, err := s.GetPropertyId(ctx, "Прогноз месяц")
				if err != nil {
					return nil, fmt.Errorf("cant get property_id: %w", err)
				}
				// Getting predict values
				predictFeed, _, err := s.GetMaterialValueForPeriod(ctx, id, predictPropertyId, predictStart, predictFinish)
				if err != nil {
					return nil, fmt.Errorf("cant get predict_feed: %w", err)
				}
				feed = append(feed, predictFeed...)
				// Calculating predict accuracy
				lastPrice, _, err := s.GetMonthlyAvgFeed(ctx, id, chartPack.PropertyId, finish, finish)
				if err != nil || len(lastPrice) == 0 {
					return nil, fmt.Errorf("cannot get last price of period for calculatin predict accuracy: %w", err)
				}
				lastPricePredict, _, err := s.GetMaterialValueForPeriod(ctx, id, predictPropertyId, finish, finish)
				dataset.PredictAccuracy = math.Round(100 - (math.Abs(lastPrice[0].Value-lastPricePredict[0].Value)/lastPrice[0].Value)*100)
			}
		case "week":
			feed, _, err = s.GetWeeklyAvgFeed(ctx, id, chartPack.PropertyId, start, finish)
			if err != nil {
				return nil, fmt.Errorf("cant get material_value: %w", err)
			}
			// Apparently not used in any charts in reports yet? Still thought it would be good to implement
			// Same logic as in month
			if chartPack.Predict {
				lastPrice, _, err := s.GetWeeklyAvgFeed(ctx, id, chartPack.PropertyId, finish, finish)
				if err != nil || len(lastPrice) == 0 {
					return nil, fmt.Errorf("cannot get last price of period for calculatin predict accuracy: %w", err)
				}
				predictStart := chartPack.Finish.AddDate(0, 0, 7).Format("2006-01-02")
				predictFinish := chartPack.Finish.AddDate(0, 3, 21).Format("2006-01-02")
				predictPropertyId, err := s.GetPropertyId(ctx, "Прогноз неделя")
				if err != nil {
					return nil, fmt.Errorf("cant get property_id: %w", err)
				}
				predictFeed, _, err := s.GetMaterialValueForPeriod(ctx, id, predictPropertyId, predictStart, predictFinish)
				if err != nil {
					return nil, fmt.Errorf("cant get predict_feed: %w", err)
				}
				feed = append(feed, predictFeed...)
				lastPricePredict, _, err := s.GetMaterialValueForPeriod(ctx, id, predictPropertyId, finish, finish)
				dataset.PredictAccuracy = math.Round(float64(100) - (math.Abs(lastPrice[0].Value-lastPricePredict[0].Value)/lastPrice[0].Value)*float64(100))
			}
		default:
			return nil, fmt.Errorf("wrong data averaging scale type: %w", err)
		}

		for _, item := range feed {
			dataset.Data = append(dataset.Data, item.Value)
			if isFirst {
				req.XLabelSet = append(req.XLabelSet, item.Date.Format("2006-01-02"))
			}
		}
		req.YDataSet = append(req.YDataSet, dataset)
		req.Options.NeedLabels = chartPack.NeedLabels
		req.Options.Type = chartPack.Type
		req.Options.XStep = chartPack.XStep
		// for now hard coding it, made more for raw chart gen, here it is usually predefined
		req.Options.TickLimit = 0
		req.Options.NeedLegend = chartPack.NeedLegend
		req.Options.ToFixed = chartPack.ToFixed
		req.Options.Predict = chartPack.Predict
		req.Options.Tall = chartPack.Tall
		isFirst = false
	}

	bytes, err := s.chart.GetChart(req)
	if err != nil {
		return nil, fmt.Errorf("cant get chart bytes: %w", err)
	}
	return bytes, nil
}

// GetChartRaw For custom report charts
func (s *Service) GetChartRaw(book []byte, tickLimit int) ([]byte, error) {
	req, err := s.ParseXlsxForChart(book)
	if err != nil {
		return nil, fmt.Errorf("cant parse book: %w", err)
	}
	req.Options.NeedLegend = true
	req.Options.Tall = true

	bytes, err := s.chart.GetChart(req)
	if err != nil {
		return nil, fmt.Errorf("cant get raw chart bytes: %w", err)
	}
	return bytes, nil
}

func (s *Service) GetCachedChart(ctx context.Context, chartPack model.ChartPack) ([]byte, error) {
	path := "./var/cache/charts/" + chartPack.ToUrl()
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		bytes, err := s.GetChart(ctx, chartPack)
		if err != nil {
			return nil, fmt.Errorf("cant get generated chart from chart_service: %w", err)
		}

		f, err := os.Create(path)
		if err != nil {
			return nil, fmt.Errorf("cant create file for chart: %w", err)
		}

		_, err = f.Write(bytes)
		if err != nil {
			return nil, fmt.Errorf("cant cant write chart to file: %w", err)
		}

		f.Close()

		return bytes, nil
	}
	return ioutil.ReadFile(path)
}
