package arangodb

type (
	UnknownAuthMethodError struct {}
)

func (e UnknownAuthMethodError) Error() string {
	return "unknown auth method"
}
