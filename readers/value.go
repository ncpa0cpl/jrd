package readers

type IJsonReader interface {
	Get(key string) *Value
	Raw() interface{}
	Type() string
}

type IJsonObjectReader interface {
	Entries() []Entry
	Get(key string) *Value
	Keys() []string
	Raw() interface{}
	Type() string
}

type IJsonArrayReader interface {
	Get(key string) *Value
	Length() int
	Raw() interface{}
	Type() string
}

type Value struct {
	value     interface{}
	valueType string
}

func (g *Value) Type() string {
	return g.valueType
}

func (g *Value) AsString() (string, bool) {
	if g.Type() != "string" {
		return "", false
	}

	return g.value.(string), true
}

func (g *Value) AsNumber() (float64, bool) {
	if g.Type() != "number" {
		return 0, false
	}

	return g.value.(float64), true
}

func (g *Value) AsBool() (bool, bool) {
	if g.Type() != "bool" {
		return false, false
	}

	return g.value.(bool), true
}

func (g *Value) AsArray() (IJsonArrayReader, bool) {
	if g.Type() != "array" {
		return nil, false
	}

	reader := &JsonArrayReader{data: g.value.([]interface{})}

	return reader, true
}

func (g *Value) AsObject() (IJsonObjectReader, bool) {
	if g.Type() != "object" {
		return nil, false
	}

	reader := &JsonObjectReader{data: g.value.(map[string]interface{})}

	return reader, true
}

func (g *Value) Raw() interface{} {
	return g.value
}

func getReaderFor(v interface{}) IJsonReader {
	switch v.(type) {
	case []interface{}:
		return &JsonArrayReader{data: v.([]interface{})}
	case map[string]interface{}:
		return &JsonObjectReader{data: v.(map[string]interface{})}
	case string:
		return &JsonStringReader{data: v.(string)}
	case float64:
		return &JsonNumberReader{data: v.(float64)}
	case bool:
		return &JsonBoolReader{data: v.(bool)}
	case nil:
		return &JsonNullReader{}
	}

	panic("Invalid type")
}
