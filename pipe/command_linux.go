//go:build linux

package pipe

import (
	"context"
	"errors"

	"github.com/nextlinux/go-pipeline/internal/ptree"
)

// On linux, we can limit or observe memory usage in command stages.
var _ LimitableStage = (*commandStage)(nil)

var (
	errProcessInfoMissing = errors.New("cmd.Process is nil")
)

func (s *commandStage) GetRSSAnon(ctx context.Context) (uint64, error) {
	if s.cmd.Process == nil {
		return 0, errProcessInfoMissing
	}

	return ptree.GetProcessTreeRSSAnon(s.cmd.Process.Pid)
}
