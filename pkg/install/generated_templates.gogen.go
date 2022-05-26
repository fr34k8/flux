// Code generated by vfsgen; DO NOT EDIT.

package install

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// templates statically implements the virtual filesystem provided to vfsgen.
var templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"/flux-account.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-account.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 826,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xa4\x52\xbb\x8e\xdb\x30\x10\xec\xf9\x15\x03\xb8\x70\x12\x58\x0a\xd2\x05\xea\x6c\x17\x29\x12\xa4\x50\x1e\x4d\x90\x62\x45\xae\xce\x3c\xd3\xa4\xc0\x87\xef\x21\xe8\xdf\x0f\x92\x7c\x07\xcb\xf6\x1d\x60\x5c\xc7\xdd\x9d\xe5\xce\xce\x4e\x96\x65\x62\x86\xdf\x1b\x46\x60\xbf\xd7\x92\x41\x52\xba\x64\xe3\x02\xd2\xa4\x10\xd9\xc3\x3b\xc3\x61\x01\xb2\x6a\x92\x42\xa5\xad\xd2\xf6\x06\xe4\x59\xcc\xe0\xac\x79\x80\x65\x56\xac\x50\x3b\x8f\xef\xa9\x62\x6f\x39\x72\xc0\x9d\x8e\x9b\xa1\x25\xab\x28\xb0\xea\x27\x70\x08\x90\xce\x46\xef\x0c\x3e\x94\xab\xe5\xfa\x63\x2e\xa8\xd1\x7f\xd9\x07\xed\x6c\x81\xfd\x17\xb1\xd5\x56\x15\xf8\x35\xb2\x5a\x8e\xa4\xc4\x8e\x23\x29\x8a\x54\x08\xc0\x50\xc5\x26\xf4\x2f\xc0\xd2\x8e\x0b\xd4\x26\xdd\x8b\xe3\xa0\x6d\xa1\x6b\xe4\x3f\x69\xc7\xa1\x21\xc9\xe8\xba\x43\x7d\x08\x0b\xb4\xed\xb4\xda\xb6\x60\xab\xba\x4e\xf4\xba\x1c\x13\xf2\x15\xc9\x9c\x52\xdc\x38\xaf\x1f\x29\x6a\x67\xf3\xed\xd7\x90\x6b\xf7\xf9\x85\xea\x7a\x14\xa7\x74\x86\xaf\xe5\x29\x7c\x32\x3c\x40\x32\x50\xa3\xbf\x79\x97\x9a\x50\xe0\xdf\xfc\xd3\xfc\xff\xd0\xe7\x39\xb8\xe4\x25\x4f\x92\x7b\xf6\xd5\x51\x22\x83\x75\xb6\x3c\x00\xff\x94\x3f\x5e\xc7\xbe\x6f\xb9\xd5\x78\xf7\xeb\x77\x74\x86\x4b\xae\x7b\xd0\xf3\x8e\x6f\x8c\x16\xc0\xb9\xac\x93\xff\x42\xaa\x6e\x59\xc6\x83\x6c\x17\xed\x72\x46\xe7\xf4\xf8\xa7\xee\xb8\xe4\x07\x13\xfa\x97\xe2\x9a\x92\x89\xa3\x41\x7a\x1f\x3d\x05\x00\x00\xff\xff\xdb\x2d\xc3\x7c\x3a\x03\x00\x00"),
		},
		"/flux-deployment.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-deployment.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 7263,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xd4\x59\xdd\x6f\x1b\x37\x12\x7f\xf7\x5f\x31\x50\x0e\x48\x0c\x48\x2b\x3b\x6e\x8b\xc3\xf6\x5c\x5c\x9a\x0f\x37\x97\x26\x35\xec\xe4\x0e\x7d\xaa\x29\xee\x48\x4b\x88\x4b\xee\x71\xb8\x52\x17\x46\xff\xf7\xc3\x90\xfb\xc1\x95\x65\xa7\xe8\xdb\xf9\x21\xb1\xc9\xe1\x70\xe6\x37\xdf\xdc\xc5\x62\x71\x22\x6a\xf5\x6f\x74\xa4\xac\xc9\x41\xd4\x35\x2d\x77\xe7\x27\x5b\x65\x8a\x1c\xde\x60\xad\x6d\x5b\xa1\xf1\x27\x15\x7a\x51\x08\x2f\xf2\x13\x00\x23\x2a\xcc\x61\xad\x9b\xdf\xef\xef\x41\xad\x21\xfb\x24\x2a\xa4\x5a\x48\x84\x3f\xfe\xe8\xf6\xc3\x9f\x39\xdc\xdf\x4f\x77\xef\xef\x01\x4d\xc1\x64\x54\xa3\x64\x66\x0e\x6b\xad\xa4\xa0\x1c\xce\x4f\x00\x08\x35\x4a\x6f\x1d\xef\x00\x54\xc2\xcb\xf2\x67\xb1\x42\x4d\x71\x21\xbd\x9b\xa9\xbd\x13\x1e\x37\x6d\xdc\xf4\x6d\x8d\x39\xdc\xa0\x74\x28\x3c\x9e\x00\x78\xac\x6a\x2d\x3c\x76\xcc\x12\x0d\xf8\x47\x18\x63\xbd\xf0\xca\x9a\x81\x39\x40\xed\x6c\x85\xbe\xc4\x86\x32\x65\x97\xb5\x75\x3e\x87\xd9\xc5\xd9\xc5\xf9\x0c\x9e\x81\x47\xad\x13\x0a\xf0\x16\x48\x3a\x51\x23\x2c\x2b\xf4\x4e\x49\x62\xe5\x6a\xab\x8c\x7f\x4e\xc0\x87\xb3\x8e\xb1\x9e\xe8\x70\xa0\x05\x40\x8f\x45\xd8\xb2\x05\xde\x4e\x50\xe0\x9f\x15\x7a\x91\x6d\x9b\x15\x3a\x83\x1e\x83\x70\x96\x72\xd0\xca\x74\x2c\x18\x3a\xb7\x53\x12\x5f\x49\x69\x1b\xe3\x3f\x4d\x6f\x00\xd8\x59\xdd\x54\x38\xc8\xb0\xe8\x64\xd8\x28\xbf\xd8\x62\x3b\x5c\x44\x0c\x9f\x1f\x2f\xee\x57\x46\x7e\x0b\x3e\x52\x04\xcf\x48\xa8\x0a\x5c\x8b\x46\xfb\x8f\xb6\xc0\x1c\xce\xbe\x39\x3b\x83\x67\xb0\x2f\xd1\x40\xc5\xd2\x60\x01\x0e\x45\xb1\xb0\x46\xb7\x73\xd8\x23\xec\xad\x79\xee\x61\x85\x20\x56\x1a\x19\x48\x59\x56\xb6\x38\xe9\x18\x3e\x83\xcf\xa5\x22\x50\x04\x02\x7c\x55\xaf\x09\x1a\xc2\x02\xd6\xd6\xc1\x06\x0d\x3a\xe1\x95\xd9\xc0\xed\xed\x4f\xb0\xc5\x96\x32\x78\x6f\xe0\xc3\xdf\x09\x7e\xb8\x84\xf3\xec\xfc\x6c\x3e\x70\xe9\xef\x8e\x2a\x10\x08\x87\xa9\x1c\x64\x59\x14\x83\x58\x80\x00\xc2\x5a\xb0\x37\x75\x40\xc1\x1e\x07\x36\x52\x18\xd8\x3b\xe5\x59\xd0\xec\x38\x7e\x1b\x34\x03\x18\x58\xd5\xbe\x7d\xa3\x5c\x0a\x62\x85\x85\x6a\xaa\x1c\x3e\x62\x65\x5d\x9b\xea\x89\xb0\xb6\x5a\xdb\x3d\x6b\xd4\x5d\xad\x28\xa8\xda\x10\xaf\x09\x90\x0d\x79\x5b\x29\x46\x60\x6b\xec\xde\xfc\x56\x5a\xf2\x34\xb0\x58\x2b\x8d\x73\xd8\x97\x4a\x96\xd0\xda\x06\xf6\x4a\xeb\xa8\x94\xb7\x50\x58\x0e\x50\x5e\xe6\x43\xfc\x8b\x03\xbb\x37\x2c\xf6\xc0\xc0\x61\x6d\xc1\x09\x5f\xa2\x03\x5f\x0a\xd3\x5d\xbc\x51\xbe\x6c\x56\x60\x79\x11\x41\xab\x2d\x66\xf0\xab\x6d\x9e\x6b\x0d\x42\x93\xed\xaf\x98\x82\x0d\xca\x83\x32\xde\x86\x33\xd2\x1a\x2f\x94\x41\x37\x87\x15\x6a\xbb\xcf\xe0\x16\x47\x54\x4b\xef\x6b\xca\x97\x4b\xf6\x29\x59\xb0\x47\x6b\xdc\x08\xd9\x86\x85\xe5\xa6\x51\x05\xd2\xb2\x21\x5c\xd4\x4e\xed\x84\xc7\xe0\x77\xac\xc5\x72\x60\xd1\x1b\x81\xa8\x5c\x48\x6b\xd6\x6a\x33\x6c\x01\xc4\x85\x8f\xa2\xce\x93\xc5\x34\xf4\x16\xc9\xb1\xbf\x6a\x90\x10\x93\xcb\xc8\x64\xf4\xbb\xaf\x1a\x63\xaf\xa8\xe4\x95\x52\xec\x10\x04\x14\x6a\xbd\x46\xc7\x69\xb6\xe7\xd0\x85\xd3\x98\x4a\x03\xf6\x91\x5d\x8a\x3e\xa7\xa3\x9d\x2a\xb0\xc7\x7b\xad\x36\x95\xa8\x47\x41\x94\x2f\x41\x18\x40\xe3\x5d\x1b\x74\xb8\x8b\x44\x77\x73\x10\xa6\x80\xc6\x48\x5b\x71\x7e\x0f\xe7\xa3\xb6\x1f\x83\x1d\x85\x29\x06\x2e\x68\x76\x81\x83\x42\xea\x0c\xf9\xc0\x02\x0c\xc3\x5f\xb0\x40\x72\xec\xab\x16\x08\x29\xc0\x5b\x50\x15\x67\x56\xb8\xba\xbe\x0a\xd1\x0f\x2f\x58\x2d\x52\x1b\xa3\xcc\x78\x39\x2b\xb7\x43\xa7\xd6\x4a\x86\x14\x0f\x75\xe3\x6a\x4b\x48\xa7\x7f\x02\xc8\x81\x4b\xcc\x1b\x11\x45\x06\x88\xef\xfb\x13\xc0\x81\x70\x9b\x31\x3e\x1f\x41\x6c\x53\x6f\x38\x71\x50\x02\xcd\x34\xf7\x3e\x7b\x24\xfb\x3e\x3c\x77\x24\xfb\xf6\x70\x0e\x21\xf8\x20\xf1\x27\xa5\xa1\x43\xdd\x61\x48\x90\xc6\xc2\x2c\xe7\xb2\x49\x7e\x06\xaa\x12\x1b\x8c\xde\xcf\x07\x32\x78\xa7\x4c\x11\x74\xae\x38\x9f\x38\x94\xa3\xd7\xc6\x5c\xa2\x51\x10\x72\xd6\x08\x47\xd9\x08\xdc\x59\x80\xf0\x43\xc0\x97\xcd\x2a\x2b\xac\xdc\xa2\xcb\xa4\xad\x96\xae\xcb\x00\x31\xee\xbd\x18\xa0\xeb\xed\xc8\x1d\x02\x77\x0f\x7c\xab\x17\x1b\x60\x49\xb3\x81\x26\x5c\x93\x43\xc7\x50\xd9\x94\x5b\x7e\x9e\xbd\xfc\x36\x7b\x39\xa5\xbd\x6e\xb4\xbe\xb6\x5a\xc9\x36\x87\xf7\xeb\x4f\xd6\x5f\x3b\xa4\x54\x0b\x87\x64\x1b\x27\x91\xd2\x04\xee\xf0\xbf\x0d\x92\x9f\xac\x01\xc8\xba\xc9\xe1\xdb\xb3\x6a\xb2\x58\x85\x1c\x9f\xc3\x77\xdf\x7c\x54\x63\x63\x61\x5d\x7a\x78\x31\x5a\xe6\x3a\x34\x19\x17\x67\x17\x5c\x32\x95\x59\x5b\x57\x05\x97\x15\x7a\xa0\xd6\x6a\x87\x06\x89\xae\x9d\x5d\x61\x2a\x01\x43\x7a\x35\x2d\xd7\xf1\xaa\xc8\x70\xba\x2c\x7c\x99\xc3\x52\xd4\x2a\x22\xbd\xfb\x6e\xa9\x0a\x34\x5e\xf9\x36\xab\x9b\x55\x42\xab\x8c\xf2\x4a\xe8\x37\xa8\x45\x7b\xcb\xf1\x59\x50\x0e\xdf\x26\x04\x5e\x55\x68\x1b\x7f\x64\x8f\xab\xab\xfa\xff\x10\x35\x09\xda\x89\x61\x8e\xf7\x45\x10\xeb\xdb\x75\x94\x0c\xbd\x0c\x92\x15\x4b\xa2\x92\x3b\x43\x1b\x7b\x55\xd0\xb6\xcb\x37\x1b\x36\x19\x28\x13\x7d\xee\x39\xc5\x33\x44\xe5\x72\x92\x26\x7b\xcc\x7e\x31\xba\xcd\xc1\xbb\x06\x99\x1b\x37\x3f\x21\x43\xad\xba\xc4\xce\x21\x55\xa3\x5b\x5b\x27\x91\x99\xc6\x6e\x87\x9b\x9d\xc7\x04\x4f\x1b\x92\xa9\xec\x3b\xe1\x3a\xd9\x23\xd9\x5f\x13\x3f\x89\xd1\xf7\x46\xea\x26\x64\x4e\xee\xd9\x62\x81\xeb\xb3\x6a\x6c\x0a\xbe\xd2\xc3\xf4\x5d\xcc\xf7\x7c\xf4\xa0\xbf\x18\xb2\x2b\x14\x28\xb5\x70\xdc\xab\xad\xec\x2e\x49\x00\x4f\xb4\x01\x31\x3d\xa6\xca\x3b\x6b\xfd\x32\x23\x2a\x1f\x55\x40\x98\xc9\xad\xb3\xb1\x44\xcd\xe2\xcd\xf3\x9e\x24\xe1\x80\x66\xa7\x9c\x35\xa1\x20\xc4\x5a\x3b\xfb\xf0\xe5\xc7\xb7\xaf\x7f\xf9\xf4\xee\xfd\xd5\x2c\x96\x80\x39\xe3\x61\x77\xe8\xdc\xb4\x5e\x27\x6c\x42\x89\x5b\xb5\xb1\x9a\x7a\x7d\x4c\xc7\x07\x85\xf6\xa1\x8e\xa3\x73\x32\xf1\xa3\x8a\x72\xcd\xe3\x51\xa5\xbf\x8d\x53\x74\xd2\x8a\x74\xd2\x05\x9b\x24\x2c\x0e\x1b\x9a\xd4\xe8\xa1\x9b\xe9\x7b\x6e\x61\x40\x68\x8f\xce\x70\x4f\xfd\x40\xe2\xb5\xb3\x15\xbb\x45\xdf\xb1\xcc\x41\x10\xbb\x5b\x57\x55\x19\x06\x6d\xe5\x96\x1e\x1a\x1b\xcd\x2e\x3f\x82\xcb\x08\xf7\x04\x97\x9d\xd0\x0d\x3e\xc0\xe4\x6b\x4e\x7c\xe8\x03\x7d\xcd\x7d\xc2\x03\xb8\xe4\x4f\x4b\xfd\x13\xc5\xfe\x11\xbf\x64\xaa\xd8\xdd\x4c\xe8\xa6\xf9\xe1\x6b\x91\xb7\x17\xdc\x94\x58\xa0\xa6\xae\x75\x0b\x3f\x7d\xfe\x7c\x0d\x2b\x41\x4a\x82\x68\x7c\x09\xd2\x61\xc8\xa4\x42\xc7\xaa\x3e\x0e\x02\xcc\x70\xa7\x44\x50\xfc\xee\xea\xfd\xe7\xdf\x5e\x7d\xf9\xfc\xd3\x97\xdb\xb7\x37\x77\x41\xdd\x61\xe9\xc3\xdb\x5f\xef\x26\x0e\xbf\x13\x4e\xf1\x18\x47\x7d\x83\x9c\x30\x8c\xed\xcb\x81\xfd\xde\x39\x5b\x4d\x6d\x18\xc9\x6e\x70\x9d\x4f\x34\x9f\xf4\x8a\x9c\xd8\x58\x85\x11\x00\xc6\x3c\x9f\xe0\x11\x21\x88\xc3\x29\x16\x5c\x89\xa5\x90\x25\x16\xec\x5a\xa9\x6f\x0f\x6d\x35\x23\xc5\xdc\xe7\x09\x17\xeb\xba\xbe\x39\x39\xd0\x0d\xd7\xe1\xe0\x3c\x5c\xc2\x43\x61\x87\xb1\x2f\x91\x52\x5f\x18\xbb\x57\xbf\xb7\x2c\x65\xc3\x38\x85\x88\x0b\x4f\x08\xc1\x11\xa1\xb4\xfb\x30\xf8\x5a\x63\x50\x06\x93\x29\x3f\xf5\x9d\xc5\x62\x50\x20\x0c\x3e\x7c\xf9\xe5\xb0\x94\x75\x4d\x5f\x46\x3b\x99\x49\xdd\x90\x47\x97\x71\x02\xd7\x29\x24\x5f\x28\xe6\x9a\x11\x8a\xd7\x91\xf4\xfd\xf5\x44\x29\x4e\x3b\x84\x3e\x0c\xd6\x53\xcf\x1e\x65\xe8\xe9\xd9\xbb\xbc\x63\xca\x30\xea\x26\x25\x28\x95\xb8\xa3\xbe\x3c\x99\x74\x99\x8a\xa0\x6a\x28\x8c\xfe\x01\x3d\x85\x45\x0c\xa7\x55\x28\x6c\xa1\xc7\x0b\x13\xff\x8b\x7e\x8c\x3e\x4d\x65\xe9\x93\x4b\x0c\x43\x76\xe0\x64\xf0\x9f\x08\xc2\xc5\x20\x16\xb8\x45\xa1\xdc\xe5\x83\xb2\x97\x8a\x75\x93\x74\x98\xa3\xf1\xbe\xdc\xfc\x1c\x5f\x26\x84\xd9\xc4\xbd\x2b\xe5\xc3\xb4\x4c\xca\x5b\xd7\x0e\xe9\xfa\x1d\x77\xc6\x09\xbb\xa7\x62\x8e\xdd\x26\xd1\xbd\x0b\x99\xa3\xe1\x94\xc6\x42\xdf\x3b\xff\xed\x45\x1a\x99\xa7\xf9\xf8\xf7\x87\xb7\xbf\x9e\xfe\x33\xce\xec\xa1\xad\x6e\x08\xdd\x72\x14\x36\x4b\x03\x9d\xf1\xe1\x70\x6a\x9c\xbe\xbc\xbf\x87\xec\x4a\x79\x56\x36\x3c\xde\x4d\x29\x56\x4e\x18\x59\xf6\x44\x3f\x86\xbf\xe2\x33\x9e\x5a\x87\x25\xce\x5f\x74\xec\x24\xf7\x70\x7c\xee\x36\x78\x0a\xfd\xcb\x2a\x93\x1c\x98\xcd\x67\xdd\x6b\xa0\x26\x4c\x8f\x3f\x9d\xd4\x1c\xb2\xe3\xc9\x38\x75\x55\xc2\xa8\x35\xf7\xe4\x1c\x43\xa4\x0a\x74\xd1\x1c\x07\x93\x4d\x78\x8c\xb0\x84\xd0\x98\x02\xdd\x81\x8d\x1d\x6a\xe1\xd5\x0e\x43\xcb\x49\xbd\x07\x6e\x26\x76\x3e\x88\xc9\x41\x39\x6a\x56\x85\x72\xe7\xf3\xf8\xff\xcb\xe1\x69\x73\x04\x27\x3c\x5d\x1e\x03\x27\xbc\x07\xf6\xa8\xf6\x54\x47\x18\x7c\x21\x74\xc7\xce\xb3\x71\x07\xcb\x31\x0d\x1c\x3f\xff\xb6\x12\xea\xa8\x00\xc8\x1b\x3d\x87\x9e\x6a\x7c\x9c\x3d\x6a\x0e\xe4\x54\xb2\xb7\x0c\x28\x9a\xf0\x6e\xc7\x38\x71\xc5\x56\xfe\x60\x00\x4f\xb1\xea\x6a\x5f\x57\xd9\x2e\x9f\x28\x75\xfd\x89\x8e\x17\x9f\xba\xfc\xc7\x16\x5b\x50\xc5\x0f\x03\xd9\x13\xed\x4c\x22\x15\xb3\x10\xbe\x71\x38\x79\x05\x38\x72\x57\xd8\x6e\x17\x03\x3d\x4d\xd2\x55\x9f\xad\x41\x79\x28\x05\x85\x52\x6c\x8d\x6e\x41\x48\x89\x14\x33\x7a\x89\xf1\x05\xed\x45\xff\x66\x73\xb7\x16\x9a\xf0\xee\xf4\xe4\xfe\x7e\xd1\x1b\xe2\xa6\xab\xe1\xc7\x6c\xd1\x33\x0d\xf4\x0f\xe3\xe1\x38\xd9\x11\x3b\x91\x77\x8d\xf4\x51\xde\x7d\x18\xe7\xb9\xc5\x6b\x3c\x50\x6b\x24\xac\xac\xdd\x6e\x11\x6b\xf6\xfa\x41\xd4\xd9\x46\xf9\xd9\x1c\x2a\x14\x0c\x38\x27\x34\x10\x61\xc6\xee\x02\xa1\xa9\xc9\x3b\x14\xd5\x10\x11\xa7\x07\x82\x31\xeb\x05\x79\xe1\xf1\x92\x13\xcc\xa3\x7e\x63\xf0\x77\xdf\x3b\x4f\x52\xf1\x84\x81\x59\x7f\xc7\xac\xaf\x47\x09\x93\x17\x98\x6d\xb2\x39\xfc\x07\xb9\xb3\x7c\xad\x6d\x53\x9c\x66\xe1\x81\xc8\xdb\x2d\xcf\x27\x04\xb5\x70\x5e\xc9\x46\x0b\xd7\x1b\xa3\xe3\x72\x58\x4a\xbb\x5b\x2f\xf7\xc4\x79\x54\x32\xaf\x6c\xcf\x7c\xb3\xbd\x75\x5b\x1a\x86\xcd\x83\x63\xe1\xa2\x4b\xb1\x92\xe7\x2f\x2f\x1e\xfe\x9b\x2a\xfc\x36\x7a\x5f\x9f\x95\x86\x97\x6a\x6b\x9e\x70\x8d\x8f\x1d\xf5\xd5\x48\x7c\xe0\x21\x3d\xbf\xc5\xc8\xef\x32\xf4\x81\x8f\x7b\xcb\xb1\x23\xe1\xe2\x47\x5c\xe7\x16\xdd\xee\xc8\x37\x0c\x1e\x08\xc6\x0e\x88\x63\xf5\xfb\xb4\x14\x8b\x2d\x97\xb1\xe8\x65\x84\x3e\xf9\x30\xf2\x3c\xf9\xb6\x92\x7c\x24\x61\xe3\x84\xa7\xbb\xd0\x94\x67\x13\x2d\xb5\x22\x8f\x66\xd1\x89\x70\x99\x5f\x9c\x5d\x9c\x0f\x20\xdd\xe0\x46\x91\x77\xed\x1b\x45\x0c\xf1\xad\x14\x26\xb8\xeb\x01\x52\xae\x23\x5b\x14\x91\x6e\x41\x1d\x61\xaa\x76\x97\x1b\x5f\x15\x85\x8a\x8f\x2c\x5c\xbc\x5f\x71\xf3\x3e\x81\x71\xdc\x1f\xfb\xb7\xfb\x7b\x70\xa1\x15\xf8\xca\xe9\x45\xf8\xea\x35\xc9\xa7\xe3\x6f\xfd\x05\xbf\xd4\x1d\xfb\x37\x9f\x6e\xfb\xc6\x8b\xe6\xdd\x40\xd4\xb8\xae\x0d\x03\x53\x58\x4f\x60\x03\x31\x54\xa2\x0d\x8f\x53\x7a\x37\x3e\x51\x1a\xd2\xd6\x6e\x9b\x1a\x14\x51\x83\x04\xd6\x00\xd9\x0a\xe1\xc3\xf0\xad\x88\xb9\x37\x35\x8d\x2f\x90\x85\xa1\xfe\xfd\x6b\xf6\xc9\x1a\x9c\xa5\x3b\xaf\x83\x00\xe9\x1b\x64\xbc\x9c\xa6\xcf\x92\xfd\x60\x13\xe4\x9b\xec\x0c\x33\xd7\xec\x7c\x76\xf2\xbf\x00\x00\x00\xff\xff\x64\x62\x6f\x3e\x5f\x1c\x00\x00"),
		},
		"/flux-secret.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-secret.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 137,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x54\xca\x31\x0a\xc2\x40\x10\x85\xe1\x7e\x4f\xf1\x2e\xb0\x82\xed\x1c\x42\x0b\xc1\x7e\xc8\xbe\xc8\x62\xb2\x19\x93\x89\x18\x86\xdc\x5d\x14\x1b\xcb\x9f\xff\xcb\x39\x27\xb5\x7a\xe5\xbc\xd4\xa9\x09\x9e\xc7\x74\xaf\xad\x08\x2e\xec\x66\x7a\x1a\xe9\x5a\xd4\x55\x12\xd0\x74\xa4\xa0\x1f\xd6\x57\xbe\x55\xcf\x85\x36\x4c\x5b\x04\x6a\x8f\xc3\x49\x47\x2e\xa6\x1d\xb1\xef\x3f\xfa\x4d\x41\xc4\xff\x8d\x00\x5b\xf9\x30\xdf\x8c\x82\xb3\xe9\x63\x65\x7a\x07\x00\x00\xff\xff\x40\x21\xa1\xbb\x89\x00\x00\x00"),
		},
		"/memcache-dep.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-dep.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 974,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x6c\x53\xcd\x6e\xdb\x3c\x10\xbc\xeb\x29\x06\xf0\xf5\x93\xf2\x29\x40\x7a\xd0\x2d\x68\xda\x22\x40\x1b\x18\x08\xd2\xfb\x9a\x5a\x29\x44\xf8\x57\x72\xe9\x5a\x15\xf2\xee\x85\x64\xc7\x96\x9a\xec\x49\xe2\xcc\xce\xce\x72\x97\x65\x59\x16\x1b\x58\xb6\x8a\xd4\x33\xb7\x68\x39\x18\x3f\x58\x76\x82\x9c\xb8\xc5\x6e\xc0\x57\x93\x0f\x10\x8f\x99\x51\x6c\xa0\xbc\x13\xd2\x8e\x23\xb4\xa5\x9e\x61\x59\xa8\x25\xa1\xaa\xa0\xa0\x7f\x72\x4c\xda\xbb\x06\x14\x42\xba\xda\xd7\xc5\x8b\x76\x6d\x83\xbb\xb3\x6c\xf1\x46\x6f\x0a\xc0\x91\xe5\xe6\x52\x7d\x1c\xa1\x3b\x54\x0f\x64\x39\x05\x52\x8c\xd7\xd7\x13\x69\xfe\x6d\x30\x8e\x6b\x74\x1c\xc1\xae\x9d\x68\x29\xb0\x9a\x14\x23\x07\xa3\x15\xa5\x06\x75\x01\x24\x36\xac\xc4\xc7\x09\x01\x2c\x89\x7a\xfe\x4e\x3b\x36\xe9\x78\xf0\xce\x40\x01\x08\xdb\x60\x48\xf8\x94\xb2\x30\x3b\x85\x59\x65\x7f\x94\x0f\xbc\x59\x99\x71\xdf\xf2\xe3\xca\xc4\x14\x3b\x16\xaa\x5e\xf2\x8e\xa3\x63\xe1\x54\x69\x7f\xe5\x53\x03\xa3\x5d\x3e\x9c\x48\xe7\x4b\x3e\x17\x2b\x3f\x2c\x36\xc5\x3c\x86\x05\xd0\xd4\xd5\xa7\xaa\xfe\xbf\x24\x13\xb4\xe3\x35\x6d\x9b\x8d\xd9\x7a\xa3\xd5\xd0\xe0\xbe\x7b\xf0\xb2\x8d\x9c\xa6\xb1\xbc\xb1\x28\xf6\x8b\xfe\x4a\x94\x16\x37\xf5\x35\x80\x0d\x7e\xd0\x41\xdb\x6c\xa7\x42\x3e\x0e\xd3\x4a\xe4\xc4\xff\x41\x3b\x58\xee\x69\x37\x08\xa7\x65\xe2\x3d\x6e\x2c\x56\x89\x49\xff\x61\x74\x3e\xc2\x3b\x86\x16\xb6\x4b\x7a\x40\x5d\x5f\xd7\x35\x36\xb8\xe3\x8e\xb2\x11\x04\x1f\x2f\xbe\x36\x13\x67\xbf\x3f\x7e\x3e\x39\xe5\xed\xbc\xa4\xe2\xd1\xb3\xc0\xf8\x3e\xc1\x77\x60\x52\xcf\x88\xfc\x2b\x73\x12\x90\x6b\x11\x39\x05\xef\x12\x57\x67\xa1\x49\x75\xd5\xe1\xf1\x5a\x95\xd1\xec\xe4\xd2\xc0\x62\x04\x5b\x1f\xa5\x39\xba\x3b\x6d\xe8\x6d\xdb\x3e\xb2\xca\x51\xcb\xf0\xd9\x3b\xe1\x83\xcc\x9b\x7a\x8c\xb4\x46\x9a\x85\x64\xcc\xee\x36\x3d\x25\x8e\x27\xb9\x7f\xa1\x6f\xd1\xe7\xf0\x1e\x23\x63\xfc\xef\x6d\xd4\x7b\x6d\xb8\xe7\x2f\x49\x91\x21\x99\x5f\x59\x47\x26\xf1\xe5\x15\xfc\x0d\x00\x00\xff\xff\x3f\x87\x20\x76\xce\x03\x00\x00"),
		},
		"/memcache-svc.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-svc.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 206,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x5c\x8c\x3d\x0e\x02\x21\x10\x46\x7b\x4e\xf1\x5d\x00\x13\x2c\x39\x84\x8d\x89\xfd\x04\x3e\x23\x51\x58\x02\x64\x9b\xc9\xde\xdd\xb0\x6b\xe3\x76\xf3\xf3\xde\xb3\xd6\x1a\xa9\xe9\xc1\xd6\xd3\x52\x3c\x56\x67\xde\xa9\x44\x8f\x3b\xdb\x9a\x02\x4d\xe6\x90\x28\x43\xbc\x01\x8a\x64\x7a\x64\xe6\x20\xe1\xc5\xa8\x8a\xf4\xc4\xe5\x26\x99\xbd\x4a\x20\xb6\xed\x07\xed\xab\x87\xea\xff\x57\x15\x2c\x71\x62\xbd\x32\xcc\x62\x5d\xda\xe8\x73\x00\xec\x39\xbf\x5f\x0f\xc4\xc3\xb9\xab\x73\x06\xe8\xfc\x30\x8c\xa5\x1d\xce\xd9\xf8\x06\x00\x00\xff\xff\x20\x2f\xef\xba\xce\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/flux-account.yaml.tmpl"].(os.FileInfo),
		fs["/flux-deployment.yaml.tmpl"].(os.FileInfo),
		fs["/flux-secret.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-dep.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-svc.yaml.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
