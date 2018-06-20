// Code generated by go-bindata.
// sources:
// override/templates/17_upsert.go.tpl
// override/templates/singleton/mysql_upsert.go.tpl
// override/templates_test/singleton/mysql_main_test.go.tpl
// override/templates_test/singleton/mysql_suites_test.go.tpl
// override/templates_test/upsert.go.tpl
// DO NOT EDIT!

package driver

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

var _templates17_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x5f\x6f\xdb\x48\x0e\x7f\xb6\x3f\x05\xd7\xe8\xee\xca\x07\x55\xed\x01\x87\x7b\xc8\x21\x0f\xcd\x9f\x76\x73\x4d\xda\x24\x6e\x2e\xc0\x05\x41\x31\x91\x28\x67\x90\xf1\x8c\x3a\x1a\x25\xf5\xe9\xf4\xdd\x0f\xe4\x48\x96\xe4\xd8\x8e\xfb\x27\x87\x7d\x8a\x35\x43\x91\x1c\xfe\xc8\x1f\x39\x4a\x59\xbe\x84\x17\x42\x49\x91\xc3\xce\x2e\x44\x6f\xe8\x17\xe6\xd1\x27\x71\xa3\x10\xfc\x9f\xe8\x83\x98\x61\x55\x0d\x59\x34\x8f\x6f\x71\x26\xfc\x36\xbd\xd0\x4a\xc0\x7f\x21\x9a\xb4\xbb\xfc\x82\x4c\x21\x7a\x93\x24\xef\x94\xb9\x11\x0a\x5e\x56\xd5\xf0\xd5\x2b\xb8\xc8\x72\xb4\xee\x1d\x08\xe7\x70\x96\xb9\x1c\x84\x06\xa9\x69\x2d\x04\xa1\x13\x48\x0c\xf2\x5a\x91\x25\xc2\x21\x18\x0b\x72\xaa\x8d\x45\x30\x1a\x62\xa3\x53\x25\x63\x17\x0d\xd3\x42\xc7\x10\x18\xf8\x4b\x59\x7a\xff\xa3\x8b\x6c\x22\xf5\xb4\x50\xc2\x56\xd5\xb8\xb1\x12\xb0\x13\xda\x38\x88\x3e\x98\x7d\xa3\x1d\x7e\x75\x55\x15\xbb\xaf\xa4\x8a\x1e\xa2\x7a\x31\x84\xb2\x44\x9d\x90\x93\xb5\xe5\x7d\xa3\x8a\x99\xce\xc3\xda\xb9\xfa\x11\x6e\x8c\x54\x51\xfd\x30\x06\xb4\xd6\x58\x28\x87\x03\x8b\xae\xb0\x1a\x4c\xe4\x0d\x7b\xbb\x5d\x9b\xfc\xde\x3b\x74\x07\x7b\xc1\xb8\x2c\x51\xe5\xc8\x7e\x84\xd0\x6c\xd4\x92\xf5\xbe\x4e\xaa\x2a\xdc\xe8\xc9\x78\x58\x0d\x87\x0b\xa7\x87\x3e\xdc\x14\xc0\x4e\xc8\xe9\xe7\xa9\xd0\x32\x5e\x0a\xfe\xe9\x8f\x45\x1f\x58\x67\x4e\x6b\x1c\x80\xad\xe1\x38\x7d\x6e\x3c\xca\xe1\x40\xa6\xe4\x14\x65\xe7\xff\x13\x8c\x7f\xb0\xd1\x5f\x76\x41\x4b\x45\x5e\x0c\x32\x0a\x51\xc0\xfa\x2e\xad\xc8\x0e\xad\x0d\xd0\xda\xf1\x78\x38\xa8\x56\x01\xb7\x06\xa9\x55\x40\x41\x91\x4b\x3d\xa5\x67\xfc\x8a\x71\xe1\x8c\xfd\x96\xc2\xe9\xa8\xce\xbe\x0f\xc5\xd3\xc7\xf1\x24\x47\x7c\xec\x0e\x6b\x97\x3a\x51\x7d\x0c\x6d\x2b\x5e\x2f\x75\xde\x7a\x3a\xd6\xdb\x43\xbe\x22\xcf\xba\x79\x45\x6e\x3c\x1f\xac\xf7\xc2\xc2\x6c\x3e\x39\x3b\x5e\x19\xcc\x0b\x2d\xbf\x14\x8d\x55\xd8\x85\xab\xeb\xdc\x59\xa9\xa7\x25\xf3\xac\x15\x7a\x8a\xf0\x42\x86\xf0\x22\x36\xaa\xc3\xb4\xcd\x0b\x64\x61\x40\x92\x32\x65\x91\xc8\xeb\xa3\xd5\x51\x59\xf2\x8a\xa7\xed\x51\xe8\xe5\x1a\xb7\xea\xdf\x15\x7b\xbb\xc8\x85\xe7\xc8\xb2\x09\x62\x0f\x29\x48\x4c\x5c\xcc\x50\x3b\xe1\xa4\xd1\x90\x1a\x0b\xb7\xe6\x01\x9c\x81\xcc\x9a\x0c\xad\x9a\x43\x91\x63\x1f\x0e\xb6\xd8\x43\x64\xdb\x24\xfd\x73\xe5\xe8\xa2\x4d\xc8\x14\x0c\xec\xb6\xe9\x54\xb7\x0d\xde\xcf\xa3\x0f\xf8\x10\x8c\xca\x32\x3a\xbd\x9b\x7a\xf4\x76\x40\x1b\x28\xcb\x5e\x23\xa6\x70\xdd\xcb\x04\x13\x0e\x61\xc1\xa7\x1d\x71\xfe\x79\xa4\x09\x48\x45\xd0\x8c\x9c\x9c\x61\xee\xc4\x2c\xfb\xec\xa5\x3e\xdf\xa2\xca\xd0\x8e\x20\x82\xca\x4b\xb7\x35\xf2\x87\x31\x77\x75\x5a\x75\xab\x29\x31\x7b\x98\x1a\x8b\x3e\xa8\x2c\xb4\x75\x69\x3d\x2e\x9e\xf6\xb4\xe4\xee\xa0\xcd\xc5\xe1\x40\xff\xe7\x00\x53\x51\x28\xc7\x83\xc8\x97\x02\xad\xc4\x3c\xfa\x60\xf4\xbf\xd1\x9a\x7a\x6b\x82\x04\x6b\x0d\xfa\x81\x79\xd0\x2d\xec\x75\xa4\x2f\xa5\xbb\xad\x85\x43\x30\x63\x52\xeb\x0b\xe3\x09\xad\x5b\xd6\x29\xeb\xe4\x00\x29\xd4\xc1\x42\xf7\x98\x10\x7d\xbd\x0e\xcf\x58\x68\x0a\x96\x87\x00\x1e\xa4\xbb\x05\x01\x8e\x27\x28\x77\x2b\x1c\xd4\xfb\x4d\xed\x50\x1d\x09\x28\x58\x33\xc4\x6c\xb7\x41\xf7\xd5\x2b\xd8\x2b\xa4\x4a\x20\x16\xf1\x2d\xc2\x1d\xce\x41\xea\x97\x4a\x6a\x84\x62\xaa\xa4\x9a\xc3\x4b\x98\xcd\xf3\x2f\x0a\xee\x73\xc8\xe8\x6f\x66\xcd\x8d\xc2\x59\x3e\x1c\xdc\x14\x29\x85\x20\x77\x76\x26\xf4\x54\x21\x35\xb9\xbd\x22\x4d\xd1\x06\x63\xde\x8d\x2e\xad\x74\x38\x61\x12\x0a\x72\x67\x63\xa3\xef\xa3\x23\x67\x44\xd0\xcb\xf3\xe8\xbd\xd4\x09\xd1\x1d\x25\xdf\xe7\x10\x62\xd2\xea\xe9\xaa\x2f\xb7\x6f\x54\xce\x21\x59\xd6\x1d\xf3\x69\xda\xe5\xbd\xb9\xc3\xe0\xf7\xe8\xf7\xa7\xdc\xe8\xd3\xc0\x7a\x37\xfa\x72\xdf\xe3\xc6\x63\x9d\x9d\xec\xfc\x09\xba\x9a\x94\xdc\xa0\x8a\xb0\xdd\xd9\x05\xda\xad\x37\xc6\xc3\x41\x0b\xde\x69\xd1\x80\x77\x53\xa4\x63\x2e\xe5\x95\x65\xe1\xcb\x76\x9f\xd2\xe5\xa4\x70\xd1\xf9\xb1\x89\xef\x48\x13\x27\x50\xe8\xf3\x28\x21\x43\x4f\xbf\x7f\x75\x87\xf3\xeb\xad\x0d\x5d\x68\xe5\x4d\x0d\x07\xd4\x07\x89\x07\xb8\x26\x7c\xf5\xfc\x52\x1b\xa6\x00\x34\xc3\xa7\x45\x47\x8e\xf4\xd1\x3b\xea\x3c\x51\x9d\x0e\x07\x83\x75\x1e\x34\x25\xfa\xb4\x48\x97\x24\xb6\x93\x36\x85\xeb\xbe\xd0\x66\x03\x3d\x8e\x87\x83\x41\xdd\x0c\x77\x76\x97\x8a\xe0\xa2\xf3\xf4\xe3\xfe\x9f\x5a\x39\x13\x76\xfe\x1e\xe7\x1d\x61\x0a\x71\xc3\x48\xde\x78\x87\x8e\x9e\xee\x2f\x85\xf6\x4c\x64\x1a\x82\x5a\xea\x36\x21\xc4\xa6\x50\x09\xf3\xfd\x0d\x93\x4f\x7d\x56\x4f\x4d\xa0\x64\xce\xdd\x87\x09\x8a\xcc\x41\x97\x64\x26\x34\x49\xcf\x32\x85\xd4\xf7\x03\x8b\x2e\x6c\xd3\x9f\x5e\xe2\x3c\x88\x88\x97\xe7\xb0\xeb\xf5\xfb\x4c\x3a\xa3\xa5\x13\x62\xe5\x20\x91\x42\x61\xec\x42\x18\x2d\xb9\x36\x6a\x5a\x70\xd3\x7b\x5b\x8d\x16\xbd\x06\xd8\x85\x74\xe6\xa2\x49\x66\xa5\x76\x29\x87\x7f\x34\x39\x3c\x3e\xdc\xff\x04\xbf\xe6\xf0\xf6\xfc\xe3\x09\x9d\xf7\xf8\xac\xaa\x96\x74\x97\x65\x74\x7e\x56\x55\x70\xf9\xc7\xe1\xf9\x21\xfc\x9a\x8f\x18\x17\x3f\xa2\xe5\xd1\x3f\x8d\xd4\x41\x7b\xca\xa3\x04\xb5\x3b\x2b\x8c\xc3\x89\x92\x31\x36\x1e\x47\xc7\x67\x21\x34\xbf\xcf\xcf\x38\xc5\xc7\x21\x8c\xc2\xd1\xb8\xd1\x56\x2b\xb8\xbc\x45\x8b\xfb\x4a\x14\x39\x32\x3e\xe4\xd0\xc8\x1f\xf8\xdc\xff\x7c\xdd\x0d\xdc\x02\x76\x7f\xd8\x7b\xa1\x0a\x3c\x11\x59\x26\xf5\x34\xe4\x52\x6b\x5b\xdd\x9e\xd4\x49\xbd\xb5\xae\x75\x7e\x9a\x67\x18\xae\x23\x80\x85\xda\x36\xc2\xf5\x78\xd0\x69\xeb\xbd\xbe\x4e\xec\xd5\xe4\x23\x1d\x98\x04\xeb\x64\x5c\x60\xf3\xdc\xce\x92\x5d\x32\xb8\xc2\xd5\xbe\xaf\xec\x6c\xe5\xbb\x2b\x87\x91\x69\x1a\x53\x86\xec\x48\x27\xd2\x62\x4c\x79\xeb\x17\xfe\x45\x12\x1f\xd3\xc0\x50\xe3\xb9\x17\xaa\x37\x54\xf0\x66\xfe\xd6\x9a\x59\x73\x04\x56\x58\x93\x6c\x0f\xa4\xb1\x27\x45\xef\x49\x0e\x57\xd7\x52\x3b\xb4\xa9\x88\xb1\xac\x16\xd3\xc5\x72\xb0\x3a\x81\x6c\x5e\x6c\x8d\x9f\x3a\xbb\xde\x74\x47\x87\x3f\xa9\x4c\xfd\x78\x7a\x80\x37\xc5\xf4\xc4\x24\xc8\x5a\xa9\x50\xde\x72\xa1\x28\x1d\xb4\xfb\xdc\x9c\x6c\xa3\x8b\x4b\x75\xfc\xb4\x34\x45\x67\x31\x93\xbe\x88\x85\x3e\x16\xb9\xf3\x6c\x7e\x74\xd0\xbd\xcf\x2c\xed\xd4\xf7\x1a\xbe\xd5\xac\xda\x1a\x2c\x8d\xf5\x7e\xd5\x62\xce\x13\x5f\x3d\xb6\xd2\xf0\xc9\x43\x7e\xd0\x71\xda\xfb\x14\x45\xd1\x98\xb5\xd0\xe4\xbf\xf9\xe5\xda\x42\xc0\x93\xed\x06\x45\xf5\xc5\xaa\xa7\x73\xb5\x9b\x9f\x9b\x84\xff\x36\x07\x1f\xbf\xf6\xed\xae\x35\x83\xf6\x8a\x92\xe8\xb7\x08\xba\xd4\xd2\x8d\xd6\xb3\xcf\xa6\x46\x41\x93\xcd\x32\x23\x2f\x20\x5f\x0b\x20\x25\xbe\xa2\xd5\x03\x90\xda\xfd\xfd\x6f\x3d\xe7\x68\xd3\x4f\xbe\x27\x22\x83\xab\xeb\xa2\x16\xa1\xf5\x86\xfe\x78\xa0\xeb\x97\xcc\x86\x9a\x59\x74\xc2\xa9\x71\x06\x78\x3e\xa9\xef\x3a\x4f\x7a\xea\xbd\x6c\x62\xef\xb3\x24\xea\x88\x25\x34\x48\xad\x0d\xe7\xa1\xb5\x93\xb9\x8e\xdf\x0a\xa9\xda\x32\x30\x8a\xbf\x94\xf2\x98\x93\xe0\xd7\xa6\x08\x4e\xdf\xe3\x7c\x71\x4b\x7e\xdd\x42\xb6\x74\xf7\xe7\xcf\x52\xdc\x74\x17\x9a\x7a\xa2\x9f\xa4\x53\x7e\x9a\xab\xd9\x71\x49\x9a\x64\x4d\xe4\xfd\xf0\xb2\x55\x05\x3c\xfa\xc5\x46\x45\xc4\xac\x55\x15\xf8\x53\xfb\x93\xd5\x38\x31\xef\xfc\xf6\xdb\xfa\x08\xff\x95\x76\x97\x77\xae\x5e\x5f\xd3\xde\x66\xaa\xbe\x1a\xb5\x61\xa9\xaa\xd1\xf5\x7a\xa8\x7a\x97\xc5\x45\x8e\x3c\x5b\x07\xe9\x4e\x29\x3f\xa1\x64\x2c\x3a\x2b\xf1\x1e\x9b\x7b\x1d\xf3\x73\xbe\xa6\x84\x80\x8e\xdb\x4b\xf7\x4d\x5d\x66\x9b\x6e\x15\xb6\x55\x35\xfe\x31\xfe\x6f\x06\xab\x2d\x5a\x40\xf7\x04\x9e\x92\x16\x05\xb7\x4c\x8c\x1d\x7a\x63\xed\xe7\xe6\x21\xe8\xdb\x7b\xac\x2e\x9a\xc4\x82\x27\x0c\x6a\x85\x5e\x7f\x97\x34\x57\xa8\x5c\xc1\x9a\xdf\xaa\xbe\x21\xd4\x9f\x90\x12\x99\xc9\x0a\xfe\x4c\x93\xf8\xbb\xc4\xe6\x9c\xa0\xd8\x75\x4b\x62\xe7\xd1\x3d\x6a\xbb\x8b\x59\x73\x01\xdc\x42\x9c\x2f\x7c\xb0\xeb\x23\xb5\xb5\x81\xc5\xc5\x6f\xb0\xe1\x0b\xd3\xe2\x9f\x25\x89\x79\x93\x3a\xb4\xdf\xf5\x75\xa9\xa6\x84\x4e\x1f\x67\xa5\x9a\x08\xb7\xfb\x95\xf3\x7f\x01\x00\x00\xff\xff\x45\x71\x40\x0b\xe3\x1a\x00\x00")

