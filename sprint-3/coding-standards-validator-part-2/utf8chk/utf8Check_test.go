package utf8chk

import "testing"

var retCt int
var retMsg, retIssue string
var result bool

var utftests = []UTF8Checker{
UTF8Checker{Path: "../test-01"},
UTF8Checker{Path: "../test-02"},
UTF8Checker{Path: "../test-03"},
UTF8Checker{Path: "../test-04"},
UTF8Checker{Path: "."},
}

func TestValidate00(t *testing.T) {
	utf := utftests[4]

	result = utf.Validate()
	retMsg = utf.GetMsg()
	retIssue = utf.GetIssues()
	retCt = utf.GetIssueCt()
}

func TestUTF8Checker_Validate(t *testing.T) {

	for _, utf := range utftests {
		if !utf.Validate() {
			t.Log("Error not UTF8! with ", utf.Path)
			t.Fail()
		}
	}
}

func TestUTF8Checker_GetIssueCt(t *testing.T) {
	for _, utf := range utftests {
		originalCount := utf.GetIssueCt()
		if !utf.Validate() {
			if utf.GetIssueCt() < originalCount {
				t.Log("Error with UTF8 Issue Counter")
				t.Fail()
			}
		}
	}
}

func TestUTF8Checker_GetMsg(t *testing.T) {
	for _, utf := range utftests {
		if utf.GetMsg() != ""{
			t.Log("Error Msg is empty")
			t.Fail()
		}
	}
}

func TestUTF8Checker_GetIssues(t *testing.T) {
	for _, utf := range utftests {
		if utf.GetIssues() != ""{
			t.Log("Error Issues found in ", utf.Path)
			t.Fail()
		}
	}
}
