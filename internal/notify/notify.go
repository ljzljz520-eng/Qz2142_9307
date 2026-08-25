package notify

import (
	"mathrush/internal/domain"
	"sync"
)

type Hub struct {
	mu     sync.Mutex
	events []domain.Event
}

func New() *Hub { return &Hub{events: []domain.Event{}} }
func (h *Hub) Publish(e domain.Event) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.events = append(h.events, e)
}
func (h *Hub) Snapshot() []domain.Event {
	h.mu.Lock()
	defer h.mu.Unlock()
	return append([]domain.Event(nil), h.events...)
}
func (h *Hub) Size() int { return len(h.Snapshot()) }
