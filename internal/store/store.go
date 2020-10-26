package store

import (
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/axard/todo-list/internal/restmodels"
	"github.com/emirpasic/gods/maps/linkedhashmap"
	"github.com/go-openapi/errors"
)

type Items struct {
	items  *linkedhashmap.Map
	lock   *sync.RWMutex
	lastID int64
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
	item.ID = newID
	i.items.Put(newID, item)

	return nil
}

func (i *Items) Update(id int64, item *restmodels.Item) error {
	if item == nil {
		return errors.New(http.StatusInternalServerError, "item must be present")
	}

	i.lock.Lock()
	defer i.lock.Unlock()

	_, exists := i.items.Get(id)
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	item.ID = id
	i.items.Put(id, item)

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

func (i *Items) Read(since int64, limit int32) []*restmodels.Item {
	i.lock.RLock()
	defer i.lock.RUnlock()

	iter := i.items.Iterator()
	result := make([]*restmodels.Item, 0, limit)
	enough := func() bool { return len(result) >= int(limit) }

	for iter.Next() && !enough() {
		if since == 0 {
			if v, ok := iter.Value().(*restmodels.Item); ok {
				result = append(result, v)
			}
		} else {
			since--
		}
	}

	return result
}
