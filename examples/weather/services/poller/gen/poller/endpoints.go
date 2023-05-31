// Code generated by goa v3.11.3, DO NOT EDIT.
//
// Poller endpoints
//
// Command:
// $ goa gen goa.design/ponos/examples/weather/services/poller/design -o
// services/poller

package poller

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "Poller" service endpoints.
type Endpoints struct {
	AddLocation  goa.Endpoint
	Subscribe    goa.Endpoint
	AddWorker    goa.Endpoint
	RemoveWorker goa.Endpoint
	Status       goa.Endpoint
}

// SubscribeEndpointInput holds both the payload and the server stream of the
// "subscribe" method.
type SubscribeEndpointInput struct {
	// Payload is the method payload.
	Payload *CityAndState
	// Stream is the server stream used by the "subscribe" method to send data.
	Stream SubscribeServerStream
}

// NewEndpoints wraps the methods of the "Poller" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		AddLocation:  NewAddLocationEndpoint(s),
		Subscribe:    NewSubscribeEndpoint(s),
		AddWorker:    NewAddWorkerEndpoint(s),
		RemoveWorker: NewRemoveWorkerEndpoint(s),
		Status:       NewStatusEndpoint(s),
	}
}

// Use applies the given middleware to all the "Poller" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.AddLocation = m(e.AddLocation)
	e.Subscribe = m(e.Subscribe)
	e.AddWorker = m(e.AddWorker)
	e.RemoveWorker = m(e.RemoveWorker)
	e.Status = m(e.Status)
}

// NewAddLocationEndpoint returns an endpoint function that calls the method
// "add_location" of service "Poller".
func NewAddLocationEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CityAndState)
		return nil, s.AddLocation(ctx, p)
	}
}

// NewSubscribeEndpoint returns an endpoint function that calls the method
// "subscribe" of service "Poller".
func NewSubscribeEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		ep := req.(*SubscribeEndpointInput)
		return nil, s.Subscribe(ctx, ep.Payload, ep.Stream)
	}
}

// NewAddWorkerEndpoint returns an endpoint function that calls the method
// "add_worker" of service "Poller".
func NewAddWorkerEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return s.AddWorker(ctx)
	}
}

// NewRemoveWorkerEndpoint returns an endpoint function that calls the method
// "remove_worker" of service "Poller".
func NewRemoveWorkerEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return nil, s.RemoveWorker(ctx)
	}
}

// NewStatusEndpoint returns an endpoint function that calls the method
// "status" of service "Poller".
func NewStatusEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return s.Status(ctx)
	}
}
