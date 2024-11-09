package storage

import "time"

type Entry struct{
	Value string
	ExpiryTime time.Time
	ExpiryTimeExists bool
}
var Store map[string]Entry

func init() {
    Store = make(map[string]Entry)
}