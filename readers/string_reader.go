package readers

import ej "encoding/json"

type JsonStringReader struct {
	data string
}

func (r *JsonStringReader) LoadFrom(data []byte) error {
	err := ej.Unmarshal(data, &r.data)
	return err
}

func (r *JsonStringReader) Type() string {
	return "string"
}

func (r *JsonStringReader) Get(key string) *Value {
	switch key {
	case "":
		return &Value{valueType: "string", value: r.data}
	default:
		panic("Invalid key, cannot read properties of a string")
	}
}

func (r *JsonStringReader) Raw() interface{} {
	return r.data
}
