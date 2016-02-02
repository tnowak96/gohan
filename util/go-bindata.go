// Code generated by go-bindata.
// sources:
// etc/schema/core.json
// etc/schema/gohan.json
// etc/extensions/donburi.js
// etc/extensions/donburi.yaml
// etc/extensions/gohan_extension.yaml
// DO NOT EDIT!

package util

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

var _etcSchemaCoreJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x57\x3f\x8f\x9b\x30\x14\xdf\xef\x53\x58\x5c\xa7\xaa\x77\x74\xe8\x74\x5b\xa4\x2e\x95\x2a\x35\xea\x5a\x65\x70\xc0\x24\x3e\x11\xdb\xb5\x8d\xda\xe8\x94\xef\x5e\x1b\x83\x13\x8c\x9f\x81\x1c\x65\x02\xfb\xfd\xf9\xbd\xff\x8f\xb7\x07\x64\x9e\xec\x83\x2a\x8e\xe4\x84\xb3\x17\x94\x1d\xb5\x16\x2f\x79\xfe\xaa\x38\x7b\x72\xa7\xcf\x5c\x1e\xf2\x52\xe2\x4a\x3f\x7d\xfe\x92\xbb\xb3\xc7\xec\x93\xe3\xc4\x75\xfd\xa3\x32\x7c\xbf\xda\x4f\xfb\xbc\xf9\x37\x27\x5a\x92\x6a\x89\x5c\xcf\x7d\x69\xdf\x76\x9d\x9e\x92\x54\x94\x51\x4d\x39\x53\x46\xda\x55\x47\x26\x89\xe2\x8d\x2c\xc8\xe0\xb4\x63\x51\x85\xa4\xc2\xf2\x58\x00\x3f\x3b\x42\x74\x15\xd5\x19\xe1\x39\x84\xe4\x82\x48\x4d\x89\x1a\x49\x8b\x48\x1c\x13\xf4\x48\x71\x53\x6b\xab\x32\x90\x0f\x21\xfb\x7a\xf3\x09\x70\x68\xaa\x6b\x6b\xe2\x80\x15\xa2\x3d\x8b\x96\x54\x69\x49\xd9\x21\x1b\x11\x5d\xc6\x7c\x19\x2d\x53\x06\x0d\xd0\x1a\xd2\x29\x90\x09\x92\x3b\xb0\x09\x2c\x09\xd3\x6b\x3a\x7c\xdb\x4a\x44\xbc\x42\xfa\x48\x15\xea\x92\x7f\xca\xaa\x0e\xc7\x9a\x96\xd5\x8d\xc4\xf5\x6c\xcf\x6f\x5b\xf2\xa5\xb0\x1d\xd7\xaa\xb0\x4d\x49\xd3\xbf\xab\x06\xc4\x49\x9c\x8c\x40\x9a\xec\x0e\x53\x7c\xdf\x03\x4c\xe9\xbb\xd7\xe3\x3c\x71\x3d\xd2\x79\xf1\xfc\x8e\xf7\x64\x71\x38\xdd\xcb\x8a\x2e\xc0\xc5\xb8\xab\xa6\x50\xfb\x3e\xba\xe9\x18\x01\x2c\x02\x6b\x4d\x24\xdb\xa6\x7b\xaa\x27\x7f\xfe\x98\xbc\x77\x22\xe7\xc9\xf2\xf4\x94\x89\x06\x6e\x1c\x23\xf2\x44\xb0\xc3\x27\xe2\xc8\x91\x38\xde\xe8\xff\xa6\x3e\x49\x31\x01\xce\x67\x09\xdf\xbf\x92\x42\xc3\xda\xe2\x5a\x00\xe9\xd7\x0c\x9d\xc8\x8b\x29\xed\x43\xad\x81\x36\x33\xee\x7f\x37\x54\x92\x72\xb0\x72\xf8\xdb\xe8\xe0\x01\x4b\x26\x03\x4b\xae\xef\xcc\x70\xf3\x1b\x5c\xec\x02\x90\xde\x15\x07\x7e\xc4\x0c\xf9\x15\x25\x24\x03\x3c\x71\x63\xb3\xa9\x64\xac\xc9\x29\x36\xff\xba\xc2\x8d\x6f\x2a\xa9\x46\x10\xba\xb4\xe0\xac\xa4\xb0\x28\x6a\xf4\x27\xba\xc3\x3d\x2d\xa7\xe7\xc1\x52\xe2\x73\x1a\x5c\xd0\x7c\x04\xaf\x69\x71\x46\xde\x2b\xa8\xe2\x12\x39\x9a\xbd\x51\x8f\xba\x7b\x7b\xba\xd9\x7e\x0b\x1d\x4e\xaa\xca\x3a\x3a\x6a\x26\x61\xcd\x29\x9a\x54\xed\xad\x59\x73\xf9\x1f\x78\x98\xb1\xf3\xd8\xf0\x5d\xc2\xf0\x39\x61\x01\x56\xb2\x59\xeb\xd8\xc4\x2a\xb6\x08\x87\x30\x24\x05\x15\xc0\x9e\xb2\x48\x14\xb8\xab\x3b\x45\x44\x9e\xa8\x52\xce\x2e\x20\x0e\x85\x24\x26\xf2\x50\x20\x1a\x51\xda\xdb\x59\xa1\x98\x31\x4b\xec\x08\x3b\xa6\xa7\xd6\x54\xf6\xdb\x07\x6a\x98\x33\xa7\xd9\x44\x01\x2e\x82\x92\x80\x33\x90\x11\x29\x4c\xcf\xbe\xa8\xc2\x63\x4d\x3e\xcc\x8a\xeb\x7a\x17\x14\x37\xd4\x11\xad\x7d\x9b\x16\xe1\xa8\x27\xc2\xbe\x5a\x94\xa8\x80\x1f\x9c\xed\x17\xff\x4b\x3a\x28\xc4\x8d\xe9\xf3\x66\x38\x28\x03\x1c\xdb\x33\xbb\xd8\x61\x34\x9c\x00\xfd\x50\x84\x62\x9f\xfc\xc3\x04\x2d\xb8\xf5\x8e\xf3\xe1\x22\xc7\xf8\xbd\x23\xbf\xf9\xc9\xce\x23\x51\x58\xe2\xa9\x1b\x48\x57\xe3\xdf\x8f\xca\x77\x90\xf7\x82\x8a\x6f\xea\xb0\x8b\x07\x91\x0f\xb3\xfb\xf2\xf0\x2f\x00\x00\xff\xff\x95\x62\x9d\x2b\x48\x11\x00\x00")

func etcSchemaCoreJsonBytes() ([]byte, error) {
	return bindataRead(
		_etcSchemaCoreJson,
		"etc/schema/core.json",
	)
}

