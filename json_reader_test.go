package jrd_test

import (
	"testing"

	"github.com/ncpa0cpl/jrd"
)

var JSON_NULL = "  null  "
var JSON_TRUE = " true"
var JSON_FALSE = "false "
var JSON_NUMBER = "123"
var JSON_STRING = "\n\"hello\""
var JSON_ARRAY = " [1,2,3] "
var JSON_OBJECT = " \n {\"a\":1,\"b\":\"This is B\"} \n "
var JSON_NESTED_OBJECT = `
{
	"foo": {
		"fooAttr": [
			{
				"type": "str",
				"value": "hello"
			},
			{
				"type": "num",
				"value": 123
			}
		],
		"description": "This is foo"
	},
	"bar": [
		{
			"baz": [
				{
					"qux": {
						"someAttr": [1, "2", null, true, false]
					}
				}
			]
		}
	]
}
`

func TestParseNullJson(t *testing.T) {
	reader, _ := jrd.NewJsonReader([]byte(JSON_NULL))
	if reader.Type() != "null" {
		t.Errorf("reader.Type() should be \"null\", but received: %s", reader.Type())
	}
	if reader.Get("").Raw() != nil {
		t.Errorf("reader.Get(\"\").Raw() should be nil")
	}
}

func TestParseBooleanJson(t *testing.T) {
	reader, _ := jrd.NewJsonReader([]byte(JSON_TRUE))
	if reader.Type() != "bool" {
		t.Errorf("reader.Type() should be \"bool\", but received: %s", reader.Type())
	}
	if val, ok := reader.Get("").AsBool(); !ok || val != true {
		t.Errorf("reader.Get(\"\").AsBool() should be true")
	}

	reader, _ = jrd.NewJsonReader([]byte(JSON_FALSE))
	if reader.Type() != "bool" {
		t.Errorf("reader.Type() should be \"bool\", but received: %s", reader.Type())
	}
	if val, ok := reader.Get("").AsBool(); !ok || val != false {
		t.Errorf("reader.Get(\"\").AsBool() should be false")
	}
}

func TestParseNumberJson(t *testing.T) {
	reader, _ := jrd.NewJsonReader([]byte(JSON_NUMBER))
	if reader.Type() != "number" {
		t.Errorf("reader.Type() should be \"number\", but received: %s", reader.Type())
	}
	if val, ok := reader.Get("").AsNumber(); !ok || val != float64(123) {
		t.Errorf("reader.Get(\"\").AsNumber() should be 123, %f", val)
	}
}

func TestParseStringJson(t *testing.T) {
	reader, _ := jrd.NewJsonReader([]byte(JSON_STRING))
	if reader.Type() != "string" {
		t.Errorf("reader.Type() should be \"string\", but received: %s", reader.Type())
	}
	if val, ok := reader.Get("").AsString(); !ok || val != "hello" {
		t.Errorf("reader.Get(\"\").AsString() should be \"hello\"")
	}
}

func TestParseArrayJson(t *testing.T) {
	reader, _ := jrd.NewJsonReader([]byte(JSON_ARRAY))
	if reader.Type() != "array" {
		t.Errorf("reader.Type() should be \"array\", but received: %s", reader.Type())
	}
	if val, ok := reader.Get("").AsArray(); ok {
		if val.Length() != 3 {
			t.Errorf("reader.Get(\"\").AsArray().Length() should be 3, but received: %d", val.Length())
		}

		if elem1, ok := val.Get("0").AsNumber(); !ok || elem1 != float64(1) {
			t.Errorf("reader.Get(\"\").AsArray().Get(\"0\").AsNumber() should be 1")
		}

		if elem2, ok := val.Get("0").AsNumber(); !ok || elem2 != float64(2) {
			t.Errorf("reader.Get(\"\").AsArray().Get(\"0\").AsNumber() should be 2")
		}

		if elem3, ok := val.Get("0").AsNumber(); !ok || elem3 != float64(3) {
			t.Errorf("reader.Get(\"\").AsArray().Get(\"0\").AsNumber() should be 3")
		}
	}
	if elem2, ok := reader.Get("1").AsNumber(); !ok || elem2 != float64(2) {
		t.Errorf("reader.Get(\"0\").Raw() should be 2")
	}
}

