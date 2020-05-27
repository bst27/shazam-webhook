// A Tracklist receives tracks from watched chartlists and uses webhooks to notify about added tracks.
package tracklist

import (
	"time"

	"github.com/bst27/shazam-webhook/common"

	"github.com/google/uuid"
)

// A Tracklist receives tracks from watched chartlists and uses webhooks to notify about added tracks.
type Tracklist struct {
	id                  string
	created             time.Time
	updated             time.Time
	tracks              []*track
	watchedShazamCharts []string
	webhooks            []common.WebhookReceiver
}

// New returns a new tracklist
func New() *Tracklist {
	now := time.Now()

	return &Tracklist{
		id:      uuid.New().String(),
		created: now,
		updated: now,
	}
}

// WatchShazamCharts makes the tracklist to start watching a chartlist.
func (tl *Tracklist) WatchShazamCharts(shazamChartKey string) {
	for _, v := range tl.watchedShazamCharts {
		if shazamChartKey == v {
			return
		}
	}

	tl.watchedShazamCharts = append(tl.watchedShazamCharts, shazamChartKey)
}

// IsWatching allows to check if a tracklist watches a chartlist.
func (tl *Tracklist) IsWatching(shazamChartKey string) bool {
	for _, v := range tl.watchedShazamCharts {
		if shazamChartKey == v {
			return true
		}
	}

	return false
}

// AddTrackIfMissing adds a track to the tracklist if it is not included yet.
func (tl *Tracklist) AddTrackIfMissing(author, title string) {
	if tl.HasTrack(author, title) {
		return
	}

	tl.addTrack(newTrack(author, title))
}

func (tl *Tracklist) addTrack(t *track) {
	tl.tracks = append(tl.tracks, t)
	tl.updated = time.Now()

	for _, w := range tl.webhooks {
		_ = w.ReceiveWebhook(t.Author, t.Title)
	}
}

// HasTrack checks if a track is included in the tracklist.
func (tl *Tracklist) HasTrack(author, title string) bool {
	for _, v := range tl.tracks {
		if author == v.Author && title == v.Title {
			return true
		}
	}

	return false
}

// RegisterWebhook makes the tracklist to start sending notifications
// about added tracks to the given webhook.
func (tl *Tracklist) RegisterWebhook(webhook common.WebhookReceiver) {
	tl.webhooks = append(tl.webhooks, webhook)
}