func templates17_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates17_upsertGoTpl,
		"templates/17_upsert.go.tpl",
	)
}

func templates17_upsertGoTpl() (*asset, error) {
	bytes, err := templates17_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/17_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSingletonMysql_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\xdf\x8f\xd3\x30\x0c\xc7\x9f\x93\xbf\xc2\x54\x3a\x5d\x2b\x45\x3d\xee\x15\x69\x0f\x77\x6c\x9c\x0a\x63\xbf\x07\x42\x88\x87\x6c\x71\xb6\x48\x5d\x3a\x12\x67\x30\xa1\xfd\xef\x28\x5d\xb7\xf5\x8e\x21\xf1\xc0\x4b\xeb\xc4\xf6\x37\xfe\xd8\xbe\xbb\x83\x45\x30\xa5\x9a\x6f\x3d\x3a\x1a\x07\x74\xfb\x8f\xfb\xe9\xb8\x7f\xbc\xf5\x20\x21\x1e\x3c\x49\xc2\x0d\x5a\x02\x4f\xce\xd8\x15\x04\x1f\xbf\xb4\x46\x08\x75\x62\x57\x92\x84\xad\xab\x76\x46\xa1\xca\xb9\x0e\x76\x79\x5d\x37\x55\x46\x82\x72\x66\x87\xce\xe7\x5d\x23\x4b\x5c\x92\x00\x92\x8b\x12\x07\x72\x83\x8d\xbe\x80\xb0\x55\x92\x50\xc0\x8f\xb5\x21\x2c\x8d\x27\xf8\xfa\xed\xe8\xcb\x4e\x35\xfc\xe2\xec\xe2\xed\xc4\xdb\x8d\xb4\xab\x12\xf3\x42\xa1\xa5\x71\xa8\x08\xa7\xa5\x59\x62\x7c\x32\xef\x8f\x05\xc4\xff\x64\xdc\xd2\xcc\x38\x67\x8b\xa0\xe1\x4d\x3b\xfb\x09\xe9\x31\x68\x8d\x2e\xcd\x38\x53\xa8\xd1\xb5\x9c\xa3\x70\x72\x2e\x82\x8e\xe9\x3b\xe9\x60\x59\x95\x61\x63\x7d\x53\x17\x67\x46\x43\x89\x36\xbd\x3c\x03\xaf\x3a\xf0\x3a\xd6\xcb\x4e\xa1\x9d\x26\xd8\xe7\xef\x2b\xd3\x0a\x15\x90\x08\x48\x32\xce\x0e\xfc\xac\x73\x6c\x45\x06\x9d\x93\x88\xde\x50\xfe\x6e\xeb\x8c\x25\x9d\x72\xc6\x22\x82\x88\xff\xa4\x18\x4c\x7b\x93\x19\x14\x4f\x83\xe1\xa4\x07\xc5\x60\x36\x84\x1b\x0f\xe9\x8d\xcf\xe0\xd3\x43\x7f\xde\x9b\xd6\x76\x52\x07\x9f\x5b\x5e\x9f\x9a\xba\x6a\xbb\x45\x5b\xca\x25\xae\xab\x52\xa1\xf3\x75\x17\xe7\x1e\x0b\xab\xf0\x67\xdb\x21\x5e\xc0\x0a\xb8\x17\x70\x9f\x45\xa9\x8c\x33\xe6\x90\x82\xb3\xb0\x08\x3a\x9f\xd6\xc8\x69\x43\xf7\x82\xa2\x81\x38\x33\xfc\xa5\x78\x18\x0e\xa0\x3b\x1f\xf5\x8b\xb7\x0f\xb3\x1e\x7c\xe8\x7d\x81\xf9\xa8\x1b\xcd\x9a\xea\x19\x54\x8b\xe9\xbf\x21\xc5\x91\xeb\xca\x81\x11\xb0\x8b\x6b\xe3\xa4\x5d\x61\xb3\xac\xf5\x6c\x8c\x06\x73\x19\x77\xa4\xca\x3f\x3b\x43\xf8\xb8\x27\x4c\x6f\xc5\x6d\x6c\xc9\x81\x33\xf6\x3d\xae\xa7\x7a\xbe\x79\x97\xbd\xfd\x63\x65\x77\x19\x6f\x89\x35\x8d\x3c\x6a\x5c\xf3\x24\xd0\x69\x9a\x96\x26\xff\x98\x79\x2c\x30\xbb\x6d\xa6\x73\x6d\x6c\x07\xfe\x3b\x00\x00\xff\xff\xa9\x3a\x4a\xd3\x2e\x04\x00\x00")

func templatesSingletonMysql_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSingletonMysql_upsertGoTpl,
		"templates/singleton/mysql_upsert.go.tpl",
	)
}

