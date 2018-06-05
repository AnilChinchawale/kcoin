// Code generated by go-bindata.
// sources:
// index.html
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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x6d\x6f\xdb\x38\x12\xfe\xec\xfc\x8a\xa9\x7a\x45\x6c\xd4\x92\x28\xbf\x24\x8e\x62\xb9\xe8\xa5\xc5\xa1\x07\xf4\x5a\xa0\x39\x5c\xef\x23\x2d\x8d\x24\x26\x14\xa9\x92\xb4\x5d\x6f\x91\xff\xbe\x20\x25\xd9\x4e\xea\xb4\xbb\x6d\x81\xdd\x2f\x89\xe6\xe5\x79\xe6\x91\x34\x33\xa2\xe7\x4f\x5e\xbd\xbb\xba\xfe\xff\xfb\xd7\x50\x9a\x8a\x2f\x4e\xe6\xf6\x1f\x70\x2a\x8a\xc4\x43\xe1\x2d\x4e\x7a\xf3\x12\x69\xb6\x38\xe9\xf5\xe6\x15\x1a\x0a\x69\x49\x95\x46\x93\x78\x2b\x93\xfb\x33\x6f\x1f\x28\x8d\xa9\x7d\xfc\xb4\x62\xeb\xc4\xfb\xe8\xff\xf7\xa5\x7f\x25\xab\x9a\x1a\xb6\xe4\xe8\x41\x2a\x85\x41\x61\x12\xef\xcd\xeb\x04\xb3\x02\x0f\x70\x82\x56\x98\x78\x6b\x86\x9b\x5a\x2a\x73\x90\xba\x61\x99\x29\x93\x0c\xd7\x2c\x45\xdf\x19\x43\x60\x82\x19\x46\xb9\xaf\x53\xca\x31\x89\xbc\xc5\x89\xe5\x31\xcc\x70\x5c\xdc\xa6\x92\x09\xb8\x92\xc2\x28\xc9\xe1\x3d\x15\xc8\xe7\x61\x13\x72\x59\x9c\x89\x5b\x28\x15\xe6\x89\x67\xb5\xea\x38\x0c\xd3\x4c\xdc\xe8\x20\xe5\x72\x95\xe5\x9c\x2a\x0c\x52\x59\x85\xf4\x86\x7e\x0e\x39\x5b\xea\xd0\x6c\x98\x31\xa8\xfc\xa5\x94\x46\x1b\x45\xeb\x70\x1c\x8c\x83\xf3\x30\xd5\x3a\xdc\xf9\x82\x8a\x89\x20\xd5\xda\x03\x85\x3c\xf1\xb4\xd9\x72\xd4\x25\xa2\xf1\x20\x5c\xfc\x58\xdd\x5c\x0a\xe3\xd3\x0d\x6a\x59\x61\x38\x09\xce\x03\xe2\x4a\x1e\xba\x7f\xa4\xaa\xc5\xeb\xa0\x90\xb2\xe0\x48\x6b\xa6\x5d\xd5\x54\xeb\x17\x39\xad\x18\xdf\x26\x6f\x6d\x1c\x95\xa2\x26\x1e\x11\x32\x1c\x13\x32\x9c\x10\x32\x9c\x12\x32\x3c\x23\xe4\x78\x25\x5b\x4a\xa7\x8a\xd5\x06\xb4\x4a\xff\xf0\x1d\xde\x7c\x5a\xa1\xda\x86\xe3\x20\x0a\xa2\xd6\x70\x77\x74\xa3\xbd\xc5\x3c\x6c\x08\x17\x3f\xc5\xed\x0b\x69\xb6\xe1\x28\x98\x04\x51\x58\xd3\xf4\x96\x16\x98\x75\x95\x6c\x28\xe8\x9c\xbf\xac\xee\x63\xdd\x72\xf3\xb0\x59\x7e\x45\xb1\x4a\x56\x28\x4c\x70\xa3\xc3\x51\x10\xcd\x02\xd2\x39\xbe\xe6\x77\x05\xec\x4b\xb3\xa5\x7a\x65\x34\x84\x72\x34\x84\x72\x0c\x5f\xac\xdd\x73\x4d\xd5\x34\x40\x0c\xfb\x0e\xb8\xdc\x07\x37\xc8\x8a\xd2\xc4\x30\x22\xc4\x79\xef\xec\x1f\xda\xa0\x53\xc9\xa5\x8a\xe1\x69\x94\x93\xf1\x38\x3b\x02\x9a\xdc\x03\xc5\xa5\x5c\xa3\xba\x0f\x9d\x46\x51\x36\x8b\xbe\x07\x0d\xd6\xa8\x0c\x4b\x29\xf7\x53\x14\x06\x55\xab\xbe\x62\xc2\x2f\xdb\xfc\x88\x90\x67\x97\xc7\xbc\xeb\xb2\x71\x67\x4c\xd7\x9c\x6e\x63\xc8\x39\x7e\x6e\x5c\x94\xb3\x42\xf8\xcc\x60\xa5\x63\x68\x98\x9b\xc0\x92\xa6\xb7\x85\x92\x2b\x91\xf9\xc7\xee\xb1\xf3\xe5\x79\x7e\xf9\x83\x8f\xf1\xa0\x84\x7d\x0a\xe3\x28\x9d\x4d\xe1\x09\xab\xec\x16\xa4\xa2\x41\x02\x1c\x26\xf9\x95\xfc\xcd\xe7\x4c\x20\x55\x7e\xa1\x68\xc6\x50\x98\xbe\x91\xf5\x70\x07\x27\xcf\xec\xf5\x79\xb4\x9c\x4d\x20\x9a\x58\xe3\x6c\x12\x5d\xcc\x08\x8c\x5d\xe4\x9c\x46\xd3\x73\x0a\x93\x33\x6b\x5c\xcc\x22\x72\x1e\xc1\xd9\xc8\x1a\x4b\x24\xf4\x6c\x0a\xe7\x17\xd6\xc0\x8c\x8c\xa6\x33\xb8\x70\x69\xf9\x05\x21\xd3\x89\x7b\xb8\x83\xef\xc8\xdb\xe0\xf2\x96\x99\xbf\xb1\xc2\xaf\x95\xc1\x52\x1a\x23\xab\xbf\x50\x5f\xce\xb8\x41\x15\x43\xad\x64\xc1\xb2\xf8\xd5\xc7\x37\x15\x2d\xf0\x5a\x51\xa1\x73\xa9\xaa\xe0\x2d\x4b\x95\xd4\x32\x37\xc1\x4e\x36\x68\x43\x95\xb9\xb2\x2d\xa8\x8d\x4a\x4e\x5b\xed\xa7\x43\x40\x91\x1d\xb8\x9b\xc2\xa7\xc3\x7f\xb5\xc0\xeb\x6d\x8d\x09\x81\xc1\xc1\x54\xd9\xaa\x0a\xb5\x6e\xc7\xa9\x96\x9a\x19\x26\x45\x6c\x77\x3d\x35\x6c\x8d\xc7\x72\x75\x4d\xc5\x57\x00\xba\xd4\x92\xaf\x0c\x3e\x18\xb5\x25\x97\xe9\x6d\xe3\x73\x9f\xf0\xc3\x31\x6d\x87\x68\x53\xb2\x16\x06\xae\x10\xd4\x0a\x5b\x7a\xa8\x69\x96\x31\x51\xc4\x70\x56\xb7\x13\x0b\x15\x55\x05\x13\x31\x90\x3d\x64\x1e\x76\x0b\x6e\x1e\x36\xa7\x95\x93\xde\x7c\x29\xb3\xad\xdb\xae\x19\x5b\x43\xca\xa9\xd6\x89\xf7\x60\x89\xb8\x53\xc8\xbd\x04\x7b\xf8\xa0\x4c\x74\xa1\x7b\x31\x25\x37\x1e\xb8\x42\x89\xd7\x88\xf0\x9b\xee\x89\x21\xb2\xf2\x5a\xc8\x03\x3e\xee\xf3\xc2\x8f\x46\x5d\xb0\x37\x2f\xa3\x8e\xc4\xe0\x67\xe3\xbb\x0d\xb4\xdb\x3d\x76\x71\xb3\xe3\x27\x99\x32\xea\xf8\xc3\x8c\xad\x5b\x79\x07\x97\xdf\x56\x6a\x64\x1d\xc3\x78\xf4\x0d\x99\xba\xf2\xa7\xde\x62\xce\x3a\x5f\x4e\x21\xa7\x7e\x46\x0d\x5d\x52\x8d\x1e\x50\xc5\xa8\x5f\xb2\x2c\x43\x91\x78\x46\xad\xb0\xd1\x0a\x73\xd7\x0e\x2c\xb3\x2c\x4c\xb8\x5c\xfb\xf5\xa9\xa9\x58\x40\xe7\x39\xd0\x79\xac\xee\xf8\xe7\xea\xba\x1e\xdb\x17\x75\xa6\x3e\xfa\x94\x76\x57\xed\xc5\xfe\xe3\xeb\xc2\x6b\xaa\x40\xa3\x5a\xdb\x8f\x80\xb5\xc3\x10\x5e\x61\xce\x04\x02\x85\x0a\x4d\x29\x33\x30\x12\x14\xa6\x52\x08\x4c\x0d\xac\x6a\x29\x5a\x00\x70\xa9\x75\xc7\xb1\xcf\x48\x20\x5f\x89\xd4\x4e\x48\x7f\xd0\xf6\x74\x9b\x9f\x80\xc0\x0d\xfc\x0f\x97\x1f\x64\x7a\x8b\xa6\xdf\xef\x6f\x98\xc8\xe4\x26\xe0\x32\xa5\x16\x60\xe7\xcd\xc8\x54\x72\x48\x92\x04\xda\x73\x81\x37\x80\x17\xe0\x6d\xb4\x3d\x21\x78\x10\xdb\x4b\x7b\x35\x80\xe7\xf0\x10\x5e\x4a\x6d\xe0\x39\x78\x21\xad\x99\x37\x68\x6e\xa8\x2d\x1e\x48\x51\xa1\xd6\xb4\xc0\x43\x81\xb8\x46\x61\x3a\x95\xee\x3e\x2a\x5d\x40\x02\xff\xfe\xf0\xee\x3f\x41\x6d\x4f\xfd\x4d\x4a\x60\xdf\x4d\xb3\x44\x7a\x3d\x96\x43\xdf\xa5\x25\x09\x88\x15\xe7\x3b\x7c\x4f\xa1\x59\x29\xd1\xa6\xdd\x9d\xdc\x4b\x0f\xdc\x3b\x82\x27\x49\x02\x2b\x91\xb9\x47\x9c\xed\x91\xff\xe8\x7b\x4f\x9b\x77\x3a\x08\xec\x8c\xec\x11\x83\x1d\xdd\x3d\xb6\xae\xcd\xbe\x45\xb8\x6b\xce\x03\xce\xce\xf7\x08\x2d\x2a\x25\xd5\x63\x9c\xf6\x0c\xd9\xff\xc2\xe9\x56\xae\x4c\x0c\xa7\x46\xd6\x57\x6e\x84\x4f\x87\x60\xf9\x63\xd8\x31\x0c\xc1\x6c\x6b\x8c\xe1\xd4\x59\x36\xce\x2a\x74\xa8\x29\x21\x64\x08\xdd\x5e\xfd\x27\x55\x31\xd8\x06\xbf\x7b\x44\x8f\x5e\xa5\xa9\xdd\xbf\x3f\xa3\xa8\xe5\xd8\x69\x6a\xed\x3f\xaf\xea\xee\x7e\x3f\xa5\x5c\x6a\x7c\xd0\xee\xa0\xd1\x5c\x37\xac\xfd\xdd\x48\x0c\x61\x4c\x08\x19\x5c\xc2\xdd\xfe\xd3\x12\x86\xf0\x5a\x1b\xba\xe4\x4c\x97\x40\x61\x83\x4b\xed\x86\x02\x5a\x0c\x93\xc2\x0e\x9e\x29\x11\x5e\xbe\x7f\xd3\x4e\xdc\x89\xeb\xb1\x36\xa3\xef\xc4\x1d\x1c\xaa\xe7\x61\xb3\xff\xe7\x61\xf3\xbb\xf6\xf7\x00\x00\x00\xff\xff\x21\xe7\x38\x11\xe8\x0e\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"index.html": indexHtml,
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
	"index.html": {indexHtml, map[string]*bintree{}},
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