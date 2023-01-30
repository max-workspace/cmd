package helloworld

import (
	"context"

	"max.workspace.com/cmd/service/page"
)

type Service struct {
	page.BaseService
}

func New(ctx context.Context) (s *Service) {
	s = &Service{}
	s.Ctx = ctx
	return s
}
