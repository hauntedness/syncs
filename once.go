package syncs

import "sync"

var once sync.Once

func Once(fn func()) {
	once.Do(fn)
}
