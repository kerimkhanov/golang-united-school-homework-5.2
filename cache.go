package cache

import "time"

type Cache struct {
	Hash map[string]user
}

func NewCache() Cache {
	return Cache{
		Hash: make(map[string]user),
	}
}

type user struct {
	value    string
	deadline time.Time
}

func (c Cache) Get(key string) (string, bool) {
	if val, ok := c.Hash[key]; ok {
		if time.Now().Before(val.deadline) || val.deadline.IsZero() {
			return val.value, true
		}
	}
	return "", false

}

func (c Cache) Put(key, value string) {
	c.Hash[key] = user{
		value:    value,
		deadline: time.Time{},
	}
}

func (c Cache) Keys() []string {
	arr := []string{}
	for key, value := range c.Hash {
		if value.deadline.IsZero() || time.Now().Before(value.deadline) {
			arr = append(arr, key)
		}
	}
	return arr
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.Hash[key] = user{
		value:    value,
		deadline: deadline,
	}
}
