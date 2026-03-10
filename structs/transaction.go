package structs

import "fmt"

type transaction struct {
	data map[string]string
}

func NewTransaction() *transaction {
	return &transaction{
		data: make(map[string]string),
	}
}

func (transaction *transaction) Set(key string, value string) {
	oldValue, ok := transaction.data[key]

    if ok {
		info := fmt.Sprintf("Preparing to overwrite value %s with key %s",oldValue,key)
        fmt.Print(info)
    }

	transaction.data[key] = value
}