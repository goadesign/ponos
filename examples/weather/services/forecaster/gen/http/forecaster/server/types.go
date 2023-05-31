// Code generated by goa v3.11.3, DO NOT EDIT.
//
// Forecaster HTTP server types
//
// Command:
// $ goa gen goa.design/ponos/examples/weather/services/forecaster/design -o
// services/forecaster

package server

import (
	goa "goa.design/goa/v3/pkg"
	forecaster "goa.design/ponos/examples/weather/services/forecaster/gen/forecaster"
)

// ForecastResponseBody is the type of the "Forecaster" service "forecast"
// endpoint HTTP response body.
type ForecastResponseBody struct {
	// Forecast location
	Location *LocationResponseBody `form:"location" json:"location" xml:"location"`
	// Weather forecast periods
	Periods []*PeriodResponseBody `form:"periods" json:"periods" xml:"periods"`
}

// ForecastTimeoutResponseBody is the type of the "Forecaster" service
// "forecast" endpoint HTTP response body for the "timeout" error.
type ForecastTimeoutResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// LocationResponseBody is used to define fields on response body types.
type LocationResponseBody struct {
	// Latitude
	Lat float64 `form:"lat" json:"lat" xml:"lat"`
	// Longitude
	Long float64 `form:"long" json:"long" xml:"long"`
	// City
	City string `form:"city" json:"city" xml:"city"`
	// State
	State string `form:"state" json:"state" xml:"state"`
}

// PeriodResponseBody is used to define fields on response body types.
type PeriodResponseBody struct {
	// Period name
	Name string `form:"name" json:"name" xml:"name"`
	// Start time
	StartTime string `form:"startTime" json:"startTime" xml:"startTime"`
	// End time
	EndTime string `form:"endTime" json:"endTime" xml:"endTime"`
	// Temperature
	Temperature int `form:"temperature" json:"temperature" xml:"temperature"`
	// Temperature unit
	TemperatureUnit string `form:"temperatureUnit" json:"temperatureUnit" xml:"temperatureUnit"`
	// Summary
	Summary string `form:"summary" json:"summary" xml:"summary"`
}

// NewForecastResponseBody builds the HTTP response body from the result of the
// "forecast" endpoint of the "Forecaster" service.
func NewForecastResponseBody(res *forecaster.Forecast2) *ForecastResponseBody {
	body := &ForecastResponseBody{}
	if res.Location != nil {
		body.Location = marshalForecasterLocationToLocationResponseBody(res.Location)
	}
	if res.Periods != nil {
		body.Periods = make([]*PeriodResponseBody, len(res.Periods))
		for i, val := range res.Periods {
			body.Periods[i] = marshalForecasterPeriodToPeriodResponseBody(val)
		}
	}
	return body
}

// NewForecastTimeoutResponseBody builds the HTTP response body from the result
// of the "forecast" endpoint of the "Forecaster" service.
func NewForecastTimeoutResponseBody(res *goa.ServiceError) *ForecastTimeoutResponseBody {
	body := &ForecastTimeoutResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewForecastPayload builds a Forecaster service forecast endpoint payload.
func NewForecastPayload(state string, city string) *forecaster.ForecastPayload {
	v := &forecaster.ForecastPayload{}
	v.State = state
	v.City = city

	return v
}
