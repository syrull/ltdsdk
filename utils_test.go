package ltdsdk

import (
	"testing"
)

func TestStructToQueryString(t *testing.T) {
	type ExampleStruct struct {
		StringField  string  `qs:"stringField"`
		IntegerField int     `qs:"integerField"`
		Float32Field float32 `qs:"float32Field"`
		Float64Field float64 `qs:"float64Field"`
	}
	exampleStruct := &ExampleStruct{
		StringField:  "test",
		IntegerField: 1,
		Float32Field: 1337.1337,
		Float64Field: 1337.1337133713,
	}
	qs := toQueryString(exampleStruct)
	if qs["stringField"] != "test" {
		t.Error("error `qs[stringField]` != 'test'")
	}
	if qs["integerField"] != "1" {
		t.Error("error `qs[integerField]` != '1'")
	}
	if qs["float32Field"] != "1337.1337" {
		t.Error("error `qs[float32Field]` != '1337.1337'")
	}
	if qs["float64Field"] != "1337.1337133713" {
		t.Error("error `qs[float64Field]` != '1337.1337133713'")
	}
}

func TestParseStringToFloat32(t *testing.T) {
	testStr := "1337.1"
	val := parseStringToFloat32(testStr, 0)
	if val != 1337.1 {
		t.Error("error val != 1337.1")
	}
}

func TestParseStringToFloat32OnError(t *testing.T) {
	testStr := "1337.1error"
	val := parseStringToFloat32(testStr, 0)
	if val != 0 {
		t.Error("error val != 0")
	}
}

func TestParseStringToInt(t *testing.T) {
	testStr := "1337"
	val := parseStringToInt(testStr, 0)
	if val != 1337 {
		t.Error("error val != 1337")
	}
}

func TestParseStringToIntOnError(t *testing.T) {
	testStr := "1337errpr"
	val := parseStringToInt(testStr, 0)
	if val != 0 {
		t.Error("error val != 0")
	}
}
