package validator

type (
	Validatable interface {
		IsValid(Error)
	}
)

// Validate validates the given interface
func Validate(val Validatable) Error {
	err := NewError()

	val.IsValid(err)

	if len(err.Errors()) > 0 {
		return err
	}

	return nil
}
