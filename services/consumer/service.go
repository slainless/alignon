package consumer

import (
	"database/sql"

	"github.com/slainless/my-alignon/pkg/auth"
	"github.com/slainless/my-alignon/pkg/platform"
)

type Service struct {
	authManager     *platform.AuthManager
	consumerManager *platform.ConsumerManager
	loanManager     *platform.LoanManager
}

func NewService(db *sql.DB, file platform.FileService, tracker platform.ErrorTracker) *Service {
	emailJwtAuth := auth.NewEmailJWTAuthService([]byte{})

	auth := platform.NewAuthManager(emailJwtAuth, tracker)
	consumer := platform.NewConsumerManager(db, auth, file, tracker)

	return &Service{
		authManager:     auth,
		consumerManager: consumer,
	}
}
