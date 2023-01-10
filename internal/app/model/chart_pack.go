package model

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const ChartRoutePostfix = ".png"

type ChartPack struct {
	MaterialIdList []int
	PropertyId     int
	Start          time.Time
	Finish         time.Time
	NeedLabels     bool
	Type           string
	Scale          string
	XStep          string
	NeedLegend     bool
}

func (c ChartPack) ToUrl() string {
	var fn string

	for i := 0; i < len(c.MaterialIdList); i++ {
		fn += strconv.Itoa(c.MaterialIdList[i])
		if i != len(c.MaterialIdList)-1 {
			fn += "-"
		}
	}

	fn += "_" + strconv.Itoa(c.PropertyId)
	fn += "_" + c.Start.Format("2006-01-02") + c.Finish.Format("2006-01-02")
	needLabels := "0"
	needLegend := "0"
	if c.NeedLabels {
		needLabels = "1"
	}
	if c.NeedLegend {
		needLegend = "1"
	}

	fn += "_" + needLabels
	fn += "_" + c.Type
	fn += "_" + c.Scale
	fn += "_" + c.XStep
	fn += "_" + needLegend

	return fn + ChartRoutePostfix
}

func NewChartPack(url string) (ChartPack, error) {
	fn := strings.TrimRight(url, ChartRoutePostfix)
	var c ChartPack
	cnt := strings.Split(fn, "_")

	var MaterialIdList []int
	idStrList := strings.Split(cnt[0], "-")
	for _, id := range idStrList {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return ChartPack{}, fmt.Errorf("cant parse material id: %w", err)
		}
		MaterialIdList = append(MaterialIdList, idInt)
	}

	c.MaterialIdList = MaterialIdList

	propertyId, err := strconv.Atoi(cnt[1])
	if err != nil {
		return ChartPack{}, fmt.Errorf("cant parse property id: %w", err)
	}
	c.PropertyId = propertyId

	start, err := time.Parse("01-02-2006", cnt[2])
	if err != nil {
		return ChartPack{}, fmt.Errorf("cant parse start time: %w", err)
	}
	c.Start = start

	finish, err := time.Parse("01-02-2006", cnt[3])
	if err != nil {
		return ChartPack{}, fmt.Errorf("cant parse finish time: %w", err)
	}
	c.Finish = finish

	if cnt[4] == "1" {
		c.NeedLabels = true
	}

	c.Type = cnt[5]

	c.Scale = cnt[6]

	c.XStep = cnt[7]

	if cnt[8] == "1" {
		c.NeedLegend = true
	}

	return c, nil
}
