package chartlist

var (
	repoInstance *repo
)

// Repository describes a collection of chartlists.
type Repository interface {
	Add(item *Chartlist)
	GetAll() []*Chartlist
}

type repo struct {
	Items []*Chartlist
}

// Get returns a singleton repository instance.
func Get() Repository {
	if repoInstance == nil {
		repoInstance = &repo{}
	}

	return repoInstance
}

// Add adds a chartlist to the repository.
func (r *repo) Add(item *Chartlist) {
	r.Items = append(r.Items, item)
}

// GetAll returns all chartlists which have been added to the repository.
func (r repo) GetAll() []*Chartlist {
	return r.Items
}
