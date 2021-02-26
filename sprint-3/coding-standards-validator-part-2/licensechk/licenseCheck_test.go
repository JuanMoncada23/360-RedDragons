package licensechk

import (
	"os"
	"testing"
)

const sep string = string(os.PathSeparator)

var retCt int
var retMsg, retIssue string
var result bool

var lc []LicenseChecker = []LicenseChecker{
	{Path: ".." + sep + "test-01"},
	{Path: ".." + sep + "test-02"},
	{Path: ".." + sep + "test-03"},
	{Path: ".." + sep + "test-04"},
	{Path: ".."}, 
}

func TestValidate00(t *testing.T) {
	lc := lc[4]

	result = lc.Validate()
	retMsg = lc.GetMsg()
	retIssue = lc.GetIssues()
	retCt = lc.GetIssueCt()
}
func TestValidate01(t *testing.T) {
	lc := lc[0]

	result = lc.Validate()
	retMsg = lc.GetMsg()
	retIssue = lc.GetIssues()
	retCt = lc.GetIssueCt()

	if !lc.Validate() {
		t.Error("No License File")
	}
	if result == false || retMsg == `` || retIssue != `` || retCt != 0 {
		t.Error("Expected true, got", result, "Excepted 0 issues, got", retCt)
	}

}

func TestValidate02(t *testing.T) {
	lc := lc[1]

	result = lc.Validate()
	retMsg = lc.GetMsg()
	retIssue = lc.GetIssues()
	retCt = lc.GetIssueCt()

	if result == true || retMsg != `` || retIssue == `` || retCt == 0 {
		t.Error("Expected false, got", result, "Excepted issues, got", retCt)
	}
}
func TestValidate03(t *testing.T) {
	lc := lc[2]

	result = lc.Validate()
	retMsg = lc.GetMsg()
	retIssue = lc.GetIssues()
	retCt = lc.GetIssueCt()

	if result == false || retMsg == `` || retIssue != `` || retCt != 0 {
		t.Error("Expected true, got", result, "Excepted 0 issues, got", retCt)
	}
}

func TestValidate04(t *testing.T) {
	lc := lc[3]

	result = lc.Validate()
	retMsg = lc.GetMsg()
	retIssue = lc.GetIssues()
	retCt = lc.GetIssueCt()

	if result == true || retMsg != `` || retIssue == `` || retCt == 0 {
		t.Error("Expected false, got", result, "Excepted 0 issues, got", retCt)
	}
}
