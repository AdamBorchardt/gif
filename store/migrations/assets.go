// Code generated by go-bindata.
// sources:
// ../db_migrations/001_initial.sql
// DO NOT EDIT!

package migrations

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _db_migrations001_initialSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x90\x4f\x4b\xc3\x30\x18\xc6\xef\xf9\x14\xcf\x71\xc5\x0d\x44\xdc\xa9\xa7\xb8\x04\x0c\x76\xe9\x0c\x89\x6c\xa7\x12\x48\x28\x01\x57\x75\xb6\xe8\xc7\x37\xeb\x1f\xa9\xb5\xbd\x94\x26\x4f\x7e\xcf\xef\xe5\xdd\x6c\x70\x73\x0e\xe5\xc5\xd6\x1e\xe6\x9d\xec\x14\xa7\x9a\x43\xd3\x87\x8c\x23\x9c\x6d\xe9\x3f\xb1\x22\x40\x70\x78\xa1\x6a\xf7\x48\xd5\xea\xfe\x36\xc1\x41\x89\x3d\x55\x27\x3c\xf1\xd3\x3a\xa6\xcd\xe5\x15\x9a\x1f\xf5\xf5\xdf\x3a\xe7\x5d\x61\x6b\xb0\x58\xa4\xc5\x9e\x43\xe6\x1a\xd2\x64\x19\x49\x52\x32\x23\x28\x6a\x5b\xf6\x92\xf6\x38\x51\x0d\xf4\xb5\x3b\xbe\xfc\xcd\xee\xb6\xdb\x64\xb6\x5a\x48\xc6\x8f\xa3\xea\x22\x54\xce\x7f\x23\x97\x7f\x74\xf1\x1b\x99\x1e\x31\x52\x3c\x9b\x19\xb2\xa9\xc2\x47\xe3\xa7\xe8\x30\xe6\x1a\x5d\x09\x19\xef\x90\xbd\x7d\x55\x84\xa9\xfc\xb0\x30\x47\xba\x10\x76\xaa\x3e\x1d\xaf\xff\xff\x55\x0b\xa4\xe4\x27\x00\x00\xff\xff\xb1\x45\xa8\xdf\xba\x01\x00\x00")

func db_migrations001_initialSqlBytes() ([]byte, error) {
	return bindataRead(
		_db_migrations001_initialSql,
		"db_migrations/001_initial.sql",
	)
}

func db_migrations001_initialSql() (*asset, error) {
	bytes, err := db_migrations001_initialSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db_migrations/001_initial.sql", size: 442, mode: os.FileMode(420), modTime: time.Unix(1463917931, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"db_migrations/001_initial.sql": db_migrations001_initialSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"db_migrations": &bintree{nil, map[string]*bintree{
		"001_initial.sql": &bintree{db_migrations001_initialSql, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
