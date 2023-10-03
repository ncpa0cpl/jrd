package jrd

import (
	"github.com/ncpa0cpl/jrd/readers"
	"github.com/ncpa0cpl/jrd/utils"
)

var JSON_TRUE = []byte("true")
var JSON_FALSE = []byte("false")
var JSON_NULL = []byte("null")

func NewJsonReader(json []byte) (readers.IJsonReader, error) {
	json = utils.Trim(json)

	if json[0] == '[' {
		reader := readers.JsonArrayReader{}
		err := reader.LoadFrom(json)
		return &reader, err
	}

	if json[0] == '{' {
		reader := readers.JsonObjectReader{}
		err := reader.LoadFrom(json)
		return &reader, err
	}

	if json[0] == '"' {
		reader := readers.JsonStringReader{}
		err := reader.LoadFrom(json)
		return &reader, err
	}

	if utils.SliceCmp(json, JSON_TRUE) || utils.SliceCmp(json, JSON_FALSE) {
		reader := readers.JsonBoolReader{}
		err := reader.LoadFrom(json)
		return &reader, err
	}

	if utils.SliceCmp(json, JSON_NULL) {
		reader := readers.JsonNullReader{}
		return &reader, nil
	}

	reader := readers.JsonNumberReader{}
	err := reader.LoadFrom(json)
	return &reader, err
}
