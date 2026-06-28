/*
Package asynqpublisher implements the output adapters responsible for
publishing background tasks using the Asynq framework.

This package translates application-specific messages into asynchronous
task payloads, serializing them to JSON and enqueueing them into Redis
via Asynq. It acts as an infrastructure layer boundary, shielding
the upper application layers from queue orchestration mechanics and
wrapping operational failures into standardized infrastructure errors.

RECOMMENDATION: Always use the factory function NewAsynqEmailPublisherAdapter to
instantiate the adapter instead of initializing the struct directly.

# Usage Example

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})

	// Recommended approach for production:
	publisher := asynqpublisher.NewAsynqEmailPublisherAdapter(client)

	err := publisher.Publish(message)
*/
package asynqpublisher
