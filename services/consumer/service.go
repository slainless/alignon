package consumer

import (
	"database/sql"

	"github.com/slainless/my-alignon/pkg/auth"
	tracker "github.com/slainless/my-alignon/pkg/error_tracker"
	"github.com/slainless/my-alignon/pkg/platform"
)

type Service struct {
	authManager     *platform.AuthManager
	consumerManager *platform.ConsumerManager
}

func NewService(db *sql.DB) *Service {
	tracker := &tracker.StdTracker{}
	emailJwtAuth := auth.NewEmailJWTAuthService([]byte{})

	auth := platform.NewAuthManager(emailJwtAuth, tracker)
	consumer := platform.NewConsumerManager(db, auth, tracker)

	return &Service{
		authManager:     auth,
		consumerManager: consumer,
	}
}