func etcSchemaCoreJson() (*asset, error) {
	bytes, err := etcSchemaCoreJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etc/schema/core.json", size: 4424, mode: os.FileMode(420), modTime: time.Unix(1445535064, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcSchemaGohanJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5d\x4b\x6f\xdc\x36\x10\xbe\xfb\x57\x10\x8b\x1e\xda\x20\x48\x9a\x6b\x6f\x41\x63\x14\x29\xdc\xd8\xb5\x93\x5c\x0c\xd7\xa0\x77\xb9\xb6\x52\x89\xda\xe8\xe1\x64\x11\xf8\xbf\x57\x24\x25\xea\xb1\xe2\xf0\x25\xc9\x92\x6b\xa2\x07\xa7\x1a\x0e\x67\x86\xe4\xf0\x9b\xe1\x63\x7f\x1c\xa1\xa2\xac\x76\x71\x18\xac\x03\x92\xae\x7e\x43\x97\xfc\xff\xb0\xf2\x43\xfe\xc5\x69\xf0\x3a\x0b\x62\x5a\x50\xac\x5e\xac\x5e\xb6\x3f\x91\xed\x96\xac\x33\xf6\x09\x87\x61\xfc\xad\xfb\x39\xd8\xf0\x4f\x9b\x28\xa0\xd7\x69\x86\x33\x12\x11\x9a\x75\x89\x76\x49\x40\xd7\xc1\x0e\x87\x92\xb6\x4b\x91\x90\x34\xce\x93\x35\x29\x08\xda\xa2\x89\xfa\x38\xbb\x63\x55\x5f\xbd\x58\xb5\x3e\x3e\x1c\xb5\xff\xba\x12\x5c\x57\xe9\xfa\x8e\x44\x18\xd4\x78\x43\xd2\x75\x12\xec\x2a\xb5\x3f\xde\x11\x14\xe1\x80\xa2\x88\x64\x58\x54\xef\xd7\xb4\xff\x1b\xab\xb5\xc1\x19\xee\x17\x3f\xdb\xef\x98\x62\xab\x06\xef\xb6\x1a\x5d\x73\x85\x79\x22\x6c\x55\x29\x72\x60\x4f\xb2\x0d\xbe\x33\x82\xd7\xb7\xf1\x1d\xa6\xaf\xef\x7f\x7d\xf5\xa6\x4b\x54\xb6\xd4\x2b\x11\xde\x6c\x02\xa6\x3a\x0e\xcf\x92\x78\x47\x92\x4c\x8c\x8f\x2d\x0e\x53\xf2\xb2\xc7\xfe\x4d\xa2\x43\x76\x3d\x06\xed\x27\x2a\x09\xb7\x38\x0f\xf9\x88\x5a\x1d\xb6\xa5\xe0\xb7\x7a\xd7\xf8\x27\x50\xab\x10\x33\x0a\xd2\x54\x54\xba\x54\xd2\x71\xda\x75\x42\x8a\xf1\x0a\x70\xe3\x54\xf9\x6e\xc3\xa8\x94\x44\x57\x80\x34\x59\x90\x85\xc4\x42\xfa\x6a\xa0\xa4\x59\x31\x61\x6e\xfb\xdb\x7c\xe8\xaf\x2f\x86\x27\x68\xf6\x96\x3d\x2f\xf8\xe8\x40\xef\xdf\x0d\x6c\x4d\x3f\x3b\xc1\xe2\x74\xcc\x03\x50\xe6\x34\xf8\x9a\x33\xda\x2c\xc9\x89\x95\x19\xc1\x99\x2c\xa9\xea\x31\xfc\x43\xc1\xa8\x24\x6b\xd9\xbc\xe2\x8d\xb6\x71\x82\xf0\x6e\x57\xf8\x65\xcc\xbe\xa1\x0d\xb9\x27\x21\x9b\x62\x90\x4e\x45\xa5\x08\xf3\x69\xb3\xc7\x51\xb8\xc0\x49\xf0\x57\x65\x5a\x83\x2e\x8e\x6f\xbe\xb0\x55\xc7\xaa\xeb\x28\x8e\x48\xba\xc3\x8a\x45\x44\x92\xb9\xf9\x9f\x0f\x15\x73\x94\x15\x4b\x85\xf0\xad\xe8\xa6\xe8\x35\x7a\x9b\xa2\x2c\x9e\xd5\x2c\x92\xb2\x8e\xe5\x6b\x76\x38\x61\xab\xfc\x08\x66\x3e\xe3\x9c\x51\xbc\x2d\xcc\x1c\xa4\xa8\x77\xb5\x6d\x8b\x32\xcf\xb1\x2e\xf4\x18\xcd\xfe\x15\x40\x30\xf7\xf7\x67\xbc\xca\x13\x30\xac\x50\x7d\x2c\xc3\x56\xc0\x6a\xf0\x81\x5d\x72\x5e\xa0\xc1\xb5\x82\xfb\x18\x1c\x00\xa9\x92\xa6\xb1\xd8\xc2\xfa\x1a\xc0\xd4\x16\xbd\x06\x32\xb5\x79\x9b\xf7\x4e\xab\x9e\xce\xa1\x37\x0b\xd0\x1f\x2d\x9e\x26\x06\x6f\x96\x07\x90\x02\x00\x30\xad\xd6\xa0\x25\x59\xc3\xa7\x2f\xda\x0a\x28\x25\x89\x3a\xdc\xea\x54\xcf\x70\x10\x5e\xdf\x07\xe4\x5b\x89\xe9\x16\x33\x8f\x2c\xc6\xa4\x2a\x2a\x33\x18\xc9\x55\x1f\xdd\xc4\x71\x48\x30\x85\x87\x84\xae\xc3\x6d\xa7\x51\x11\x9d\x67\x24\xa1\x96\x32\xf3\xaa\xff\xbc\x7a\xf1\x93\x31\x35\xaf\x61\x17\xb9\xaa\x65\xb6\x17\x56\x27\x82\x1d\x17\xce\x49\x3a\xd9\xb7\x92\x25\x6a\x48\x66\xae\x4f\xcd\xd1\x66\x20\x74\x8b\x66\x60\x1c\x34\x66\xea\x98\x7b\x2b\x63\xba\x3f\xdd\x5a\xf9\xd2\x66\xb1\x6f\x50\x36\x5c\x99\x08\x27\x09\xde\xdb\x19\xa8\x2a\x96\x86\xaa\xca\x00\x42\x3b\xf5\x6b\x55\x1e\x4f\xec\x80\x66\xe4\xb6\x08\x6f\x17\x26\x36\xcd\xa3\x9b\xe5\x49\xad\x5b\xa8\xa1\xf2\x78\x52\x9b\x82\x99\xbe\x02\x03\x9c\xbe\x62\x08\xb5\x9a\xa5\x99\xcd\x13\x7e\x6f\x64\xe7\x6a\x96\x4f\x55\x32\xb0\x4c\x3e\xaa\xf9\x78\x74\x90\xbd\xd2\x4d\xa0\xe7\xa7\x34\xe3\x84\x3e\x17\x9c\xd0\x29\x0d\xf7\xb3\x5f\x4c\x49\xe1\x6c\xdc\x74\x8e\x02\xfa\x3e\x23\x11\xc3\x20\x6f\x7c\x86\xf5\x31\x93\xc0\xc3\x4c\x62\x41\x75\x60\x20\xb2\xb5\x95\x0e\xca\x94\xad\xaa\xd8\x5a\x5a\x26\x52\x7d\xc6\x57\xc9\x64\xe6\xf3\xc9\x22\xc2\x6d\xd5\x33\x4b\xc8\xab\xab\x4f\xa9\x62\x39\x6c\x1c\xb4\x34\xce\xa8\x2b\x39\xd4\x76\xe2\x52\x78\x98\xca\x65\xd1\xb6\x35\x55\x19\xa8\xf9\x1a\x2b\x29\xa0\x1c\x94\x10\x52\xb2\x68\x64\x47\x85\x20\x33\x1f\x5a\xad\x4c\x82\x83\xc9\xdc\x87\x26\xaf\x5e\xae\x08\x6e\x01\x12\xe7\x60\x94\xde\x00\x39\xe8\x52\x1f\x50\x71\xc0\x58\xbc\x4d\x5f\x44\xe8\x80\x61\x1b\x43\xb3\xee\x73\xef\xa5\x70\xdc\xc1\xe9\x99\x7b\x18\xd0\xf9\x9d\x0d\x92\xb2\x98\xd0\x03\xfa\xe6\x6d\x86\xb4\x9d\x90\x08\x2d\xc7\x86\xcf\xe9\xaf\xb2\xb1\x84\x84\x78\x80\xf0\xec\xbc\x62\x33\xf3\xd5\xb0\x52\xf7\xba\xec\xa2\xfd\x40\x7a\x57\x43\xdf\x2b\x44\x9b\xc6\x00\x5f\xf3\x20\x21\x8e\x90\xda\x13\x0c\x3c\xea\xa2\x78\x5e\x6a\x8e\xb6\x01\x09\x37\x5e\x33\x73\x29\x41\x62\xfa\x15\xde\x5f\x57\x56\x94\x46\xbb\xf8\xfb\x04\x71\xad\x67\x3e\xac\x2b\x81\x7d\x74\xfd\xc8\xff\x98\xbb\xa2\xa2\x2d\x07\x3d\x63\x4a\x1e\x67\xaf\xc2\x3b\x04\xe0\x5c\x5c\xe7\x5c\x8b\x49\xb5\xa4\x7a\xb2\xa9\x76\x02\x3c\xd9\xd0\x3c\x74\xc1\x5d\x1d\x1e\x3c\xbb\xef\xc9\xa5\x84\x5c\x9e\x5c\x7c\xb2\xdf\xac\x38\xc6\x5a\xbc\xed\x7a\x16\xbb\x79\xab\x9a\xd1\x10\xa9\xfc\xc9\x37\x20\xfc\x16\x66\xc9\x66\x90\xa9\xca\x39\x0d\x31\x5d\x39\xa3\x61\xa6\x2c\x67\x35\xcc\xb4\xe5\xac\x06\x98\xba\x25\x9f\x01\xa6\x2f\xe7\x34\xc8\x14\xe6\x9c\x7c\xa7\x31\x2b\x57\xce\xb5\x1d\xe7\x0e\x2b\xc3\x3b\x01\xdf\x51\xec\x0b\xf8\xaa\xe2\xb0\x37\x38\x2a\x0a\x91\x47\xd3\x7d\xf0\xd6\x27\xc1\x64\xee\x11\x72\xbb\x0b\xfd\x15\x46\x81\x6f\x9a\x7d\x12\xb5\x07\xd8\xbe\x64\x2c\xd0\x3a\xa6\xdb\xe0\x36\x4f\x30\x8b\x92\x8b\xff\x3e\xbd\x9f\x3a\x27\x69\x4c\x6d\x61\xa2\x46\xe6\xf2\x34\xd9\x14\x1e\xdc\x76\xd1\x64\xbb\x5a\x4e\x31\x8e\x4b\xbc\xe0\xb6\xb7\xe8\x78\x90\xca\x7d\x7f\xc2\xcd\x09\xb9\xa7\x9c\xbd\xf3\xae\x4e\xbb\x7c\x1e\xee\xc7\x71\xab\xcc\x65\x13\xdf\x35\x8d\xe8\x71\x56\xc2\x21\xb9\xe1\x9a\xfc\x32\xae\x66\x11\x92\x34\x53\x6b\x96\xbe\x60\xe3\x7c\xb8\xc5\x69\x7e\x0b\x47\x32\x8a\x0d\x5c\x72\xf8\x7a\x07\x6d\xe0\x98\x1d\x37\x72\xac\x04\x36\x3f\x73\x5c\xad\x09\x26\xe7\xf7\xad\x80\x85\xfd\xb1\x79\x37\xd3\x21\xa1\x81\x85\x01\x4d\xa1\xf2\x1c\x14\xb6\x44\xe5\xba\x7e\xb7\xca\xa9\xcf\x41\xff\xc3\x9c\xb8\xd5\xaa\xf8\xdc\xe7\x16\x79\x58\xeb\xa0\xd4\x4a\x7b\xb5\xef\x84\x6e\xb2\x48\x91\xfe\xbc\x38\xfd\x60\x70\x8d\xce\xe7\x42\x6b\x5a\xe8\x90\x87\x18\xf6\x86\x07\x17\xbb\xcb\x4a\x88\x5d\x87\x5d\xfe\x7d\xbf\x4a\x9d\xb1\x2e\xa0\xe9\x77\x3f\xba\x06\x3e\xc1\x37\xe4\x09\x5c\xa4\xd4\x6d\xd9\x78\x59\x55\xbc\x27\x02\xbb\x2d\xc7\x4b\xf4\xe7\xe5\x4b\x21\xa8\x6a\x63\x31\xa6\xb7\x3f\x7c\xc2\x1e\x3c\x31\xbb\x82\xe5\x70\xf7\x2a\xa0\xbb\xdc\xee\x14\xac\x1b\x44\x35\x83\xc0\xab\x38\xcf\xa6\x91\x67\x18\x18\x60\x83\x7d\xfd\x16\x9a\xb7\xfa\x71\x5e\x4a\x73\x29\x33\xd9\x65\x96\xbd\x3f\x93\x79\x28\x4f\x8f\x1c\xc6\x59\x1a\x75\x36\xa6\x5e\xbf\x14\xdf\x77\xd0\x5d\x6e\x30\x67\x03\x5f\xa7\xee\x7f\xa4\x47\x7e\xd5\x87\x8c\xd5\x03\x07\x8a\xaf\x54\xf3\xbe\x42\xfd\x88\x88\xc6\x3d\x1e\x7c\xed\x71\x1f\xfa\xd8\x78\xb4\x1e\x80\xcd\xa8\x08\x85\xfb\x54\xa8\x57\x73\xf1\x6e\xc6\xbb\xa2\xf3\x68\xa0\xe8\x00\x70\x6e\x75\x9f\x4a\x6a\x40\x24\xc5\xd3\x4c\xb2\xed\x3f\xd8\x4b\x49\xe8\xa2\xf3\xfe\x52\x83\xa1\xfe\x85\x28\xfe\x94\xd6\x5e\xfb\x46\x94\x20\x9b\xea\x8d\x28\xf9\xbe\x57\x97\xc2\xf7\x91\x28\x93\x57\x9f\xe4\xc3\x61\xe6\xf8\xa9\xac\xb2\x98\x65\xbb\xe3\x88\xc7\x82\x4c\xeb\x98\x8a\x43\x90\xa6\xa0\xe9\x12\x12\xba\x63\xf3\x9a\x39\x50\xc7\x2c\xd2\x34\x8e\xb1\xa0\xc5\x6d\xa6\x3d\xfc\xbb\x89\x99\x0c\xf6\x53\x54\x7d\x2c\xdf\xd2\x33\x9f\x2d\x65\x15\x40\x1e\xa3\x83\x07\xbd\x8f\xf7\xf5\xb4\x4d\x81\x1d\x22\x10\xdb\xce\xb3\x43\x8f\xb5\xd6\x9b\xf0\xc9\x35\x70\xdb\x6a\xe6\x6f\xad\x59\xd9\xa5\xf9\xdc\xa3\xb9\x79\xea\x5a\xcb\x1b\x67\x67\x26\xb2\xfb\x98\x14\x7c\x1f\x53\x52\x75\x2c\x2a\x2b\x2d\xc6\xa0\x36\xaf\xa5\x94\x8f\x82\xda\xa4\x11\xcf\x58\x95\xa1\xd3\x88\xc6\x3b\x1b\xcb\xcb\x71\xbb\xed\x07\x19\x1c\x34\xf0\x0c\x89\xcf\x0d\x06\xb6\x51\xe2\x75\xaa\x48\x58\xe7\xdb\x74\x53\x15\x06\xcd\x30\x48\x68\x80\x3f\xb3\x60\xcd\x31\xfc\xea\x8f\x7a\x3a\xe1\xd7\x99\x20\x72\x0b\xbf\xc8\xf7\x8c\x50\xe6\xaa\xb4\x11\x98\xa4\x9c\x2a\x08\x93\x0d\x3e\x4e\x18\xb6\x8e\x37\xa3\xbc\x7a\xc9\xf9\x02\xe4\xf5\x45\xb8\x2f\xf8\x1e\x8b\xaa\xcb\x59\x6d\x1a\xb8\x1f\x56\xd3\x2f\xae\xdb\x90\x6b\xed\x96\x57\xb3\x7f\xcc\x4c\xd9\xd3\x53\xba\xeb\x26\x73\x36\x3f\xd2\x6c\xf1\x3d\x03\x75\x15\x99\xd7\x3b\xae\x1a\x08\x75\x80\xd1\x61\x00\x35\xd7\x01\xa6\xc1\x7d\x5e\x36\xcc\x13\x6d\x98\xe3\xe2\x79\x3f\x9d\x9f\x98\x39\x5e\xd6\xfe\xf2\x7a\x44\xa3\x9e\x51\x87\x4c\x05\xde\x6a\x0f\x0e\x0c\x01\x68\x86\x01\x7c\xc7\x85\x64\x4a\x18\xd4\x41\x65\xc7\x92\xce\x11\x98\xdd\xb3\x47\x9b\xb5\xa0\xec\xbe\xe7\xa7\x21\x60\x40\x46\xe3\x74\x4f\xd7\x2b\xd5\x9b\x9f\x3e\x80\x8d\x09\xf3\x38\x60\xed\x26\xde\xc0\x77\xa7\xbb\x26\xe6\x15\x8c\x9c\xc1\x42\x5f\x86\x3f\xe6\xe3\x47\xa7\xa6\xcf\x51\x9a\x7b\x92\x68\xdf\x71\xe9\x1b\xd9\x65\x3d\x71\xdc\x83\x94\x27\xf2\xc5\xc9\x0f\x31\xe8\x85\x95\x86\x06\x0c\x13\x25\xbb\x99\x2e\x95\x65\x0c\x0c\x0f\xbe\xe2\xf8\xa4\xb0\x97\xb8\xfd\x5d\x29\x8c\x8a\xf0\x3d\xc2\xc9\x1e\xfd\x4b\xf6\x08\xe7\x59\x7c\x5d\x44\xf3\x09\xff\xa5\x1b\x04\xda\xcd\x0e\xc1\x39\x99\xd7\x04\xc2\xb9\xe0\x0f\x31\x25\x9f\xf1\x5e\x4f\x23\x11\x49\x33\x1c\xed\xc6\xb3\xba\x6c\x02\xfd\x9c\xd3\xe0\x3b\xfb\xe7\x2f\x0b\xec\x85\x8f\xd2\x52\x63\x0d\x7e\xab\xb0\xda\xbe\x1b\x96\x19\x4d\x0f\x12\x48\x4f\x85\xad\x21\x58\x0d\x61\xe7\x4c\x33\xb6\x56\xf0\xb2\x26\x20\xd5\xb8\xd8\xbb\x0f\xed\x76\x71\x37\x1f\x67\x27\xf1\xad\x23\xee\x96\x07\x91\xfa\x8f\xa1\x96\xb8\x5b\x75\x5c\xa9\x81\x86\x25\x89\x1b\x22\x9e\x16\xc4\xfb\xe2\xef\x91\x7f\xa9\xcc\xec\x46\xd2\x5c\xbd\xc7\xfc\x7e\xa9\x6c\x1e\x88\x70\x82\x6c\xdc\x14\x3f\x3d\xf6\xa4\xc3\xc8\x29\x7e\x60\x6c\x0c\x8f\xc1\xf9\x2e\xcf\xdc\x1f\x34\x62\x2f\xe2\x17\xc6\x74\x67\x79\x85\x34\x53\x3b\x19\xf0\xfc\xb1\x20\xfd\xff\xfc\xca\x95\xa7\x29\x47\xfd\xfd\x2a\xc3\xbc\x92\x8b\x2d\x0d\x12\x33\x33\x75\x0c\x9f\xcd\x53\x4a\xb3\x08\x42\x00\xff\x6b\x74\x59\x00\xbc\x8b\x00\x5f\x25\xd0\x84\x29\x72\xd9\x1e\x35\x54\x51\x06\x08\x9d\x70\xa5\xfe\x5d\xc9\x3a\x5c\xe1\x7f\x5d\x1d\x3d\x1c\xfd\x17\x00\x00\xff\xff\x51\x0d\x4c\xf7\x6e\x7a\x00\x00")

