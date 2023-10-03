package readers

import (
	ej "encoding/json"
	"strconv"

	"github.com/ncpa0cpl/jrd/utils"
)

type JsonArrayReader struct {
	data []interface{}
}

func (r *JsonArrayReader) LoadFrom(data []byte) error {
	err := ej.Unmarshal(data, &r.data)
	return err
}

func (r *JsonArrayReader) Type() string {
	return "array"
}

func (r *JsonArrayReader) Get(key string) *Value {
	indexStr, rest := utils.SplitKey(key)

	index, err := strconv.Atoi(indexStr)

	if err != nil {
		return &Value{valueType: "null", value: nil}
	}

	if index >= len(r.data) {
		return &Value{valueType: "null", value: nil}
	}

	element := r.data[index]

	if rest == "" {
		return &Value{valueType: utils.TypeOf(element), value: element}
	}

	subReader := getReaderFor(element)
	return subReader.Get(rest)
}

func (r *JsonArrayReader) Length() int {
	return len(r.data)
}

func (r *JsonArrayReader) Raw() interface{} {
	return r.data
}
