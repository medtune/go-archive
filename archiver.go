package archiver

import (
	"errors"
	"os"
)

var (
	errArchiverExist      = errors.New("error archiver already registred")
	errArchiveUnsupported = errors.New("error archive unsupported format")
	errUnknownCommand     = errors.New("error unknown command")
	errLargeSize          = errors.New("error large size")

	supportPool = map[string]Archiver{}
)

type archiverMetadata struct {
	Ext string
}

// Archiver as defined
type Archiver interface {
	Meta() archiverMetadata
	Check(file string) bool
	Compress(file string, dest string) error
	Decompress(file, dest string) error
}

func registerArchiver(name string, archiver Archiver) error {
	_, ok := supportPool[name]
	if ok {
		return errors.New("Archiver already registred")
	}

	supportPool[name] = archiver
	return nil
}

// Pick one archiver for kind
func Pick(kind string) (Archiver, error) {
	v, ok := supportPool[kind]
	if ok {
		return v, nil
	}
	return nil, errArchiveUnsupported
}

func extractHeader(file string, size int) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	buf := make([]byte, size)
	lenght, err := f.Read(buf)

	if err != nil {
		return nil, err
	}

	if lenght < size {
		return buf, errLargeSize
	}

	return buf, nil
}
