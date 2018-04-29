package archiver

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	_ = registerArchiver("zip", Zip)
}

var (
	// Zip default instance
	Zip = &zipper{}

	zipHeader = []byte("PK\x03\x04")
	zipExt    = "zip"
)

type zipper struct{}

func (z *zipper) Meta() archiverMetadata {
	return archiverMetadata{
		Ext: zipExt,
	}
}

func (z *zipper) Check(file string) bool {
	if filepath.Ext(file) == "."+zipExt {
		return true
	}

	header, err := extractHeader(file, 4)
	if err != nil {
		return false
	}

	return bytes.Compare(header, zipHeader) == 0
}

func (z *zipper) Compress(file, dest string) error {
	return zipFile(file, dest)
}

func (z *zipper) Decompress(file, dest string) error {
	return unzipFile(file, dest)
}

func unzipFile(file, dest string) error {
	r, err := zip.OpenReader(file)
	if err != nil {
		return err
	}

	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}

		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}

			_, err = io.Copy(outFile, rc)
			outFile.Close()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func zipFile(fname, dest string) error {
	zipfile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(fname)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(fname)
	}

	filepath.Walk(fname, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, fname))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}
