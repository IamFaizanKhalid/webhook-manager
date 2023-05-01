package logic

import (
	"github.com/IamFaizanKhalid/webhook-api/internal/errors"
	"github.com/IamFaizanKhalid/webhook-api/server/repo"
)

type CoreLogic struct {
	repo *repo.Repo
}

func New(repo *repo.Repo) *CoreLogic {
	return &CoreLogic{repo}
}

func wrap(err error) error {
	if err == nil {
		return nil
	}

	return errors.InternalError(err)
}
