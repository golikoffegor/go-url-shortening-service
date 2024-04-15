package postgresql

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

//lint:ignore SA1003 should not use underscores in Go names
func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _bindata_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindata_go() ([]byte, error) {
	return bindata_read(
		_bindata_go,
		"bindata.go",
	)
}

var _internal_storage_sql_postgresql_migrations_default_1_init_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x42\xe6\x06\x97\x24\x96\xa4\xe6\xa6\xe6\x95\x38\xa5\xa6\x67\xe6\x71\x39\x07\xb9\x3a\x86\xb8\x2a\x84\x38\x3a\xf9\xb8\x2a\x14\x67\xe4\x17\x95\xa4\xe6\xa5\x16\x95\x16\xe5\x14\x2b\x68\x70\x29\x28\x28\x28\x64\xa6\x28\x04\xbb\x06\x79\x3a\xfa\x28\x04\x04\x79\xfa\x3a\x06\x45\x2a\x78\xbb\x46\xea\x80\xa5\x4a\x8b\x72\x14\x4a\x52\x2b\x4a\x14\xfc\xfc\x43\x14\xfc\x42\x7d\x7c\xe0\xc2\xf1\xd9\xa9\x95\xa8\x52\x5c\x9a\xd6\xd8\x9d\xe1\x9a\x97\x02\x08\x00\x00\xff\xff\xb7\x22\xe8\x03\xae\x00\x00\x00")

func internal_storage_sql_postgresql_migrations_default_1_init_sql() ([]byte, error) {
	return bindata_read(
		_internal_storage_sql_postgresql_migrations_default_1_init_sql,
		"internal/storage/sql/postgresql/migrations/default/1_init.sql",
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
	"bindata.go": bindata_go,
	"internal/storage/sql/postgresql/migrations/default/1_init.sql": internal_storage_sql_postgresql_migrations_default_1_init_sql,
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

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"bindata.go": &_bintree_t{bindata_go, map[string]*_bintree_t{}},
	"internal": &_bintree_t{nil, map[string]*_bintree_t{
		"storage": &_bintree_t{nil, map[string]*_bintree_t{
			"sql": &_bintree_t{nil, map[string]*_bintree_t{
				"postgresql": &_bintree_t{nil, map[string]*_bintree_t{
					"migrations": &_bintree_t{nil, map[string]*_bintree_t{
						"default": &_bintree_t{nil, map[string]*_bintree_t{
							"1_init.sql": &_bintree_t{internal_storage_sql_postgresql_migrations_default_1_init_sql, map[string]*_bintree_t{}},
						}},
					}},
				}},
			}},
		}},
	}},
}}
