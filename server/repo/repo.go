package repo

import (
	"fmt"
	"github.com/IamFaizanKhalid/webhook-api/internal/services/file"
	"github.com/IamFaizanKhalid/webhook-api/server/dao"
	"gopkg.in/yaml.v2"
	"sync"
)

// Repo is an array of Repo objects
type Repo struct {
	file  *file.File
	mutex sync.RWMutex
	ids   map[string]int
	hooks []*dao.Hook
}

func NewRepo(hooksFile string) (*Repo, error) {
	h := &Repo{
		file:  file.New(hooksFile, 0666),
		ids:   make(map[string]int),
		hooks: []*dao.Hook{},
	}

	b, err := h.file.Load()
	if err != nil {
		return nil, fmt.Errorf("error reading hooks file: %s", err)
	}

	err = yaml.Unmarshal(b, &h.hooks)
	if err != nil {
		return nil, fmt.Errorf("error decoding hooks file: %s", err)
	}

	// keeps last one for duplicates
	for i, hook := range h.hooks {
		h.ids[hook.ID] = i
	}

	return h, nil
}

func (h *Repo) saveToFile() error {
	b, err := yaml.Marshal(h.hooks)
	if err != nil {
		return fmt.Errorf("error encoding hooks: %s", err)
	}

	err = h.file.Save(b)
	if err != nil {
		return fmt.Errorf("error writing hooks file: %s", err)
	}

	return nil
}
