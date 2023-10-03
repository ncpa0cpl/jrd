package readers

import (
	ej "encoding/json"

	"github.com/ncpa0cpl/jrd/utils"
)

type JsonObjectReader struct {
	data map[string]interface{}
}

type Entry struct {
	Key   string
	Value *Value
}

func (r *JsonObjectReader) LoadFrom(data []byte) error {
	err := ej.Unmarshal(data, &r.data)
	return err
}

func (r *JsonObjectReader) Type() string {
	return "object"
}

func (r *JsonObjectReader) Get(key string) *Value {
	elementKey, rest := utils.SplitKey(key)

	element, ok := r.data[elementKey]

	if !ok {
		return &Value{valueType: "null", value: nil}
	}

	if rest == "" {
		return &Value{valueType: utils.TypeOf(element), value: element}
	}

	subReader := getReaderFor(element)
	return subReader.Get(rest)
}

func (r *JsonObjectReader) Keys() []string {
	keys := make([]string, len(r.data))

	i := 0
	for key := range r.data {
		keys[i] = key
		i++
	}

	return keys
}

func (r *JsonObjectReader) Entries() []Entry {
	entries := make([]Entry, len(r.data))

	i := 0
	for key, value := range r.data {
		entries[i] = Entry{
			Key: key,
			Value: &Value{
				valueType: utils.TypeOf(value),
				value:     value,
			},
		}
		i++
	}

	return entries
}

func (r *JsonObjectReader) Raw() interface{} {
	return r.data
}
