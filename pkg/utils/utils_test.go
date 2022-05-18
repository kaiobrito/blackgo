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

func TestSeparate(t *testing.T) {
	slice := []string{
		"foo",
		"bar",
		"foo",
	}

	newSlice, removed := Separate(slice, func(str string) bool { return str == "bar" })

	if !reflect.DeepEqual(newSlice, []string{"bar"}) {
		t.Fail()
	}
	if !reflect.DeepEqual(removed, []string{"foo", "foo"}) {
		t.Fail()
	}
}

func TestReduce(t *testing.T) {
	slice := []string{
		"foo",
		"bar",
		"foo",
	}

	output := Reduce(slice, func(previous string, newValue string) string { return previous + newValue }, "abc")

	if output != "abcfoobarfoo" {
		t.Fail()
	}
}

func TestReduceSum(t *testing.T) {
	slice := []int{
		1, 2, 3, 4, 5,
	}

	output := Reduce(slice, func(previous int, newValue int) int { return previous + newValue }, 10)

	if output != 25 {
		t.Fail()
	}
}

func TestSum(t *testing.T) {
	slice := []int{
		1, 2, 3, 4, 5,
	}

	output := Sum(slice, 10)

	if output != 25 {
		t.Fail()
	}
}
