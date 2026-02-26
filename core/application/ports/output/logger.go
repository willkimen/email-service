package outputport

type Logger interface {
	Info(msg string, fields ...any)
	Error(msg string, err error, fields ...any)
}
