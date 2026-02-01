package renderer

var renderer = &HTMLEmailRenderer{}

type FakeEmailMessageWithTemplateIDNotExist struct{}

func (FakeEmailMessageWithTemplateIDNotExist) ValidateData() error { return nil }
func (FakeEmailMessageWithTemplateIDNotExist) TemplateID() string  { return "TemplateNotExist" }

type FakeEmailMessageWithDataInvalid struct {
	FieldNotExist string
}

func (FakeEmailMessageWithDataInvalid) ValidateData() error { return nil }
func (FakeEmailMessageWithDataInvalid) TemplateID() string  { return "activation_code" }