func templatesSingletonMysql_upsertGoTpl() (*asset, error) {
	bytes, err := templatesSingletonMysql_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/singleton/mysql_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonMysql_main_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x5b\x6f\xdb\x3a\x12\x7e\x96\x7e\xc5\xac\x81\x9c\x95\x72\x14\xe6\x00\x05\xf6\x21\x85\x11\x34\x8e\x53\x04\x6d\x2e\xb5\xb3\x5b\x2c\x9a\xa2\x65\xa4\x71\x42\x54\x22\x15\x92\x8a\xe3\x0d\xf2\xdf\x17\x43\xea\xe6\xc4\x32\x5a\xa0\xfb\xb6\x4f\xad\x38\x1f\xe7\xf2\x0d\xe7\xe2\x3c\x70\x0d\xfa\xf6\xf1\x6c\x35\xff\xf4\xf1\x07\xae\x60\x0c\x1a\x6f\xf1\xb1\x64\x67\x95\xb1\x13\x55\x94\x22\xc7\xe8\x7b\x74\x58\xc4\x51\x94\x5c\xcb\xf8\xf0\xda\xfc\x39\xb9\x38\x9f\x5f\xcd\xde\x9d\x9e\x5f\xb1\xdd\xc3\x93\x8b\xd9\xf4\xf4\xfd\x39\x7c\x98\xfe\x9b\xed\x1e\x5e\xcb\xf8\xcf\xef\x71\x18\xda\x55\x89\x50\xac\xcc\x7d\x7e\x85\xc6\xa2\x06\x63\x75\x95\x5a\x78\x0a\x83\xec\x66\xa2\xa4\x84\x5d\x73\x9f\xb3\xe3\xa3\x90\x0e\xce\x79\x81\x81\xb1\x5a\xc8\xdb\x30\xb8\x53\xc6\xb6\x1f\x95\x41\xdd\x7e\x94\xdc\x98\xf6\xc3\x98\xbc\x50\x59\x77\xad\x54\xda\x06\x42\xda\x30\x0c\x54\x69\x85\x92\x27\x22\x47\xa8\xa5\x61\x60\xd1\xd8\xe3\x23\x32\xd4\x9c\x3d\x87\xe1\xa2\x92\x29\x08\x29\x6c\x14\x7b\xcf\xce\xb8\x90\x30\x86\x3f\x7a\x9e\x3f\x3d\xb7\xc8\xa8\x80\xdd\x9e\x24\x06\x83\xb6\x2a\xa3\x18\x50\x6b\xa5\x49\x03\xb1\x89\x5a\xfb\x83\x30\x0c\x1e\x44\x89\x9a\xcd\xd1\x1e\xe3\x82\x57\xb9\x8d\x46\xee\x3e\xab\x9d\x1f\x25\x30\xb2\xba\xc2\x51\x3c\x0c\xa5\xb8\x46\x09\xbc\x79\xf3\xd7\x3f\xe2\x30\x0c\x0a\xe6\xe9\x82\x31\xf8\x1b\xef\xd1\xce\x5d\x40\xcd\x85\xec\x46\xf2\xc2\xa9\x2c\x18\x71\x39\x8c\x24\xa9\xc7\x11\xcd\xc3\x38\x92\x7a\x1c\x65\x60\x18\x47\xd2\x1a\xa7\xf4\x9a\xdd\x53\xb9\x1e\x8f\x03\xd5\x24\x0c\xeb\x6b\x58\xa2\xb8\x89\xd5\x31\x3c\xf0\x9c\xb3\x23\xbc\x15\xf2\x5f\x3c\x17\x19\xa7\x3c\x47\x31\xab\x3f\x30\x0a\x83\xc0\x41\xbc\x9e\x73\x65\xa7\x45\x69\x57\x91\x0f\x30\x81\xb5\x78\x92\x41\x30\xf1\xd2\x82\x3d\x49\x2d\xf8\x5c\xd9\xc8\xfd\x67\x7a\x5f\xf1\xdc\x44\x3e\xd6\x04\xfe\x6a\x2f\xf8\x00\xb7\xa8\xf7\x09\x6c\xf1\x4d\xbe\x86\x2f\xd4\x3c\xb4\x37\x5a\x5e\x92\x30\x88\xd9\xe4\x0e\xd3\x1f\x11\x71\x24\x16\xee\xf1\xfd\x6d\x0c\x52\xe4\xf4\x1c\x03\x8d\xb6\xd2\x92\x4e\xc3\xe0\x39\x0c\x83\xfd\x7d\x98\x68\xe4\x16\x81\x83\xe6\x32\x53\x85\xf8\x0f\x66\x90\xdd\x00\xf9\xc0\x28\x2b\xbd\x42\x19\x77\x18\x36\xb7\xfc\x26\x47\x2f\x68\x63\xe8\x19\x1d\x43\xc1\x0a\xfe\x03\x2f\xda\xda\x8b\xe2\xb7\xc3\xee\x28\x6d\xd8\x67\xcd\xcb\x08\x35\xe5\x25\x55\x55\x9e\xc9\xbf\x5b\x20\x15\xe0\xeb\x17\x16\x22\x77\xcf\xf8\x79\xdd\x4a\xa6\x55\x79\xe5\x9c\xdc\x6a\x81\xee\xf5\xaf\xa5\x2e\xee\x9f\xbc\x18\x06\x59\x55\x94\x93\x22\x83\x83\x31\xe0\x23\xa6\x6c\xa2\x8a\x82\xcb\xac\x7e\x9a\x24\x1d\x25\xe4\x8c\x2f\x56\xe3\x03\x4e\x60\xb4\xb7\x27\xd5\x5e\xc6\x2d\xf7\xe2\x9a\xa6\xc0\x5b\x1f\x56\x38\xa4\x8c\x34\xdd\x70\x83\x4e\xde\xa5\x86\x88\xd7\x09\x2c\x49\x9b\x50\xec\x52\x94\x18\xc5\xad\xd3\x6c\x6e\x33\x55\x51\x01\x2e\x7b\x96\xe9\xd4\x75\x37\x89\xcb\x93\x0f\xb8\x3a\x46\x63\xb5\x5a\xa1\x8e\x7a\xfd\x3f\x01\xbd\x96\xd4\x4e\x23\xd7\xf6\x57\x32\xba\xe0\x22\xc7\x0c\xac\x02\x43\x57\xa1\xa5\x0d\x52\x1f\xb8\xcf\x6c\x67\xa9\xef\xe7\xef\xb0\xb5\x6e\x67\x43\x48\x9f\xb9\xd8\x64\x65\x51\x58\x76\xa9\x85\xb4\xb9\x24\xf5\xf1\xcf\x19\x5e\x72\x61\x61\xa1\xf4\x50\x9c\x61\xf0\x8d\xd2\xc1\x26\xb9\x32\x18\xc5\xb0\xbf\x0f\xef\x16\x34\x17\x9b\x77\x26\x0c\x64\x4a\x62\x02\x29\x21\xc0\xde\x21\x2c\xb5\xb0\x08\x28\x33\x50\x0b\x77\x50\x8a\x12\xc3\x8d\x94\xfd\x0f\x63\x79\x19\x47\xad\x40\x8a\x7c\xcb\x5c\x34\xf9\x99\xca\x30\x72\xed\xdd\x8f\xdb\xb8\xfe\x97\xfc\x32\x4b\x61\xd3\x3b\x70\xd2\xa7\x30\x48\xb9\xc1\x7a\x0e\x1e\x74\x1e\x8e\x66\xd3\x4f\xff\x3c\x9d\x4d\x8f\x47\x0d\x62\xc1\x73\xb3\x0e\x39\x3e\x9d\xbf\x3b\xfa\xe8\x20\x75\xed\xf4\xa5\x97\xb3\xe9\xc9\x74\xe6\x35\x6c\x19\xe2\xeb\x55\xd7\x73\xb3\xd6\x43\x24\xce\x4b\x62\x71\x11\x51\x45\xd6\xf0\x3d\x6a\x4f\xe3\x1d\xe3\x2a\xb3\xdb\x38\xe2\x61\x43\x2f\xdb\x63\xb7\x36\xd8\xa2\x4c\x5c\xee\x5c\x39\x57\x56\xe4\xec\x0a\x8b\xd2\xc1\x46\xb4\x24\x78\xfd\x4d\x43\xdc\xd6\xe7\x07\xb3\xea\x1f\xcb\xc6\xde\x6a\xae\x26\x97\x64\xda\x11\x1c\x06\xdf\x92\xfa\x7d\x29\x43\xb5\x68\xeb\x89\xe8\x0d\x2b\xc3\x4e\x0d\xcd\xa6\x47\x61\xac\x7b\x54\xce\x01\xaf\x63\x0c\x94\xc5\x30\x78\x06\xcc\x0d\xc2\x2f\xf8\xe9\x06\x00\x48\x65\xa9\x84\x2d\x14\xed\x6e\x42\x0e\x52\x06\x4e\xca\xfa\x1d\x3b\xae\x46\x5f\xd2\x5c\xa0\xb4\x5f\x09\xd2\x89\x17\xb5\x94\x2e\x8f\x77\xcc\xb5\x74\xc9\xa9\x9d\x7f\x0d\xa3\x49\x3d\xde\xc9\x6a\x18\x7d\x6d\x84\xd1\xba\xd0\x69\xa3\xaf\xcd\xda\xb8\x31\x4b\xa5\xb3\x0e\x4a\x27\x1b\xa1\xc6\xe4\x7b\xf4\xfa\x3b\x68\x5b\x31\xcd\x70\x8f\x3d\xdd\x9e\xd7\xa6\x96\x5f\x90\x50\x6a\x65\x55\xaa\xf2\xb1\x4d\xcb\x6d\x5c\xb5\x3d\xe9\xff\x74\xbd\xa4\xab\x5f\xba\xf4\x7c\x8b\x92\xb9\x0d\x27\xee\x3a\x1d\x9d\xd5\x6d\x7b\xb8\xb6\xd7\xb7\x8b\xae\xb2\xa9\x81\x52\x65\xf5\x7b\x48\x5d\x89\xcd\x68\x87\x1d\xf3\xf6\xd5\x78\x6f\x8c\x17\x4c\x57\x72\x52\x64\x91\xb9\xcf\x9b\xfd\x6f\xb4\xc5\x8f\xfe\x72\xb4\xdd\x0b\x42\x76\x3e\x50\xa9\x52\x45\x9b\xdf\xea\x8d\x45\xae\x33\xb5\x94\x7d\x5f\xc4\xc2\xed\x45\xee\x37\xe0\xeb\xce\xd0\x88\x5a\xc6\xfb\xb3\xfb\xe0\x17\xd7\xbf\xd6\x6f\x65\xd8\x0c\x0b\xf5\x40\x0f\xe6\xa7\x5a\x75\x13\x26\x6d\x4c\x49\x33\x01\xeb\xd1\x90\x00\xd7\xb7\x06\x18\x63\xcd\x64\x6b\x63\x73\x82\x31\xf0\xb2\x44\x99\x45\x5f\xbe\x7a\xc0\xd3\xcb\xcd\xee\xd9\xab\x60\x8c\xd1\x33\x4b\x37\x2c\x85\xb5\xc5\x1e\x8e\x60\xed\x06\xe7\xf5\x1a\x76\x8e\xcb\x19\xf2\x0c\xb5\xf7\x94\xb4\x19\xbf\xfa\x1d\x8c\xe1\x8f\x9b\x95\x45\xc3\x8e\xaa\xc5\xc2\xfd\x8a\x25\x51\xcd\xe2\x2b\x51\xda\x5f\x1a\xbd\x8a\xf6\xd0\xcf\x02\x7f\xb9\x9f\x0a\x12\xcf\x2a\xf9\x3a\x0b\xfd\xad\xa3\x19\x40\xba\x92\x52\xc8\xdb\x83\x51\xcb\xa6\x8f\x2d\x5e\x87\x7b\xd3\xf5\x2f\xa0\x28\x7e\x2d\x45\xad\xfb\xd2\x81\x7c\x6f\x5d\x4e\x52\x25\xe9\x41\x46\xf5\x1f\x20\x12\x9f\xbe\x78\xf8\x6d\xbe\x78\x9a\x89\x53\xef\xcc\xad\xff\xdc\x0f\x3a\x44\xcd\xd9\x7d\xce\x2e\x4a\x94\xdd\x8e\x9f\x69\xf1\x80\x9a\xb9\x6d\xfb\xa8\x12\x79\xf6\xa9\x42\xbd\xaa\x03\x6a\x7e\xa5\xfa\x4e\xb8\x5e\x83\x4d\x63\x6e\x3a\x6f\xdd\x04\x7b\xad\x6f\x3d\x07\x1d\x11\xc9\x2b\x76\xd6\x03\x79\x0e\xff\x1b\x00\x00\xff\xff\x01\x8f\x47\x4d\x03\x12\x00\x00")

func templates_testSingletonMysql_main_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMysql_main_testGoTpl,
		"templates_test/singleton/mysql_main_test.go.tpl",
	)
}

