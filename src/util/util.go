package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

var Uname string

func ExecCommand(strCommand string) string {
	log.Println("cmd:", strCommand)
	cmd := exec.Command("/bin/bash", "-c", strCommand)
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		fmt.Println("Execute failed when Start:" + err.Error())
		return ""
	}

	outBytes, _ := ioutil.ReadAll(stdout)
	if err := stdout.Close(); err != nil {
		log.Println("stdout close error", err.Error())
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Execute failed when Wait:" + err.Error())
		return ""
	}
	return strings.TrimRight(string(outBytes), "\n")
}
func GetKernelInfo() string {
	cmd := "uname"
	res := ExecCommand(cmd)
	Uname = res
	return res
}
