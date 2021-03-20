package main //todo rename

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//Todo: add log

func main() {

	configFile, err := os.Open("config.json")
	CheckError(err, "Could not open config file", "Critical")
	defer configFile.Close()

	byteValue, err := ioutil.ReadAll(configFile)
	CheckError(err, "Could not read config file", "Critical")

	var configuration map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &configuration)
	CheckError(err, "Could not unmarshal config file", "Critical")

	PkgManagers, ok := configuration["pkg_mngrs"].(map[string]interface{})
	CheckStatus(ok, "Could not assert configuration", "Critical")

	var packages = GetPKGs(PkgManagers)
	file, _ := json.MarshalIndent(packages, "", " ")
	err = ioutil.WriteFile("test.json", file, 0644)
	CheckError(err, "Could not marshal packages json", "Critical")
	//jsonString, err := json.Marshal(packages)

	// enc := json.NewEncoder(writer)
	// err = enc.Encode(&jsonString)
	// CheckError(err, "Could not write packages json to file", "Critical")

	// pkgJsonFile, err := os.Create("packages.json")
	// CheckError(err, "Could not create packages json file", "Critical")
	// defer pkgJsonFile.Close()
	// pkgFileBytes, err := pkgJsonFile.Write(jsonString)
	// CheckError(err, "Could not write packages json file", "Critical")
	// fmt.Printf("wrote %d bytes\n", pkgFileBytes)
	// pkgJsonFile.Flush()

	// file, _ := json.MarshalIndent(data, "", " ")

	// _ = ioutil.WriteFile("test.json", file, 0644)
}
