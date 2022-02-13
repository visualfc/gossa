// export by github.com/goplus/gossa/cmd/qexp

//+build go1.16,!go1.17

package suffixarray

import (
	q "index/suffixarray"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "suffixarray",
		Path: "index/suffixarray",
		Deps: map[string]string{
			"bytes":           "bytes",
			"encoding/binary": "binary",
			"errors":          "errors",
			"io":              "io",
			"math":            "math",
			"regexp":          "regexp",
			"sort":            "sort",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"Index": {reflect.TypeOf((*q.Index)(nil)).Elem(), "", "Bytes,FindAllIndex,Lookup,Read,Write,at,lookupAll"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(q.New),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}