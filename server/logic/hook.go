package logic

import (
	"context"
	"fmt"
	"github.com/IamFaizanKhalid/webhook-api/server/dao"
	"github.com/IamFaizanKhalid/webhook-api/server/logic/output"
)

func (l *CoreLogic) GetAllHooks(ctx context.Context) []*dao.Hook {
	return l.repo.GetAllHooks()
}

func (l *CoreLogic) GetHook(ctx context.Context, id string) (*dao.Hook, error) {
	hook, err := l.repo.GetHook(id)
	return hook, wrap(err)
}

func (l *CoreLogic) DeleteHook(ctx context.Context, id string) error {
	if !l.repo.HookExists(id) {
		return output.ErrNotFound
	}

	return wrap(l.repo.DeleteHook(id))
}

func (l *CoreLogic) AddHook(ctx context.Context, h dao.Hook) error {
	if l.repo.HookExists(h.ID) {
		return output.ErrConflict(fmt.Errorf("hook with id `%s` already exists", h.ID))
	}

	return wrap(l.repo.AddHook(h))
}

func (l *CoreLogic) UpdateHook(ctx context.Context, id string, h dao.Hook) error {
	if !l.repo.HookExists(id) {
		return output.ErrNotFound
	}
	if h.ID != id && l.repo.HookExists(h.ID) {
		return output.ErrConflict(fmt.Errorf("another hook with id `%s` exists", h.ID))
	}

	return wrap(l.repo.UpdateHook(id, h))
}
