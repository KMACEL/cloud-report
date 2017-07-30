package cases

type Config struct {
}

var Path string

func (c Config) SetConfig(path string) {
	Path = path
}
