package cmd

import (
	"testing"
	"io"
)

func TestCmd_Exec(t *testing.T) {
	test_cmd := Cmd{
		Cmd:"python",
		Args:[]string{"-u", "cmd_test.py"},
	}
	std_out, _ := test_cmd.cmd.StdoutPipe()
	p := make([]byte, 9)
	test_cmd.cmd.Output()
	test_cmd.Exec()
	test_cmd.cmd.Wait()
	std_out.Read(p)
	if string(p) != "test data" {
		t.Error("Wrong data executed or didn't execute properly")
		t.Fail()
	}
}

func TestMultiWriter(t *testing.T){
	io.MultiWriter()
}