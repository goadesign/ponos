// Code generated by goa v3.11.3, DO NOT EDIT.
//
// HTTP request path constructors for the Poller service.
//
// Command:
// $ goa gen goa.design/pulse/examples/weather/services/poller/design -o
// services/poller

package client

// AddLocationPollerPath returns the URL path to the Poller service add_location HTTP endpoint.
func AddLocationPollerPath() string {
	return "/poller/location"
}

// SubscribePollerPath returns the URL path to the Poller service subscribe HTTP endpoint.
func SubscribePollerPath() string {
	return "/poller/subscribe"
}

// AddWorkerPollerPath returns the URL path to the Poller service add_worker HTTP endpoint.
func AddWorkerPollerPath() string {
	return "/poller/workers"
}

// RemoveWorkerPollerPath returns the URL path to the Poller service remove_worker HTTP endpoint.
func RemoveWorkerPollerPath() string {
	return "/poller/workers"
}

// StatusPollerPath returns the URL path to the Poller service status HTTP endpoint.
func StatusPollerPath() string {
	return "/poller/status"
}
