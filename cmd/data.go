package cmd

import (
	//"os/exec"
	//"log"
	//"os"
	//"path/filepath"
)

var LOG_ROOT = "log/"

type CmdConfig struct {
	Cmd     string
	Args    []string
	LogFile string
}
//
//type Cmd struct {
//	Conf   CmdConfig
//	Logger log.Logger
//	cmd    *exec.Cmd
//}
//
//func NewCmd(c CmdConfig) *Cmd {
//	if c.LogFile == "" {
//		c.LogFile=""
//	}
//	filepath.Dir(c.LogFile)
//	os.MkdirAll()
//	return &Cmd{
//		Conf:c,
//		Logger:log.New()
//	}
//}
//
//func (c *Cmd) Exec() error {
//	if c.cmd != nil {
//		return c.cmd.Start()
//	}
//	c.cmd = exec.Command(c.Cmd, c.Args...)
//	c.cmd.Start()
//	return nil
//}
