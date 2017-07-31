package cases

import "example/cloud/raport/lib"

type Config struct {
}

var (
	readLib lib.ReadExel
	Path    string
)

func (c Config) SetConfig(path string) {
	Path = path
}

//StartCase StartNoneDevice
func (c Config) StartCase() {
	readLib.SetReadFilePath(Path)
}