func TestParseObjectJson(t *testing.T) {
	reader, _ := jrd.NewJsonReader([]byte(JSON_OBJECT))
	if reader.Type() != "object" {
		t.Errorf("reader.Type() should be \"object\", but received: %s", reader.Type())
	}
	if val, ok := reader.Get("").AsObject(); ok {
		if aVal, ok := val.Get("a").AsNumber(); !ok || aVal != float64(1) {
			t.Errorf("reader.Get(\"\").AsObject().Get(\"a\").AsNumber() should be 1")
		}

		if bVal, ok := val.Get("b").AsString(); !ok || bVal != "This is B" {
			t.Errorf("reader.Get(\"\").AsObject().Get(\"b\").AsString() should be \"This is B\"")
		}
	}
}

func TestParseNestedObjectJson(t *testing.T) {
	reader, _ := jrd.NewJsonReader([]byte(JSON_NESTED_OBJECT))

	fooArr1 := reader.Get("foo.fooAttr.0")
	if fooArr1.Type() != "object" {
		t.Errorf("reader.Get(\"foo.fooAttr.0\").Type() should be \"object\", but received: %s", fooArr1.Type())
	}
	if obj, ok := fooArr1.AsObject(); !ok {
		t.Errorf("reader.Get(\"foo.fooAttr.0\").AsObject() should return the object")
	} else {
		if ovjType, ok := obj.Get("type").AsString(); !ok || ovjType != "str" {
			t.Errorf("reader.Get(\"foo.fooAttr.0\").AsObject().Get(\"type\") should be \"str\"")
		}
		if ovjValue, ok := obj.Get("value").AsString(); !ok || ovjValue != "hello" {
			t.Errorf("reader.Get(\"foo.fooAttr.0\").AsObject().Get(\"value\") should be \"hello\"")
		}
	}
	if elem2Type, ok := reader.Get("foo.fooAttr.1.type").AsString(); !ok || elem2Type != "num" {
		t.Errorf("reader.Get(\"foo.fooAttr.1.type\") should be \"num\"")
	}
	if elem2Value, ok := reader.Get("foo.fooAttr.1.value").AsNumber(); !ok || elem2Value != float64(123) {
		t.Errorf("reader.Get(\"foo.fooAttr.1.value\") should be 123")
	}
	qux := reader.Get("bar.0.baz.0.qux")
	if qux.Type() != "object" {
		t.Errorf("reader.Get(\"bar.0.baz.0.qux\").Type() should be \"object\", but received: %s", qux.Type())
	}
	if dict, ok := qux.AsObject(); !ok {
		t.Errorf("reader.Get(\"bar.0.baz.0.qux\").AsObject() should return the object")
	} else {
		if arr, ok := dict.Get("someAttr").AsArray(); !ok {
			t.Errorf("reader.Get(\"bar.0.baz.0.qux\").AsObject().Get(\"someAttr\").AsArray() should return the array")
		} else {
			if elem0, ok := arr.Get("0").AsNumber(); !ok || elem0 != float64(1) {
				t.Errorf("reader.Get(\"bar.0.baz.0.qux\").AsObject().Get(\"someAttr\").AsArray().Get(\"0\").AsNumber() should be 1")
			}
			if elem1, ok := arr.Get("1").AsString(); !ok || elem1 != "2" {
				t.Errorf("reader.Get(\"bar.0.baz.0.qux\").AsObject().Get(\"someAttr\").AsArray().Get(\"1\").AsString() should be \"2\"")
			}
			if elem2Type := arr.Get("2").Type(); elem2Type != "null" {
				t.Errorf("reader.Get(\"bar.0.baz.0.qux\").AsObject().Get(\"someAttr\").AsArray().Get(\"2\").Type() should be \"null\", but received: %s", elem2Type)
			}
			if elem3, ok := arr.Get("3").AsBool(); !ok || elem3 != true {
				t.Errorf("reader.Get(\"bar.0.baz.0.qux\").AsObject().Get(\"someAttr\").AsArray().Get(\"3\").AsBool() should be true")
			}
			if elem4, ok := arr.Get("4").AsBool(); !ok || elem4 != false {
				t.Errorf("reader.Get(\"bar.0.baz.0.qux\").AsObject().Get(\"someAttr\").AsArray().Get(\"4\").AsBool() should be false")
			}
		}
	}
}
