package drivers

type (
	UnconfiguredStorageBackend struct {}
)

func (e *UnconfiguredStorageBackend) Error() string {
	return "No configured Storage Backends"
}
