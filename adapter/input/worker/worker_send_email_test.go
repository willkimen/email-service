package worker_test

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"

	"emailservice/adapter/input/worker"
	"emailservice/core/application/email_errors"
	"emailservice/core/application/email_message"

	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/require"
)

func TestProcessSendEmail_InvalidPayload_ReturnsSkipRetry(t *testing.T) {
	handler := &worker.SendEmailTaskHandler{
		UseCase: &mockUseCase{},
		Logger:  slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}

	task := asynq.NewTask("send:email", []byte("invalid-json"))

	err := handler.ProcessSendEmail(context.Background(), task)

	require.Equal(t, asynq.SkipRetry, err,
		"expected SkipRetry when payload cannot be converted to EmailMessage")
}

func TestProcessSendEmail_TemporaryFailure_ReturnsError(t *testing.T) {
	handler := &worker.SendEmailTaskHandler{
		UseCase: &mockUseCase{
			executeFn: func(emailmessage.EmailMessage) error {
				return emailerrors.ErrTemporaryFailure
			},
		},
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}

	task := validTask(t)

	err := handler.ProcessSendEmail(context.Background(), task)

	require.ErrorIs(t, err, emailerrors.ErrTemporaryFailure,
		"expected temporary failure to be returned to allow retry")
}

func TestProcessSendEmail_NonTemporaryError_ReturnsSkipRetry(t *testing.T) {
	handler := &worker.SendEmailTaskHandler{
		UseCase: &mockUseCase{
			executeFn: func(emailmessage.EmailMessage) error {
				return errors.New("permanent failure")
			},
		},
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}

	task := validTask(t)

	err := handler.ProcessSendEmail(context.Background(), task)

	require.Equal(t, asynq.SkipRetry, err,
		"expected SkipRetry when error is not temporary")
}

func TestProcessSendEmail_Success_ReturnsNil(t *testing.T) {
	handler := &worker.SendEmailTaskHandler{
		UseCase: &mockUseCase{
			executeFn: func(emailmessage.EmailMessage) error {
				return nil
			},
		},
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}

	task := validTask(t)

	err := handler.ProcessSendEmail(context.Background(), task)

	require.NoError(t, err,
		"expected no error when email is successfully processed")
}
