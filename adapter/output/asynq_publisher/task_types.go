package emailpublisher

type Payload struct {
	To        string
	Subject   string
	EmailType string
	BodyData  any
}
