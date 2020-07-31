package storage

type (
	StorageItem interface {
		GetData() interface{}
	}

	StorageItems []StorageItem

	StorageMeta interface {
		GetMeta() interface{}
	}

	StorageDriver interface {
		Create(StorageItem) (StorageMeta, error)
		Read(string, string) (StorageItem, StorageMeta, error)
		Update(string, map[string]interface{}, string) (StorageMeta, error)
		Delete(string, string) (StorageMeta, error)
		List(itemType string, offset int, limit int) (StorageItems, error)

		InitDB() error
	}
)
