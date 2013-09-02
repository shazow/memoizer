package memoizer

import (
	"testing"
)

type Counter int

func (i *Counter) Increment(by int) int {
	*i += Counter(by)
	return int(*i)
}

func TestMemoizeMember(t *testing.T) {
	memoize := Memoize{NewMemoryCache()}

	var i Counter = 0

	r, _ := memoize.Call(i.Increment, 2)
	if i != 2 {
		t.Error("Failed to call uncached function.")
	}
	if r != 2 {
		t.Error("Failed to return correct uncached value:", r)
	}

	r, _ = memoize.Call(i.Increment, 2)
	if i != 2 {
		t.Error("Failed to recall cached function.")
	}
	if r != 2 {
		t.Error("Failed to return correct cached value:", r)
	}

	memoize.Call(i.Increment, 3)
	if i != 5 {
		t.Error("Failed to call uncached function.")
	}

	memoize.Call(i.Increment, 3)
	if i != 5 {
		t.Error("Failed to recall cached function.")
	}
	
	r, _ = memoize.Call(i.Increment, 2)
	if i != 5 {
		t.Error("Failed to recall cached function.")
	}
	if r != 2 {
		t.Error("Failed to return correct cached value:", r)
	}
}
