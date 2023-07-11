// Code generated by goa v3.11.3, DO NOT EDIT.
//
// HTTP request path constructors for the Forecaster service.
//
// Command:
// $ goa gen goa.design/pulse/examples/weather/services/forecaster/design -o
// services/forecaster

package server

import (
	"fmt"
)

// ForecastForecasterPath returns the URL path to the Forecaster service forecast HTTP endpoint.
func ForecastForecasterPath(state string, city string) string {
	return fmt.Sprintf("/forecast/%v/%v", state, city)
}
