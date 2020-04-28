package storage

type (
	Storage struct {
		driver StorageDriver
	}

	StorageDriver interface {

	}
)

func NewStorage(driver StorageDriver) *Storage {
	s := new(Storage)
	s.driver = driver

	return s
}
