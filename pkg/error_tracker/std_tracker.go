package tracker

import (
	"context"
	"fmt"
	"os"
)

type StdTracker struct{}

func (s *StdTracker) Report(ctx context.Context, err error) {
	fmt.Fprintf(os.Stderr, "[ERROR] %+v", err)
}
