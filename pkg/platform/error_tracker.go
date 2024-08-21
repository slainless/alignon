package platform

import (
	"context"
)

type ErrorTracker interface {
	Report(ctx context.Context, err error)
}
