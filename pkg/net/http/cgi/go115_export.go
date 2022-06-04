// export by github.com/goplus/igop/cmd/qexp

//+build go1.15,!go1.16

package cgi

import (
	q "net/http/cgi"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "cgi",
		Path: "net/http/cgi",
		Deps: map[string]string{
			"bufio":                                 "bufio",
			"crypto/tls":                            "tls",
			"errors":                                "errors",
			"fmt":                                   "fmt",
			"io":                                    "io",
			"io/ioutil":                             "ioutil",
			"log":                                   "log",
			"net":                                   "net",
			"net/http":                              "http",
			"net/textproto":                         "textproto",
			"net/url":                               "url",
			"os":                                    "os",
			"os/exec":                               "exec",
			"path/filepath":                         "filepath",
			"regexp":                                "regexp",
			"runtime":                               "runtime",
			"strconv":                               "strconv",
			"strings":                               "strings",
			"vendor/golang.org/x/net/http/httpguts": "httpguts",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"Handler": {reflect.TypeOf((*q.Handler)(nil)).Elem(), "", "ServeHTTP,handleInternalRedirect,printf,stderr"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Request":        reflect.ValueOf(q.Request),
			"RequestFromMap": reflect.ValueOf(q.RequestFromMap),
			"Serve":          reflect.ValueOf(q.Serve),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
