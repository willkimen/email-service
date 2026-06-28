package asynqpublisher_test

import (
	"context"
	"emailservice/core/application/message"
	"errors"
	"math"
	"testing"

	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// ======================= Fake Message dependency ============================

const (
	id               = "fake-id"
	to               = "user@email.com"
	typeMessage      = message.MessageTypeEmailVerificationCode
	verificationCode = "123456"
)

var fakeMessageToMarshalError = message.Message{
	Id:   id,
	To:   to,
	Type: typeMessage,
	Variables: map[string]any{
		"invalid_number": math.Inf(1),
	},
}

var fakeMessageCorrect = message.Message{
	Id:   id,
	To:   to,
	Type: typeMessage,
	Variables: map[string]any{
		"verification_code": "123456",
	},
}

// ========================= Fake Client dependency =======================

type fakeClient struct {
	enqueueFunc func(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error)
}

func (f *fakeClient) Enqueue(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	return f.enqueueFunc(task, opts...)
}

var fakeClientSuccess = &fakeClient{
	enqueueFunc: func(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
		return &asynq.TaskInfo{ID: "123"}, nil
	},
}

var fakeErrorEnqueue = errors.New("enqueue failed")
var fakeClientFail = &fakeClient{
	enqueueFunc: func(task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
		return nil, fakeErrorEnqueue
	},
}

// ========================== Configration to integration tests ==============================

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

func NewAsynqClient(addr string) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{
		Addr: addr,
	})
}
