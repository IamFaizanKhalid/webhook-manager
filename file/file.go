package file

import (
	"fmt"
	"github.com/IamFaizanKhalid/webhook-api/model"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"sync"
)

// File is an array of Hook objects
type File struct {
	mutex sync.RWMutex
	path  string
	ids   map[string]int
	hooks []*model.Hook
}

func (f *File) GetAllHooks() []*model.Hook {
	return f.hooks
}

func (f *File) GetHook(id string) (*model.Hook, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	i, ok := f.ids[id]
	if !ok {
		return nil, fmt.Errorf("hook with id `%s` not found", id)
	}

	return f.hooks[i], nil
}

func (f *File) AddHook(h model.Hook) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	f.hooks = append(f.hooks, &h)
	f.ids[h.ID] = len(f.hooks) - 1

	return f.flush()
}

func (f *File) DeleteHook(id string) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	i, ok := f.ids[id]
	if !ok {
		return fmt.Errorf("hook with id `%s` not found", id)
	}

	f.hooks = append(f.hooks[:i], f.hooks[i+1:]...)

	delete(f.ids, id)

	return f.flush()
}

func (f *File) UpdateHook(id string, newHook model.Hook) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	i, ok := f.ids[id]
	if !ok {
		return fmt.Errorf("hook with id `%s` not found", id)
	}

	f.hooks[i] = &newHook

	delete(f.ids, id)
	f.ids[newHook.ID] = i

	return f.flush()
}

func (f *File) Exists(id string) bool {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	_, ok := f.ids[id]

	return ok
}

func Parse(path string) (*File, error) {
	file, err := os.OpenFile("hooks.yml", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, fmt.Errorf("Error reading hooks file: %s", err)
	}

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Error reading hooks file: %s", err)
	}

	f := &File{path: path, ids: make(map[string]int)}
	err = yaml.Unmarshal(b, &f.hooks)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling hooks file: %s", err)
	}

	// keeps last one for duplicates
	for i, h := range f.hooks {
		f.ids[h.ID] = i
	}

	return f, nil
}

func (f *File) flush() error {
	b, err := yaml.Marshal(f.hooks)
	if err != nil {
		return fmt.Errorf("Error marshalling hooks to file: %s", err)
	}

	err = os.WriteFile(f.path, b, 0644)
	if err != nil {
		return fmt.Errorf("Error writing hooks file: %s", err)
	}

	return nil
}
