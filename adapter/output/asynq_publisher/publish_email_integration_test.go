package emailpublisher_test

import (
	"emailservice/adapter/output/asynq_publisher"
	"encoding/json"
	"testing"

	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
)

// TestPublish_Integration verifies the full integration flow of the
// AsynqEmailPublisherAdapter.
//
// This test:
//  1. Spins up a real Redis container using Testcontainers.
//  2. Creates a real Asynq client connected to that Redis instance.
//  3. Calls Publish to enqueue a task.
//  4. Uses Asynq Inspector to assert that the task was actually enqueued.
//  5. Validates the serialized payload contents.
func TestPublish_Integration(t *testing.T) {

	// Start a real Redis container for integration testing.
	// Returns the container instance and its connection address.
	redis, addr := RunRedisContainer(t)

	// Create a real Asynq client connected to the test Redis instance.
	client := NewAsynqClient(addr)
	defer client.Close()

	// Create the adapter under test, injecting the real Asynq client.
	adapter := emailpublisher.AsynqEmailPublisherAdapter{
		Client: client,
	}

	// Execute the method under test.
	// This should serialize the message and enqueue a task.
	err := adapter.Publish(fakeMessage{})

	// Ensure the Redis container is cleaned up after the test finishes.
	defer testcontainers.CleanupContainer(t, redis)

	// Assert that Publish returned no error.
	require.NoError(t, err)

	// Create an Asynq Inspector connected to the same Redis instance.
	// The Inspector allows querying the state of queues.
	inspector := asynq.NewInspector(asynq.RedisClientOpt{
		Addr: addr,
	})
	defer inspector.Close()

	// Retrieve all pending tasks from the default queue.
	tasks, err := inspector.ListPendingTasks("default")
	require.NoError(t, err)

	// Assert that exactly one task was enqueued.
	require.Len(t, tasks, 1)

	// Assert that the task type matches the expected identifier.
	require.Equal(t, "email:send", tasks[0].Type)

	// Ensure that the task payload is not empty.
	require.NotEmpty(t, tasks[0].Payload)

	// Deserialize the task payload back into the expected struct.
	// This verifies that the JSON serialization was correct.
	var p emailpublisher.Payload
	err = json.Unmarshal(tasks[0].Payload, &p)
	require.NoError(t, err)

	// Validate that all fields were correctly serialized and preserved.
	require.Equal(t, to, p.To)
	require.Equal(t, subject, p.Subject)
	require.Equal(t, emailType, p.EmailType)
	require.Equal(t, bodyData, p.BodyData)
}
