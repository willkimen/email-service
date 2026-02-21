package worker_test

import (
	"context"
	"emailservice/adapter/input/worker"
	"emailservice/core/application/email_message"
	"encoding/json"
	"testing"

	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const taskType = "email:send"

type mockUseCase struct {
	executeFn func(emailmessage.EmailMessage) error
}

func (m *mockUseCase) ExecuteSend(message emailmessage.EmailMessage) error {
	return m.executeFn(message)
}

// validTask builds a well-formed Asynq task payload representing
// a valid send email request.
//
// This helper centralizes payload creation to ensure all integration
// tests enqueue structurally correct tasks.
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
	require.NoError(
		t,
		err,
		"failed to marshal asynq task payload for send email task",
	)

	return asynq.NewTask(taskType, payload)
}

// upClient creates an Asynq client and enqueues a single send email task
// into the default queue.
//
// This helper is intentionally minimal and exists to trigger the
// background processing flow during integration tests.
func upClient(t *testing.T, addr string) {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: addr,
	})

	task := validTask(t)

	_, err := client.Enqueue(task)
	require.NoError(
		t,
		err,
		"failed to enqueue send email task into asynq",
	)

}

// upServer starts an Asynq server configured with the send email task
// handler and runs it asynchronously.
//
// The server lifecycle is tied to the test execution and is not
// explicitly shut down, as the Redis container teardown implicitly
// stops task processing.
func upServer(t *testing.T, addr string, handler worker.SendEmailTaskHandler) *asynq.Server {

	mux := asynq.NewServeMux()
	mux.HandleFunc(taskType, handler.ProcessSendEmail)

	server := asynq.NewServer(asynq.RedisClientOpt{Addr: addr}, asynq.Config{
		Concurrency: 1,
	})

	go func() {
		if err := server.Run(mux); err != nil {
			require.NoError(
				t,
				err,
				"asynq server failed to start or stopped unexpectedly",
			)
		}
	}()

	return server
}

// RunRedisContainer starts a disposable Redis container for integration
// testing and waits until it is ready to accept connections.
//
// Each invocation returns an isolated Redis instance to prevent
// cross-test interference.
func RunRedisContainer(t *testing.T) (*testcontainers.DockerContainer, string) {
	ctx := context.Background()
	redis, err := testcontainers.Run(
		ctx, "redis:7-alpine",
		testcontainers.WithExposedPorts("6379/tcp"),
		testcontainers.WithWaitStrategy(
			wait.ForListeningPort("6379/tcp"),
			wait.ForLog("Ready to accept connections"),
		),
	)
	require.NoError(
		t,
		err,
		"expected Redis test container to start successfully",
	)

	addr, err := redis.Endpoint(ctx, "")
	require.NoError(
		t,
		err,
		"expected to retrieve Redis container endpoint without error",
	)

	return redis, addr
}

// setupIntegration orchestrates the full integration test setup by starting
// a Redis container, running the Asynq server, and enqueueing
// an initial send email task.
//
// This helper provides a clean and deterministic environment
// for each integration test scenario.
func setupIntegration(t *testing.T, handler worker.SendEmailTaskHandler) (
	*testcontainers.DockerContainer, string) {
	redis, addr := RunRedisContainer(t)
	upServer(t, addr, handler)
	upClient(t, addr)
	return redis, addr
}
