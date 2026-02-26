//go:build slow

package emailpublisher_test

import (
	"emailservice/adapter/output/asynq_publisher"
	"encoding/json"
	"testing"

	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/assert"
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
		Logger: logger,
	}

	// Execute the method under test.
	// This should serialize the message and enqueue a task.
	err := adapter.Publish(fakeMessage{})

	// Ensure the Redis container is cleaned up after the test finishes.
	defer testcontainers.CleanupContainer(t, redis)

	// Assert that Publish returned no error.
	require.NoError(
		t,
		err,
		"expected Publish to enqueue task without returning error",
	)

	// Create an Asynq Inspector connected to the same Redis instance.
	// The Inspector allows querying the state of queues.
	inspector := asynq.NewInspector(asynq.RedisClientOpt{
		Addr: addr,
	})
	defer inspector.Close()

	// Retrieve all pending tasks from the default queue.
	tasks, err := inspector.ListPendingTasks("default")
	require.NoError(
		t,
		err,
		"expected inspector to list pending tasks without error",
	)

	// Assert that exactly one task was enqueued.
	require.Len(
		t,
		tasks,
		1,
		"expected exactly one task to be enqueued in the default queue",
	)

	// Assert that the task type matches the expected identifier.
	assert.Equal(
		t,
		"email:send",
		tasks[0].Type,
		"expected enqueued task type to be 'email:send'",
	)

	// Ensure that the task payload is not empty.
	require.NotEmpty(
		t,
		tasks[0].Payload,
		"expected enqueued task payload to be non-empty",
	)

	// Deserialize the task payload back into the expected struct.
	// This verifies that the JSON serialization was correct.
	var p emailpublisher.Payload
	err = json.Unmarshal(tasks[0].Payload, &p)
	require.NoError(
		t,
		err,
		"expected task payload to deserialize without error",
	)

	// Validate that all fields were correctly serialized and preserved.
	assert.Equal(t, to, p.To, "expected payload To field to match input")
	assert.Equal(t, subject, p.Subject, "expected payload Subject field to match input")
	assert.Equal(t, emailType, p.EmailType, "expected payload EmailType field to match input")
	assert.Equal(t, bodyData, p.BodyData, "expected payload BodyData to match input")
}
