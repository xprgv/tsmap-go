package tsmap_test

import (
	"testing"

	"github.com/xprgv/tsmap-go"
)

func TestTsMap(t *testing.T) {
	m := tsmap.NewThreadSafeMap[string, string]()

	m.Set("one", "data1")
	m.Set("two", "data2")

	if res1, exist := m.Get("one"); exist {
		if res1 != "data1" {
			t.Fatal("incorrect data")
		}
	} else {
		t.Fatal("failed to get key")
	}
}
