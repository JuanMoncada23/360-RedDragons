Information
===========
The `test-03` directory demonstrates only one validation failing. The files in this directory are
    designed to fail only the line feed and tabs check (in linefmtchk.go). All other checks should
    be successful.

To manually test this, run `val` and provide a path to the `test-03` directory. To automatically
    test this, navigate to the `test-03` directory and run `go test`.