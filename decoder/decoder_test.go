package decoder

import (
	"reflect"
	"testing"
)

func TestDecodeInt(t *testing.T) {
	d := New([]byte("i-32e"))
	expected := -32
	found := d.decodeInt()
	if found != expected {
		t.Errorf("TestDecodeInt: Expected %+v, got %+v", expected, found)
	}
}

func TestDecodeString(t *testing.T) {
	d := New([]byte("4:spam"))
	expected := "spam"
	found := d.decodeString()
	if found != expected {
		t.Errorf("TestDecodeString: Expected %+v, got %+v", expected, found)
	}
}

func TestDecodeList(t *testing.T) {
	d := New([]byte("l4:spam4:eggse"))
	expected := []interface{}{"spam", "eggs"}
	found := d.decodeList()
	if !reflect.DeepEqual(found, expected) {
		t.Errorf("TestDecodeList: Expected %+v, got %+v", expected, found)
	}
}

func TestDecodeDictionary(t *testing.T) {
	d := New([]byte("d4:spaml1:a1:bee"))
	expected := map[string]interface{}{
		"spam": []interface{}{
			"a",
			"b",
		},
	}
	found := d.decodeDictionary()
	if !reflect.DeepEqual(found, expected) {
		t.Errorf("TestDecodeDictionary: Expected %+v, got %+v", expected, found)
	}
}

func TestIsDigit(t *testing.T) {
	found := isDigit('5')
	if !found {
		t.Errorf("TestIsDigit: Expected %+v, got %+v", true, found)
	}
}
