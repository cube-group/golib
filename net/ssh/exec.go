package ssh

import "os/exec"

//执行shell命令
func Exec(shells ...string) (string, error) {
	cmd := exec.Command(shells[0], shells[1:]...)
	out, err := cmd.Output()
	return string(out), err
}
