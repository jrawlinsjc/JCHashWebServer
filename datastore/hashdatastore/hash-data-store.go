package hashdatastore

import "github.com/jmrawlins/JCHashWebServer/hash"

type HashDataStore interface {
	GetNextId() (id hash.HashId, err error)
	StoreHash(id hash.HashId, hash string) error
	GetHash(id hash.HashId) (string, error)
	GetAllHashes() *map[hash.HashId]string
}