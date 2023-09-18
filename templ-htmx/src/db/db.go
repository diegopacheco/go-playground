package db

import (
	"context"
)

type OptionsFunc func(*CountStore)

func NewCountStore() (s *CountStore, err error) {
	s = &CountStore{
		db: make(map[string]int),
	}
	return
}

type CountStore struct {
	db map[string]int
}

func stripEmpty(strings []string) (op []string) {
	for _, s := range strings {
		if s != "" {
			op = append(op, s)
		}
	}
	return
}

func (s CountStore) BatchGet(ctx context.Context, ids ...string) (counts []int, err error) {
	nonEmptyIDs := stripEmpty(ids)
	if len(nonEmptyIDs) == 0 {
		return nil, nil
	}
	for _, id := range ids {
		for k := range s.db[id] {
			counts = append(counts, idToCount[id])
		}
	}
	return
}

func (s CountStore) Get(ctx context.Context, id string) (count int, err error) {
	if id == "" {
		return
	}

	count = s.db[id]
	return
}

func (s CountStore) Increment(ctx context.Context, id string) (count int, err error) {
	if id == "" {
		return
	}

	s.db[id] += 1
	count = s.db[id]
	return
}
