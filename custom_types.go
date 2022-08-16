package fortnox

import (
	"encoding/json"
	"strconv"
)

// StringIsh exists because Fortnox sends back integers unquoted even if underlying type is string
type StringIsh string

func (f *StringIsh) UnmarshalJSON(data []byte) error {
	var receiver string

	if len(data) == 0 {
		return nil
	}
	if data[0] != '"' {
		quoted := strconv.Quote(string(data))
		data = []byte(quoted)
	}

	if err := json.Unmarshal(data, &receiver); err != nil {
		return err
	}

	*f = StringIsh(receiver)

	return nil
}

// FloatIsh type to allow unmarshalling from either string or float
type FloatIsh float64

func unmarshalIsh(data []byte, receiver interface{}) error {
	if len(data) == 0 {
		return nil
	}
	if data[0] == '"' {
		data = data[1:]
		data = data[:len(data)-1]
	}

	if len(data) < 1 {
		return nil
	}
	return json.Unmarshal(data, receiver)
}

// Float64 gets the value as float64
func (f *FloatIsh) Float64() float64 {
	if f == nil {
		return 0.0
	}
	return float64(*f)
}

// UnmarshalJSON to allow unmarshalling from either string or float
func (f *FloatIsh) UnmarshalJSON(data []byte) error {
	var newF float64
	err := unmarshalIsh(data, &newF)
	*f = FloatIsh(newF)
	return err
}

// IntIsh type to allow unmarshalling from either string or int
type IntIsh int

// Int gets the value as int
func (f *IntIsh) Int() int {
	if f == nil {
		return 0
	}
	return int(*f)
}

// UnmarshalJSON to allow unmarshalling from either string or int
func (f *IntIsh) UnmarshalJSON(data []byte) error {
	var newI int
	err := unmarshalIsh(data, &newI)
	*f = IntIsh(newI)
	return err
}
