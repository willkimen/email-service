package renderer_test

import (
	"emailservice/adapter/output/content_renderer/html"
	"log/slog"
	"os"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

var rendererAdapter = &renderer.HTMLEmailContentRendererAdapter{Logger: logger}

type FakeEmailMessageWithEmailTypeNotExist struct{}

func (FakeEmailMessageWithEmailTypeNotExist) ValidateData() error  { return nil }
func (FakeEmailMessageWithEmailTypeNotExist) GetEmailType() string { return "EmailTypeNotExist" }
func (FakeEmailMessageWithEmailTypeNotExist) GetTo() string        { return "to" }
func (FakeEmailMessageWithEmailTypeNotExist) GetSubject() string   { return "subject" }
func (FakeEmailMessageWithEmailTypeNotExist) GetBodyData() any     { return nil }

type FakeEmailMessageWithDataInvalid struct {
	FieldNotExist string
}

var emailTypeExists = "email_verification_code"

func (FakeEmailMessageWithDataInvalid) ValidateData() error  { return nil }
func (FakeEmailMessageWithDataInvalid) GetEmailType() string { return emailTypeExists }
func (FakeEmailMessageWithDataInvalid) GetTo() string        { return "to" }
func (FakeEmailMessageWithDataInvalid) GetSubject() string   { return "subject" }
func (FakeEmailMessageWithDataInvalid) GetBodyData() any     { return nil }
