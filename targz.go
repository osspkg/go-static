package static

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"io/ioutil"
	"os"
)

//FromBase64TarGZ ...
func (c *Cache) FromBase64TarGZ(v string) error {
	b64, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		return err
	}
	return c.FromTarGZArchive(bytes.NewBuffer(b64))
}

//FromTarGZArchive ...
func (c *Cache) FromTarGZArchive(r io.Reader) error {
	gzf, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	return c.FromTarArchive(gzf)
}

//FromTarArchive ...
func (c *Cache) FromTarArchive(r io.Reader) error {
	tr := tar.NewReader(r)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		b, err := ioutil.ReadAll(tr)
		if err != nil {
			return err
		}
		c.Set("/"+hdr.Name, b)
	}
	return nil
}

//ToTarGZArchive ...
func (c *Cache) ToTarGZArchive(w io.Writer) error {
	gw := gzip.NewWriter(w)
	defer gw.Close()
	err := c.ToTarArchive(gw)
	return err
}

//ToTarArchive ...
func (c *Cache) ToTarArchive(w io.Writer) error {
	tw := tar.NewWriter(w)
	defer tw.Close()

	c.lock.RLock()
	defer c.lock.RUnlock()

	for name, v := range c.files {
		hdr := &tar.Header{
			Name: name,
			Mode: int64(os.ModePerm),
			Size: int64(len(v)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}
		if _, err := tw.Write(v); err != nil {
			return err
		}
	}
	return nil
}

//ToBase64TarGZ ...
func (c *Cache) ToBase64TarGZ() (string, error) {
	buf := &bytes.Buffer{}
	if err := c.ToTarGZArchive(buf); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}
