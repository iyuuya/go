package pathname

import (
	"os"
	"path/filepath"
)

type Pathname struct {
	path string
}

func New(path string) Pathname {
	return Pathname{path}
}

func Pwd() Pathname {
	d, _ := os.Getwd()
	return New(d)
}

func Glob(pattern string) []Pathname {
	var ps []Pathname
	return ps
}

func (p Pathname) Join(other ...string) Pathname {
	return New(filepath.Join(append([]string{p.path}, other...)...))
}

func (p Pathname) IsExist() bool {
	_, err := os.Stat(p.path)
	return err == nil
}

func (p Pathname) IsAbsolute() bool {
	return filepath.IsAbs(p.path)
}

func (p Pathname) IsRelative() bool {
	return !p.IsAbsolute()
}

func (p Pathname) IsDirectory() bool {
	st, err := os.Stat(p.path)
	return err == nil && st.IsDir()
}

func (p Pathname) ExpandPath() Pathname {
	s, e := filepath.Abs(p.path)
	if e != nil {
		panic(e)
	}
	return New(s)
}

func (p Pathname) Basename() Pathname {
	return New(filepath.Base(p.path))
}

func (p Pathname) Extname() string {
	return filepath.Ext(p.path)
}

func (p Pathname) Dirname() Pathname {
	return New(filepath.Dir(p.path))
}

func (p Pathname) MkdirAll(perm os.FileMode) error {
	return os.MkdirAll(p.path, perm)
}

func (p Pathname) Clean() Pathname {
	return New(filepath.Clean(p.path))
}

func (p Pathname) Stat() (os.FileInfo, error) {
	return os.Stat(p.path)
}

func (p Pathname) String() string {
	return p.path
}

func (p Pathname) Split() []string {
	return filepath.SplitList(p.path)
}
