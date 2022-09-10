package validator

import "github.com/dena-gohost/gohost-server/gen/api"

const (
	defaultListSpotsLimit = 3
)

type Spot struct {
	*api.Spot
}

func ValidateListSpotsLimit(l *int) int {
	if l == nil {
		return defaultListSpotsLimit
	}

	return *l
}
