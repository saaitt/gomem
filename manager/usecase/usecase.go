package usecase

import (
	"github.com/saaitt/gomem/manager/domain"
	"sync"
)

type useCase struct {
	blocks map[string]*domain.Block
	m      *sync.RWMutex
}

func New() domain.UseCase {
	uc := useCase{
		blocks: map[string]*domain.Block{},
		m:      new(sync.RWMutex),
	}
	return &uc
}
