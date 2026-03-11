package structs

import "errors"

type kv struct {
	transactions []transaction
	store        *store
}

func NewKv() *kv {
	return &kv{
		transactions: make([]transaction, 0),
		store:        NewStore(),
	}
}

func (kv *kv) Set(key string, value string) error {

	if len(kv.transactions) == 0 {
		err := errors.New("kv not long enough")
		return err
	}

	kv.transactions[len(kv.transactions)-1].data[key] = value

	return nil
}

func (kv *kv) Get(key string) *string {

	for _,transaction := range kv.transactions {

		transactionValue, found := transaction.data[key]

		if found {
			return &transactionValue
		}

	}

	v,ok := NewStore().data[key]
	
	if(ok) {
		return &v
	}
	
	return nil
}

func (kv *kv)begin() *[]transaction {
	trx := NewTransaction()
	kv.transactions = append(kv.transactions,*trx)
	return &kv.transactions
}
