package gimme

import "testing"

func TestAbs(t *testing.T) {
	img := Image{URL: "something.com/name.qqq"}
	img.createName()
	if img.Name != "name.qqq" {
		t.Errorf("Wrong name")
	}
}
