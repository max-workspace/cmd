package helloworld

import (
	"context"

	"max.workspace.com/cmd/service/data"
)

type Service struct {
	data.BaseService
}

func New(ctx context.Context) (s *Service) {
	s = &Service{}
	s.Ctx = ctx
	return s
}
