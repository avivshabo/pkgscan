package main

var GetPKGsCommand = map[string][]string{
	"powershell": {"(Get-WmiObject -Class Win32_Product | select Name, Version | ConvertTo-Csv -Delimiter '|' -NoTypeInformation | Select-Object -Skip 1)", "-replace '\"',''"},
	"apt":        {"dpkg-query", "--list", "-f='${binary:Package}\t${Version}\t${Architecture}\n'"},
	"yum":        {"rpm", "-qa"}}

var ExecutionShells = map[string]string{
	"powershell": "powershell",
	"apt":        "bash",
	"yum":        "bash"}
