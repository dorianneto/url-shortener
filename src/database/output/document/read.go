package document

import "strings"

// TODO: add a method to interact with data
type ReadOutput struct {
	Data map[string]interface{}
}

func (r *ReadOutput) GetByKey(key string) string {
	var result string

	for k := range r.Data {
		if found := strings.EqualFold(key, k); found {
			result = k
			break
		}
	}

	return r.Data[result].(string)
}
