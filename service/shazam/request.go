package shazam

import (
	"github.com/bst27/shazam"
)

// Request describes a web request to fetch details about tracks from Shazam.
type Request interface {
	Send() ([]shazam.Track, error)
}

type request struct {
	chartKey string
}

func newRequest(chartKey string) *request {
	return &request{
		chartKey: chartKey,
	}
}

func (r request) Send() ([]shazam.Track, error) {
	shazam.Cities().GermanyBerlin()
	return shazam.FetchCityCharts(r.chartKey)
}
