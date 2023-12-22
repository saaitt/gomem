package usecase

import (
	"github.com/saaitt/gomem/config"
)

func (uc *useCase) Set(block, key string, value interface{}) error {
	if block == "" {
		block = config.DefaultBlock
	}
	uc.m.Lock()
	uc.blocks[block].Content[key] = value
	uc.m.Unlock()
	return nil
}