func templates_testSingletonMysql_main_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonMysql_main_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mysql_main_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonMysql_suites_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xc1\x0a\xc2\x30\x10\x44\xef\xfd\x8a\xa5\xe4\xd0\x4a\x9b\x0f\x10\x3c\x78\xd4\x83\x88\xb4\x1f\x10\xed\xb6\x04\xe2\x5a\xba\x5b\x10\x42\xfe\x5d\xd2\x46\xe9\xc1\xdb\x0c\x6f\x32\x99\xed\x67\x7a\x40\x83\x2c\xed\xc8\x38\x49\x21\xb0\x13\x64\xb1\x34\xe8\xa6\x04\x9f\x01\x78\x5f\xc3\x64\x68\x40\x50\x96\x3a\x7c\x57\xa0\xc4\xdc\x1d\xc2\xfe\x00\xba\x89\x8a\x43\x48\x39\xdb\x27\xa8\x4f\x7c\x7e\x59\x5a\x30\xd4\x3f\x8e\x8e\xb7\x56\x19\x67\x0d\xc7\x22\xa5\x8f\x51\x22\xaf\x8d\xdf\x96\x8b\x79\xe2\x92\x16\x7d\x9b\xa9\xc8\xbd\x5f\x9f\xe8\x76\xbc\xba\x79\x32\x2e\x84\xbc\x82\x38\xf8\x0f\x59\x2f\x2a\x97\xbf\x90\xba\xed\x8c\xe4\x42\xf6\x09\x00\x00\xff\xff\x11\x5d\x4c\xce\xff\x00\x00\x00")

