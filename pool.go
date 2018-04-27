package dec

import "errors"

var (
	zip  = "zip"
	gzip = "gzip"
	gz   = "gz"

	supportPool = map[string]Archiver{}
)

func RegisterArchiver(name string, archiver Archiver) error {
	_, ok := supportPool[name]
	if ok {
		return errors.New("Archiver already registred")
	}
	supportPool[name] = archiver
	return nil
}
