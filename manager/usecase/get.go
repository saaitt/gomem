package usecase

import (
	"github.com/saaitt/gomem/config"
)

func (uc *useCase) Get(block, key string) (interface{}, error) {

	if block == "" {
		block = config.DefaultBlock
	}
	uc.m.Lock()
	value, ok := uc.blocks[block].Content[key]
	if !ok {
		return nil, config.ErrEntityNotFound
	}
	uc.m.Unlock()
	return value, nil
}
