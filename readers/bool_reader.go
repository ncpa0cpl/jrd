package readers

import ej "encoding/json"

type JsonBoolReader struct {
	data bool
}

func (r *JsonBoolReader) LoadFrom(data []byte) error {
	err := ej.Unmarshal(data, &r.data)
	return err
}

func (r *JsonBoolReader) Type() string {
	return "bool"
}

func (r *JsonBoolReader) Get(key string) *Value {
	switch key {
	case "":
		return &Value{valueType: "bool", value: r.data}
	default:
		panic("Invalid key, cannot read properties of a boolean")
	}
}

func (r *JsonBoolReader) Raw() interface{} {
	return r.data
}
