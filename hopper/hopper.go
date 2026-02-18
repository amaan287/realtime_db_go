package hopper

import (
	"fmt"
	"os"

	"go.etcd.io/bbolt"
)

const (
	defaultDbName = "default"
	ext           = "hopper"
)

type Map map[string]string

type Hopper struct {
	currentDatabase string
	*Options
	db *bbolt.DB
}

func New(options ...OptFunc) (*Hopper, error) {
	opts := &Options{
		Encoder: JSONEncoder{},
		Decoder: JSONDecoder{},
		DBName:  defaultDbName,
	}
	for _, fn := range options {
		fn(opts)
	}
	dbname := fmt.Sprintf("%s.hopper", defaultDbName)
	db, err := bbolt.Open(dbname, 066, nil)
	if err != nil {
		return nil, err
	}

	return &Hopper{
		db: db,
	}, nil
}

func (h *Hopper) DropDatabase(name string) error {
	dbname := fmt.Sprintf("%s.%s", name, ext)
	return os.Remove(dbname)
}

func (h *Hopper) CreateCollection(name string) (*bbolt.Bucket, error) {
	tx, err := h.db.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	bucket, err := tx.CreateBucketIfNotExists([]byte("users"))
	if err != nil {
		return nil, err
	}

	return bucket, nil

}
func (h *Hopper) Coll(name string) *Filter {
	return NewFilter(h, name)
}
