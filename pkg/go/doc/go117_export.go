// export by github.com/goplus/igop/cmd/qexp

//go:build go1.17 && !go1.18
// +build go1.17,!go1.18

package doc

import (
	q "go/doc"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "doc",
		Path: "go/doc",
		Deps: map[string]string{
			"bytes":               "bytes",
			"fmt":                 "fmt",
			"go/ast":              "ast",
			"go/token":            "token",
			"internal/lazyregexp": "lazyregexp",
			"io":                  "io",
			"path":                "path",
			"sort":                "sort",
			"strconv":             "strconv",
			"strings":             "strings",
			"text/template":       "template",
			"unicode":             "unicode",
			"unicode/utf8":        "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"Example": {reflect.TypeOf((*q.Example)(nil)).Elem(), "", ""},
			"Filter":  {reflect.TypeOf((*q.Filter)(nil)).Elem(), "", ""},
			"Func":    {reflect.TypeOf((*q.Func)(nil)).Elem(), "", ""},
			"Mode":    {reflect.TypeOf((*q.Mode)(nil)).Elem(), "", ""},
			"Note":    {reflect.TypeOf((*q.Note)(nil)).Elem(), "", ""},
			"Package": {reflect.TypeOf((*q.Package)(nil)).Elem(), "", "Filter"},
			"Type":    {reflect.TypeOf((*q.Type)(nil)).Elem(), "", ""},
			"Value":   {reflect.TypeOf((*q.Value)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"IllegalPrefixes": reflect.ValueOf(&q.IllegalPrefixes),
		},
		Funcs: map[string]reflect.Value{
			"Examples":      reflect.ValueOf(q.Examples),
			"IsPredeclared": reflect.ValueOf(q.IsPredeclared),
			"New":           reflect.ValueOf(q.New),
			"NewFromFiles":  reflect.ValueOf(q.NewFromFiles),
			"Synopsis":      reflect.ValueOf(q.Synopsis),
			"ToHTML":        reflect.ValueOf(q.ToHTML),
			"ToText":        reflect.ValueOf(q.ToText),
		},
		TypedConsts: map[string]igop.TypedConst{
			"AllDecls":    {reflect.TypeOf(q.AllDecls), constant.MakeInt64(int64(q.AllDecls))},
			"AllMethods":  {reflect.TypeOf(q.AllMethods), constant.MakeInt64(int64(q.AllMethods))},
			"PreserveAST": {reflect.TypeOf(q.PreserveAST), constant.MakeInt64(int64(q.PreserveAST))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