func etcSchemaGohanJsonBytes() ([]byte, error) {
	return bindataRead(
		_etcSchemaGohanJson,
		"etc/schema/gohan.json",
	)
}

func etcSchemaGohanJson() (*asset, error) {
	bytes, err := etcSchemaGohanJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etc/schema/gohan.json", size: 31342, mode: os.FileMode(420), modTime: time.Unix(1454114428, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcExtensionsDonburiJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x1c\x5d\x6f\xdb\xc8\xf1\x39\xfe\x15\x7b\x04\x7a\x91\x10\x99\x4e\x0e\xe8\x8b\xd5\x3c\xb8\x89\xdb\xba\x0d\xec\x22\x56\x7a\x3d\x08\x06\x41\x51\x2b\x89\x17\x9a\x54\xf9\x11\xd9\xf0\xe9\xbf\x77\x66\xbf\xb8\xbb\x5c\x52\xe4\x39\x4e\xfc\x70\xc6\xc5\x26\xf7\x63\xbe\x77\x66\x76\x76\x79\x27\x27\xe4\x5d\xb6\xbd\xcf\xe3\xf5\xa6\x24\xa3\x77\x63\xf2\xd3\xeb\x37\x7f\x26\x97\xb3\x19\xb9\x48\xd3\xec\x4b\x58\xc6\x59\x0a\x8f\x45\x19\x97\x55\x49\x27\xf0\x18\xf9\x47\x27\x27\xf0\x1f\xf9\x10\x47\x34\x2d\xe8\x92\x54\xe9\x92\xe6\xa4\xdc\x50\x72\xb6\x0d\x23\xf8\x23\x7a\x26\xe4\x3f\x34\x2f\x10\xc0\x4f\xfe\x6b\x32\xc2\x01\x9e\xe8\xf2\xc6\x53\x04\x71\x9f\x55\xe4\x36\xbc\x27\x69\x56\x92\xaa\xa0\x00\x23\x2e\xc8\x2a\x4e\x28\xa1\x77\x11\xdd\x96\x24\x4e\x49\x94\xdd\x6e\x93\x38\x4c\x23\x4a\x76\x71\xb9\x61\x78\x04\x14\xa4\x84\xfc\x22\x60\x64\x8b\x32\x84\xe1\x21\x4c\xd8\xc2\xdb\x4a\x1f\x48\xc2\x52\x10\x0d\x3f\x9b\xb2\xdc\x9e\x9e\x9c\xec\x76\x3b\x3f\x64\xf4\xfa\x59\xbe\x3e\x49\xf8\xc8\xe2\xe4\xc3\xc5\xbb\xf3\xcb\xeb\xf3\x63\xa0\x59\xcc\xf9\x94\x26\xb4\x28\x48\x4e\xff\x57\xc5\x39\xf0\xbb\xb8\x27\xe1\x16\x68\x8a\xc2\x05\x50\x9a\x84\x3b\x92\xe5\x24\x5c\xe7\x14\xfa\xca\x0c\x69\xde\xe5\x71\x19\xa7\xeb\x09\x29\xb2\x55\xb9\x0b\x73\x8a\x60\x96\x71\x51\xe6\xf1\x02\xc4\xa8\x8b\x4c\x52\x08\x8c\xeb\x03\x40\x68\x61\x4a\xbc\xb3\x6b\x72\x71\xed\x91\xbf\x9e\x5d\x5f\x5c\x4f\x10\xc8\xcf\x17\xb3\x7f\x5c\x7d\x9a\x91\x9f\xcf\x3e\x7e\x3c\xbb\x9c\x5d\x9c\x5f\x93\xab\x8f\xe4\xdd\xd5\xe5\xfb\x8b\xd9\xc5\xd5\x25\xbc\xfd\x8d\x9c\x5d\xfe\x42\xfe\x75\x71\xf9\x7e\x42\x28\x08\x0c\xf0\xd0\xbb\x6d\x8e\x1c\x64\x39\x82\x88\x51\x9e\x74\xc9\x84\x77\x4d\xa9\x41\xc5\x2a\xe3\x54\x15\x5b\x1a\xc5\xab\x38\x02\xee\xd2\x75\x15\xae\x29\x59\x67\x5f\x68\x9e\x02\x53\x64\x4b\xf3\xdb\xb8\x40\xbd\x16\x40\xe3\x12\xc1\x24\xf1\x6d\x5c\x32\x5b\x29\x9a\xac\xf9\x47\x47\x5f\xc2\x9c\x2c\xb3\x74\x51\xe5\x31\x79\x4b\x1e\x8e\x08\xd9\x26\xd5\x3a\x4e\x4f\xc9\xc3\x7e\x02\x6f\x69\x16\x94\x14\xc8\x0a\x4b\x1a\x7c\xa6\xf7\xbb\x2c\x5f\x16\xa7\x64\xfe\x92\x7e\x09\x93\x97\x13\xf2\x72\x91\x64\xd1\x67\x7c\x58\xd2\x55\x9c\xd2\x97\x38\xa7\xe5\xe7\x25\x70\x9a\x55\x79\x44\x8b\xce\x51\x29\x2d\xa3\x2c\x5d\x05\x51\x92\x15\x14\x21\xcb\x06\x7a\x47\xa3\xce\x99\x45\xb1\xa9\x67\xe1\x0b\x9b\x71\x83\x53\x72\xba\x06\x15\xd2\x3c\x90\xdc\xad\xaa\x34\x42\xa9\x8c\xd2\xf0\x16\x96\x03\x6f\x1e\x3f\x30\xe8\x68\xeb\x3e\x6f\x99\x63\xf7\x0d\x48\x86\xbf\x42\x37\x93\x0a\x1a\xd9\xbd\x12\x8c\x06\x0d\x08\x2d\xe9\x5d\x39\x21\xcb\xb0\x0c\x05\x38\x14\x71\x41\x93\x15\x40\x41\xc8\x53\xd6\x18\xaf\x46\x81\x1f\x17\xd7\x60\x56\xe9\x7a\xc4\x46\x8b\xe1\x48\x6c\x59\xe5\x29\xe8\x75\x13\xa6\x0a\x09\x1b\x33\x21\x02\xc1\x98\x43\xd9\xeb\xb0\xce\xf2\x3c\xbc\xb7\x40\x21\x6e\x10\x7b\x95\x94\x80\x7d\x7e\x33\x15\xcd\x81\x4f\x61\x79\x09\x90\x8a\xf8\xa2\x5a\x04\x1a\xdd\x9c\x14\x9c\xeb\x6f\xab\x62\x33\x42\x1e\x7c\x93\xf3\x9a\x5f\x35\x77\x2c\x71\xec\xd5\x93\xe0\x87\xc3\x72\x10\x7e\xb5\xf8\x95\x46\x65\x17\xe5\x0f\xfb\x7e\x94\x4f\x08\xd8\x68\x83\xfc\x39\x34\xa2\x0e\xfb\x31\x30\x80\x7e\xd1\x81\xd3\x8e\x78\x5b\x5e\x81\xc6\xc2\xe2\xb3\x66\x12\xf4\x0b\x4d\xcb\xa0\xbc\xdf\x52\xa5\xbd\x09\xc1\x31\x07\xcc\xe3\x07\x14\xcd\xa7\x94\xaf\xab\xe5\x08\x67\xf8\xbb\x0d\x4d\x6b\x11\xa1\xd3\x95\x0c\x68\x4c\x97\xf9\xfd\x83\xb6\x4a\x00\x14\x2e\x56\x0d\x00\x79\xfb\x96\xac\xc2\xa4\xa0\xe3\x07\x63\x35\xb9\x91\xce\x3d\x0a\x43\xbd\x9b\xb1\x35\xba\x81\x88\xff\x30\x21\x4b\x31\x14\xed\xdc\x2b\xb0\x16\x84\x7d\x14\x96\xa0\x60\x9a\xe7\x0d\x7c\x44\x42\xf0\xa1\x17\xdc\xe1\x5b\x02\x7f\xa7\xf6\xfc\xa3\xf6\x37\xae\x2f\x7d\x46\xdd\xdf\x82\xf7\x00\x46\x1b\xa2\x84\xb7\xd7\xac\xa4\xb6\xe4\xa9\xf6\x0e\xc2\x0b\xa2\xac\x4a\xd1\xbc\x99\x6a\x58\x93\xe1\x1c\x6a\x3d\x68\xc3\x0d\x37\xa1\xc1\x78\xa3\x1b\x26\x04\x8b\x91\x42\x03\x9d\xaf\xa7\xe2\xf1\x2f\xfa\x2c\xd1\xf8\xea\x95\x02\x69\x68\x54\x2c\xb5\xc0\x87\xe5\x53\x30\x53\x18\x6b\x4b\xce\x5c\x68\x9c\x2b\xee\x24\xe5\x4a\x13\x1e\x14\x17\xdf\xd4\xb4\x47\x93\x39\xe1\x7b\x2d\x75\x77\xe9\x8a\x63\x43\xf3\x0f\x62\x58\xc9\x85\x14\x61\xdd\x32\x75\x8d\x5d\xc6\x51\x69\x0c\xc5\x86\x4e\xd2\x6a\x80\x63\xf2\xe3\x8f\xc4\xd1\x89\x20\x9a\xa4\x0b\xb7\xa5\xd6\x02\x67\x51\x70\x0a\xc1\xdf\xb5\x28\x40\x4e\x62\x65\xa0\xc4\xc6\x06\xe7\xb8\x54\x1a\x6b\x55\x8b\x1f\x1a\x9d\x8d\x55\xd3\xe2\x25\xf4\x6e\x25\x45\xe6\x28\x34\x60\x43\x16\x97\x11\x85\xba\x08\x72\xc4\x23\xf9\x23\x4c\xae\x9e\xad\x19\x1c\xbe\x77\xb8\x04\xec\x06\x98\xf8\xc7\x86\x2a\x7c\x2c\x38\xf8\xaf\xae\x19\x9d\x23\x11\x25\x15\x9a\xa6\xf4\xac\x16\xa7\xfc\x74\x8d\xba\x8c\xab\x9f\x42\x85\xa9\xd7\xfa\x64\xa0\x7e\x87\x3a\x45\x6c\xee\x20\xc7\x11\xa4\xe5\x8f\xae\x4e\x9c\xac\x69\x13\x08\xab\x40\xc6\xa8\xae\xc0\x76\x25\xfc\xc7\x52\xec\x03\x8c\x3a\x55\xe3\x27\x84\x01\x38\x65\xbf\x1b\x78\xbf\x91\xc2\x55\x5e\xa1\x10\x0d\xd1\xb8\x16\x7a\xb4\x51\x2d\x81\x5f\xe6\xaf\xb6\xf8\x05\xc9\x73\x63\x10\xd2\x64\xd3\xa3\x02\x93\x2b\xcc\x75\x06\xb9\x56\x8a\x8a\xa8\xa2\x26\x3d\x8d\x84\xa0\x6f\x32\x20\xa1\xe9\xe2\x69\xcb\x03\x0e\x44\x64\x4d\xac\xb6\xd7\x2c\x37\x79\xb6\x33\xc7\x9b\xf1\xba\x95\xd9\x30\xd9\x85\xf7\x86\x2b\xb3\x58\xed\xcd\xa8\x80\xf4\xa8\xb4\xc3\x95\x64\x34\x72\x54\xb6\x4f\xa9\x0d\x5e\x4b\x48\x0f\x9b\x3e\x5b\x55\x07\xd2\x53\xc6\xb1\x6b\x77\xe8\xc7\x20\xbb\xbb\xab\x15\xcb\x10\x30\xd5\x3c\x7e\xa3\x25\xf5\x00\xf8\x50\x26\xce\xb1\x3b\xf2\x6c\xb1\x74\x9d\x74\xd7\x93\xf4\x34\xbc\xe8\x93\x87\x17\x5d\x9c\xda\xc9\x9b\x70\x69\x6c\x9e\xe6\xce\xb4\x7c\xde\x15\xff\xb1\xbb\x63\x23\x20\x78\x15\x7f\x5b\x55\x69\xec\x33\x97\x9d\x0a\xc2\x46\x86\x0f\xa3\xfa\xdc\xdb\x66\x45\x19\x44\x39\x05\x39\x7b\xae\x1d\x34\x1f\x50\x6d\x97\xed\x03\x72\x1a\x2c\x69\x42\xdb\xfa\xd3\xac\xc4\xaa\x04\xab\x34\x78\x37\xba\xa8\x38\x19\x13\x97\x22\x94\xc4\xf8\x56\x57\x6d\xd2\xe1\x65\x99\xd0\xdc\x90\x98\xbd\xc5\x7e\xdc\x5a\x44\xf9\xf9\x5c\xf9\xd3\xaf\xb3\x16\x15\x1c\x25\x7a\x86\x33\x88\x81\x8e\x3c\x4c\x8b\x90\x51\xcf\x62\xb2\xec\x23\xaf\x88\x67\xf5\x7b\x7d\x05\x62\xcd\xfb\x2a\xf2\x69\xc0\x34\xc5\xb5\x5c\x04\x4f\x21\xb1\xbd\x58\xb5\xfb\xa3\x23\x51\x88\xf2\xad\x6a\xcd\xc8\xe3\xce\xd8\x73\x1a\x51\xc3\x09\x30\x4a\x50\x0b\x38\x38\xc0\x02\x0e\x90\xc0\xba\x7c\x7c\x99\xca\x6e\xc6\x8c\xea\x62\x6f\xac\xaf\x8d\x0a\x05\x6e\x00\x19\x1c\x53\x4a\x77\x81\xe8\x05\x7c\xb0\x2e\xee\x4a\x9a\x2e\x6d\x8f\xc7\xc7\xcb\x9a\x82\x24\xc2\x69\xc7\x1a\x40\xe9\xc3\x18\xe9\x28\x4b\xfc\xd7\x25\xc9\x45\xb5\x1e\x20\x48\x68\x2a\xb2\x84\xfa\x49\xb6\x1e\x49\x1f\xdb\x8d\xa1\x48\x28\xdd\x0e\xc0\xc0\x4d\x9d\xcd\xea\x89\x81\xd5\x1c\x07\x60\xe8\x96\x65\x23\x84\x74\x23\x97\xd5\xc8\x6c\x4b\xd3\xde\x34\xbc\x30\x4a\x7a\x3a\x08\xce\xb2\xbf\x01\xff\x2b\x86\xfb\x55\x41\x73\x34\xb3\x71\x4f\x4a\xb0\xca\x39\x64\x69\x30\x93\x8c\xb2\xdb\x5b\x70\x2a\x60\x8e\x12\x7e\x67\x38\xf6\xc5\x78\x20\xe8\x85\xb5\xff\x78\xe1\x66\x0e\xa9\xe2\x85\x27\x09\x20\x4d\x29\x23\x70\x3c\x91\xd8\xc7\x47\x2f\xf6\x3d\x99\x64\x75\xdd\xfe\xf2\x6e\xd0\x68\x12\xc7\xa0\x69\xd4\x8d\x7b\x50\x82\x05\xe5\x47\x28\x5d\x4e\x7f\x84\xc2\x65\x49\xfb\xf9\x28\x5b\x52\xf4\x15\x14\xad\x8a\xf7\x8f\x56\xb2\x82\x34\x50\xc1\xea\x78\x62\x80\x7c\x59\x71\x55\xc5\x73\x48\x76\xf5\x44\x89\xfc\xf6\x1b\x69\xed\xb5\x03\xbf\xa6\xae\x1c\x26\xe5\x78\x7c\x27\x63\x53\x81\x47\x60\xa3\xb1\xc8\x0e\x79\xa7\x2f\x1e\x46\x22\x8c\xf6\x74\x72\x72\x3a\x0f\xba\xf5\x0e\x69\xa0\x8f\x64\x01\xbb\x53\x9a\x42\xf3\x03\x64\x69\x58\x16\xb3\xaa\x3a\x68\x4b\xcb\x0c\xf3\x75\x71\xd0\x45\xa3\xda\x07\xe0\x75\xd4\x53\x04\x29\x9a\xfd\xf4\x61\x19\x74\x37\xc4\x76\x44\x7e\x2c\x2a\x21\x76\x65\xa4\x2e\x8a\xc8\x3d\xbe\xa8\x36\xb0\x7e\x46\xce\x41\x41\x70\x36\x86\xab\x40\xa0\xe8\x06\x9e\xc0\xeb\x00\xd0\x2c\xe3\x32\x12\x61\x99\x25\x6a\xad\x53\x31\x70\x15\x27\x80\x4a\xd9\xbf\x78\x85\xe5\xc4\x6b\x4c\xfc\xdd\x87\x24\x2a\x04\x94\x71\xbd\x50\x54\xcb\xd4\xb6\x27\x48\x5e\x91\xe0\x91\x91\xdf\x8a\xd5\x15\x6d\xe8\x2d\x9e\x26\x31\xa8\x07\x85\xba\x82\x18\xb2\x79\x2a\xc6\xe5\x06\x59\x11\xcd\xb0\x75\x51\xcd\xdf\xe2\xe5\xc4\x16\xc1\x78\x6a\x29\xb4\xa7\xfb\x7b\x0a\xd6\x1c\x6e\xb2\xde\x91\xda\x9e\xd0\xe5\x36\xb5\xfd\xab\x58\x15\x8d\x5a\xbd\x14\xc4\xd8\xac\x36\xf8\xcc\x3c\xb8\x3c\xab\x2a\x5e\x8e\x8c\xca\x02\x1f\xb2\xcd\x21\x32\xe7\x65\x4c\x0b\x5f\x33\x26\x6e\x45\xd6\x49\xa4\x52\x0c\xa7\xe6\xb0\x66\x6a\xd8\xad\xdb\x7b\x21\xc9\x06\xea\x16\xa9\xf1\x6d\x7a\x5f\xa9\x89\x4d\xbd\xe9\xd8\x14\x17\xbc\x77\x30\x17\xfb\x03\xa1\xaf\x0f\x71\x75\x9c\x6c\xa1\x8d\xf7\x76\xd1\x26\x8c\x7c\xaf\x0c\xfd\x80\x89\xcb\x02\xc7\xd3\xac\x5d\xb9\x80\x7e\xc7\xf2\x35\x0c\x70\x42\x3c\x6f\x2c\xf4\x6f\x1f\x03\x72\x0c\x63\x43\x62\x42\x02\x7a\x38\xd1\x00\x76\x45\x16\x09\xcf\x19\x5a\x7e\x8f\xb1\x28\x02\x0f\x6f\x43\x79\x29\xe9\xeb\x6b\x62\xb8\x15\x29\xc7\x71\x88\x68\xc4\x98\x87\xf1\x90\xdc\x42\x94\xe7\x18\x26\x3d\x50\xf1\x96\xa9\xdb\x93\x49\x3c\x81\xa0\x64\xae\xe6\xdf\x34\x2e\x89\x34\xc6\x7a\x6b\x9a\xd2\x3c\x8e\xbc\x9b\x43\x19\x9c\x55\x5d\xed\xc0\xfa\x98\xfd\xb2\x02\x3b\x34\x6e\xf6\x11\x9d\xd3\x6d\x0a\x63\xc9\x3e\x53\xdd\x4c\xc2\xaa\xdc\x04\xac\x71\xda\xc5\xb5\xbf\xa6\x65\x80\xf7\xc9\x28\x26\x0b\x38\x7c\x52\xd3\x20\x3c\x4e\x4f\x86\x07\x26\x48\x7d\xf9\x7d\x0a\xc6\x0e\x32\xb5\x81\x60\x37\x88\x17\x41\x05\xce\xab\x4d\xb3\x28\xc3\xe8\xf3\x21\xc3\x64\x84\xf0\x5d\xab\xc9\x81\xb8\x9f\x16\x6e\xe3\xa0\xca\x13\xad\x18\x6e\xea\xa7\x25\x3d\xd0\x13\x03\x41\xdc\xbb\xab\xcb\xd9\xc7\xb3\x8b\x0f\xc1\xa7\x8f\x1f\xc8\x2b\x4d\xf4\xaf\x88\x57\x78\x8e\x25\xd2\x35\xe1\xc4\x83\xdf\xdc\x00\x59\xbd\x5e\xe4\x36\x42\xe2\x6e\x6a\x85\x1a\x0e\xdd\x15\xe3\xc1\x85\xc1\xd1\x8e\x59\x45\x8b\xb6\x4c\x71\x7f\x0f\xa0\x8c\xac\x65\x9b\xe1\xe5\x41\x19\x91\xf0\x82\xe5\xc8\xfb\xf7\xd5\xf5\x0c\x94\x29\x0e\x61\x98\x34\x6b\xaa\xc6\x6d\x37\xec\x1e\xbc\xff\x1e\x9f\x81\xad\x1d\xcf\x90\x6a\xef\x54\x50\xdf\x32\x9a\x78\xef\x50\xb1\x69\x79\x3c\x03\x35\xc3\x68\x4f\x5c\xcc\x44\x19\x9c\xfc\x5a\xd4\xa5\x6e\xfb\x67\xdf\x06\x52\x30\x3c\x6e\x26\x65\x0f\x52\x4b\xec\x18\x1e\xac\xac\xac\x8a\x00\x6b\xd6\x78\x71\x0f\x36\x65\xf4\x22\x2d\x47\x52\x18\x7a\xbf\x4a\xc9\xd8\xc4\x45\xb6\xbc\xe7\x47\xa7\x7c\x20\xbe\x2b\x17\xed\x02\xfe\x96\xfc\xf4\xfa\xb5\x75\xec\x33\xf7\x50\x07\x1e\x2a\xe3\x9f\xd7\x57\x97\x3e\xc3\x3f\x32\x60\xba\xbc\xaf\x75\xda\x23\x72\xbc\x3e\xd6\x83\x8f\xdf\xdc\x82\x3e\xb5\x1b\x10\x5b\x84\x7f\x18\xd1\x33\x30\x22\x91\x8b\xf7\x33\xa2\x1e\x87\xa0\x0e\x4b\x10\x54\x7b\xef\xcf\x3f\x9c\xcf\xce\xd5\xd1\x60\x0f\xcb\x70\x5b\xc2\x7e\x42\xd2\x2a\x49\xc6\x8e\x9d\xd7\x93\xeb\xc7\x2d\x44\x2d\x6c\x6a\x12\x74\x66\x06\x8f\x92\xe0\xdf\xcf\x67\xdf\x4d\x7c\x2e\xa9\xe9\x36\xdb\xe8\x64\x3b\x38\x30\x5b\xef\x49\xec\x56\x64\xb0\x87\xce\xee\x0f\xde\x52\xe8\xbd\x67\xf8\x5a\xa9\x17\xbb\x4c\x0f\x02\x50\x00\xea\x6d\x18\xe4\x34\x42\x9e\x87\x4b\x1b\x4a\x2e\xcc\x0e\xcc\x44\xa2\x2d\x7d\x90\x52\xed\x4a\x8b\x5d\x39\x11\xee\xcc\x71\x28\x92\x62\xe3\xd7\x35\xd5\x59\x8b\xf0\x8c\x5b\xdb\x22\x76\x09\x41\x34\xae\x6e\x8b\x42\x6e\x92\x64\x3b\x31\xbd\xf5\x56\xa9\x06\x49\x6e\x51\xd5\xb3\x84\x5a\x1f\xbc\x03\x89\x06\x6a\xc4\x6d\x1c\x4b\x9b\x57\x4a\xf7\x0e\x59\x9b\x61\xb7\x25\xd8\x6a\x48\xa4\xd8\xb1\x8e\xde\x7d\x16\xe0\x54\xad\xe9\xa0\xdd\x6e\xd9\x50\x6c\xcd\x01\x5b\x28\xde\x97\x38\x2f\xab\x30\x39\x4e\x69\xb9\xcb\xf2\xcf\xc7\x45\xb5\x80\x47\xaf\x5e\x39\xe4\x31\x4b\xa7\x87\xb5\x0b\xc4\xfa\xa2\xa9\x9b\x7a\x2c\x0a\x1b\x8e\xe6\x23\x99\x84\x1c\x9b\x16\x9b\x69\x48\x44\x34\x32\xc6\xca\xd4\xad\xdb\xba\x36\x06\xdf\xf5\xcd\x85\x66\xf0\xe2\x06\x21\x4c\x0b\x0a\x76\x25\xb3\x10\xd6\xe7\x8b\xd6\x28\x5e\xe6\x7e\x01\xe9\x48\x39\x82\xe4\x5f\x9e\xc2\xd4\x03\xc4\x76\x85\x51\xb3\x0d\xc0\x14\x56\xf1\xdd\xa9\x05\x72\xfe\xfa\x66\x62\x8f\x09\x12\x9a\x9e\xd6\x11\xcd\x9a\xf0\x46\x5c\x98\xdf\xdb\xa2\x03\x74\x4e\x16\xe7\x0d\x79\xdd\x38\xd4\xb7\x0d\x6f\x61\xda\xaa\xa8\xa1\xcc\xbd\x46\xa7\x9c\xd9\xf0\x23\x8d\x91\xb5\x5c\x5d\x18\xe6\x6a\x4d\xea\xab\xd3\x0b\xcb\x32\x07\xd3\x7d\xf0\xd8\x58\xce\x77\x01\x0d\xf3\x1b\x23\x9b\xf3\xca\x0c\x1b\xf1\x4e\x48\x08\x51\xe3\x78\x99\xdd\x86\x31\x1e\x72\xa8\x16\x30\x30\xbc\xb4\xaa\x37\xc9\x05\x82\xa0\xc5\xf5\x24\x26\x45\xe9\x46\xba\x98\xd6\x24\xab\x5a\xa7\x96\x12\x74\x92\x5d\xe3\x41\xd1\x73\xce\x20\xfc\x35\xf8\xab\x65\xda\x27\x30\xe8\x53\xf9\x7d\x63\xc3\xfb\x77\x40\xb2\x3d\xb5\xf0\xc6\x3a\x40\xf3\x83\x1a\x68\xd1\xbc\x27\x2e\x28\xde\x28\xed\x1f\x8b\xe9\xe4\x07\x73\x4d\xb0\x36\x87\xcf\xad\x05\xae\x9e\xdc\x1f\xf2\xd8\x57\x81\x39\xdc\x56\xe7\xaf\xbb\xff\xfd\x21\x29\x38\x9c\xb1\xbc\x9c\x63\xa9\x6f\x7e\xf3\x6d\x85\x64\x93\xa0\x6e\x92\x23\xf8\x69\x43\x78\x35\xcf\x43\xec\x8c\x99\xa5\x89\xc7\xb6\x62\xf3\xe2\xa2\x15\x0c\x1d\x7e\x57\x78\xe4\xda\xfb\xaa\xe7\xf6\xb3\x0e\x48\x07\x8e\x4e\x4e\x66\x57\xef\xaf\x46\x29\xec\xe3\xc6\x64\x86\x5f\x96\xc2\x7f\xf4\x0e\xa2\x42\x7c\x0b\x3a\x0b\x13\xfc\x02\xf2\x67\x0a\xc0\xf8\x77\x9b\xf8\x01\x6a\x56\x95\xdb\xaa\xc4\xaf\x2f\xf1\x73\x94\x32\x4b\xa9\x0c\x4c\xbc\x5a\xa4\x57\x9a\x5a\x2b\x45\xf0\xd6\x9e\xb4\x2f\xc2\x82\xe2\x78\x98\x8d\xbf\x65\x4d\xc7\x59\x16\xea\x53\x5e\x92\xf0\x1c\x09\xaf\x42\x35\xa4\x70\x64\x13\x3f\xa8\x7a\xd4\xbb\x10\x64\x62\x79\xbe\x1b\x79\xc6\xf6\xf3\xda\xc5\xbf\xf9\xb6\xa5\xa0\x16\x7b\x78\x26\xf5\x20\xc7\x52\x7b\xb6\xb6\xf4\x74\x45\xa1\xe7\x53\xf3\x69\xb5\x96\xa7\x2b\xfc\x74\x98\xc0\x77\xa8\xfe\xf4\x2f\xee\x3c\x91\xa8\xbe\x67\x85\xa7\xbf\xad\x3a\x6b\x3d\x4f\x65\xad\x28\x7f\x9a\x2e\xb7\x59\x9c\x96\xae\xaf\xe8\x65\x5f\x00\xff\x68\xbe\x0a\x23\xd8\xc4\x6a\x1f\x0e\xb0\x6a\x03\x0b\xd8\x9e\x88\xcb\x22\x61\x93\x5b\x56\x70\x0a\x61\x92\xad\xf5\xdb\xf1\xbc\xa5\x4e\xa8\x57\xb2\xc9\x9f\xf1\x4c\xf1\x2d\x31\x3e\x4d\xa8\x61\x8a\x61\xe7\x82\x26\xe3\x93\x06\xd1\x66\x64\x75\x98\x82\x8a\x76\xff\x42\xd2\xcf\x10\x34\xd9\xb2\xbf\x7f\xe3\x5c\xa9\xe9\x9f\x3e\x7e\xf0\x73\xba\x4d\x60\xe4\xc8\xfb\xd3\xa8\xbe\xe3\x83\xf7\xcd\x54\x51\x8b\xb5\x9a\x1f\x8e\xba\xf2\x64\xae\x95\xb1\xae\x16\x91\xa5\x30\x9d\xb0\x43\xba\xef\x59\x7b\x13\xc7\x84\x1a\x1c\xed\x9e\x97\xc9\xec\xb0\x62\x05\x83\x6b\x7e\x10\x50\x37\x4d\x2d\x8b\x52\xe5\x07\xa9\x84\xda\x2a\xbd\x6d\xb5\x80\x80\x83\xfb\xcb\x2c\x07\xc2\x61\x6f\xce\xbf\x7b\x19\x37\x2a\x1e\x52\xad\x35\x9e\x53\xed\x59\x2e\xf5\xfa\x7f\x24\x21\xef\x70\xf1\x77\xde\xbf\x1f\x56\x3b\xb4\x1c\x91\xab\x86\xd8\x95\x49\x4e\xcd\x24\x56\x42\x72\xd7\x17\x35\x89\xca\x4c\xf6\x8f\x92\xe3\xa3\x4a\x8e\x1d\x49\xdd\x93\xd5\x1d\xdb\xe3\x5d\x7b\xf1\x11\x44\xfc\xff\x00\x00\x00\xff\xff\x40\x08\x1d\x7c\xa1\x48\x00\x00")

