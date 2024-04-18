package vars

const BINARY_NAME = "mmtime"
const VERSION = "v0.1"

func GetDirName() string {
	return BINARY_NAME + "-" + VERSION
}
