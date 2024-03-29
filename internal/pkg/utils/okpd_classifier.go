package utils

import (
	"fmt"
	"metallplace/internal/app/model"
)

func OkpdUnitClassifier(codeStr string) (model.Unit, error) {
	unit, ok := unitClassification[codeStr]
	if !ok {
		return model.Unit{}, fmt.Errorf("unknown ОКПД unit id: %s", codeStr)
	}
	return unit, nil
}

var unitClassification = map[string]model.Unit{
	"168": {Name: "тонна", ValueMultiplication: 1},
	"169": {Name: "тонна", ValueMultiplication: 1000},
}
