package worker_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"emailservice/adapter/input/worker"
	"emailservice/core/application/email_errors"
	"emailservice/core/application/email_message"

	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/require"
)

type mockUseCase struct {
	executeFn func(emailmessage.EmailMessage) error
}

func (m *mockUseCase) ExecuteSend(msg emailmessage.EmailMessage) error {
	return m.executeFn(msg)
}

func validTask(t *testing.T) *asynq.Task {
	payload, err := json.Marshal(map[string]any{
		"To":        "user@test.com",
		"Subject":   "Activation",
		"EmailType": emailmessage.EmailTypeActivationCode,
		"BodyData": map[string]any{
			"VerificationCode":       "123456",
			"ActivationLink":         "http://link",
			"CodeExpirationHours":    "2",
			"ActivationDeadlineDays": "3",
		},
	})
	require.NoError(t, err)

	return asynq.NewTask("send:email", payload)
}

func TestProcessSendEmail_InvalidPayload_ReturnsSkipRetry(t *testing.T) {
	handler := &worker.SendEmailTaskHandler{
		UseCase: &mockUseCase{},
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
	}

	task := validTask(t)

	err := handler.ProcessSendEmail(context.Background(), task)

	require.NoError(t, err,
		"expected no error when email is successfully processed")
}

