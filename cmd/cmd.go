// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package cmd

import (
	"os/exec"
)

type Cmd struct {
	err error
	out []byte
}

func Run(cmd string) *Cmd {
	c := &Cmd{}
	x := exec.Command(cmd)
	c.out, c.err = x.Output()
	return c
}

func (c *Cmd) Failed() bool {
	return c.err != nil
}

func (c *Cmd) Error() error {
	return c.err
}

func (c *Cmd) Output() string {
	return string(c.out)
}
