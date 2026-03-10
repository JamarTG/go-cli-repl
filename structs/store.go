package structs

type store struct{ 
	data map[string]string 
}

func NewStore() *store { 
	return &store{
		data: make(map[string]string),
	} 
}

func (store *store) Set(key string, value string) {
	store.data[key] = value
}

func (store *store) Get(value string) string{
	return store.data[value]
}

