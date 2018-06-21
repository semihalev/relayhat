// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/dashboard.html

package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _templatesDashboardHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5d\xfb\x6f\xe3\x36\xb6\xfe\xd9\xf9\x2b\x38\x9a\x16\x4e\x30\xb1\x64\xe7\x31\x93\x71\x6c\x63\x77\xdb\xe9\xdd\x05\xee\x6e\x8b\xb6\x7b\x81\x8b\x60\x70\x41\x4b\x94\xc5\x89\x44\xaa\x14\x95\xc7\x0e\xf2\xbf\x5f\x90\x94\x6c\xc9\x96\x22\x52\x54\xd2\xdd\x22\x45\xc7\x16\x79\xbe\xf3\xe0\xc7\xc3\xcf\x8f\x44\x8b\x37\xdf\xff\xf8\xdd\xaf\xff\xfb\xd3\x27\x10\xf1\x24\x5e\x1d\x2d\xc4\x3f\x80\x6c\x26\x30\x4d\x97\x0e\x43\x31\x7c\x8c\x20\x77\x56\x47\x47\x8b\x08\xc1\x60\x75\x34\x5a\x24\x88\x43\x10\x71\x9e\x4e\xd0\x6f\x39\xbe\x5b\x3a\x3e\x25\x1c\x11\x3e\xe1\x8f\x29\x72\x40\xf1\x6c\xe9\x70\xf4\xc0\x3d\x01\x77\x0d\xfc\x08\xb2\x0c\xf1\xe5\x3f\x7f\xfd\x61\x72\xe5\x6c\x41\xca\xcb\x4e\xce\xc3\xe2\x3a\xc7\x3c\x46\xab\x9f\x85\xdf\xbf\xfe\xf9\x57\x30\x01\xdf\xc3\x2c\x5a\x53\xc8\x82\x85\xa7\xc6\x4a\x63\x02\x13\xb4\x74\x36\x88\x20\x06\x39\x65\x15\xc7\x7f\xa1\x94\xa7\xf1\xa3\x03\xbc\xbd\xc9\x77\x18\xdd\xa7\x94\xf1\xca\xdc\x7b\x1c\xf0\x68\x19\xa0\x3b\xec\xa3\x89\x7c\x72\x0a\x30\xc1\x1c\xc3\x78\x92\xf9\x30\x46\xcb\xd9\x29\x48\xe0\x03\x4e\xf2\xa4\xbc\x20\x03\x8d\x31\xb9\x05\x11\x43\xe1\xd2\xf1\x3c\x3f\x20\x5f\x32\xd7\x8f\x69\x1e\x84\x31\x64\xc8\xf5\x69\xe2\xc1\x2f\xf0\xc1\x8b\xf1\x3a\xf3\xd6\x79\x9c\x40\x6f\xea\x7e\x70\x67\x9e\x9f\x15\xcf\xdd\x04\x13\xd7\xcf\x32\x07\x30\x14\x2f\x9d\x8c\x3f\xc6\x28\x8b\x10\xe2\x12\xfe\xcd\x64\x72\x83\x43\x10\x73\xf0\xb7\x4f\xe0\xe3\x67\x71\x29\xf3\x19\x4e\x39\xc8\x98\x2f\x5c\x8a\xc2\x5e\x66\x11\x4e\xdc\x0d\xa5\x9b\x18\xf9\x34\x50\x6e\xb3\x3b\xe2\x71\x96\x93\x5b\x35\xc5\xfd\x92\x39\xab\x85\xa7\x8c\x25\xf2\x0d\x22\x01\x0e\x3f\x4f\x26\x12\x54\xb8\x5d\x1d\x8d\x46\x62\xf2\xe9\xd1\x68\xb4\xa6\xc1\x23\xf8\x7a\x34\x1a\x8d\x42\x4a\xf8\x24\x84\x09\x8e\x1f\xe7\x60\xfc\xe9\x81\x82\xb3\xf1\xf5\xd1\x68\xf4\x74\x74\x34\x1a\xb9\x11\x62\x74\xe5\x66\xc8\xe7\x98\x12\x65\x90\x40\xb6\xc1\x64\xc2\x69\x3a\x07\xb3\xf3\x69\xfa\xa0\x66\x8f\x16\x5e\xe9\x65\xe7\xee\x4f\x93\x7b\xb4\xbe\xc5\x7c\x72\x8b\x1e\x43\x06\x13\x94\x81\x2c\xc5\xe4\xcf\x8c\xe6\x24\x28\xfc\x33\x9a\xa8\x47\xa3\x72\x32\x67\x90\x64\x21\x65\xc9\x1c\x30\xca\x21\x47\xc7\xd3\x93\x6b\x39\xa3\x69\x44\x0c\x3c\x89\xff\x71\xda\x89\x73\x7e\xf9\x31\x40\x9b\x56\xb0\x62\xb8\x44\x94\x25\xf8\xd3\x7f\x6e\xe8\x6e\x76\x8f\xb9\x1f\xdd\x88\x1d\xbb\xf4\x23\xe4\xdf\xae\xe9\xc3\x67\xe5\x89\xe6\x3c\xc6\x04\xcd\xc1\x54\x22\x96\x7e\xf3\x0c\xb1\x49\x86\x62\xe4\xf3\x39\x20\x94\x20\x35\x9a\xd0\x7f\xb5\x0e\x65\x2d\x23\xcd\x57\x03\x9c\xa5\x31\x7c\x9c\x03\x4c\x84\xff\xc9\x3a\xa6\xfe\xad\x1c\x49\x69\x86\x05\xcb\xe6\x00\xae\x33\x1a\xe7\x5c\x19\xd0\x14\xfa\x98\x3f\xce\xc1\xf4\xf9\xb4\xe6\x21\xf5\xf3\xec\x5d\x0c\xd7\x28\x9e\xcf\x61\xc8\x11\x3b\xd5\x9c\xbc\x46\x21\x65\x48\x73\xb6\x09\xb2\x02\xde\x2b\xf8\x2c\x7d\x00\x01\xe5\x1c\x05\xe0\xed\xfa\x52\xfc\x3c\x9f\xd8\x4d\x80\x33\xb8\x8e\x51\x50\xac\x9c\x9f\xb3\x8c\x32\x51\x52\x3e\x81\x71\x4c\xef\x51\xa0\x0b\xa0\xc2\x2a\xe2\x29\xeb\xea\x6a\xfb\xd7\xad\xd7\x81\x45\xad\x0e\xfd\xfd\x76\x94\xfe\xc0\x40\xce\xb7\xf5\x1a\xd1\xbb\x12\xc4\xb8\xf2\xd5\x7a\xef\xe8\x2d\x8e\x5b\x8e\xef\x0e\xf6\x83\x3c\x8e\xae\xb7\x3d\x39\xc3\xff\x12\x6c\x61\x28\x91\xd7\xe4\x6e\x89\x10\xde\x44\xbc\x3e\x39\x85\x41\x80\xc9\x66\x12\xa3\x90\xcf\xc1\xb9\x7b\x59\x5a\x94\x03\xb2\x53\xbb\x67\xe5\xe5\x32\x8b\x94\x62\xc2\x11\xd3\xc9\xa0\x7b\xc9\x1b\x16\xba\x65\x3f\x6f\x13\xde\xed\x7c\x19\xa0\xea\x43\x2a\x09\xf5\x58\x9e\xd3\x73\x70\x5e\x06\x5e\x26\x3f\xdb\xa5\xb8\xa6\x2c\x40\x6c\x0e\x5c\x51\x26\x90\xd1\x18\x07\x40\xf6\xc6\x14\x32\x44\x78\x65\xce\x84\xc1\x00\xe7\xd9\x1c\x5c\xa8\x03\x6b\xb4\x86\xfe\xed\x46\x36\xf3\x79\xb9\x0f\x55\x79\x94\x66\x98\x83\xf1\x58\xaf\x34\x1d\xac\x3c\xe4\x62\x43\x05\x5a\x6a\x55\xac\xdc\x36\x5d\x55\x9d\xca\x85\xa2\x44\xb3\x83\x12\x95\x17\x1a\xce\x13\xf9\x30\x86\x1c\x9d\x07\xc7\xd3\x53\x20\xfe\x53\xa7\x8a\xce\x1c\x9d\x62\x86\x61\xb8\xc3\x2b\xd3\x8a\x63\x11\x77\x06\x10\xcc\xd0\x84\xe6\xdc\xac\xd6\x2e\xce\x26\x8c\xc7\xb5\xfd\x54\xa3\xfd\xb4\xc6\x78\xa6\xaa\xa0\xf6\x82\x01\x72\x37\xcf\xeb\xd3\xab\x74\x57\x71\xc0\x9c\x53\x19\x4a\x11\x42\xc7\xa1\xb5\xe7\xbd\x83\x4a\xf5\xd9\x15\x46\x35\xfb\x56\x3c\xe9\x38\x35\xe5\x23\x14\x68\x1f\x85\xf5\xf9\xd5\x02\xd4\x38\x30\x9d\x06\xb3\xf5\x99\x99\xef\x83\x84\x66\xee\x7b\x83\x14\xcc\x6a\xd9\x6c\xd5\x59\x53\xbd\x90\x04\x68\x71\xde\x6b\x57\xb6\xc1\xa6\xb9\xba\x13\x9f\xc6\xa2\x7f\xb7\xf4\xb9\x62\x54\x4b\x5a\x34\xc5\xa9\xc1\xc1\x3d\x93\x4a\xd1\x1a\x9a\xaa\x7e\x04\xa6\x5c\x7c\xc6\xd6\xaa\x72\x3a\xdc\x7d\x2e\x6e\x83\x0a\xee\x99\xb6\x55\x52\x37\x22\x1e\x61\x62\x52\xbd\xca\xfc\x6a\xc5\xd4\xb1\x73\x79\x71\x79\x71\xb9\x7f\xb0\xb8\xe7\x1f\xf4\xe8\x5f\x0d\x45\xa3\x20\x95\xe9\xd5\x22\xd0\x87\x49\x16\xc1\x80\xde\xcf\xc1\x14\x4c\xc1\x79\xfa\x00\xde\x7e\x80\xe2\x47\xa3\xaf\x8a\xe2\x99\xb1\xa9\x6e\x52\x23\x51\xfd\xdc\x3b\xbb\x48\x1f\x8c\x23\xd0\xe9\xee\x35\x8b\x5a\x25\x6a\xfe\x2f\xa7\xdf\x76\xbb\xcf\x12\x18\xc7\x03\x48\x50\xf7\xc3\xa5\xa1\x08\x3d\xab\x98\x0c\xa1\x42\xf7\xb2\xd1\x5b\xcc\xaa\xc1\x4b\xe8\xd2\xb3\x8a\x0e\xdb\x49\xd3\xd9\xd9\xbf\x85\x38\x3d\xa8\x98\x06\xf9\xaa\xf3\x5f\x54\xae\x16\x67\x68\xbd\xb1\x54\xae\xfd\x81\x44\xab\xac\x69\x3f\xf5\x5a\x6c\xa2\x3e\x3e\x0c\x36\xc8\x0b\x08\xda\xa6\x78\x74\xe9\xf7\x02\x12\x77\x8b\xdd\x47\x5f\x34\x19\x5a\x89\xde\xd6\x68\x1a\xd4\x6f\x9f\xe4\x7a\xd6\xbd\xb7\x1c\x36\x89\xb1\xaf\x2c\x6e\x33\x7e\x15\x7d\xdc\x1a\xb9\x09\xa5\x07\x56\xcc\x07\xb8\xbd\xa9\xfd\x7b\x6b\xe8\xce\x4c\xfa\x54\x79\x60\x55\xbd\xc5\x37\x95\xd7\x4d\x86\x87\x3a\xfb\x62\xfa\x71\xfa\xf1\xe0\x38\x3c\xbb\x9a\x99\x6e\x2d\x43\xc9\xdd\x64\x37\x94\xf6\xde\xf5\x72\x73\x11\xde\x62\x6b\xad\xc6\xdb\x62\x32\x3a\x9a\x86\xd4\xe7\x09\x0a\x70\x9e\x0c\xf1\x1e\x71\x55\x73\xe9\x29\xf4\x8b\x8a\xc9\x50\x0a\xbd\x9a\x8f\xde\x52\xd7\x2c\x5e\xe4\xbd\xe3\xca\x0b\x91\x9d\x46\xbf\xfa\xf0\x6f\xa3\xd1\xeb\x35\xd3\xa0\x62\xcd\xe0\x65\xdf\x54\x2e\x5e\xeb\xef\x15\xaf\x72\xf1\x0f\xa4\xd3\x55\x59\xfb\x09\xf5\x0b\x5d\x11\xd4\xe0\xc4\x64\x9f\xbc\x80\x54\x6f\x8c\x48\x9b\x85\x2f\x24\xd6\x15\x78\x1f\x49\xd3\x68\x69\x2d\xd7\x9b\xe3\x39\x48\xf7\xac\x47\x72\x7d\x2b\xdf\x53\xae\x9b\x84\xd8\x57\xac\xb7\x5a\xbf\x8a\x5a\x6f\x8f\xdd\x88\xd5\x03\xeb\xf5\x43\xe0\xfe\xec\xfe\xbd\x15\x7b\x77\x2e\xbd\x2a\x3d\xb0\x66\xdf\x39\x30\x15\xed\x8d\x96\x87\xaa\xfd\xfd\xd5\xec\xea\xec\x40\xb5\x5f\xbc\xbf\xd2\x7c\xd3\xa6\x39\x40\xa3\xd2\xbd\x80\x6e\xaf\xf4\x75\x73\xe1\xde\x66\x6c\xad\xdc\x5b\xa3\x32\x3b\xa9\x86\xd4\xee\x31\x64\x1b\x34\x88\x74\x37\x55\xee\x83\xcb\xf6\x4a\x2a\x7a\xeb\x5c\x35\x78\x09\xd1\x7e\xe1\x1e\xc8\xce\xea\x7b\xed\xbf\xb7\x64\xaf\xd5\x4b\x83\x81\xd5\xf9\x2f\x2c\xd8\x9b\xf4\xfa\x1f\x52\xae\xcb\x9a\xf6\x53\xeb\x7a\xed\xf9\xd0\x81\xc1\xde\x78\x01\xa1\xde\x14\x8f\x2e\xf7\x5e\x48\xa6\x4b\xec\x3e\x3a\xa6\xc9\xd0\x5a\xa4\x37\x46\x73\xa8\xd1\xb5\x3f\xbf\xae\xe1\xf5\x2c\x7c\x5f\x95\x6e\x16\x64\x5f\xa1\xde\x66\xfc\x2a\x3a\xbd\x35\x72\x13\x52\x0f\xac\xd2\x0f\x70\x7b\x93\xfb\xf7\xd6\xe8\x9d\x99\xf4\xa9\xf2\xc0\x0a\x7d\x8b\x6f\x2a\xd0\x9b\x0c\x0f\xf5\xf9\xd5\xec\x6a\x76\x75\xa0\xcf\x2f\x75\xbf\xbd\xd5\x18\x9d\x49\xd5\x5e\x40\x9c\xef\x9a\xb9\xb9\x36\x6f\xb1\xb5\x96\xe6\x6d\x31\x19\x9d\x4d\x43\x0a\xf3\xfb\x08\xf3\x5e\xa7\x52\x93\x61\xeb\xa9\x14\x86\xa1\x66\x28\xb6\xfd\xa4\x0b\xc4\xaa\x9f\x84\x61\x08\xde\xe0\x24\xa5\x8c\x43\xc2\x07\xc8\x48\x63\xd9\xbb\x30\xda\xfa\x8a\x59\xc9\xc5\xf6\xeb\x79\xbc\x74\x61\xb4\x6c\x67\x42\x09\xea\x0e\x30\x27\x45\xb6\xca\x8d\x09\x13\x9a\x4d\xed\x28\xba\x87\xd9\x57\x4a\x74\xc3\xbc\x2a\x4d\x35\xb2\xd2\xe0\x41\x37\x8a\x15\x55\xd7\x31\xf4\x6f\xfb\x74\x84\x26\xc3\x76\xf9\x0c\xc5\x8f\x66\x34\xb6\xbd\xaa\x0b\xc4\x4e\xfb\xc8\x4c\x8c\x78\xd0\x99\x94\x06\x0b\xba\x30\x5a\x65\x90\x69\xe1\x2d\x3a\x56\x17\xc6\x40\x1d\x4b\xba\xe9\xd7\x16\xaa\xa6\xd6\x5c\xdd\x83\xb5\x6f\x5a\x6d\x30\xaf\xcd\x57\x8d\xc4\x8c\xfa\x56\x1b\x8a\x2d\x67\x63\xa9\xa6\xfb\xbc\x38\x6a\x30\x6c\x3f\xc0\x2e\xc5\x8f\x66\x34\xd6\x2f\xdb\x3a\x40\xec\xce\x2f\x99\x89\x11\x15\x3a\x93\xd2\x11\xd8\x1d\x18\xad\xc7\x97\x69\xe1\x2d\x5a\x57\x17\xc6\x40\xad\x4b\xba\xe9\xd7\x1c\xaa\xa6\xd6\x5c\xdd\x83\xb5\x6f\x5d\x6d\x30\xaf\xcd\x57\x8d\xc4\x8c\x5a\x57\x1b\x8a\x2d\x67\x03\xc8\x7a\x89\xae\x06\xbb\x56\x32\x9c\xbf\x17\x3f\x7a\xb1\xd8\xf6\xad\x0e\x0c\x2b\x1a\xa8\x3c\x8c\x68\xd0\x95\x92\x06\x07\x3a\x20\xda\x08\x60\x5a\x74\x8b\x9e\xd5\x01\x31\x50\xcb\x12\x5e\xfa\xf5\x84\x8a\xa5\x35\x47\xeb\xa8\xf6\xfd\xaa\x05\xe5\xb5\x79\xda\x9d\x96\x51\xb7\x6a\x01\xb1\xe5\x6a\xca\x70\x02\xd9\x63\x9f\xbe\xd0\x6c\x6a\xfd\x29\x4b\x01\x6b\xdb\xb5\xba\x61\x06\x78\x9b\xdc\x88\x10\x1a\x89\x69\x30\xa2\x1b\xc5\xf6\x5d\xf3\x8a\x07\x8b\x0e\xd6\x8d\x32\x50\x13\x2b\x1c\xf5\xeb\x15\x75\x63\x6b\xee\x1e\x00\xdb\x77\xb3\x76\xa0\xd7\xe6\xaf\x56\x72\x46\x3d\xad\x1d\xc7\xfa\x93\x1f\x4c\x7a\x69\xb0\x06\xbb\xf6\xf3\xed\xec\xc3\x79\xe0\xeb\xc5\x62\xff\xda\xf1\x59\x0c\xbb\xb3\x4d\xe6\x61\xf8\xd2\xf1\xf9\x94\xb4\x5e\x95\x3d\x0b\xd1\x7a\xae\x19\x16\xdd\xea\x75\xe3\xb3\x10\x83\xbd\x6c\x24\x3d\x35\x58\xc5\xd2\x9a\xa3\x75\xd4\x21\x5e\x33\x36\xa2\xbc\x36\x4f\xbb\xd3\x32\x7c\xc5\xd8\x08\x62\xcb\x55\x4c\x42\xda\xa7\x29\x34\xd8\xb5\x12\xe1\x6c\xfa\xd1\x47\x1a\xc4\x14\x98\xb6\xcd\xaa\x03\xc3\x8a\x04\x2a\x0f\x23\x12\x74\xa5\xa4\x41\x81\x0e\x88\x36\x02\x98\x16\xdd\xa2\x59\x75\x40\x0c\xd4\xac\x84\x97\x7e\x0d\xa1\x62\x69\xcd\xd1\x3a\xaa\x7d\xb3\x6a\x41\x79\x6d\x9e\x76\xa7\x65\xd4\xac\x5a\x40\x6c\xb9\x9a\xe5\xbe\x8f\xb2\xac\xd7\x6f\x03\x37\x9a\xb6\xd3\xe1\x3c\x98\xbd\xd7\xf9\x5b\x00\x0a\xd6\xb6\x6b\x75\xc3\xd8\x11\x42\x66\x63\x44\x08\x8d\xc4\x74\x7e\xb1\xb4\x13\xa5\x95\x12\xe6\x0b\x60\xd1\xc1\xba\x51\x06\x6a\x62\x85\xa3\x7e\xbd\xa2\x6e\x6c\xcd\xdd\x03\x60\xfb\x6e\xd6\x0e\xf4\xda\xfc\xd5\x4a\xce\xa8\xa7\xb5\xe3\xd8\x72\xf8\x1e\x32\x82\xc9\xa6\xd7\xd7\xa7\x1a\x4d\x9f\xf9\xd2\x4c\x10\x5c\x7e\xd0\x8e\xc8\xb6\xad\x75\xc3\x58\x7e\x6f\x46\x64\x63\xf6\x0d\xaf\xee\xc4\x74\xbe\x42\xd5\x89\xd2\xfe\xd5\x19\xe3\x05\xb0\xf9\xaa\x57\x27\xca\x50\x5f\xf6\x52\x8e\x7a\x7e\xa7\xaa\x66\x6c\xcd\xdd\x03\x60\xfb\xb6\xd6\x0e\xf4\xda\xfc\xd5\x4a\xce\xec\xcb\x5f\xad\x38\xb6\x1c\x0e\x20\xd9\x20\xd6\xef\x53\xbf\x06\xcb\x67\x88\x71\x7e\xa5\xd3\x66\x15\xaa\x6d\x4f\xeb\x44\xb1\xa4\x84\xc8\xc5\xf0\x33\xc9\xae\xb4\xb4\x3e\xf3\xeb\x00\x69\x27\x83\x69\xf1\xad\x3e\x99\xec\x00\x19\xec\xb3\x49\xe1\xa7\xef\x27\x80\x15\x5b\x6b\xce\xee\xe3\xda\xf7\xb2\x56\x9c\xd7\xe6\xad\x4e\x6a\x86\x9f\x52\xb6\xc0\x68\x70\xb7\x76\x77\x87\xda\x1d\x2a\x3a\x6e\x8a\x01\xc9\x26\x8f\x21\x73\xbf\x64\xde\xcc\x3d\x73\xcf\xa6\xdb\x2b\x09\x26\x7b\x77\xad\x58\x78\xea\xe6\x23\x47\x8b\x35\x0d\x1e\x8b\xfb\x63\x80\x1f\xf0\x03\x0a\x00\x81\x77\x6b\xc8\x80\xba\x9d\x05\x81\x77\xc0\x8f\x61\x96\x2d\x9d\xe2\x7a\x04\xb3\x82\xcd\x00\x67\x93\x50\x98\x4c\x38\x4d\xc1\x76\x53\x38\x80\xd1\x18\xc9\xf9\x78\x03\x39\xa6\xc4\x01\x90\x61\x38\x91\x25\x58\x3a\x09\xc4\x04\x54\x06\x57\x47\xa3\xd1\x22\xc0\x5b\x3f\x3e\x25\x1c\x62\x82\x98\x1c\xa9\x0d\xa9\x10\x26\x6b\x06\x49\xa0\x46\x47\x0b\xb8\x37\x88\x39\x4a\x9c\xf2\x46\x22\xc5\xa4\xd1\x62\xbd\xbd\x05\xca\xc2\x5b\x17\x96\x1e\x5c\x1d\xa9\x47\xc2\x85\x0a\x7a\x9d\x73\x2e\x02\xde\x73\x98\x33\x99\x58\x2d\x0d\x44\xf2\xe2\x0a\x7a\x48\x21\x09\x50\xb0\x74\x42\x18\x67\xc8\x01\x01\xe4\x70\xc2\x21\xdb\x20\x5e\x62\xfc\x5d\x4c\x2f\xa3\xc9\x52\x48\x94\x69\x84\x83\x00\x91\xa5\xc3\x59\x8e\xe4\x02\xa5\x90\x0c\x3f\x6b\xe1\x05\xf8\x4e\x55\x53\x3d\x6a\x29\xac\xca\x09\x07\x0d\x41\x37\xcc\x46\xdb\x45\xe8\x5c\x85\xca\x4d\x67\xa0\x8e\xc9\xdb\x08\xc5\xa9\xb3\xfa\x2b\x8a\xd3\x8a\xc5\x61\x08\xd2\xa6\x18\x1e\x2d\xe4\x1f\x8f\x5a\x1d\xdf\xdc\xb8\xff\x83\x58\x86\x29\xf9\xfc\xf9\x64\xe1\xa9\xab\x05\xc4\xb6\x10\x8d\x35\x29\xaf\x2d\x3c\x02\xe5\x95\xaa\xcb\x08\x31\x2a\x29\x9f\xc7\xb1\xfa\x35\x27\x45\x5d\xb1\x73\xfe\x82\x36\x98\x80\x14\x6e\x50\x79\x27\x1c\xb5\x7f\x46\x8b\xf2\x86\x2e\x05\x48\xf1\xf4\x90\xda\x3e\x8d\xf3\x84\x64\x0d\xd5\x56\x23\x4e\x43\x11\xf6\x76\x8a\x18\x8d\xce\xca\x41\x79\x73\x1f\x67\xf5\x13\x65\x3c\x03\xbf\x70\xc8\xf3\xcc\x75\xdd\x85\x17\x9d\x6d\x27\x73\xb8\x8e\xd1\x76\xbe\x78\xe2\x00\xb2\x99\x08\x58\x46\xe3\x18\xb1\xa5\x23\x7a\xe8\x77\x9c\xc5\x5b\x0f\xa3\x05\x2f\xee\x5b\xb4\x7d\xce\x76\x4f\xc4\x28\x50\xb7\x00\x72\xce\xbf\x75\x56\x6f\x17\x1e\x8f\x9a\x87\x2f\xa6\xdf\xaa\xe8\xc0\x3f\x60\x82\x3a\xe6\xa9\xf8\x0f\x27\xad\xea\x97\x16\x5e\x25\x18\x31\x54\x8d\x74\xc1\x8b\x5e\xb7\x0b\x5c\x64\x9b\x45\xf4\x7e\xe9\xe0\xec\x53\x92\xf2\xc7\x63\x91\x6f\x76\xe2\xd4\xbc\x04\xc0\xa7\xb1\xd8\x4c\x4b\xe7\xa2\x3a\x32\x5a\xa4\xdb\xe2\xa1\x07\x5e\xb4\x3f\x20\x1f\xfb\x88\x70\xb1\x2e\xff\xa0\x80\x21\x9f\xb2\x20\x03\xa1\xe8\xf5\x6f\x16\x5e\x5a\x05\xf7\x78\xd0\x12\x7e\x19\x1f\x43\x29\x82\x5c\xad\x04\x10\x24\x93\x0b\xca\x19\xf4\x6f\xc1\xfa\x11\x7c\x83\x49\x80\x1e\x9c\xaa\x59\xb0\xfa\xfa\x55\x5d\x7e\x37\x7b\x7a\xda\x73\xc1\x03\x01\x8a\xc3\xa5\xf3\x46\x20\xb9\x28\xc0\xdc\x59\x7d\xfd\x2a\x9f\x10\x98\xa0\x67\x2c\x2a\x06\x0b\x4c\xd2\x9c\x97\xe9\xcb\x27\x0e\x90\x47\xa2\xac\x85\x03\xee\x60\x9c\xa3\xa5\x53\x03\x96\xec\x4a\x68\x20\xda\xe7\xf6\xb2\xe8\x54\x7b\xfe\xaa\x05\xaa\xd0\x3d\xc4\x28\x0e\xea\x0b\xa0\xa2\x10\x0d\x4b\x9d\xca\xdf\x09\x11\xf0\x8b\x7a\x7b\xe4\xff\xca\x2a\x08\xc7\x2a\xb4\xf2\xb8\x76\x8a\xfb\x57\x75\x18\x95\x7b\x56\xce\x02\xbb\x37\xe7\x0e\xf2\xc8\x24\x3d\xd5\xe6\x89\xb1\x7f\x2b\xf6\x39\x57\x9c\x3d\x56\x70\x27\xf5\xb8\xe5\x19\x02\x42\xca\xba\x62\x28\x56\xa6\xc0\x12\x0f\x4f\xc4\x02\x49\xfb\x1a\x8f\xb6\x7d\xad\x81\x56\xcd\x6b\xbe\x6b\xc0\xea\xd0\x03\xc5\xe7\x27\xa0\x22\x5b\xaa\x29\x6d\x4d\x97\xc5\x01\x23\xcf\x97\x4f\x01\xe6\xc5\x51\x23\x3a\xb5\x16\x77\x1a\x1c\x17\x95\xed\xf6\x2d\x4f\xd8\xeb\x0c\x71\xd1\x35\x8a\xe2\x9e\x82\x2d\x9b\x4e\xca\xb8\x7e\x81\x77\xa8\x3d\xae\xbd\x3e\x51\x6d\x0c\x0b\x4f\x76\x42\x8d\x03\x43\x3c\x28\xfa\xb9\x3a\x3d\x42\x4a\x85\xbc\x2b\xe9\x2a\x9f\x39\x40\xaa\xb9\xa5\x53\xfc\x71\x84\xf9\xac\x3c\xb1\x9e\xeb\xe4\x8b\x74\x77\xcb\xb6\xfd\xf3\x2c\xdd\x8f\x41\x39\x92\xc7\x56\x71\x90\x2d\xb6\x77\x27\x1b\x1d\x87\x39\x51\x27\xd0\xf1\x89\x12\x9e\x77\x90\x01\x25\x68\xc0\x12\x04\xd4\xcf\x13\x44\xb8\xfb\x5b\x8e\xd8\xe3\x2f\xf2\x1e\x52\x94\x1d\x8f\xdd\x9a\xf2\x19\xab\x3f\x43\x21\x2c\x85\x44\x78\xc6\xee\xed\x18\xbc\x2b\xd0\x5d\xa1\x82\x32\xc4\x5d\x25\x84\x8a\xbf\x64\xa1\x86\x60\x10\x7c\xba\x43\x84\xff\x37\xce\x38\x22\x88\x1d\x8f\xe5\x42\x8f\x4f\xc1\x41\xb8\xa5\x89\xac\x94\x98\xef\x72\xba\xd9\xc4\xe8\x78\x8c\xb3\x09\xf4\x39\xbe\x43\x45\x78\x23\x11\x9b\xc6\xbc\x27\xf9\xcf\xd3\xc9\xf1\xc9\xb5\x58\x36\x91\x15\x4c\x53\xb0\x04\x5b\xbd\x4c\x83\x5c\x18\x96\x77\xeb\x1b\x9f\x82\x9b\xcf\x6a\x32\x4c\x53\x77\x77\x3e\x1e\x8f\xcb\xf3\xb1\x16\xf9\x37\x99\x4f\x53\x74\x0a\xbe\x89\x38\x4f\x4f\xc1\x37\x1c\x27\x88\xe6\xbc\x48\x48\x8d\xba\x7f\x87\x3c\x02\x4b\x70\x8f\x49\x40\xef\xe5\xb3\xeb\xca\xa8\xea\xf2\x4b\x70\xf3\x59\xba\x2d\x2f\x6f\x10\xff\x1e\x72\x08\x96\x0d\x75\x12\x79\x08\x87\x3f\xa3\xdf\x72\x94\x71\xb0\x54\xfe\x8f\xd5\xe8\x28\x41\x3c\xa2\xc1\x1c\x8c\xff\xeb\xd3\xaf\xe3\x53\x75\x2d\x67\xf1\x1c\x8c\x3d\x98\x62\xef\x6e\xe6\xa9\x2e\x36\x56\x9a\xf8\xe9\xc4\x2d\xf6\x64\x85\x41\x62\x45\x4f\x81\x9a\x57\xba\xdd\x8f\x58\xcc\xb9\x2e\x20\x54\xb9\x9b\x32\x28\x6a\x5f\x5e\x95\x7f\x40\xe7\x0e\xc6\x3f\x94\xae\x9a\x12\x2c\xeb\x78\x48\xe9\x26\xf4\xea\xe5\x7d\xf8\x72\xfc\xe9\x14\x5c\x4e\xa7\xc5\xdd\xe1\x6a\x01\xed\x9a\x6e\x2d\x94\x7a\xea\x38\xdc\x5d\x61\x88\xe7\x8c\x00\xe7\x47\xe2\x28\xa0\xd1\xf6\x4a\x18\x3a\xaa\x10\x55\x07\x45\x07\xab\xa0\x1f\x17\xbd\x4c\xb6\xb1\xc2\x43\x9f\xf5\x44\xc4\xa7\x01\xfa\xe7\xcf\x7f\x3b\xde\xad\xac\x72\xe6\x8d\xdf\x29\x59\x30\xf6\xc6\xef\xa4\x9b\x21\xd7\x5a\xfc\xd3\x94\xe6\x41\x19\x55\xa2\x03\xa7\xa8\x36\x7b\x99\xe1\xb0\x89\x1d\x64\x55\xa8\xc5\x1a\x35\xe8\xfa\x4b\x95\x17\x74\xfd\x05\x2c\x97\x80\xe4\x71\xbc\x25\x87\x38\x37\xaf\xab\x33\xdc\x18\x91\x0d\x8f\xc0\x0a\x4c\xb7\x93\xd4\x01\xd7\x34\x6b\xb9\x5c\x56\xe6\xed\xc0\x42\xca\xc0\xb1\xd8\xfd\xb7\xe8\x51\x28\xc4\x4a\x24\x12\x22\x82\xd9\x8f\xf7\xe4\x27\x46\x53\xc4\xf8\xa3\xeb\xc3\x38\x16\xb0\xa7\x62\xfe\x49\x93\xdf\xa7\x2a\x7f\xb7\x7e\x54\x05\x3c\xef\xd9\x3d\x25\xf7\x7c\xf5\x7d\x06\x75\xb4\x1e\x2d\xe4\xcd\x33\x57\xff\x1f\x00\x00\xff\xff\x2b\xb3\xff\xe8\x1e\x75\x00\x00")

func templatesDashboardHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesDashboardHtml,
		"templates/dashboard.html",
	)
}

func templatesDashboardHtml() (*asset, error) {
	bytes, err := templatesDashboardHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/dashboard.html", size: 29982, mode: os.FileMode(420), modTime: time.Unix(1529590546, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4e, 0x7b, 0x55, 0x3e, 0xf0, 0x15, 0x19, 0x48, 0x28, 0xfb, 0xa9, 0xe2, 0x79, 0x3c, 0xc5, 0x16, 0x7a, 0x73, 0x4b, 0x71, 0x2c, 0x11, 0x88, 0xa1, 0xbb, 0xb8, 0x40, 0x4d, 0x78, 0xee, 0x5, 0xe0}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"templates/dashboard.html": templatesDashboardHtml,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
		"dashboard.html": &bintree{templatesDashboardHtml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
