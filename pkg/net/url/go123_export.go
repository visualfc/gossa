// export by github.com/goplus/igop/cmd/qexp

//go:build go1.23
// +build go1.23

package url

import (
	q "net/url"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "url",
		Path: "net/url",
		Deps: map[string]string{
			"errors":  "errors",
			"fmt":     "fmt",
			"path":    "path",
			"slices":  "slices",
			"strconv": "strconv",
			"strings": "strings",
			"unsafe":  "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Error":            reflect.TypeOf((*q.Error)(nil)).Elem(),
			"EscapeError":      reflect.TypeOf((*q.EscapeError)(nil)).Elem(),
			"InvalidHostError": reflect.TypeOf((*q.InvalidHostError)(nil)).Elem(),
			"URL":              reflect.TypeOf((*q.URL)(nil)).Elem(),
			"Userinfo":         reflect.TypeOf((*q.Userinfo)(nil)).Elem(),
			"Values":           reflect.TypeOf((*q.Values)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"JoinPath":        reflect.ValueOf(q.JoinPath),
			"Parse":           reflect.ValueOf(q.Parse),
			"ParseQuery":      reflect.ValueOf(q.ParseQuery),
			"ParseRequestURI": reflect.ValueOf(q.ParseRequestURI),
			"PathEscape":      reflect.ValueOf(q.PathEscape),
			"PathUnescape":    reflect.ValueOf(q.PathUnescape),
			"QueryEscape":     reflect.ValueOf(q.QueryEscape),
			"QueryUnescape":   reflect.ValueOf(q.QueryUnescape),
			"User":            reflect.ValueOf(q.User),
			"UserPassword":    reflect.ValueOf(q.UserPassword),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
