package store

import "time"

type (
	//IStore defines a simple KV storage interface
	IStore interface {
		//Put store bytes against a key for a given time
		Put(key string, value []byte, TTLSecs time.Duration) error

		//Get get keyed content if exists and has not expired
		Get(key string) (value []byte, exists bool, err error)

		//Del delete keyed content
		Del(key string) error
	}
)
