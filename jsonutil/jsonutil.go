package jsonutil

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func UnmarshalString(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}

func MarshalString(v interface{}) (string, error) {
	bs, err := json.Marshal(v)
	return string(bs), err
}

func MarshalIndentString(v interface{}) (string, error) {
	bs, err := json.MarshalIndent(v, "", "  ")
	return string(bs), err
}

func MustMarshalString(v interface{}) string {
	s, err := MarshalString(v)
	if err != nil {
		panic(err)
	}
	return s
}

func MustMarshalIndentString(v interface{}) string {
	s, err := MarshalIndentString(v)
	if err != nil {
		panic(err)
	}
	return s
}

func UnmarshalReader(r io.Reader, v interface{}) error {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, v)
}
