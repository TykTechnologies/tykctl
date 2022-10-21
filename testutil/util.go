package testutil

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Equal[T comparable](t *testing.T, expected, actual T) {
	t.Helper()
	if expected != actual {
		t.Errorf("want: %v; got: %v", expected, actual)
	}
}

func EqualError(a, b error) bool {
	return a == nil && b == nil || a != nil && b != nil && a.Error() == b.Error()
}

func CheckDeepEqual(t *testing.T, expected, actual interface{}, opts ...cmp.Option) {
	t.Helper()
	if diff := cmp.Diff(actual, expected, opts...); diff != "" {
		t.Errorf("%T differ (-got, +want): %s", expected, diff)
		return
	}
}
