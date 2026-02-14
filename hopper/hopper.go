package hopper

import (
	"fmt"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

const (
	defaultDbName = "default"
)

type M map[string]string

type Collection struct {
	*bbolt.Bucket
}
type Hopper struct {
	db *bbolt.DB
}

func New() (*Hopper, error) {
	dbname := fmt.Sprintf("%s.hopper", defaultDbName)
	db, err := bbolt.Open(dbname, 066, nil)
	if err != nil {
		return nil, err
	}

	return &Hopper{
		db: db,
	}, nil
}

func (h *Hopper) CreateCollection(name string) (*Collection, error) {
	tx, err := h.db.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	bucket, err := tx.CreateBucketIfNotExists([]byte("users"))
	if err != nil {
		return nil, err
	}

	return &Collection{Bucket: bucket}, nil

}
