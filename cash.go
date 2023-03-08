package goRepa

type Cache struct {
	userId int
}

func New() Cache {
	return Cache{}
}

func (cache *Cache) Get() int {
	return cache.userId
}

func (cache *Cache) Set(userId int) {
	cache.userId = userId
}

func (cache *Cache) Delete() {
	cache.userId = 0
}
