package store

import (
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/axard/todo-list/internal/restmodels"
	"github.com/emirpasic/gods/maps/linkedhashmap"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

type Items struct {
	items  *linkedhashmap.Map
	lock   *sync.RWMutex
	lastID int64
}

type Item struct {
	Completed   bool
	Description string
}

func NewItems() *Items {
	return &Items{
		items: linkedhashmap.New(),
		lock:  &sync.RWMutex{},
	}
}

func (i *Items) newItemID() int64 {
	return atomic.AddInt64(&i.lastID, 1)
}

func (i *Items) Create(item *restmodels.Item) error {
	if item == nil {
		return errors.New(http.StatusInternalServerError, "item must be present")
	}

	i.lock.Lock()
	defer i.lock.Unlock()

	newID := i.newItemID()
	i.items.Put(newID, &Item{
		Completed:   swag.BoolValue(item.Completed),
		Description: swag.StringValue(item.Description),
	})

	item.ID = newID

	return nil
}

func (i *Items) Update(id int64, item *restmodels.Item) error {
	if item == nil {
		return errors.New(http.StatusInternalServerError, "item must be present")
	}

	i.lock.Lock()
	defer i.lock.Unlock()

	v, exists := i.items.Get(id)
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	stored, ok := v.(*Item)
	if !ok {
		return errors.New(http.StatusInternalServerError, "ivalid stored item")
	}

	if item.Completed != nil {
		stored.Completed = *item.Completed
	}

	if item.Description != nil {
		stored.Description = *item.Description
	}

	item.ID = id

	return nil
}

func (i *Items) Delete(id int64) error {
	i.lock.Lock()
	defer i.lock.Unlock()

	_, exists := i.items.Get(id)
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	i.items.Remove(id)

	return nil
}

func (i *Items) Read(since int64, limit int32) ([]*restmodels.Item, error) {
	i.lock.RLock()
	defer i.lock.RUnlock()

	iter := i.items.Iterator()
	result := make([]*restmodels.Item, 0, limit)
	enough := func() bool { return len(result) >= int(limit) }

	for iter.Next() && !enough() {
		if since == 0 {
			v, ok := iter.Value().(*Item)
			if !ok {
				return nil, errors.New(http.StatusInternalServerError, "ivalid stored item")
			}

			k, ok := iter.Key().(int64)
			if !ok {
				return nil, errors.New(http.StatusInternalServerError, "ivalid stored item")
			}

			result = append(result, &restmodels.Item{
				ID:          k,
				Completed:   &v.Completed,
				Description: &v.Description,
			})
		} else {
			since--
		}
	}

	return result, nil
}

func (i *Items) Size() int64 {
	return int64(i.items.Size())
}
