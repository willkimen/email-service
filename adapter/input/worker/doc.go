/*
Package worker implements the background task consumers and background processing orchestration.

It acts as a bridge between the asynchronous task execution framework (Asynq/Redis) and the core
application layer, listening for queued events and translating them into domain execution calls.

# Architectural Role & Responsibility

The worker package is strictly an edge component. Its structural responsibilities are limited to:
  - Consuming raw tasks from infrastructure queues.
  - Deserializing wire-format payloads (e.g., JSON) into strongly typed internal application messages.
  - Routing execution control immediately to the application core via input ports (driving interfaces).
  - Translating domain or infrastructure error policies into framework-specific retry strategies
    (such as signaling temporary retries or skipping permanent failures via asynq.SkipRetry).

By isolating queue mechanics here, the core use cases remain completely unaware of the worker
infrastructure, task states, or underlying message distribution technology.
*/
package worker
