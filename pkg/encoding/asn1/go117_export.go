// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.17
// +build go1.17

package asn1

import (
	q "encoding/asn1"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "asn1",
		Path: "encoding/asn1",
		Deps: map[string]string{
			"bytes":         "bytes",
			"errors":        "errors",
			"fmt":           "fmt",
			"math":          "math",
			"math/big":      "big",
			"reflect":       "reflect",
			"sort":          "sort",
			"strconv":       "strconv",
			"strings":       "strings",
			"time":          "time",
			"unicode/utf16": "utf16",
			"unicode/utf8":  "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"BitString":        {reflect.TypeOf((*q.BitString)(nil)).Elem(), "At,RightAlign", ""},
			"Enumerated":       {reflect.TypeOf((*q.Enumerated)(nil)).Elem(), "", ""},
			"Flag":             {reflect.TypeOf((*q.Flag)(nil)).Elem(), "", ""},
			"ObjectIdentifier": {reflect.TypeOf((*q.ObjectIdentifier)(nil)).Elem(), "Equal,String", ""},
			"RawContent":       {reflect.TypeOf((*q.RawContent)(nil)).Elem(), "", ""},
			"RawValue":         {reflect.TypeOf((*q.RawValue)(nil)).Elem(), "", ""},
			"StructuralError":  {reflect.TypeOf((*q.StructuralError)(nil)).Elem(), "Error", ""},
			"SyntaxError":      {reflect.TypeOf((*q.SyntaxError)(nil)).Elem(), "Error", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"NullBytes":    reflect.ValueOf(&q.NullBytes),
			"NullRawValue": reflect.ValueOf(&q.NullRawValue),
		},
		Funcs: map[string]reflect.Value{
			"Marshal":             reflect.ValueOf(q.Marshal),
			"MarshalWithParams":   reflect.ValueOf(q.MarshalWithParams),
			"Unmarshal":           reflect.ValueOf(q.Unmarshal),
			"UnmarshalWithParams": reflect.ValueOf(q.UnmarshalWithParams),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"ClassApplication":     {"untyped int", constant.MakeInt64(int64(q.ClassApplication))},
			"ClassContextSpecific": {"untyped int", constant.MakeInt64(int64(q.ClassContextSpecific))},
			"ClassPrivate":         {"untyped int", constant.MakeInt64(int64(q.ClassPrivate))},
			"ClassUniversal":       {"untyped int", constant.MakeInt64(int64(q.ClassUniversal))},
			"TagBMPString":         {"untyped int", constant.MakeInt64(int64(q.TagBMPString))},
			"TagBitString":         {"untyped int", constant.MakeInt64(int64(q.TagBitString))},
			"TagBoolean":           {"untyped int", constant.MakeInt64(int64(q.TagBoolean))},
			"TagEnum":              {"untyped int", constant.MakeInt64(int64(q.TagEnum))},
			"TagGeneralString":     {"untyped int", constant.MakeInt64(int64(q.TagGeneralString))},
			"TagGeneralizedTime":   {"untyped int", constant.MakeInt64(int64(q.TagGeneralizedTime))},
			"TagIA5String":         {"untyped int", constant.MakeInt64(int64(q.TagIA5String))},
			"TagInteger":           {"untyped int", constant.MakeInt64(int64(q.TagInteger))},
			"TagNull":              {"untyped int", constant.MakeInt64(int64(q.TagNull))},
			"TagNumericString":     {"untyped int", constant.MakeInt64(int64(q.TagNumericString))},
			"TagOID":               {"untyped int", constant.MakeInt64(int64(q.TagOID))},
			"TagOctetString":       {"untyped int", constant.MakeInt64(int64(q.TagOctetString))},
			"TagPrintableString":   {"untyped int", constant.MakeInt64(int64(q.TagPrintableString))},
			"TagSequence":          {"untyped int", constant.MakeInt64(int64(q.TagSequence))},
			"TagSet":               {"untyped int", constant.MakeInt64(int64(q.TagSet))},
			"TagT61String":         {"untyped int", constant.MakeInt64(int64(q.TagT61String))},
			"TagUTCTime":           {"untyped int", constant.MakeInt64(int64(q.TagUTCTime))},
			"TagUTF8String":        {"untyped int", constant.MakeInt64(int64(q.TagUTF8String))},
		},
	})
}
