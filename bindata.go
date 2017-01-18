// Code generated by go-bindata.
// sources:
// pieces.png
// tile.png
// DO NOT EDIT!

package main

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

var _piecesPng = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xea\x0c\xf0\x73\xe7\xe5\x92\xe2\x62\x60\x60\xe0\xf5\xf4\x70\x09\x02\xd2\x06\x40\x2c\xc0\xc1\x06\x24\x03\xd6\xfd\xd9\x08\xa4\x66\x7b\xba\x38\x86\x54\xdc\xba\x3a\xb5\x91\xeb\x80\x02\x87\xcb\xdf\xff\xf3\xcb\x78\xd8\x56\xaf\x7d\x3d\x45\x7b\x95\x42\x93\xed\xea\xe0\xf4\x59\x57\x0f\xcf\xef\x6c\x13\x96\xfe\xf1\x57\xee\xe9\xaf\x8d\x5f\x8f\xbd\xea\xfc\x7c\xdf\xbb\x8b\x5b\x81\xdd\xcc\xe3\x4e\x7e\x7e\x72\xfd\x11\x46\x31\x1b\xee\xee\xd7\x37\x15\xd8\xd2\xae\x1c\x0c\x60\xd3\xf9\x51\xbd\x80\xe3\x1b\xcb\x32\xa3\x3b\x8f\x4d\xb8\xd7\xb0\x85\x30\x0a\xff\x2c\x62\x14\x5a\xee\xba\xcd\xf2\x39\x10\xcf\xbf\xcf\x7e\x20\xd8\x72\x1b\xbf\x40\xe3\x97\xa9\x9e\x1b\x2c\x6b\x1d\x04\x0c\x0e\x31\xf2\xb3\x05\xbe\x62\xee\x6a\xc8\x69\xdb\x1c\xb3\x8e\x31\xfe\xcc\x74\xee\x09\x3f\x3e\x9e\x7f\xb1\x91\x2f\x46\xf2\x4d\xed\xfc\xab\x37\xe7\xfe\x01\x3a\x93\xc1\xd3\xd5\xcf\x65\x9d\x53\x42\x13\x20\x00\x00\xff\xff\xc8\x64\xc9\x07\xd4\x00\x00\x00")

func piecesPngBytes() ([]byte, error) {
	return bindataRead(
		_piecesPng,
		"pieces.png",
	)
}

func piecesPng() (*asset, error) {
	bytes, err := piecesPngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pieces.png", size: 212, mode: os.FileMode(420), modTime: time.Unix(1484641973, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tilePng = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xea\x0c\xf0\x73\xe7\xe5\x92\xe2\x62\x60\x60\xe0\xf5\xf4\x70\x09\x02\xd2\x1c\x20\xcc\xc1\x06\x24\x8f\xf0\xef\xeb\x06\x52\x22\x9e\x2e\x8e\x21\x15\xb7\x92\x13\x12\x12\x7e\xf4\x8b\xff\x3f\xce\xdc\xd6\xcf\xc4\xc0\xd2\x74\x49\xef\x6b\xd5\x74\x6e\xa0\x3c\x83\xa7\xab\x9f\xcb\x3a\xa7\x84\x26\x40\x00\x00\x00\xff\xff\xfd\x23\x46\xe5\x4d\x00\x00\x00")

func tilePngBytes() ([]byte, error) {
	return bindataRead(
		_tilePng,
		"tile.png",
	)
}

func tilePng() (*asset, error) {
	bytes, err := tilePngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tile.png", size: 77, mode: os.FileMode(420), modTime: time.Unix(1484642985, 0)}
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
	"pieces.png": piecesPng,
	"tile.png": tilePng,
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
	"pieces.png": &bintree{piecesPng, map[string]*bintree{}},
	"tile.png": &bintree{tilePng, map[string]*bintree{}},
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

