package structs
import "errors";

type kv struct {
	transactions []map[string]string 
}

func Newkv() *kv {
	return &kv{
		transactions: make([]map[string]string, 0),
	}
}

func (kv *kv) Set(key string, value string) error {
	
	if(len(kv.transactions) == 0) {
		err := errors.New("kv not long enough")
		return err
	}
	
	kv.transactions[len(kv.transactions)-1][key] = value
	
	return nil
}

func (kv *kv) Get() {
	
}