func etcExtensionsDonburiJsBytes() ([]byte, error) {
	return bindataRead(
		_etcExtensionsDonburiJs,
		"etc/extensions/donburi.js",
	)
}

func etcExtensionsDonburiJs() (*asset, error) {
	bytes, err := etcExtensionsDonburiJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etc/extensions/donburi.js", size: 18593, mode: os.FileMode(420), modTime: time.Unix(1454435312, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcExtensionsDonburiYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\x55\xc8\x4c\xb1\x52\x48\xc9\xcf\x4b\x2a\x2d\xca\xe4\x52\x50\x28\x48\x2c\xc9\xb0\x52\xd0\xd3\x02\x32\x4b\x8b\x72\xac\x14\x52\x73\x93\x52\x53\xac\xf4\xf5\x53\x4b\x92\xf5\x53\x2b\x4a\x52\xf3\x8a\x33\xf3\xf3\x8a\xf5\xa1\x1a\xf4\xb2\x8a\x01\x01\x00\x00\xff\xff\x79\x6b\xd1\x13\x41\x00\x00\x00")

func etcExtensionsDonburiYamlBytes() ([]byte, error) {
	return bindataRead(
		_etcExtensionsDonburiYaml,
		"etc/extensions/donburi.yaml",
	)
}

func etcExtensionsDonburiYaml() (*asset, error) {
	bytes, err := etcExtensionsDonburiYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etc/extensions/donburi.yaml", size: 65, mode: os.FileMode(420), modTime: time.Unix(1454435312, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcExtensionsGohan_extensionYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\xad\x28\x49\xcd\x2b\xce\xcc\xcf\x2b\xb6\xe2\x52\x00\x02\x5d\x85\xcc\x14\x2b\x85\xe2\xe4\x8c\xd4\xdc\x44\xb0\x80\x82\x42\x72\x7e\x4a\x6a\x7c\x49\x65\x41\xaa\x95\x42\x7a\x3e\x92\x98\x95\x42\x46\x62\x5e\x4a\x4e\x6a\x3c\x8a\xea\x82\xc4\x92\x0c\x2b\x05\xfd\xf4\x7c\xa0\xa4\x7e\x99\x81\x9e\xa1\x3e\x58\x5a\x4f\x0b\x10\x00\x00\xff\xff\xc8\x52\x27\xbc\x6a\x00\x00\x00")

