package vars

const binary_name = "mmtime"
const version = "v0.1"

func GetDirName() string {
	return binary_name + "-" + version
}
