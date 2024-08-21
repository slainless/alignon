package platform

import (
	"context"
	"database/sql"
	"errors"
	"mime/multipart"
	"sync"
	"time"

	"github.com/go-jet/jet/v2/qrm"
	"github.com/google/uuid"
	"github.com/slainless/my-alignon/pkg/internal/artifact/database/my_alignon/public/table"
	"github.com/slainless/my-alignon/pkg/internal/query"

	. "github.com/go-jet/jet/v2/postgres"
)

var (
	ErrConsumerNotFound          = errors.New("consumer not found")
	ErrConsumerAlreadyRegistered = errors.New("consumer already registered")
	ErrPhotosUploadFailed        = errors.New("photos upload failed")
	ErrNoLimitSetYet             = errors.New("no limit set yet")
)

type ConsumerManager struct {
	db *sql.DB

	authManager  *AuthManager
	errorTracker ErrorTracker

	fileService FileService
}

func NewConsumerManager(db *sql.DB, auth *AuthManager, file FileService, tracker ErrorTracker) *ConsumerManager {
	return &ConsumerManager{
		authManager: auth,
		db:          db,

		errorTracker: tracker,
		fileService:  file,
	}
}

func (m *ConsumerManager) GetByEmail(ctx context.Context, email string) (*Consumer, error) {
	var consumer Consumer
	err := query.GetConsumerByEmail(ctx, m.db, email, &consumer)
	if err != nil {
		if err == qrm.ErrNoRows {
			return nil, ErrConsumerNotFound
		}

		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	consumer.P_BirthDate = consumer.BirthDate.Format(time.DateOnly)
	return &consumer, nil
}

func (m *ConsumerManager) Register(ctx context.Context, payload *ConsumerRegisterInput, ktp, selfie *multipart.FileHeader) (*Consumer, error) {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	insertCustomerStmt := table.Consumers.
		INSERT(table.Consumers.MutableColumns).
		MODEL(payload).
		RETURNING(table.Consumers.ID.AS("id"))

	var id struct{ id uuid.UUID }
	err = insertCustomerStmt.QueryContext(ctx, tx, &id)
	if err != nil {
		m.errorTracker.Report(ctx, tx.Rollback())
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	ktpFileId := id.id.String() + "_ktp"
	selfieFileId := id.id.String() + "_selfie"

	wg := sync.WaitGroup{}
	wg.Add(2)

	uploadCtx, cancelUpload := context.WithDeadline(ctx, time.Now().Add(2*time.Second))
	var ktpErr, selfieErr error
	go func() {
		defer wg.Done()
		ktpErr = m.fileService.Upload(uploadCtx, ktpFileId, ktp)
		if ktpErr != nil {
			cancelUpload()
		}
	}()

	go func() {
		defer wg.Done()
		selfieErr = m.fileService.Upload(uploadCtx, selfieFileId, selfie)
		if selfieErr != nil {
			cancelUpload()
		}
	}()

	wg.Wait()
	if ktpErr != nil || selfieErr != nil {
		err := errors.Join(ktpErr, selfieErr, ErrPhotosUploadFailed)
		m.errorTracker.Report(ctx, tx.Rollback())
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	updatePhotosStmt := table.Consumers.UPDATE().
		SET(
			table.Consumers.KtpPhoto.SET(String(ktpFileId)),
			table.Consumers.SelfiePhoto.SET(String(selfieFileId)),
		).
		WHERE(table.Consumers.ID.EQ(UUID(id.id))).
		RETURNING(table.Consumers.AllColumns)

	var consumer Consumer
	err = updatePhotosStmt.QueryContext(ctx, tx, &consumer)
	if err != nil {
		wg := sync.WaitGroup{}
		wg.Add(2)

		go func() {
			defer wg.Done()
			m.errorTracker.Report(ctx, m.fileService.Delete(ctx, ktpFileId))
		}()

		go func() {
			defer wg.Done()
			m.errorTracker.Report(ctx, m.fileService.Delete(ctx, selfieFileId))
		}()

		wg.Wait()

		m.errorTracker.Report(ctx, tx.Rollback())
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		m.errorTracker.Report(ctx, tx.Rollback())
		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	consumer.P_BirthDate = consumer.BirthDate.Format(time.DateOnly)
	return &consumer, nil
}

func (m *ConsumerManager) GetLimit(ctx context.Context, id uuid.UUID) (*Limit, error) {
	var limit Limit
	err := query.GetConsumerLimit(ctx, m.db, id, &limit)
	if err != nil {
		if err == qrm.ErrNoRows {
			return nil, ErrNoLimitSetYet
		}

		m.errorTracker.Report(ctx, err)
		return nil, err
	}

	return &limit, nil
}
