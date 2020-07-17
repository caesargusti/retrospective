package storage

import "log"

type StorageType int

const (
	// IOTA AWAL = 1, Auto Increment ++
	StorageTypeMemory StorageType = iota
	StorageTypeMock
	StorageTypeDatabase
)

func GetStorage(t StorageType) Storage {
	var s Storage
	switch t {
	case StorageTypeMemory:
		s = newMemory()
	case StorageTypeMock:
		s = mock{}
	case StorageTypeDatabase:
		s = database{}
	default:
		log.Fatal("STORAGE DOESN'T EXIST")
	}
	return s
}
