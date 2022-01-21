package command

import (
	"os/exec"

	"github.com/rohanthewiz/rerr"
)

func ExecCmd(cmd string, params ...string) (bytesOut []byte, err error) {
	c := exec.Command(cmd, params...)
	bytesOut, err = c.Output()
	if err != nil {
		return bytesOut, rerr.Wrap(err, "Command finished with error")
	}
	return
}
