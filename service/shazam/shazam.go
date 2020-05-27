// Shazam offers a web client to fetch details from Shazam.
package shazam

// Client describes a web client to fetch details from Shazam.
type Client interface {
	CreateFetchChartsRequest(chartKey string) Request
}

// New returns a new web client.
func New() Client {
	return &client{}
}

type client struct {
}

// CreateFetchChartsRequest creates a request to fetch tracks for a chartlist from Shazam.
func (c client) CreateFetchChartsRequest(chartKey string) Request {
	return newRequest(chartKey)
}
