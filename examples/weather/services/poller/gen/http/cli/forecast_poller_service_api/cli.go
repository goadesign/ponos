// Code generated by goa v3.11.3, DO NOT EDIT.
//
// Forecast Poller Service API HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/pulse/examples/weather/services/poller/design -o
// services/poller

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	pollerc "goa.design/pulse/examples/weather/services/poller/gen/http/poller/client"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `poller (add-location|subscribe|add-worker|remove-worker|status)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` poller add-location --city "Santa Barbara" --state "CA"` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
	dialer goahttp.Dialer,
	pollerConfigurer *pollerc.ConnConfigurer,
) (goa.Endpoint, any, error) {
	var (
		pollerFlags = flag.NewFlagSet("poller", flag.ContinueOnError)

		pollerAddLocationFlags     = flag.NewFlagSet("add-location", flag.ExitOnError)
		pollerAddLocationCityFlag  = pollerAddLocationFlags.String("city", "REQUIRED", "")
		pollerAddLocationStateFlag = pollerAddLocationFlags.String("state", "REQUIRED", "")

		pollerSubscribeFlags     = flag.NewFlagSet("subscribe", flag.ExitOnError)
		pollerSubscribeCityFlag  = pollerSubscribeFlags.String("city", "REQUIRED", "")
		pollerSubscribeStateFlag = pollerSubscribeFlags.String("state", "REQUIRED", "")

		pollerAddWorkerFlags = flag.NewFlagSet("add-worker", flag.ExitOnError)

		pollerRemoveWorkerFlags = flag.NewFlagSet("remove-worker", flag.ExitOnError)

		pollerStatusFlags = flag.NewFlagSet("status", flag.ExitOnError)
	)
	pollerFlags.Usage = pollerUsage
	pollerAddLocationFlags.Usage = pollerAddLocationUsage
	pollerSubscribeFlags.Usage = pollerSubscribeUsage
	pollerAddWorkerFlags.Usage = pollerAddWorkerUsage
	pollerRemoveWorkerFlags.Usage = pollerRemoveWorkerUsage
	pollerStatusFlags.Usage = pollerStatusUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "poller":
			svcf = pollerFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "poller":
			switch epn {
			case "add-location":
				epf = pollerAddLocationFlags

			case "subscribe":
				epf = pollerSubscribeFlags

			case "add-worker":
				epf = pollerAddWorkerFlags

			case "remove-worker":
				epf = pollerRemoveWorkerFlags

			case "status":
				epf = pollerStatusFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "poller":
			c := pollerc.NewClient(scheme, host, doer, enc, dec, restore, dialer, pollerConfigurer)
			switch epn {
			case "add-location":
				endpoint = c.AddLocation()
				data, err = pollerc.BuildAddLocationPayload(*pollerAddLocationCityFlag, *pollerAddLocationStateFlag)
			case "subscribe":
				endpoint = c.Subscribe()
				data, err = pollerc.BuildSubscribePayload(*pollerSubscribeCityFlag, *pollerSubscribeStateFlag)
			case "add-worker":
				endpoint = c.AddWorker()
				data = nil
			case "remove-worker":
				endpoint = c.RemoveWorker()
				data = nil
			case "status":
				endpoint = c.Status()
				data = nil
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// pollerUsage displays the usage of the poller command and its subcommands.
func pollerUsage() {
	fmt.Fprintf(os.Stderr, `Service that polls weather forecasts and notifies subscribers.
Usage:
    %[1]s [globalflags] poller COMMAND [flags]

COMMAND:
    add-location: Adds a location to the polling list.
    subscribe: Subscribes to weather forecast notifications for a location.
    add-worker: Adds a worker to the poller worker pool.
    remove-worker: Removes a worker from the poller worker pool.
    status: Provides poller status and statistics.

Additional help:
    %[1]s poller COMMAND --help
`, os.Args[0])
}
func pollerAddLocationUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] poller add-location -city STRING -state STRING

Adds a location to the polling list.
    -city STRING: 
    -state STRING: 

Example:
    %[1]s poller add-location --city "Santa Barbara" --state "CA"
`, os.Args[0])
}

func pollerSubscribeUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] poller subscribe -city STRING -state STRING

Subscribes to weather forecast notifications for a location.
    -city STRING: 
    -state STRING: 

Example:
    %[1]s poller subscribe --city "Santa Barbara" --state "CA"
`, os.Args[0])
}

func pollerAddWorkerUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] poller add-worker

Adds a worker to the poller worker pool.

Example:
    %[1]s poller add-worker
`, os.Args[0])
}

func pollerRemoveWorkerUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] poller remove-worker

Removes a worker from the poller worker pool.

Example:
    %[1]s poller remove-worker
`, os.Args[0])
}

func pollerStatusUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] poller status

Provides poller status and statistics.

Example:
    %[1]s poller status
`, os.Args[0])
}
