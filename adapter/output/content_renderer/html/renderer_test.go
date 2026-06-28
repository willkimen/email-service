package htmlrenderer_test

import (
	"emailservice/core/application/apperrors"
	"emailservice/core/application/message"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRender_ParseError(t *testing.T) {
	message := message.Message{
		Id:        "123456",
		Type:      message.MessageTypeNotifyAccountDeleted,
		To:        "email@email.com",
		Variables: make(map[string]any),
	}
	subject, html, err := rendererAdapterParseError.Render(message)

	require.Empty(t, html)
	require.Empty(t, subject)
	require.Error(t, err)

	var targetError *apperrors.InfrastructureError
	if assert.ErrorAs(t, err, &targetError) {
		assert.Contains(t, targetError.Message, "failed to parse email template")
		assert.Error(t, targetError.OriginalCause)
	}
}
