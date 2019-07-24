package test

import (
	"testing"

	"go.uber.org/zap"
	zapobserver "go.uber.org/zap/zaptest/observer"
)

func TestS_FooBar(t *testing.T) {
	core, obs := zapobserver.New(zap.InfoLevel) //
	// core, obs := zapobserver.New(zap.FatalLevel) // the test will be failed

	logger := zap.New(core)
	s := S{
		logger: logger,
	}

	s.FooBar()

	if obs.Len() != 1 {
		t.Fatal("the log should be emitted")
	}
}
