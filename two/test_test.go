package two

import "testing"

func TestTwo(t *testing.T) {
	expected := "test"
	actual := Test()
	if expected != actual {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}