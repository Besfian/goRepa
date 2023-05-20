package goRepa

type Cache struct {
	Email    string
	Password string
	Username string
	Id       int
}

func New() Cache {
	return Cache{}
}

func (cache *Cache) Get() *Cache {
	return cache
}

func (cache *Cache) Set(Email string, Password string, Username string, Id int) {
	cache.Email = Email
	cache.Password = Password
	cache.Username = Username
	cache.Id = Id
}

func (cache *Cache) Delete() {
	cache.Email = ""
	cache.Password = ""
	cache.Username = ""
	cache.Id = 0
}
