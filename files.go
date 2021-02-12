package static

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

var (
	//ErrUndefinedFormat error
	ErrUndefinedFormat = errors.New("undefined format")
)

func (c *Cache) fileCreate(filename string, call func(r io.Writer) error) error {
	v, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer v.Close()

	err = call(v)
	return err
}

func (c *Cache) fileOpen(filename string, call func(r io.Reader) error) error {
	v, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer v.Close()

	err = call(v)
	return err
}

//FromFile ...
func (c *Cache) FromFile(filename string) error {
	switch true {
	case strings.HasSuffix(filename, ".tar.gz"):
		return c.fileOpen(filename, func(v io.Reader) error {
			return c.FromTarGZArchive(v)
		})

	case strings.HasSuffix(filename, ".tar"):
		return c.fileOpen(filename, func(v io.Reader) error {
			return c.FromTarArchive(v)
		})

	default:
		return ErrUndefinedFormat
	}
}

//FromDir ...
func (c *Cache) FromDir(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, _ error) error {
		if info.IsDir() {
			return nil
		}
		v, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		c.Set(path, v)
		return nil
	})
}

//ToFile ...
func (c *Cache) ToFile(filename string) error {
	switch true {
	case strings.HasSuffix(filename, ".tar.gz"):
		return c.fileCreate(filename, func(w io.Writer) error {
			return c.ToTarGZArchive(w)
		})

	case strings.HasSuffix(filename, ".tar"):
		return c.fileCreate(filename, func(w io.Writer) error {
			return c.ToTarArchive(w)
		})

	default:
		return ErrUndefinedFormat
	}
}
