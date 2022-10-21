package testutil

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func EqualError(a, b error) bool {
	return a == nil && b == nil || a != nil && b != nil && a.Error() == b.Error()
}

func CheckDeepEqual(t *testing.T, expected, actual interface{}, opts ...cmp.Option) {

	if diff := cmp.Diff(actual, expected, opts...); diff != "" {
		t.Errorf("%T differ (-got, +want): %s", expected, diff)
		return
	}
}
