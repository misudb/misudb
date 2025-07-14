package kv

import "github.com/huandu/skiplist"

type KV struct {
	store skiplist.SkipList
}

func (k *KV) Get(key string) (string, bool) {
	value, ok := k.store.GetValue(key)
	if !ok {
		return "", false
	}
	return value.(string), true
}

func (k *KV) Set(key, value string) {
	k.store.Set(key, value)
}

func (k *KV) Delete(key string) {
	k.store.Remove(key)
}

func New() *KV {
	return &KV{
		store: *skiplist.New(skiplist.String),
	}
}
