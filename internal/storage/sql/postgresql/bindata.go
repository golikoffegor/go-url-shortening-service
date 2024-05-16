package postgresql

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindataGo() ([]byte, error) {
	return bindataRead(
		_bindataGo,
		"bindata.go",
	)
}

var _internalStorageSQLPostgresqlMigrationsDefault1InitSQL = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x42\xe6\x06\x97\x24\x96\xa4\xe6\xa6\xe6\x95\x38\xa5\xa6\x67\xe6\x71\x39\x07\xb9\x3a\x86\xb8\x2a\x84\x38\x3a\xf9\xb8\x2a\x14\x67\xe4\x17\x95\xa4\xe6\xa5\x16\x95\x16\xe5\x14\x2b\x68\x70\x29\x28\x28\x28\x64\xa6\x28\x04\xbb\x06\x79\x3a\xfa\x28\x04\x04\x79\xfa\x3a\x06\x45\x2a\x78\xbb\x46\xea\x80\xa5\x4a\x8b\x72\x14\x4a\x52\x2b\x4a\x14\xfc\xfc\x43\x14\xfc\x42\x7d\x7c\xe0\xc2\xf1\xd9\xa9\x95\xa8\x52\x5c\x9a\xd6\xd8\x9d\xe1\x9a\x97\x02\x08\x00\x00\xff\xff\xb7\x22\xe8\x03\xae\x00\x00\x00")

func internalStorageSQLPostgresqlMigrationsDefault1InitSQL() ([]byte, error) {
	return bindataRead(
		_internalStorageSQLPostgresqlMigrationsDefault1InitSQL,
		"internal/storage/sql/postgresql/migrations/default/1_init.sql",
	)
}

var _internalStorageSQLPostgresqlMigrationsDefault2SQL = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x42\xe6\x06\x97\x24\x96\xa4\xe6\xa6\xe6\x95\x38\xa5\xa6\x67\xe6\x71\x39\x07\xb9\x3a\x86\xb8\x2a\x84\xfa\x79\x06\x86\xba\x2a\x78\xfa\xb9\xb8\x46\x28\x94\x16\xe5\xc4\x67\xa6\x54\x28\xf8\xfb\x29\x14\x67\xe4\x17\x95\xa4\xe6\xa5\x16\x95\x16\xe5\x14\x2b\x68\x94\x16\xe5\x68\x5a\x63\x37\xce\x35\x2f\x05\x10\x00\x00\xff\xff\x49\x5d\x57\x61\x76\x00\x00\x00")

func internalStorageSQLPostgresqlMigrationsDefault2SQL() ([]byte, error) {
	return bindataRead(
		_internalStorageSQLPostgresqlMigrationsDefault2SQL,
		"internal/storage/sql/postgresql/migrations/default/2.sql",
	)
}

var _internalStorageSQLPostgresqlMigrationsDefault3SQL = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x42\xe6\x06\x97\x24\x96\xa4\xe6\xa6\xe6\x95\x38\xa5\xa6\x67\xe6\x71\x39\xfa\x84\xb8\x06\x29\x84\x38\x3a\xf9\xb8\x2a\x14\x67\xe4\x17\x95\xa4\xe6\xa5\x16\x95\x16\xe5\x14\x2b\x38\xba\xb8\x28\x38\xfb\xfb\x84\xfa\xfa\x29\x94\x16\xa7\x16\xc5\x67\xa6\x28\x38\x7b\x38\x06\x69\x18\x1a\x59\x68\x2a\xb8\xb8\xba\x39\x86\xfa\x84\x28\xf8\x85\xfa\xf8\x58\x63\x37\xdf\x35\x2f\x05\x10\x00\x00\xff\xff\xcb\x2a\xb7\x61\x87\x00\x00\x00")

func internalStorageSQLPostgresqlMigrationsDefault3SQL() ([]byte, error) {
	return bindataRead(
		_internalStorageSQLPostgresqlMigrationsDefault3SQL,
		"internal/storage/sql/postgresql/migrations/default/3.sql",
	)
}

var _internalStorageSQLPostgresqlMigrationsDefault4SQL = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcd\x31\xae\xc2\x30\x0c\x06\xe0\xbd\xa7\xf8\xf7\xa7\x9e\xe0\x4d\x0e\x71\x27\x93\x48\x34\x99\x51\xa5\x58\x25\x52\x1b\x50\x62\xee\xcf\xca\xc0\xf8\x4d\xdf\x3c\xe3\xef\xac\x7b\xdf\x4c\x91\x5f\xd3\x37\x57\xdb\x4c\x4f\x6d\xe6\x74\xaf\x6d\x22\x49\x7c\x43\x22\x27\x8c\xf1\x78\x76\xd3\xa6\xfd\xdd\x8f\x01\xf2\x1e\x97\x28\xf9\x1a\x50\xc7\xbd\xe8\xa1\xa6\x05\x2e\x46\x61\x0a\x08\x31\x21\x64\x11\x78\x5e\x28\x4b\xc2\x42\xb2\xf2\xff\xef\x8a\x5b\xf9\x04\x00\x00\xff\xff\x65\xf6\xe5\xed\x92\x00\x00\x00")

func internalStorageSQLPostgresqlMigrationsDefault4SQL() ([]byte, error) {
	return bindataRead(
		_internalStorageSQLPostgresqlMigrationsDefault4SQL,
		"internal/storage/sql/postgresql/migrations/default/4.sql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"bindata.go": bindataGo,
	"internal/storage/sql/postgresql/migrations/default/1_init.sql": internalStorageSQLPostgresqlMigrationsDefault1InitSQL,
	"internal/storage/sql/postgresql/migrations/default/2.sql":      internalStorageSQLPostgresqlMigrationsDefault2SQL,
	"internal/storage/sql/postgresql/migrations/default/3.sql":      internalStorageSQLPostgresqlMigrationsDefault3SQL,
	"internal/storage/sql/postgresql/migrations/default/4.sql":      internalStorageSQLPostgresqlMigrationsDefault4SQL,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintreeT struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintreeT
}

var _bintree = &_bintreeT{nil, map[string]*_bintreeT{
	"bindata.go": &_bintreeT{bindataGo, map[string]*_bintreeT{}},
	"internal": &_bintreeT{nil, map[string]*_bintreeT{
		"storage": &_bintreeT{nil, map[string]*_bintreeT{
			"sql": &_bintreeT{nil, map[string]*_bintreeT{
				"postgresql": &_bintreeT{nil, map[string]*_bintreeT{
					"migrations": &_bintreeT{nil, map[string]*_bintreeT{
						"default": &_bintreeT{nil, map[string]*_bintreeT{
							"1_init.sql": &_bintreeT{internalStorageSQLPostgresqlMigrationsDefault1InitSQL, map[string]*_bintreeT{}},
							"2.sql":      &_bintreeT{internalStorageSQLPostgresqlMigrationsDefault2SQL, map[string]*_bintreeT{}},
							"3.sql":      &_bintreeT{internalStorageSQLPostgresqlMigrationsDefault3SQL, map[string]*_bintreeT{}},
							"4.sql":      &_bintreeT{internalStorageSQLPostgresqlMigrationsDefault4SQL, map[string]*_bintreeT{}},
						}},
					}},
				}},
			}},
		}},
	}},
}}
