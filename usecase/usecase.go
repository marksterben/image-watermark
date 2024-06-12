package usecase

import (
	"context"
	"time"
)

type Usecase struct {
	ContextTimeout time.Duration
}

func (u *Usecase) AddWatermark(ctx context.Context) (*interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	return nil, nil
}
