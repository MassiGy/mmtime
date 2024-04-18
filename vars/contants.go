package vars

import (
	"io/ioutil"
	"mmtime/utils"
	"os"
)

func GetDirName() string {
	res := ""
	file, err := os.Open("./BINARY_NAME")

	utils.Check(err)
	defer file.Close()

	chars, err := ioutil.ReadAll(file)
	utils.Check(err)

	fileContent := string(chars)

	if len(fileContent) <= 1 {
		panic("BINARY_NAME file is empty !.")
	}
	res += fileContent
	res += "-"

	file, err = os.Open("./VERSION")

	utils.Check(err)
	defer file.Close()

	chars, err = ioutil.ReadAll(file)
	utils.Check(err)

	fileContent = string(chars)

	if len(fileContent) <= 1 {
		panic("BINARY_NAME file is empty !.")
	}

	res += fileContent
	return res
}
