package static

import (
	"net/http"
	"sort"
	"sync"
)

type Reader interface {
	Get(filename string) ([]byte, string)
	List() []string

	ResponseWrite(w http.ResponseWriter, filename string) error
}

//Cache model
type Cache struct {
	files map[string][]byte
	lock  sync.RWMutex
}

//New init cache
func New() *Cache {
	c := &Cache{}
	c.Reset()
	return c
}

//Reset clean cache
func (c *Cache) Reset() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.files = make(map[string][]byte)
}

//Set setting data to cache
func (c *Cache) Set(filename string, v []byte) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.files[filename] = v
}

//Get getting file by name
func (c *Cache) Get(filename string) ([]byte, string) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	b, ok := c.files[filename]
	if !ok {
		return nil, ""
	}
	return b, DetectContentType(filename, b)
}

//Write write file to response
func (c *Cache) ResponseWrite(w http.ResponseWriter, filename string) error {
	b, ct := c.Get(filename)
	if b == nil {
		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	w.Header().Set("Content-Type", ct)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(b)
	return err
}

//List getting all files list
func (c *Cache) List() []string {
	c.lock.RLock()
	defer c.lock.RUnlock()

	result := make([]string, 0, len(c.files))
	for name := range c.files {
		result = append(result, name)
	}
	sort.Strings(result)
	return result
}
