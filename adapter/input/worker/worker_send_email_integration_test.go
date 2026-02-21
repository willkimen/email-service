//go:build slow

package worker_test

import (
	"emailservice/adapter/input/worker"
	"emailservice/core/application/email_errors"
	"emailservice/core/application/email_message"
	"errors"
	"testing"
	"time"

	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
)

func TestProcessSendEmail_PermanentFailure_archiveTask(t *testing.T) {
	handler := &worker.SendEmailTaskHandler{
		UseCase: &mockUseCase{
			executeFn: func(emailmessage.EmailMessage) error {
				return errors.New("permanent failure")
			},
		},
	}

	redis, addr := setupIntegration(t, *handler)

	defer testcontainers.CleanupContainer(t, redis)

	inspector := asynq.NewInspector(asynq.RedisClientOpt{
		Addr: addr,
	})
	defer inspector.Close()

	require.Eventually(t, func() bool {
		tasks, err := inspector.ListArchivedTasks("default")
		if err != nil {
			return false
		}

		return len(tasks) == 1
	}, 5*time.Second, 100*time.Millisecond)

	tasks, err := inspector.ListArchivedTasks("default")
	require.NoError(
		t,
		err,
		"expected inspector to list pending tasks without error",
	)

	require.Len(
		t,
		tasks,
		1,
		"expected exactly one archived task in the default queue",
	)

	assert.Equal(
		t,
		taskType,
		tasks[0].Type,
		"expected archived task type to match send email task type",
	)

	require.NotEmpty(
		t,
		tasks[0].Payload,
		"expected archived task payload to be non-empty",
	)

}

func TestProcessSendEmail_TemporaryFailure_retryTask(t *testing.T) {
	handler := &worker.SendEmailTaskHandler{
		UseCase: &mockUseCase{
			executeFn: func(emailmessage.EmailMessage) error {
				return emailerrors.ErrTemporaryFailure
			},
		},
	}

	redis, addr := setupIntegration(t, *handler)
	defer testcontainers.CleanupContainer(t, redis)

	inspector := asynq.NewInspector(asynq.RedisClientOpt{
		Addr: addr,
	})
	defer inspector.Close()

	require.Eventually(t, func() bool {
		tasks, err := inspector.ListRetryTasks("default")
		if err != nil {
			return false
		}

		return len(tasks) == 1
	}, 5*time.Second, 100*time.Millisecond)

	tasks, err := inspector.ListRetryTasks("default")
	require.NoError(
		t,
		err,
		"expected inspector to list retry tasks without error",
	)

	require.Len(
		t,
		tasks,
		1,
		"expected exactly one retry task in the default queue",
	)

	assert.Equal(
		t,
		taskType,
		tasks[0].Type,
		"expected retry task type to match send email task type",
	)

	require.NotEmpty(
		t,
		tasks[0].Payload,
		"expected retry task payload to be non-empty",
	)
}

func TestProcessSendEmail_Success_taskIsProcessed(t *testing.T) {
	handler := &worker.SendEmailTaskHandler{
		UseCase: &mockUseCase{
			executeFn: func(emailmessage.EmailMessage) error {
				return nil
			},
		},
	}

	redis, addr := setupIntegration(t, *handler)
	defer testcontainers.CleanupContainer(t, redis)

	inspector := asynq.NewInspector(asynq.RedisClientOpt{
		Addr: addr,
	})
	defer inspector.Close()

	// In asynq, successful tasks are not persisted.
	// A task is considered "completed" when it does NOT appear
	// in retry or archived queues after processing.
	require.Eventually(t, func() bool {
		archived, err := inspector.ListArchivedTasks("default")
		if err != nil {
			return false
		}

		retry, err := inspector.ListRetryTasks("default")
		if err != nil {
			return false
		}

		return len(archived) == 0 && len(retry) == 0
	}, 5*time.Second, 100*time.Millisecond)

	archived, err := inspector.ListArchivedTasks("default")
	require.NoError(
		t,
		err,
		"expected inspector to list archived tasks without error",
	)
	require.Len(
		t,
		archived,
		0,
		"expected no archived tasks after successful processing",
	)

	retry, err := inspector.ListRetryTasks("default")
	require.NoError(
		t,
		err,
		"expected inspector to list retry tasks without error",
	)
	require.Len(
		t,
		retry,
		0,
		"expected no retry tasks after successful processing",
	)
}