func templates_testSingletonMysql_suites_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMysql_suites_testGoTpl,
		"templates_test/singleton/mysql_suites_test.go.tpl",
	)
}

func templates_testSingletonMysql_suites_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonMysql_suites_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mysql_suites_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testUpsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\xcd\x6e\xdb\x3c\x10\x3c\x8b\x4f\xb1\x9f\xf1\xb5\xa0\x0a\x85\x69\xaf\x29\x7c\xc8\xdf\x21\x68\x6b\xb8\xb1\x7c\x2e\x18\x69\xe5\x10\xa6\x49\x95\x5c\xd5\x76\x05\xbe\x7b\x41\xc9\x76\x9c\xd8\x69\x73\x68\x0f\x39\xf8\x87\x8b\xd9\x9d\x99\x5d\x2e\xdb\xf6\x04\xfe\x97\x5a\x49\x0f\x67\x43\x10\xe7\xf1\x1f\x7a\x91\xcb\x3b\x8d\xd0\xff\x88\x91\x5c\x60\x08\xac\x6a\x4c\x01\x84\x9e\xda\xb6\xcf\x10\xd3\x7a\xac\x1b\x27\x75\x08\xd3\xda\xa3\x23\x4e\xf0\x2e\x02\x94\x99\x89\x3c\x85\x96\x25\x24\xc6\xd2\x49\xad\x51\xf3\x94\xb1\x44\x55\xa0\xd1\xf0\x5d\x81\x2b\xbb\x34\x13\x65\x66\x8d\x96\x2e\x84\x4b\xab\x9b\x85\xf1\x29\x0c\x87\xbf\x83\x8d\x9d\x5a\x48\xb7\xfe\x84\xeb\x5d\x42\xcb\x92\x84\xc4\x64\xae\x6a\x3e\x88\xdf\xb5\x32\x33\xa0\xce\xc3\x52\xd1\x3d\x58\xa3\xd7\x50\xf7\x79\x30\xc7\x35\x14\x7d\xe6\x20\x65\x49\xd8\xc9\x5a\xac\x27\x5f\x3f\xef\x99\x7b\xa0\x9c\x1a\xf5\xbd\xc1\x7d\x7d\xef\xff\xc8\x69\x2c\x34\x5d\xda\x96\x0c\xc8\x42\x61\x4d\xa5\x55\x41\x60\x4d\xcf\xcd\x12\x8f\x58\xc6\xde\x3b\x69\x4a\xbb\x50\x3f\x51\x8c\x70\x39\x41\x2c\x79\xca\x92\x1f\xd2\x01\xba\xee\x63\x1d\x4b\x4e\x4f\xe1\x9c\x08\x17\x35\x01\xdd\x23\xdc\x8c\x26\xd7\xb7\x39\x78\x55\x22\xd8\x0a\xa4\x81\xe9\x38\x46\x58\x62\x63\xc5\xa3\x56\xda\xde\x6f\x2c\xba\xcf\x39\x21\xd7\x14\xc4\xa3\x98\x0c\xde\xda\x0c\x9e\x69\xfe\xd5\x45\xbe\xae\xd1\x67\x50\x49\xed\x31\xfd\xd8\x15\xfa\x6f\x08\x46\xe9\x4d\x47\xae\xa3\xd4\x8a\x0f\xa6\xa6\xeb\x05\xd9\x07\x96\xe3\x8a\xc0\x77\xdc\x67\xf0\xc6\x0f\xb2\x58\x6f\xd3\x98\xb6\x55\x15\x18\x4b\x20\x46\xf6\xd2\x1a\xc2\x15\x85\x50\xd0\x2a\x5a\x2b\xfa\xb3\xb8\x90\xc5\x7c\xe6\x6c\x63\x4a\x9e\xb6\x2d\x9a\x32\x04\x96\xf4\x90\x2f\x8d\xa7\x7c\xc5\xbb\x2a\xfb\x15\x0e\x02\x77\x56\x69\x71\x81\x33\x65\xba\x1a\xda\xe3\x7e\x2c\x5f\xf1\x82\x56\x59\x34\xb8\x65\x78\x11\x28\x65\x49\x89\x15\x3a\x88\x6b\xc3\x53\x68\xe1\x1b\x0c\x81\x56\xe2\xd6\x6a\x7d\x27\x8b\x39\x4f\x21\xc4\x11\xef\x86\x61\xc5\x66\x8b\x9e\x33\x1e\x87\x82\xa6\x84\x93\x10\x20\x9e\x3a\xfe\x1b\x53\xa1\xe3\xe9\xe3\xd3\xcb\xe6\xd2\x74\x74\xc7\x87\x72\x30\x8d\xc2\x36\x86\xba\xc0\x93\xab\xb5\x7d\x02\x78\x2a\x2e\x23\xe6\x85\xf2\x1f\x9c\x1f\xaa\xe4\x5b\xda\x08\xe9\x88\x23\xe8\xc3\x23\xc8\x60\x29\x4d\x5c\x23\x04\x87\x85\x75\x65\x06\x33\x4b\x67\x83\xac\xc7\x6f\x44\x3f\xd9\x97\xe9\xf8\xea\x3c\xbf\x3e\xb6\x2f\x7f\x6d\x23\x9e\x85\x1d\xbc\x5a\x42\x88\x7f\xba\x3e\xaf\xef\x5e\xbd\x92\x6b\x15\xd8\xaf\x00\x00\x00\xff\xff\x05\xd5\xa1\x1c\x35\x07\x00\x00")

