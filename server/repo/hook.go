package repo

import (
	"fmt"
	"github.com/IamFaizanKhalid/webhook-manager/server/dao"
)

func (h *Repo) GetAllHooks() []*dao.Hook {
	return h.hooks
}

func (h *Repo) GetHook(id string) (*dao.Hook, error) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	i, ok := h.ids[id]
	if !ok {
		return nil, fmt.Errorf("hook with id `%s` not found", id)
	}

	return h.hooks[i], nil
}

func (h *Repo) AddHook(hook dao.Hook) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	h.hooks = append(h.hooks, &hook)
	h.ids[hook.ID] = len(h.hooks) - 1

	return h.saveToFile()
}

func (h *Repo) DeleteHook(id string) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	i, ok := h.ids[id]
	if !ok {
		return fmt.Errorf("hook with id `%s` not found", id)
	}

	h.hooks = append(h.hooks[:i], h.hooks[i+1:]...)

	delete(h.ids, id)

	return h.saveToFile()
}

func (h *Repo) UpdateHook(id string, newHook dao.Hook) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	i, ok := h.ids[id]
	if !ok {
		return fmt.Errorf("hook with id `%s` not found", id)
	}

	h.hooks[i] = &newHook

	delete(h.ids, id)
	h.ids[newHook.ID] = i

	return h.saveToFile()
}

func (h *Repo) HookExists(id string) bool {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	_, ok := h.ids[id]

	return ok
}
