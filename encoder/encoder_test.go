package encoder

import (
	"reflect"
	"testing"
)

func TestEncodeInt(t *testing.T) {
	e := New(123)
	expected := "i123e"
	found := e.encodeInt()
	if found != expected {
		t.Errorf("TestEncodeInt: Expected %+v, got %+v", expected, found)
	}
}

func TestEncodeString(t *testing.T) {
	e := New("abcd")
	expected := "4:abcd"
	found := e.encodeString()
	if found != expected {
		t.Errorf("TestEncodeString: Expected %+v, got %+v", expected, found)
	}
}

func TestEncodeList(t *testing.T) {
	e := New([]interface{}{
		"spam",
		"eggs"})
	expected := "l4:spam4:eggse"
	found := e.encodeList()
	if found != expected {
		t.Errorf("TestEncodeList: Expected %+v, got %+v", expected, found)
	}
}

func TestEncodeDictionary(t *testing.T) {
	e := New(map[string]interface{}{
		"spam": []interface{}{
			"a",
			"b",
		},
	})
	expected := "d4:spaml1:a1:bee"
	found := e.encodeDictionary()
	if found != expected {
		t.Errorf("TestEncodeDictionary: Expected %+v, got %+v", expected, found)
	}
}

func TestGetSortedKeys(t *testing.T) {
	expected := []string{"m", "n", "o"}
	found := getSortedKeys(map[string]interface{}{
		"m": "aaa",
		"n": "ccc",
		"o": "rrr",
	})
	if !reflect.DeepEqual(found, expected) {
		t.Errorf("TestEncodeDictionary: Expected %+v, got %+v", expected, found)
	}
}
