package common

// A TrackHolder can have tracks added to it.
type TrackHolder interface {
	AddTrackIfMissing(author, title string)
}
