// export by github.com/goplus/igop/cmd/qexp

//go:build go1.17 && !go1.18
// +build go1.17,!go1.18

package fstest

import (
	q "testing/fstest"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "fstest",
		Path: "testing/fstest",
		Deps: map[string]string{
			"errors":         "errors",
			"fmt":            "fmt",
			"io":             "io",
			"io/fs":          "fs",
			"path":           "path",
			"reflect":        "reflect",
			"sort":           "sort",
			"strings":        "strings",
			"testing/iotest": "iotest",
			"time":           "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"MapFS":   {reflect.TypeOf((*q.MapFS)(nil)).Elem(), "Glob,Open,ReadDir,ReadFile,Stat,Sub", ""},
			"MapFile": {reflect.TypeOf((*q.MapFile)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"TestFS": reflect.ValueOf(q.TestFS),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
