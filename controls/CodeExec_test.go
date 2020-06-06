package controls

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

func TestAllCodeExec(t *testing.T) {
	scriptName := "../scripts/" + "tomcat_7.0.68.go"
	var outInfo bytes.Buffer
	var cmd = exec.Command("go","run",scriptName,"http://xxx:4180/")
	cmd.Stdout = &outInfo
	cmd.Run()
	fmt.Println(outInfo.String())
}
