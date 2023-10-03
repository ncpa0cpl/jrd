package readers

type JsonNullReader struct {
}

func (r *JsonNullReader) Type() string {
	return "null"
}

func (r *JsonNullReader) Get(key string) *Value {
	return &Value{valueType: "null", value: nil}
}

func (r *JsonNullReader) Raw() interface{} {
	return nil
}
