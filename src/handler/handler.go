package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"util"
)

func GetServerInfo(c *gin.Context) {
	res := make(map[string]interface{})
	uname := util.ExecCommand(GetCmd("uname"))
	procVersion := util.ExecCommand(GetCmd("procVersion"))

	log.Println("uname", uname)

	res["uname"] = uname
	res["procVersion"] = procVersion

	c.JSON(http.StatusOK, res)
}
func GetCmd(intent string) string {
	kernelInfo := util.GetKernelInfo()
	switch kernelInfo {
	case "Darwin":
		return MacCmd(intent)
	case "Linux":
		return LinuxCmd(intent)
	default:
		return ""
	}
}
func LinuxCmd(intent string) string {
	switch intent {
	case "cpuInfo":
		return "cat /etc/cpuinfo"
	case "procVersion":
		return "cat /proc/version"
	case "uname":
		return "uname"
	case "releaseInfo":
		return "cat /etc/issue"
	default:
		return ""
	}
}
func MacCmd(intent string) string {
	switch intent {
	case "cpuInfo":
		return "sysctl machdep.cpu"
	case "procVersion":
		return "cat /proc/version"
	case "uname":
		return "uname"
	case "releaseInfo":
		return "cat /etc/issue"
	default:
		return ""
	}
}
