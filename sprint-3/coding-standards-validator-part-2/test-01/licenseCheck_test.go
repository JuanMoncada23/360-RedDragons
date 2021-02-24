package licensechk

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestLicenseChk(t *testing.T) {
	bytes, err := ioutil.ReadFile("LICENSE")
	if err != nil {
		t.Error("Failed to open file")
	}
	content := strings.ToUpper(string(bytes))
	licMIT := `PERMISSION IS HEREBY GRANTED, FREE OF CHARGE,`
	licReserved := `ALL RIGHTS RESERVED`
	if strings.Contains(content, `GNU`) || strings.Contains(content, licMIT) || strings.Contains(content, licReserved) {
		return
	}
}

func TestLicenseChk1(t *testing.T) {
	bytes, err := ioutil.ReadFile("LICENSE1")
	if err != nil {
		t.Error("Failed to open file")
	}
	content := strings.ToUpper(string(bytes))
	licMIT := `PERMISSION IS HEREBY GRANTED, FREE OF CHARGE,`
	licReserved := `ALL RIGHTS RESERVED`
	if strings.Contains(content, `GNU`) || strings.Contains(content, licMIT) || strings.Contains(content, licReserved) {
		return
	}
}