func etcExtensionsGohan_extensionYamlBytes() ([]byte, error) {
	return bindataRead(
		_etcExtensionsGohan_extensionYaml,
		"etc/extensions/gohan_extension.yaml",
	)
}

func etcExtensionsGohan_extensionYaml() (*asset, error) {
	bytes, err := etcExtensionsGohan_extensionYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etc/extensions/gohan_extension.yaml", size: 106, mode: os.FileMode(420), modTime: time.Unix(1454103480, 0)}
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
	"etc/schema/core.json":                etcSchemaCoreJson,
	"etc/schema/gohan.json":               etcSchemaGohanJson,
	"etc/extensions/donburi.js":           etcExtensionsDonburiJs,
	"etc/extensions/donburi.yaml":         etcExtensionsDonburiYaml,
	"etc/extensions/gohan_extension.yaml": etcExtensionsGohan_extensionYaml,
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
	"etc": &bintree{nil, map[string]*bintree{
		"extensions": &bintree{nil, map[string]*bintree{
			"donburi.js":           &bintree{etcExtensionsDonburiJs, map[string]*bintree{}},
			"donburi.yaml":         &bintree{etcExtensionsDonburiYaml, map[string]*bintree{}},
			"gohan_extension.yaml": &bintree{etcExtensionsGohan_extensionYaml, map[string]*bintree{}},
		}},
		"schema": &bintree{nil, map[string]*bintree{
			"core.json":  &bintree{etcSchemaCoreJson, map[string]*bintree{}},
			"gohan.json": &bintree{etcSchemaGohanJson, map[string]*bintree{}},
		}},
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