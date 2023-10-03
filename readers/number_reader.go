package readers

import ej "encoding/json"

type JsonNumberReader struct {
	data float64
}

func (r *JsonNumberReader) LoadFrom(data []byte) error {
	err := ej.Unmarshal(data, &r.data)
	return err
}

func (r *JsonNumberReader) Type() string {
	return "number"
}

func (r *JsonNumberReader) Get(key string) *Value {
	switch key {
	case "":
		return &Value{valueType: "number", value: r.data}
	default:
		panic("Invalid key, cannot read properties of a number")
	}
}

func (r *JsonNumberReader) Raw() interface{} {
	return r.data
}
