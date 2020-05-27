// A Chartlist represents a collection of tracks at Shazam.
package chartlist

import (
	"time"

	"github.com/bst27/shazam-webhook/common"
)

// A Chartlist represents a collection of tracks at Shazam. Shazam has many different chartlists
// for different cities and countries. Each chartlist is identified by a key.
type Chartlist struct {
	shazamKey string
	tracks    []*track
	created   time.Time
	updated   time.Time
}

// New creates and returns an empty chartlist.
func New(shazamKey string) *Chartlist {
	now := time.Now()
	return &Chartlist{
		shazamKey: shazamKey,
		created:   now,
		updated:   now,
	}
}

// HasTrack checks if a track is included in the chartlist.
func (c *Chartlist) HasTrack(author, title string) bool {
	for _, v := range c.tracks {
		if author == v.Author && title == v.Title {
			return true
		}
	}

	return false
}

// AddTrackIfMissing adds a track to the chartlist if it is not included yet.
func (c *Chartlist) AddTrackIfMissing(author, title string) {
	if c.HasTrack(author, title) {
		return
	}

	c.addTrack(newTrack(author, title))
}

// WriteTracks writes the tracks included in the chartlist to the given track holder.
func (c *Chartlist) WriteTracks(trackHolder common.TrackHolder) {
	for _, v := range c.tracks {
		trackHolder.AddTrackIfMissing(v.Author, v.Title)
	}
}

// ShazamKey returns the key of a chartlist as known to Shazam
func (c *Chartlist) ShazamKey() string {
	return c.shazamKey
}

func (c *Chartlist) addTrack(t *track) {
	c.tracks = append(c.tracks, t)
	c.updated = time.Now()
}
