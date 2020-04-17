package main

import (
	"reflect"
	"testing"
)

func TestLen(t *testing.T) {
	unsortedMap := map[string]int{}
	unsortedMap["hello"] = 1
	unsortedMap["hello"] = 1
	unsortedMap["hello"] = 1
	unsortedMap["hello there"] = 1

	sm := &sortedMap{
		unsortedMap: unsortedMap,
	}

	expected := 2

	actual := sm.Len()
	if actual != expected {
		t.Errorf("got: %v, want: %v", actual, expected)
	}
}

func TestLess(t *testing.T) {
	unsortedMap := map[string]int{}
	unsortedMap["hello"] = 1
	unsortedMap["hello there"] = 2

	sm := &sortedMap{
		unsortedMap: unsortedMap,
		keys:        []string{"hello", "hello there"},
	}

	expected := false
	actual := sm.Less(0, 1)
	if actual != expected {
		t.Errorf("got: %v, want: %v", actual, expected)
	}

	expected = true
	actual = sm.Less(1, 0)
	if actual != expected {
		t.Errorf("got: %v, want: %v", actual, expected)
	}
}

func TestSwap(t *testing.T) {
	sm := &sortedMap{
		keys: []string{"hello", "hello there"},
	}

	expected := []string{"hello there", "hello"}
	sm.Swap(0, 1)
	if !reflect.DeepEqual(sm.keys, expected) {
		t.Errorf("got: %v, want: %v", sm.keys, expected)
	}
}

func TestSortedKeys(t *testing.T) {
	unsortedMap := map[string]int{}
	unsortedMap["hello"] = 1
	unsortedMap["hello there"] = 2

	expected := []string{"hello there", "hello"}
	actual := sortedKeys(unsortedMap)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got: %v, want: %v", actual, expected)
	}
}
