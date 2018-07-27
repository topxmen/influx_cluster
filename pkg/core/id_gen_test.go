package core

import (
	"testing"
)

func TestIDGen(t *testing.T) {
	idGen := NewIDGenerator(100)
	existedMap := make(map[uint64]bool)
	for i := 0; i < 100000; i++ {
		id := idGen.New()
		if _, ok := existedMap[id]; ok {
			t.Fatal("id duplicate")
		} else {
			existedMap[id] = true
		}
	}
}
