package storage

type (
	StorageItem interface {
		GetData() interface{}
	}

	StorageMeta interface {
		GetMeta() interface{}
	}

	StorageDriver interface {
		Create(StorageItem) (StorageMeta, error)
		Read(string, string) (StorageItem, StorageMeta, error)
		Update(string, map[string]interface{}, string) (StorageMeta, error)
		Delete(string, string) (StorageMeta, error)

		Init() error
	}
)
