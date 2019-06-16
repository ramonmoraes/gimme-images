package gimme

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {
	img := Image{URL: "something.com/name.qqq"}
	img.createName()
	if img.Name != "name.qqq.png" {
		t.Errorf(fmt.Sprintf("Wrong name: %s and should be name.qqq.png", img.Name))
	}

	img = Image{URL: "sogmething.com/name.png"}
	img.createName()

	if img.Name != "name.png" {
		t.Errorf(fmt.Sprintf("Wrong name: %s and it should be name.png", img.Name))
	}
}
