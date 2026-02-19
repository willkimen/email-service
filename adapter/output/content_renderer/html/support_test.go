package renderer_test

import "emailservice/adapter/output/content_renderer/html"

var rendererAdapter = &renderer.HTMLEmailContentRendererAdapter{}

type FakeEmailMessageWithEmailTypeNotExist struct{}

func (FakeEmailMessageWithEmailTypeNotExist) ValidateData() error  { return nil }
func (FakeEmailMessageWithEmailTypeNotExist) GetEmailType() string { return "EmailTypeNotExist" }
func (FakeEmailMessageWithEmailTypeNotExist) GetTo() string        { return "to" }
func (FakeEmailMessageWithEmailTypeNotExist) GetSubject() string   { return "subject" }
func (FakeEmailMessageWithEmailTypeNotExist) GetBodyData() any     { return nil }

type FakeEmailMessageWithDataInvalid struct {
	FieldNotExist string
}

func (FakeEmailMessageWithDataInvalid) ValidateData() error  { return nil }
func (FakeEmailMessageWithDataInvalid) GetEmailType() string { return "activation_code" }
func (FakeEmailMessageWithDataInvalid) GetTo() string        { return "to" }
func (FakeEmailMessageWithDataInvalid) GetSubject() string   { return "subject" }
func (FakeEmailMessageWithDataInvalid) GetBodyData() any     { return nil }
