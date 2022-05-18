package utils

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	slice := []string{
		"foo",
		"bar",
		"foo",
	}

	newSlice := Filter(slice, func(str string) bool { return str == "bar" })

	if !reflect.DeepEqual(newSlice, []string{"bar"}) {
		t.Fail()
	}
}
