/*
Package message provides the Message type, which represents the data
required to send an email to a user, as well as the data needed to
be enqueued.

It is strongly recommended to use this factory function (or workflow-specific
factories) instead of instantiating the Message struct directly, as it centralizes
and encapsulates all validation rules required for a Message entity.

Example:

	variables := map[string]any{"code": "123456"}
	msg, err := message.NewMessage(
		"msg-123",
		"user@example.com",
		message.MessageTypeEmailVerificationCode,
		variables
	)
	if err != nil {
	    // handle validation error
	}

Any provided message type must strictly match one of the system's pre-existing
constants (e.g., MessageTypeEmailVerificationCode).

The term "message" was taken from the context of Message Queueing or Message Brokers.
In software architectures, a message is essentially the data packet
sent from one system (the producer) to another (the consumer).
*/
package message
