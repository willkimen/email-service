package renderer

var renderer = &HTMLEmailContentRendererAdapter{}

type FakeEmailMessageWithTemplateIDNotExist struct{}

func (FakeEmailMessageWithTemplateIDNotExist) ValidateData() error { return nil }
func (FakeEmailMessageWithTemplateIDNotExist) TemplateID() string  { return "TemplateNotExist" }
func (FakeEmailMessageWithTemplateIDNotExist) GetTo() string       { return "to" }
func (FakeEmailMessageWithTemplateIDNotExist) GetSubject() string  { return "subject" }

type FakeEmailMessageWithDataInvalid struct {
	FieldNotExist string
}

func (FakeEmailMessageWithDataInvalid) ValidateData() error { return nil }
func (FakeEmailMessageWithDataInvalid) TemplateID() string  { return "activation_code" }
func (FakeEmailMessageWithDataInvalid) GetTo() string       { return "to" }
func (FakeEmailMessageWithDataInvalid) GetSubject() string  { return "subject" }
