package searcher

import (
	"bytes"
	"math"
	"reflect"
	"testing"
)

func TestEncodeOneByte(t *testing.T) {
	input := int(5)
	buff := new(bytes.Buffer)
	expected := []byte{5 + 128}

	err := Encode(buff, input)

	if !reflect.DeepEqual(expected, buff.Bytes()) {
		t.Errorf("\nexpected: %v\nactual: %v", expected, buff.Bytes())
	}
	if err != nil {
		t.Errorf("\nerror: %v", err)
	}
}

func TestEncodeTwoByte(t *testing.T) {
	input := int(130)
	buff := new(bytes.Buffer)
	expected := []byte{1, 2 + 128}

	err := Encode(buff, input)

	if !reflect.DeepEqual(expected, buff.Bytes()) {
		t.Errorf("\nexpected: %v\nactual: %v", expected, buff.Bytes())
	}
	if err != nil {
		t.Errorf("\nerror: %v", err)
	}
}

func TestEncodeInt64Max(t *testing.T) {
	input := int(math.MaxInt64)
	buff := new(bytes.Buffer)
	expected := []byte{127, 127, 127, 127, 127, 127, 127, 127, 255}

	err := Encode(buff, input)

	if !reflect.DeepEqual(expected, buff.Bytes()) {
		t.Errorf("\nexpected: %v\nactual: %v", expected, buff.Bytes())
	}
	if err != nil {
		t.Errorf("\nerror: %v", err)
	}
}

func TestEncodeZero(t *testing.T) {
	input := int(0)
	buff := new(bytes.Buffer)
	expected := []byte{128}

	err := Encode(buff, input)

	if !reflect.DeepEqual(expected, buff.Bytes()) {
		t.Errorf("\nexpected: %v\nactual: %v", expected, buff.Bytes())
	}
	if err != nil {
		t.Errorf("\nerror: %v", err)
	}
}

func TestDecodeOneByte(t *testing.T) {
	input := bytes.NewReader([]byte{5 + 128})
	expected := int(5)

	actual, err := Decode(input)

	if actual != expected {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
	if err != nil {
		t.Errorf("\nerror: %v", err)
	}
}

func TestDecodeTwoByte(t *testing.T) {
	input := bytes.NewReader([]byte{1, 2 + 128})
	expected := int(130)

	actual, err := Decode(input)

	if actual != expected {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
	if err != nil {
		t.Errorf("\nerror: %v", err)
	}
}

func TestDecodeInt64Max(t *testing.T) {
	input := bytes.NewReader([]byte{127, 127, 127, 127, 127, 127, 127, 127, 255})
	expected := int(math.MaxInt64)

	actual, err := Decode(input)

	if actual != expected {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
	if err != nil {
		t.Errorf("\nerror: %v", err)
	}
}

func TestDecodeZero(t *testing.T) {
	input := bytes.NewReader([]byte{128})
	expected := int(0)

	actual, err := Decode(input)

	if actual != expected {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
	if err != nil {
		t.Errorf("\nerror: %v", err)
	}
}

func TestDecodeEOF(t *testing.T) {
	input := bytes.NewReader([]byte{10})
	expected := int(10)

	actual, err := Decode(input)

	if actual != expected {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
	if err == nil || err.Error() != "EOF" {
		t.Errorf("\nexpected: %v\nactual: %v", "EOF", err)
	}
}
