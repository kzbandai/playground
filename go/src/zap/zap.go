package test

import "go.uber.org/zap"

type S struct {
	logger *zap.Logger
}

func Main() {
	s := S{}
	s.FooBar()
}

func (s *S) FooBar() {
	s.logger.Warn("warning")
}
