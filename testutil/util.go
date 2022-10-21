package testutil

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

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

func Equal[T comparable](t *testing.T, description string, expected, actual T) {
	t.Helper()
	if expected != actual {
		t.Errorf("%s want: %v; got: %v", description, expected, actual)
	}
}
