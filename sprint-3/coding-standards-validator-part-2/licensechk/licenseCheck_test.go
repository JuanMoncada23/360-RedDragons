package licensechk

import (
	"os"
	"testing"
)

const sep string = string(os.PathSeparator)

func TestValidate01(t *testing.T) {
	lc := LicenseChecker{Path: `..` + sep + `test-01`}
	if !lc.Validate() {
		t.Error(`No License File`)
	}
}
