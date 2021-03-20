// +build !windows

package main

import (
	"os/exec"
)

func GetPKGs(PkgManagers map[string]interface{}){
	unixDistro = GetDistro() //todo
	currentPkgMngr = PkgManagers[unixDistro].(string)
	currentGetPkgsCmd = GetPKGsCommand[currentPkgMngr]
	cmd := exec.Command(currentGetPkgsCmd)
    stdout, err := cmd.Output()
	CheckError(err, "Could not run package manager command", "Critical")
	return stdout
}