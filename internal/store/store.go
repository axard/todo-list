package store

import (
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/axard/todo-list/internal/restmodels"
	"github.com/go-openapi/errors"
)

type Items struct {
	items  map[int64]*restmodels.Item
	lock   *sync.RWMutex
	lastID int64
}

func NewItems() *Items {
	return &Items{
		items: make(map[int64]*restmodels.Item),
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
	i.items[newID] = item

	return nil
}

func (i *Items) Update(id int64, item *restmodels.Item) error {
	if item == nil {
		return errors.New(http.StatusInternalServerError, "item must be present")
	}

	i.lock.Lock()
	defer i.lock.Unlock()

	_, exists := i.items[id]
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	item.ID = id
	i.items[id] = item

	return nil
}

func (i *Items) Delete(id int64) error {
	i.lock.Lock()
	defer i.lock.Unlock()

	_, exists := i.items[id]
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	delete(i.items, id)

	return nil
}

func (i *Items) Read(since int64, limit int32) (result []*restmodels.Item) {
	i.lock.RLock()
	defer i.lock.RUnlock()

	result = make([]*restmodels.Item, 0)

	for id, item := range i.items {
		if len(result) >= int(limit) {
			return
		}

		if since == 0 || id > since {
			result = append(result, item)
		}
	}

	return
}
