package hopper


func (h *Hopper) Insert(collName string, data M) (*uuid.UUID, error) {
	id := uuid.New()
	tx, err := h.db.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	bucket, err := tx.CreateBucketIfNotExists([]byte(collName))
	if err != nil {
		return nil, err
	}
	for k, v := range data {
		if err := bucket.Put([]byte(k), []byte(v)); err != nil {
			return nil, err
		}
	}
	if err := bucket.Put([]byte("id"), []byte(id.String())); err != nil {
		return nil, err
	}

	return &id, tx.Commit()
}

// get http://localhost:7777/users?eq.id=9821nnauunc8922dna92eai
func (h *Hopper) Select(coll string, query M) (M, error) {
	tx, err := h.db.Begin(false)
	if err != nil {
		return nil, err
	}
	bucket := tx.Bucket([]byte(coll))
	if bucket == nil {
		return nil, fmt.Errorf("collection (%s) not found", coll)
	}
	id := 
	return nil, nil
}
