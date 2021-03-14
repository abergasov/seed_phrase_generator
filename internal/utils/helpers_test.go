package utils

import (
	"fmt"
	"testing"
)

func TestReverseSlice(t *testing.T) {
	src := []string{"a", "b", "c", "d"}
	expect := []string{"d", "c", "b", "a"}
	res := ReverseSlice(src)
	if len(res) != len(expect) {
		t.Fatal(fmt.Sprintf("len mismatch. expect %v, got %v", expect, res))
	}
	for i := 0; i < len(res); i++ {
		if res[i] != expect[i] {
			t.Fatal(fmt.Sprintf("slices mismatch. expect %v, got %v", expect, res))
		}
	}
}

func TestPositionInArray(t *testing.T) {
	table := []struct {
		src []string
		val string
		pos int
	}{
		{src: []string{"a", "b", "c"}, val: "b", pos: 1},
		{src: []string{"a", "b", "c"}, val: "d", pos: -1},
	}
	for _, v := range table {
		if PositionInArray(v.val, v.src) != v.pos {
			t.Fatalf("value exist in array %s", v.val)
		}
	}
}

func TestCutText(t *testing.T) {
	table := map[string]string{
		"12":  "12",
		"123": "123...",
	}
	for s, e := range table {
		if r := CutText(s, 3); r != e {
			t.Fatalf("expect %s got %s", e, r)
		}
	}
}
