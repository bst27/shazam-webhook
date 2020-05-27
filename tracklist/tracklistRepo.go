package tracklist

var (
	repoInstance *repo
)

// Repository describes a collection of tracklists.
type Repository interface {
	Add(tracklist *Tracklist)
	GetAll() []*Tracklist
}

type repo struct {
	Tracklists []*Tracklist
}

// Get returns a singleton repository instance.
func Get() Repository {
	if repoInstance == nil {
		repoInstance = &repo{}
	}

	return repoInstance
}

// Add adds a tracklist to the repository.
func (r *repo) Add(tl *Tracklist) {
	r.Tracklists = append(r.Tracklists, tl)
}

// GetAll returns all tracklists which have been added to the repository.
func (r repo) GetAll() []*Tracklist {
	return r.Tracklists
}
