package chartlist

import (
	"time"
)

type track struct {
	Author  string
	Title   string
	Created time.Time
}

func newTrack(author, title string) *track {
	return &track{
		Author:  author,
		Title:   title,
		Created: time.Now(),
	}
}
