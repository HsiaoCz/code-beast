package cache

import "time"

type Cacher interface {
	Set([]byte, []byte, time.Duration) error
	Hash([]byte) bool
	Get([]byte) ([]byte, error)
	Delete([]byte) error
}
