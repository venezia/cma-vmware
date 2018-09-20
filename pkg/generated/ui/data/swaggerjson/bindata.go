// Code generated by go-bindata.
// sources:
// assets/generated/swagger/api.swagger.json
// DO NOT EDIT!

package swaggerjson

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
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

var _apiSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\xdd\x8f\xdb\xb8\x11\x7f\xdf\xbf\x62\xa0\x16\xe8\x4b\x6e\x9d\x4b\x5f\x0e\xfb\xd2\x06\xbb\xc0\x65\x71\xd9\x6b\x70\x69\x93\x87\x36\x58\x8c\xa5\xb1\x34\x59\x8a\x54\x48\xca\xae\x5b\xec\xff\x5e\x0c\x25\xd9\x92\x2c\xf9\x33\x9b\x38\xe9\xe5\x25\x89\xc9\xe1\x7c\xfd\x38\x5f\xb4\xff\x7b\x01\x10\xb9\x05\xa6\x29\xd9\xe8\x0a\xa2\x17\x97\xcf\xa3\x67\xf2\x19\xeb\x99\x89\xae\x40\xd6\x01\x22\xcf\x5e\x91\xac\x5f\xab\xd2\x79\xb2\x70\x87\x1a\x53\xb2\xf0\xee\xee\x3d\x5a\x82\x57\xa4\x0a\xb2\xf0\xf2\xcd\x6d\xa0\x06\x88\xe6\x64\x1d\x1b\x2d\x34\xf3\xe7\x97\x3f\xd6\xc7\x02\x44\xb1\xd1\x1e\x63\xbf\x3a\x1b\x20\xd2\x98\x87\xc3\xef\x38\xce\x90\x14\xbc\x23\x4d\xff\x61\xac\x29\x00\xa2\xd2\x2a\x59\xcf\xbc\x2f\xdc\xd5\x64\x92\xb2\xcf\xca\xe9\x65\x6c\xf2\x89\xc3\xdc\x95\x3a\xfd\x21\xd6\xb1\x9f\xc4\x39\xfe\x30\xcf\x17\x68\x69\x4d\x4a\x39\x72\x20\xce\xe7\xd5\xa9\x7f\x4d\xe5\x13\x21\x8e\xc2\x9e\xc7\x0b\x80\xc7\xa0\xb2\x8b\x33\xca\xc9\x45\x57\xf0\xcf\x4a\xd4\xc0\xaf\x91\x5b\xfe\x23\x14\x1f\xc2\xde\xd8\x68\x57\x76\x36\x63\x51\x28\x8e\xd1\xb3\xd1\x93\x8f\xce\xe8\xf5\xde\xc2\x9a\xa4\x8c\xf7\xdc\x8b\x3e\x73\x6b\xbb\x4f\xb0\xe0\xc9\xfc\xc7\x49\x5c\x99\xbd\x6d\xb4\x94\xda\x36\x14\xf1\xcb\x3c\x47\xbb\x14\x5d\xdf\xb3\x52\x60\xc9\x5b\xa6\x39\x81\xcf\x08\x9c\x47\x5f\x3a\x30\x33\x40\xa8\x0f\x03\xd4\x09\xb0\x77\xf0\x50\x4e\x29\x36\x7a\xc6\x29\xcc\x8c\x85\xd8\x68\x4d\xb1\xe7\x39\xfb\xe5\xca\x8e\x00\x91\x29\xc8\x06\x91\x6f\x13\xe1\xf1\x33\xf9\x1a\x0c\xed\x4d\x96\x5c\x61\xb4\x23\xd7\x91\x0d\x20\x7a\xf1\xfc\x79\xef\x23\x80\x28\x21\x17\x5b\x2e\x7c\x0d\x94\xd6\x41\x95\x46\xe2\x10\xdc\x20\x03\x88\xfe\x68\x69\x26\x14\x7f\x98\x24\x34\x63\xcd\x72\x82\x13\xff\x57\xee\x5f\xcb\xf6\x1b\x15\x6a\x19\x75\xc8\x1f\x2f\x86\xfe\xfd\xd8\x52\xa2\x40\x8b\x39\x79\xb2\x6b\x97\x55\x7f\x7a\xe2\x37\xb8\x0d\x7f\x3f\xdb\xaa\xda\xaf\x98\x93\x58\x5f\x7c\xd1\xd8\xdf\x1b\x98\x12\x28\x63\x1e\x28\x81\xb2\xb8\xec\x1f\xc1\x81\xf2\x53\x49\x76\xd9\x5f\xb2\xf4\xa9\x64\x4b\xe2\x88\x19\x2a\x47\xbd\x65\xbf\x2c\x82\x60\xce\x5b\xd6\x69\x34\xa8\xf0\x87\x96\xc2\x1e\xd3\xbe\xaa\xcd\x4d\x5f\x13\x7f\xb8\xe8\x59\x2a\x4a\x48\x91\xa7\xed\x28\xac\xf6\xac\x51\xb7\x05\x51\x37\x61\xeb\xd9\x82\xaa\x23\xde\xb9\xe0\xea\x7d\x86\x1e\xd8\xb5\x71\xf5\x27\x07\x42\x28\xf0\x4a\xc8\x79\x6b\x96\xdf\x1e\xb2\x8a\x72\x47\x70\xc3\xe4\x63\xe9\x3c\x20\x14\xd6\xcc\x59\x32\xcd\x5e\x10\x7b\x19\xc8\x6a\x01\x7e\x35\x09\xb9\xf3\xc3\x59\x47\xc6\x2f\x82\xb3\xa9\x49\x36\x70\x50\x41\x64\x68\xa5\x85\x10\x6f\xcb\x3e\x40\x3e\x8f\xda\x77\x2e\xdd\x47\xe9\xe3\x91\x76\xd1\xb2\x59\x3f\xcb\x4e\x14\x3b\x7f\x5c\xaa\x45\x10\x5a\x09\xf4\xf5\x59\x6e\xaf\x0c\xfa\x5a\x18\x9e\x1d\x10\xbb\xf2\x1d\x85\xc4\xcf\xe8\x94\xb2\x48\x2d\x26\x74\xa8\x5f\x4a\xab\xa1\x26\x05\x13\x6c\xe4\x42\x95\x83\x90\xf2\x9c\xf4\x1e\x31\xe3\x67\xf2\xff\xa8\x0e\xa8\x25\xbf\xd5\x33\x63\xf3\xb0\xe3\x2c\x9d\x36\x2a\xed\x19\x27\x2d\xf0\xf2\xd9\x82\x40\xba\x09\x29\xae\x39\x21\xc9\x30\xc1\x57\xb5\xff\xbe\xc7\x34\xe6\x3d\xe5\x85\x97\x5c\xdd\x80\x74\x9f\x34\xd6\xf5\xf0\xf9\x81\xb0\x2b\xdf\xff\x4f\x0e\xeb\xea\xfd\x75\x92\xd8\xba\xf1\x3e\x38\x4e\xd6\xa4\xc0\xeb\x90\x01\x38\x35\xa5\x07\x2c\x18\x1c\xd9\xf9\xae\x40\xf9\xae\x3a\xe1\xdc\x23\x64\x2d\xe6\x17\x4b\x69\xab\x11\x43\x4b\xa0\x75\x93\xdf\xaf\x7c\xaa\x99\xca\x6f\x94\x9b\x39\xdd\x61\x9c\xb1\xa6\xb7\x05\xc5\x6d\x8f\x36\xf1\xcb\x4c\x3f\x52\xbc\x2e\x1d\xa2\xc2\x8a\x4f\x3c\xf7\x4c\x1c\x65\xc6\xf9\xbe\xd1\x7b\x31\xf0\x59\x67\xad\x99\xf7\xfc\x3d\x23\x10\xe2\x10\x87\xdf\xbe\x7d\x05\x18\xc7\xe4\xdc\x5a\xd3\xc7\x41\x30\xf6\x6c\x3c\x00\x8b\x13\x94\x49\xd9\xdf\x6f\xa2\xfc\x30\x9d\x3c\xa6\x60\x74\xc8\x40\x29\x7b\xb0\x54\x18\xc7\xde\xd8\x16\x1c\xda\x4e\x17\x96\xb1\xc9\x73\x3e\xc1\x8a\xe8\xb2\x66\x02\x20\x2c\xeb\xe3\x46\xd9\x79\x4b\x74\xef\x3c\xf6\x1a\xec\x7d\x59\xbe\xcf\xc8\x67\x64\xc1\x58\xd0\xc6\x07\xae\x72\x22\x2c\xd0\x41\xac\x08\x35\x2c\x32\xd2\x30\x2d\x59\x8d\x08\x21\x4b\xc9\x7d\x72\xac\x00\x37\xe8\xc3\xc4\x23\x1c\x33\xa2\xa6\x39\xc9\x8f\x35\xaa\x84\x49\x6a\xa0\x74\x94\x48\x1e\x8d\x4d\x5e\xb0\xa2\x61\x8e\xf5\xa2\x3d\x8a\xdf\x75\x4d\x1c\x58\x0d\x9f\x5f\x28\xf4\x82\xf1\xa3\xce\x7f\x53\x13\x03\xfb\xca\x4d\x15\xbf\x24\xdc\xbd\x09\xd8\x52\x6b\xa9\x88\xaa\x71\x61\xcd\x7b\xf0\xf6\x8d\xf7\x54\x27\xdc\xba\x3a\xd1\x1e\x83\xc4\xa3\xaa\xbd\x61\x03\x63\x92\xdc\xeb\xd0\xb7\x8f\x88\x82\xd6\x62\x37\xdd\x47\xec\x29\xef\xef\xdf\x9d\x26\xaa\x28\xdc\x8e\xbf\xed\xd4\x30\xac\x6a\xbd\xdb\xc1\x22\xe3\x38\x13\x05\x17\xa8\x43\x75\x87\x49\x00\x67\xcb\x06\xc3\xfa\xd9\x10\xf5\x9f\x56\xc5\x7d\xf3\xcd\x29\xfa\x56\x7a\xc0\xcc\x9a\x7c\x44\xe9\x03\x90\x5b\xe5\xe9\x13\xb0\x6b\x1e\xc6\x6c\x39\x35\x46\x82\x61\xd7\x9a\x55\x96\x1a\x5d\x5e\x23\x1b\x05\xd5\xec\x00\xc1\x95\x21\x2f\xce\x4a\x29\xa6\x3e\x95\xe4\xfc\xde\xaa\xd6\x4a\xde\x90\x47\x56\xb7\x9e\xf2\x53\x34\xe5\xe4\xa8\x3b\x7a\x7b\xd3\x1b\x4e\x0f\xa3\xf3\xe8\x18\x30\x30\xfe\x1e\xe6\x50\xbd\x53\x9c\x1c\x67\xd6\xcf\x1d\x3b\x39\xae\x5f\x3f\x4e\xe6\xda\x7a\x48\x09\x89\x28\xbc\xa3\x8c\x5f\xfb\xbd\x60\xf1\x3b\x20\xbe\x08\x20\x76\xf8\xc2\x12\x7a\xfa\xca\x79\x74\xfc\x09\x69\x35\x03\x1f\x2b\x4b\x1e\x7e\x72\xc7\x54\x5a\xbd\x0e\x4c\xaa\xd8\xf9\xba\xe6\xfa\xa5\x9c\x92\xd5\xe4\xa9\x1a\xa8\x2d\x8c\x7d\x20\x29\x36\x13\x72\x97\x70\x6d\xb4\xb7\x46\x41\xa1\x50\xaf\xa8\x5c\xc8\xf7\x89\x74\xf5\x39\x6b\x4a\x60\xba\x0c\xda\xb4\x92\xce\xe5\xb0\x02\x19\xa7\xd9\x3d\xce\x91\x15\x4e\x59\xb1\x5f\x3e\x4d\x3c\xdf\xac\x99\x1b\x43\xb3\x83\x57\x2f\x47\xee\x00\x79\xd1\xfd\x7e\x86\x53\xcb\xf1\xd1\xfd\x41\x45\x5e\x7b\x74\xbc\xc2\x8c\x2b\xc3\xde\x07\xc3\x7e\x3b\xa5\x90\xd4\xb2\x96\x1d\xed\xbe\xf0\x15\x8e\xbe\x47\xcd\xb0\xe0\x7b\xd2\x49\x61\x58\x1f\xdb\x46\xb2\x03\x97\x99\x52\x25\x02\x12\x84\x39\xaa\x92\x40\xf1\x03\x01\x17\x57\x85\xb1\xbe\xae\xac\x59\xa9\x7a\x07\x5b\x5f\xa2\x82\xdb\x37\x13\x59\xfe\x97\x7e\x83\x4e\x1a\xa5\x29\xc6\x0f\x82\x35\xfa\xb7\x27\xab\x51\x41\x5c\x3a\x6f\x72\xb2\xae\x46\x20\x4e\x15\xd5\xed\x54\x5e\x6a\x8e\xa5\x9b\x3b\x36\x93\xb5\xa3\xe7\xd9\xd6\x72\xe3\x77\x5f\x7a\xb1\x56\x88\xad\xa2\x16\xbb\x81\x42\xaf\x7d\x4f\x37\xbe\xb1\xb1\x0f\x1e\xdb\x49\x7f\x3c\x54\x24\xa1\x54\xec\x67\xb3\x46\x1e\x68\xc6\x6c\x7b\x3b\x68\xe0\x65\xfb\xfc\x1c\x74\x1d\x60\xdf\xd6\x77\x4a\xcd\x03\xf7\x58\xb0\x3c\xa1\x78\x78\xdb\x29\x18\x0e\xad\xe9\x87\x9e\xce\xce\xcf\xa4\xb7\xdd\xa6\x9c\xab\x79\x98\x5b\x3a\x81\xdf\x36\x54\x3f\x7d\x60\x6e\x5f\x84\xdd\x21\xf9\x75\xff\xd1\xf5\x70\x3f\x7d\x77\x3e\x3a\x2e\xf2\xb4\xba\xd0\x03\x6c\xb8\xe3\xc5\xf1\xfc\xac\x7a\x8d\xba\x1f\x4a\xea\xb1\xd3\x48\x24\x69\x2a\xd8\xd3\x81\xbf\xe5\xe9\x71\x4b\x1b\xd3\xa9\xa0\x0b\xe3\x1c\x4b\x7e\xb6\x9c\x66\x1e\xb4\x59\x1c\xe2\xac\xce\xe3\xc7\xf9\xb9\xe6\x76\x06\xab\x77\xa5\x90\x7a\xff\xf6\xcb\x56\x97\xdc\xf3\xe0\xab\x02\x8c\x63\x7e\xf7\xe3\xc4\xf6\x79\x73\x7b\xe7\xa6\xdd\x57\x76\x5b\xd1\x05\x2e\xa1\x47\x6a\xba\xa7\x4e\x3a\xd9\xf0\xd1\xba\xb1\x7a\x8d\x53\x52\x5f\xa5\xd9\x94\x32\x43\xd7\x0d\x27\x82\x0a\x72\x0c\x7b\x41\x4a\xd0\xa3\x59\x54\x05\xec\x20\x8f\xed\x38\xde\x7c\x69\xfd\xe6\x47\xdb\xa7\x3c\x85\xb4\x63\x44\xe0\xbc\x68\x02\x9a\xb0\xf6\xe6\x2f\x47\x1a\xf6\x5c\x83\xc4\xe7\x99\xbc\x6e\x76\x7f\x27\x68\x5a\x3a\x69\xa4\x4e\xb8\x6e\xcd\x01\xe3\x4f\xad\xed\x81\xc8\xd3\x3d\xe8\xb6\x9f\xb3\x8c\x1d\xe5\xc2\xda\x53\xda\xf9\x3e\x40\xc7\x7f\xac\xfd\x9f\x5f\x6c\x91\x21\x34\xab\x7b\xc9\x80\xce\x2d\x8c\x3d\x68\x8c\x39\x30\xb8\x7a\xf8\xc9\xad\x2e\x88\xb0\x0d\x57\xb6\x3d\xa4\xba\x84\xd0\x5e\x6b\x93\x90\x5c\x69\xa3\xd5\x12\x10\x72\xac\xca\xae\x59\x05\xb7\x19\x93\x4a\x64\x39\xe4\x13\x4a\x46\xc6\x55\x21\x96\x3d\x7d\x91\xdc\x4f\x14\xbb\x6b\x09\xb1\x44\x25\xdc\xca\x08\x79\x3d\xce\xd8\x2b\x95\x09\xbd\x2b\x28\xe6\x59\xfd\xf3\x85\xfa\xbb\x74\xcd\x67\xc1\x7c\xfd\xdf\x55\x34\x33\x86\x1b\x13\xb7\xbe\xf5\xd0\xf3\xd1\x9d\xb1\x54\x7f\xd9\x64\xef\xdf\x98\x1c\xfe\xb3\x10\x91\xe9\xe2\xf1\xe2\x7f\x01\x00\x00\xff\xff\x83\xc3\x4b\xa5\xf6\x32\x00\x00")

func apiSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_apiSwaggerJson,
		"api.swagger.json",
	)
}

func apiSwaggerJson() (*asset, error) {
	bytes, err := apiSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.swagger.json", size: 13046, mode: os.FileMode(420), modTime: time.Unix(1537404638, 0)}
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
	"api.swagger.json": apiSwaggerJson,
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
	"api.swagger.json": {apiSwaggerJson, map[string]*bintree{}},
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

func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
