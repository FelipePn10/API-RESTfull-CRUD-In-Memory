package repository

import (
	"api-restfull-crud-in-memory/models"
	"sync"
)

type ItemRepository struct {
	mu    sync.RWMutex
	items map[string]models.Item
}

func NewItemRepository() *ItemRepository {
	return &ItemRepository{
		items: make(map[string]models.Item),
	}
}

func (r *ItemRepository) Create(item models.Item) models.Item {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[item.ID] = item
	return item
}

func (r *ItemRepository) GetAll() []models.Item {
	r.mu.RLock()
	defer r.mu.RUnlock()
	items := []models.Item{}
	for _, item := range r.items {
		items = append(items, item)
	}
	return items
}

func (r *ItemRepository) GetByID(id string) (models.Item, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	item, found := r.items[id]
	return item, found
}

func (r *ItemRepository) Update(id string, item models.Item) (models.Item, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, found := r.items[id]; !found {
		return models.Item{}, false
	}
	r.items[id] = item
	return item, true
}

func (r *ItemRepository) Delete(id string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, found := r.items[id]; !found {
		return false
	}
	delete(r.items, id)
	return true
}
