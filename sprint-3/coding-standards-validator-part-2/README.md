INFORMATIONAL
=============
Author: Rich Goluszka, Bryan Gabe, Juan Moncada  
Last Updated: 2/28/2021  
Project: Coding Standards Validator - Part 2  
Course: Applied Programming Languages (CPSC360-1) with Professor Eric Pogue  

Contact
-------
Please direct any and all comments/concerns/inquiries to richardjgoluszka@lewisu.edu

ORIGINALITY
===========
Credit to Chapter 8 of _Introducing Go_ by Caleb Doxsey for the code structure to open files 
	and directories using the io/ioutil package.

Credit to [Open Source Initiative](opensource.org/licenses/MIT) for the standard contents of an 
	MIT License.

All other code is the original work of the authors and may be used in accordance with the 
	specifications laid out in the LICENSE file.

BUILD / EXECUTE / DEPENDENCY
============================
Required files
--------------
coding-standards-validator-part-2
	-val directory  
		-val.go  
	-directorychk directory  
		-dirCheck.go  
		-dirCheck_test.go  
	-licensechk directory  
		-licenseCheck.go  
		-licenseCheck_test.go    
	-linefmtchk directory  
		-lineFmtCheck.go  
		-lineFmtChk_test.go   
	-utf8chk directory  
		-utf8Check.go  
		-utf8Check_test.go

_Note: The GitHub repository https://github.com/JuanMoncada23/360-RedDragons contains all_  
_required files plus README.md and LICENSE files.

Build instructions
------------------
To compile an executable:
1. Open the command-line or terminal
2. Navigate to .../go/src/360-RedDragons/sprint-3/coding-standards-validator-part-2
3. Run `go install` within each subdirectory (directorychk / licensechk / linefmtchk / utf8chk)
4. Run `go build` within the val subdirectory (`coding-standards-validator-part-2/val`)
You should now have a `val.exe` executable to call in order to run the program

Execution instructions
----------------------
1. Build the program _(using above instructions)_
2. Run `val.exe` and specify path to project when prompted
3. Optionally use `val.exe detail` to view detailed validation information
4. Optionally use `val.exe help` to view help instructions


Manual Testing Instructions
========================

test-01
-------
1. Run `val` and give the path `coding-standards-validator-part-2/test-02`
2. If all validations return true the test is considered PASS
3. Otherwsie, the test is considered a FAIL

test-02
-------
1. Run `val` and give the path `coding-standards-validator-part-2/test-02`
2. If all validations fail multiple times, the test is considered a PASS
3. Otherwsie, the test is considered a FAIL
_Note: The LicenseCheck Validator is incapable of failing multiple times._

test-03
-------
1. Run `val` and give the path `coding-standards-validator-part-2/test-03`
2. If only one validation fails, the test is considered a PASS
3. Otherwise, the test is considered a FAIL

Automated Testing Instructions
==============================

linefmtchk
----------
1. Navigate to `coding-standards-validator-part-2/linefmtchk`
2. Run `go test` with optional flags `-v` for verbose or `-cover` for coverage
