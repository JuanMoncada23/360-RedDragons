package directorychk

import "testing"

var dirchks = []DirChecker{
	DirChecker{Path: "../test-01"},
	DirChecker{Path: "../test-02"},
	DirChecker{Path: "../test-03"},
	DirChecker{Path: "../test-04"},

}
func TestDirChecker_Validate(t *testing.T) {
	for _, dir := range dirchks {
		if !dir.Validate() {
			t.Fail()
			t.Log("DirCheck failed in ", dir.Path)
		}
	}
}
func TestUTF8Checker_GetIssueCt(t *testing.T) {
	for _, dir := range dirchks {
		originalCount := dir.GetIssueCt()
		if !dir.Validate() {
			if dir.GetIssueCt() < originalCount {
				t.Log("Error with DirCheck Issue Counter")
				t.Fail()
			}
		}
	}
}

func TestUTF8Checker_GetMsg(t *testing.T) {
	for _, dir := range dirchks {
		if dir.GetMsg() != ""{
			t.Log("Error Msg is empty")
			t.Fail()
		}
	}
}

func TestUTF8Checker_GetIssues(t *testing.T) {
	for _, dir := range dirchks {
		if dir.GetIssues() != "" {
			t.Log("Error Issues found in ", dir.Path)
			t.Fail()
		}
	}
}