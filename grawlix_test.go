package gograwlix

import "testing"

func TestGrawlix(t *testing.T) {
	for _, v := range []string{
		"shit", "yoink",
		"bark", "tree",
		"bass boom", "boing",
		"gutter",
	} {
		grawr := Grawlix(v)
		if grawr == v {
			t.Log(v, "failed")
			t.Fail()
		} else {
			t.Log(grawr)
		}
	}
}
