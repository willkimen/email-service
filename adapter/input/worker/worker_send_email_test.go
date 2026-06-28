package worker_test

import (
	"context"
	"errors"
	"testing"

	"emailservice/adapter/input/worker"
	"emailservice/core/application/apperrors"
	"emailservice/core/application/message"

	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/require"
)

func TestProcessSendEmail_Success(t *testing.T) {
	handler := &worker.SendEmailTaskHandler{
		UseCase: &mockUseCase{
			executeFn: func(message.Message) error {
				return nil
			},
		},
		Logger: LoggerDummy,
	}

	err := handler.ProcessSendEmail(context.Background(), validTask(t))

	require.NoError(t, err)
}

func TestProcessSendEmail_InvalidPayload_ReturnsSkipRetry(t *testing.T) {
	handler := &worker.SendEmailTaskHandler{
		UseCase: &mockUseCase{},
		Logger:  LoggerDummy,
	}

	task := asynq.NewTask("send:email", []byte("invalid-json"))

	err := handler.ProcessSendEmail(context.Background(), task)

	require.ErrorIs(t, err, asynq.SkipRetry)
}

func TestProcessSendEmail_TemporaryFailure_ReturnsError(t *testing.T) {
	handler := &worker.SendEmailTaskHandler{
		UseCase: &mockUseCase{
			executeFn: func(message.Message) error {
				return apperrors.ErrTemporaryFailure
			},
		},
		Logger: LoggerDummy,
	}

	err := handler.ProcessSendEmail(context.Background(), validTask(t))

	require.ErrorIs(t, err, apperrors.ErrTemporaryFailure)
}

func TestProcessSendEmail_NonTemporaryError_ReturnsSkipRetry(t *testing.T) {
	handler := &worker.SendEmailTaskHandler{
		UseCase: &mockUseCase{
			executeFn: func(message.Message) error {
				return errors.New("permanent failure")
			},
		},
		Logger: LoggerDummy,
	}

	err := handler.ProcessSendEmail(context.Background(), validTask(t))

	require.ErrorIs(t, err, asynq.SkipRetry)
}
