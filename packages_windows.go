// +build windows

package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func GetPKGs(PkgManagers map[string]interface{}) map[string]string {
	currentPkgMngr := PkgManagers[runtime.GOOS].(string)
	currentGetPkgsCmd := GetPKGsCommand[currentPkgMngr]
	fmt.Println("Running command: " + strings.Join(currentGetPkgsCmd, " "))
	cmd := exec.Command(ExecutionShells[currentPkgMngr], currentGetPkgsCmd...)
	stdout, err := cmd.Output()
	fmt.Println("command output: " + string(stdout))
	CheckError(err, "Could not run package manager command", "Critical")

	scanner := bufio.NewScanner(stdout)
	scanner.Split(ScanLinesR)

	packages := make(map[string]string)
	for scanner.Scan() {
		pkgLine := strings.Split(scanner.Text(), "|")
		if len(pkgLine) == 2 {
			packages[pkgLine[0]] = pkgLine[1]
		} else {
			fmt.Println("Error parsing pkg line" + string(i) + ": " + strings.Join(pkgLine, ", "))
		}
	}
	// packages := make(map[string]string)
	// pkgArr := strings.Split(string(stdout), "\n")

	// for i := 0; i < len(pkgArr); i++ {
	// 	pkgLine := strings.Split(pkgArr[i], "|")
	// 	if len(pkgLine) == 2 {
	// 		packages[pkgLine[0]] = pkgLine[1]
	// 	} else {
	// 		fmt.Println("Error parsing pkg line" + string(i) + ": " + strings.Join(pkgLine, ", "))
	// 	}
	// }

	return packages
}
