package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/rtbaker/GoToDo/password"
)

func TestOK(t *testing.T) {
	os.Args = []string{"./hashpassword", "newpassword"}
	out = bytes.NewBuffer(nil)

	status := run()

	passwordOut := out.(*bytes.Buffer).String()

	if !password.VerifyPassword("newpassword", passwordOut) {
		t.Errorf("Outputted password hash did not verify")
	}

	if status != 0 {
		t.Errorf("Return status not 0")
	}
}

func TestTooFewArgs(t *testing.T) {
	os.Args = []string{"./hashpassword"}
	errOut = bytes.NewBuffer(nil)

	status := run()

	error := errOut.(*bytes.Buffer).String()

	if error != "Usage: ./hashpassword <password>" {
		t.Errorf("Outputted error wrong for too few args")
	}

	if status != 1 {
		t.Errorf("Return status not 1")
	}
}
