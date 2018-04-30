// +build !windows,!darwin

package open

import (
	"os/exec"
)

// http://sources.debian.net/src/xdg-utils/1.1.0~rc1%2Bgit20111210-7.1/scripts/xdg-open/
// http://sources.debian.net/src/xdg-utils/1.1.0~rc1%2Bgit20111210-7.1/scripts/xdg-mime/

//Open() opens url in defualt browser
func Open(input string) *exec.Cmd {
	return exec.Command("xdg-open", input)
}

//OpenWith() opens url in desired app
func OpenWith(input string, appName string) *exec.Cmd {
	return exec.Command(appName, input)
}
