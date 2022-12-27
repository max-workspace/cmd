package helloworld

import (
	"max.workspace.com/cmd/service/page"
)

type Service struct {
	page.BaseService
}

func New() *Service {
	return &Service{}
}
