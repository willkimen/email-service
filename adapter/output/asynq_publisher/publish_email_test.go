package asynqpublisher_test

import (
	"emailservice/adapter/output/asynq_publisher"
	"emailservice/core/application/apperrors"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPublish_Success(t *testing.T) {
	adapter := &asynqpublisher.AsynqEmailPublisherAdapter{
		Client: fakeClientSuccess,
	}

	err := adapter.Publish(fakeMessageCorrect)

	require.Nil(t, err)
}

func TestPublish_Return_EnqueueError(t *testing.T) {
	adapter := &asynqpublisher.AsynqEmailPublisherAdapter{
		Client: fakeClientFail,
	}

	actualErr := adapter.Publish(fakeMessageCorrect)

	require.Error(t, actualErr)

	var targetErr *apperrors.InfrastructureError
	require.ErrorAs(t, actualErr, &targetErr)
	require.Contains(t, targetErr.Message, "failed to enqueue message")
	require.ErrorIs(t, targetErr.Unwrap(), fakeErrorEnqueue)
}

func TestPublish_Return_MarshalError(t *testing.T) {
	adapter := &asynqpublisher.AsynqEmailPublisherAdapter{
		Client: fakeClientSuccess,
	}

	actualErr := adapter.Publish(fakeMessageToMarshalError)

	require.Error(t, actualErr)

	var targetErr *apperrors.InfrastructureError
	require.ErrorAs(t, actualErr, &targetErr)
	require.Contains(t, targetErr.Message, "failed to serialize message in JSON")

	var jsonErr *json.UnsupportedValueError
	require.ErrorAs(t, targetErr.Unwrap(), &jsonErr)
}
