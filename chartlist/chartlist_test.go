package chartlist

import "testing"

func TestChartlist_AddTrackIfMissing(t *testing.T) {
	cl := New("123")

	cl.AddTrackIfMissing("John Doe", "Blue or green")
	if got := len(cl.tracks); got != 1 {
		t.Errorf("Got %d tracks; want 1", got)
	}

	cl.AddTrackIfMissing("John Doe", "Hello")
	if got := len(cl.tracks); got != 2 {
		t.Errorf("Got %d tracks; want 2", got)
	}

	cl.AddTrackIfMissing("John Doe", "Blue or green")
	if got := len(cl.tracks); got != 2 {
		t.Errorf("Got %d tracks; want 2", got)
	}

	cl.AddTrackIfMissing("Fred", "Hello")
	if got := len(cl.tracks); got != 3 {
		t.Errorf("Got %d tracks; want 3", got)
	}
}

func TestChartlist_HasTrack(t *testing.T) {
	cl := New("123")

	if cl.HasTrack("John Doe", "Hello") {
		t.Errorf("Has track; want false")
	}

	cl.AddTrackIfMissing("John Doe", "Blue or green")
	if cl.HasTrack("John Doe", "Hello") {
		t.Errorf("Has track; want false")
	}

	cl.AddTrackIfMissing("John Doe", "Hello")
	if !cl.HasTrack("John Doe", "Hello") {
		t.Errorf("Has no track; want track")
	}
}

func TestChartlist_ShazamKey(t *testing.T) {
	cl := New("123")

	if got := cl.ShazamKey(); got != "123" {
		t.Errorf("Got %s key; want 123", got)
	}
}

func TestChartlist_WriteTracks(t *testing.T) {
	cl := New("123")
	trackHolder := New("ABC")

	cl.WriteTracks(trackHolder)

	if len(trackHolder.tracks) != len(cl.tracks) {
		t.Errorf("Got %d tracks; want %d", len(trackHolder.tracks), len(cl.tracks))
	}

	cl.AddTrackIfMissing("John Doe", "Hello")
	cl.WriteTracks(trackHolder)
	if len(trackHolder.tracks) != len(cl.tracks) {
		t.Errorf("Got %d tracks; want %d", len(trackHolder.tracks), len(cl.tracks))
	}
}
