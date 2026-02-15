package emailpublisher_test

import (
	"context"
	"testing"

	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	to        = "user@email.com"
	subject   = "subject fake"
	emailType = "activation"
	bodyData  = "body data fake"
)

type fakeMessage struct{}

func (fakeMessage) GetTo() string {
	return to
}

func (fakeMessage) GetSubject() string {
	return subject
}

func (fakeMessage) GetEmailType() string {
	return emailType
}

func (fakeMessage) GetBodyData() any {
	return bodyData
}

func (fakeMessage) ValidateData() error { return nil }

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
	require.NoError(t, err)

	addr, err := redis.Endpoint(ctx, "")
	require.NoError(t, err)

	return redis, addr
}

func NewAsynqClient(addr string) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{
		Addr: addr,
	})
}
