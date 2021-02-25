package linefmtchk

import (
	"os"
	"testing"
)

//setup const to make creating paths easier
const sep string = string(os.PathSeparator)

//setup vars to hold method return values
var result bool
var retMsg, retIssue string
var retCt int

//setup struct slice to test methods with
var lfcs []LineFmtChecker = []LineFmtChecker{
	{Path: `..` + sep + `test-01`},
	{Path: `..` + sep + `test-02`},
	{Path: `..` + sep + `test-03`},
	{Path: `..` + sep + `test-04`},
}

//test on folder `test-01`, it should pass
func TestValidate01(t *testing.T) {
	//get appropriate lfc struct, call methods
	lfc := lfcs[0]

	result = lfc.Validate()
	retMsg = lfc.GetMsg()
	retIssue = lfc.GetIssues()
	retCt = lfc.GetIssueCt()

	//expecting result = true AND retMsg != `` AND retIssue = `` AND retCt = 0
	//error when result = false OR retMsg = `` OR retIssue != `` OR retCt != 0
	if result == false || retMsg == `` || retIssue != `` || retCt != 0 {
		t.Error(`Expected true, got`, result,
			`Expected 0 issues, got`, retCt)
	}
}

//test on folder `test-02`, it should fail
func TestValidate02(t *testing.T) {
	//get appropriate lfc struct, call methods
	lfc := lfcs[1]

	result = lfc.Validate()
	retMsg = lfc.GetMsg()
	retIssue = lfc.GetIssues()
	retCt = lfc.GetIssueCt()

	//expecting result = false AND retMsg != `` AND retIssue != `` AND retCt != 0
	//error when result = true OR retMsg = `` OR retIssue = `` OR retCt = 0
	if result == true || retMsg == `` || retIssue == `` || retCt == 0 {
		t.Error(`Expected false, got`, result,
			`Expected issues, got`, retCt)
	}
}

//test on folder `test-03`, it should fail
func TestValidate03(t *testing.T) {
	//get appropriate lfc struct, call methods
	lfc := lfcs[2]

	result = lfc.Validate()
	retMsg = lfc.GetMsg()
	retIssue = lfc.GetIssues()
	retCt = lfc.GetIssueCt()

	//expecting result = false AND retMsg != `` AND retIssue != `` AND retCt != 0
	//fail when result = true OR retMsg = `` OR retIssue = `` OR retCt = 0
	if result == true || retMsg == `` || retIssue == `` || retCt == 0 {
		t.Error(`Expected false, got`, result,
			`Expected issues, got`, retCt)
	}
}

//test on folder `test-04`, it should fail
func TestValidate04(t *testing.T) {
	//get appropriate lfc struct, call methods
	lfc := lfcs[3]

	result = lfc.Validate()
	retMsg = lfc.GetMsg()
	retIssue = lfc.GetIssues()
	retCt = lfc.GetIssueCt()

	//expecting result = false AND retMsg != `` AND retIssue != `` AND retCt != 0
	//fail when result = true OR retMsg = `` OR retIssue = `` OR retCt = 0
	if result == true || retMsg == `` || retIssue == `` || retCt == 0 {
		t.Error(`Expected false, got`, result,
			`Expected issues, got`, retCt)
	}
}
