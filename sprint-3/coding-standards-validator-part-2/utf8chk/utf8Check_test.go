package utf8chk

import "testing"

func TestUTF8Checker_Validate(t *testing.T) {
	utf := UTF8Checker{
		Path: ".",
	}
	if !utf.Validate() {
		t.Log("Error not UTF8!")
	}
}