func templates_testUpsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testUpsertGoTpl,
		"templates_test/upsert.go.tpl",
	)
}

func templates_testUpsertGoTpl() (*asset, error) {
	bytes, err := templates_testUpsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"templates/17_upsert.go.tpl": templates17_upsertGoTpl,
	"templates/singleton/mysql_upsert.go.tpl": templatesSingletonMysql_upsertGoTpl,
	"templates_test/singleton/mysql_main_test.go.tpl": templates_testSingletonMysql_main_testGoTpl,
	"templates_test/singleton/mysql_suites_test.go.tpl": templates_testSingletonMysql_suites_testGoTpl,
	"templates_test/upsert.go.tpl": templates_testUpsertGoTpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"17_upsert.go.tpl": &bintree{templates17_upsertGoTpl, map[string]*bintree{}},
		"singleton": &bintree{nil, map[string]*bintree{
			"mysql_upsert.go.tpl": &bintree{templatesSingletonMysql_upsertGoTpl, map[string]*bintree{}},
		}},
	}},
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"mysql_main_test.go.tpl": &bintree{templates_testSingletonMysql_main_testGoTpl, map[string]*bintree{}},
			"mysql_suites_test.go.tpl": &bintree{templates_testSingletonMysql_suites_testGoTpl, map[string]*bintree{}},
		}},
		"upsert.go.tpl": &bintree{templates_testUpsertGoTpl, map[string]*bintree{}},
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
