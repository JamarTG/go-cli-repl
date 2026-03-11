package structs

type Kv struct {
	transactions []transaction
	store        *store
}

func NewKv() *Kv {
	return &Kv{
		transactions: make([]transaction, 0),
		store:        NewStore(),
	}
}

func (kv *Kv) Set(key string, value string) error {

	if len(kv.transactions) == 0 {
		kv.store.data[key] = value
	}

	kv.transactions[len(kv.transactions)-1].data[key] = value

	return nil
}

func (kv *Kv) Get(key string) *string {

	for _, transaction := range kv.transactions {

		transactionValue, found := transaction.data[key]

		if found {

			return &transactionValue
		}

	}

	v, ok := kv.store.data[key]

	if ok {
		return &v
	}

	return nil
}

func (kv *Kv) begin() []transaction {
	kv.transactions = append(kv.transactions, *NewTransaction())
	return kv.transactions
}

func (kv *Kv) delete(key string) {

	// r := kv.Get(key)

	// we look at the end of the transactions

	// if len(kv.transactions) == 0 {
	// 	v, ok := kv.store.data[key]

	// 	if ok {
	// 		return kv[]
	// 	}

	// }
	// we find the first occurence of the transactions with the desired key
	// traversing in reverse order
	//
}
