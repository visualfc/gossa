// export by github.com/goplus/gossa/cmd/qexp

//+build go1.15,!go1.16

package http

import (
	q "net/http"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "http",
		Path: "net/http",
		Deps: map[string]string{
			"bufio":                                  "bufio",
			"bytes":                                  "bytes",
			"compress/gzip":                          "gzip",
			"container/list":                         "list",
			"context":                                "context",
			"crypto/rand":                            "rand",
			"crypto/tls":                             "tls",
			"encoding/base64":                        "base64",
			"encoding/binary":                        "binary",
			"errors":                                 "errors",
			"fmt":                                    "fmt",
			"io":                                     "io",
			"io/ioutil":                              "ioutil",
			"log":                                    "log",
			"math":                                   "math",
			"math/rand":                              "rand",
			"mime":                                   "mime",
			"mime/multipart":                         "multipart",
			"net":                                    "net",
			"net/http/httptrace":                     "httptrace",
			"net/http/internal":                      "internal",
			"net/textproto":                          "textproto",
			"net/url":                                "url",
			"os":                                     "os",
			"path":                                   "path",
			"path/filepath":                          "filepath",
			"reflect":                                "reflect",
			"runtime":                                "runtime",
			"sort":                                   "sort",
			"strconv":                                "strconv",
			"strings":                                "strings",
			"sync":                                   "sync",
			"sync/atomic":                            "atomic",
			"time":                                   "time",
			"unicode/utf8":                           "utf8",
			"vendor/golang.org/x/net/http/httpguts":  "httpguts",
			"vendor/golang.org/x/net/http/httpproxy": "httpproxy",
			"vendor/golang.org/x/net/http2/hpack":    "hpack",
			"vendor/golang.org/x/net/idna":           "idna",
		},
		Interfaces: map[string]reflect.Type{
			"CloseNotifier":  reflect.TypeOf((*q.CloseNotifier)(nil)).Elem(),
			"CookieJar":      reflect.TypeOf((*q.CookieJar)(nil)).Elem(),
			"File":           reflect.TypeOf((*q.File)(nil)).Elem(),
			"FileSystem":     reflect.TypeOf((*q.FileSystem)(nil)).Elem(),
			"Flusher":        reflect.TypeOf((*q.Flusher)(nil)).Elem(),
			"Handler":        reflect.TypeOf((*q.Handler)(nil)).Elem(),
			"Hijacker":       reflect.TypeOf((*q.Hijacker)(nil)).Elem(),
			"Pusher":         reflect.TypeOf((*q.Pusher)(nil)).Elem(),
			"ResponseWriter": reflect.TypeOf((*q.ResponseWriter)(nil)).Elem(),
			"RoundTripper":   reflect.TypeOf((*q.RoundTripper)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"Client":        {reflect.TypeOf((*q.Client)(nil)).Elem(), "", "CloseIdleConnections,Do,Get,Head,Post,PostForm,checkRedirect,deadline,do,makeHeadersCopier,send,transport"},
			"ConnState":     {reflect.TypeOf((*q.ConnState)(nil)).Elem(), "String", ""},
			"Cookie":        {reflect.TypeOf((*q.Cookie)(nil)).Elem(), "", "String"},
			"Dir":           {reflect.TypeOf((*q.Dir)(nil)).Elem(), "Open", ""},
			"HandlerFunc":   {reflect.TypeOf((*q.HandlerFunc)(nil)).Elem(), "ServeHTTP", ""},
			"Header":        {reflect.TypeOf((*q.Header)(nil)).Elem(), "Add,Clone,Del,Get,Set,Values,Write,WriteSubset,get,has,sortedKeyValues,write,writeSubset", ""},
			"ProtocolError": {reflect.TypeOf((*q.ProtocolError)(nil)).Elem(), "", "Error"},
			"PushOptions":   {reflect.TypeOf((*q.PushOptions)(nil)).Elem(), "", ""},
			"Request":       {reflect.TypeOf((*q.Request)(nil)).Elem(), "", "AddCookie,BasicAuth,Clone,Context,Cookie,Cookies,FormFile,FormValue,MultipartReader,ParseForm,ParseMultipartForm,PostFormValue,ProtoAtLeast,Referer,SetBasicAuth,UserAgent,WithContext,Write,WriteProxy,closeBody,expectsContinue,isH2Upgrade,isReplayable,multipartReader,outgoingLength,requiresHTTP1,wantsClose,wantsHttp10KeepAlive,write"},
			"Response":      {reflect.TypeOf((*q.Response)(nil)).Elem(), "", "Cookies,Location,ProtoAtLeast,Write,bodyIsWritable,closeBody,isProtocolSwitch"},
			"SameSite":      {reflect.TypeOf((*q.SameSite)(nil)).Elem(), "", ""},
			"ServeMux":      {reflect.TypeOf((*q.ServeMux)(nil)).Elem(), "", "Handle,HandleFunc,Handler,ServeHTTP,handler,match,redirectToPathSlash,shouldRedirectRLocked"},
			"Server":        {reflect.TypeOf((*q.Server)(nil)).Elem(), "", "Close,ListenAndServe,ListenAndServeTLS,RegisterOnShutdown,Serve,ServeTLS,SetKeepAlivesEnabled,Shutdown,closeDoneChanLocked,closeIdleConns,closeListenersLocked,doKeepAlives,getDoneChan,getDoneChanLocked,idleTimeout,initialReadLimitSize,logf,maxHeaderBytes,newConn,numListeners,onceSetNextProtoDefaults,onceSetNextProtoDefaults_Serve,readHeaderTimeout,setupHTTP2_Serve,setupHTTP2_ServeTLS,shouldConfigureHTTP2ForServe,shuttingDown,trackConn,trackListener"},
			"Transport":     {reflect.TypeOf((*q.Transport)(nil)).Elem(), "", "CancelRequest,Clone,CloseIdleConnections,RegisterProtocol,RoundTrip,alternateRoundTripper,cancelRequest,connectMethodForRequest,customDialTLS,decConnsPerHost,dial,dialConn,dialConnFor,getConn,hasCustomTLSDialer,maxIdleConnsPerHost,onceSetNextProtoDefaults,putOrCloseIdleConn,queueForDial,queueForIdleConn,readBufferSize,removeIdleConn,removeIdleConnLocked,replaceReqCanceler,roundTrip,setReqCanceler,tryPutIdleConn,useRegisteredProtocol,writeBufferSize"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"DefaultClient":           reflect.ValueOf(&q.DefaultClient),
			"DefaultServeMux":         reflect.ValueOf(&q.DefaultServeMux),
			"DefaultTransport":        reflect.ValueOf(&q.DefaultTransport),
			"ErrAbortHandler":         reflect.ValueOf(&q.ErrAbortHandler),
			"ErrBodyNotAllowed":       reflect.ValueOf(&q.ErrBodyNotAllowed),
			"ErrBodyReadAfterClose":   reflect.ValueOf(&q.ErrBodyReadAfterClose),
			"ErrContentLength":        reflect.ValueOf(&q.ErrContentLength),
			"ErrHandlerTimeout":       reflect.ValueOf(&q.ErrHandlerTimeout),
			"ErrHeaderTooLong":        reflect.ValueOf(&q.ErrHeaderTooLong),
			"ErrHijacked":             reflect.ValueOf(&q.ErrHijacked),
			"ErrLineTooLong":          reflect.ValueOf(&q.ErrLineTooLong),
			"ErrMissingBoundary":      reflect.ValueOf(&q.ErrMissingBoundary),
			"ErrMissingContentLength": reflect.ValueOf(&q.ErrMissingContentLength),
			"ErrMissingFile":          reflect.ValueOf(&q.ErrMissingFile),
			"ErrNoCookie":             reflect.ValueOf(&q.ErrNoCookie),
			"ErrNoLocation":           reflect.ValueOf(&q.ErrNoLocation),
			"ErrNotMultipart":         reflect.ValueOf(&q.ErrNotMultipart),
			"ErrNotSupported":         reflect.ValueOf(&q.ErrNotSupported),
			"ErrServerClosed":         reflect.ValueOf(&q.ErrServerClosed),
			"ErrShortBody":            reflect.ValueOf(&q.ErrShortBody),
			"ErrSkipAltProtocol":      reflect.ValueOf(&q.ErrSkipAltProtocol),
			"ErrUnexpectedTrailer":    reflect.ValueOf(&q.ErrUnexpectedTrailer),
			"ErrUseLastResponse":      reflect.ValueOf(&q.ErrUseLastResponse),
			"ErrWriteAfterFlush":      reflect.ValueOf(&q.ErrWriteAfterFlush),
			"LocalAddrContextKey":     reflect.ValueOf(&q.LocalAddrContextKey),
			"NoBody":                  reflect.ValueOf(&q.NoBody),
			"ServerContextKey":        reflect.ValueOf(&q.ServerContextKey),
		},
		Funcs: map[string]reflect.Value{
			"CanonicalHeaderKey":    reflect.ValueOf(q.CanonicalHeaderKey),
			"DetectContentType":     reflect.ValueOf(q.DetectContentType),
			"Error":                 reflect.ValueOf(q.Error),
			"FileServer":            reflect.ValueOf(q.FileServer),
			"Get":                   reflect.ValueOf(q.Get),
			"Handle":                reflect.ValueOf(q.Handle),
			"HandleFunc":            reflect.ValueOf(q.HandleFunc),
			"Head":                  reflect.ValueOf(q.Head),
			"ListenAndServe":        reflect.ValueOf(q.ListenAndServe),
			"ListenAndServeTLS":     reflect.ValueOf(q.ListenAndServeTLS),
			"MaxBytesReader":        reflect.ValueOf(q.MaxBytesReader),
			"NewFileTransport":      reflect.ValueOf(q.NewFileTransport),
			"NewRequest":            reflect.ValueOf(q.NewRequest),
			"NewRequestWithContext": reflect.ValueOf(q.NewRequestWithContext),
			"NewServeMux":           reflect.ValueOf(q.NewServeMux),
			"NotFound":              reflect.ValueOf(q.NotFound),
			"NotFoundHandler":       reflect.ValueOf(q.NotFoundHandler),
			"ParseHTTPVersion":      reflect.ValueOf(q.ParseHTTPVersion),
			"ParseTime":             reflect.ValueOf(q.ParseTime),
			"Post":                  reflect.ValueOf(q.Post),
			"PostForm":              reflect.ValueOf(q.PostForm),
			"ProxyFromEnvironment":  reflect.ValueOf(q.ProxyFromEnvironment),
			"ProxyURL":              reflect.ValueOf(q.ProxyURL),
			"ReadRequest":           reflect.ValueOf(q.ReadRequest),
			"ReadResponse":          reflect.ValueOf(q.ReadResponse),
			"Redirect":              reflect.ValueOf(q.Redirect),
			"RedirectHandler":       reflect.ValueOf(q.RedirectHandler),
			"Serve":                 reflect.ValueOf(q.Serve),
			"ServeContent":          reflect.ValueOf(q.ServeContent),
			"ServeFile":             reflect.ValueOf(q.ServeFile),
			"ServeTLS":              reflect.ValueOf(q.ServeTLS),
			"SetCookie":             reflect.ValueOf(q.SetCookie),
			"StatusText":            reflect.ValueOf(q.StatusText),
			"StripPrefix":           reflect.ValueOf(q.StripPrefix),
			"TimeoutHandler":        reflect.ValueOf(q.TimeoutHandler),
		},
		TypedConsts: map[string]gossa.TypedConst{
			"SameSiteDefaultMode": {reflect.TypeOf(q.SameSiteDefaultMode), constant.MakeInt64(int64(q.SameSiteDefaultMode))},
			"SameSiteLaxMode":     {reflect.TypeOf(q.SameSiteLaxMode), constant.MakeInt64(int64(q.SameSiteLaxMode))},
			"SameSiteNoneMode":    {reflect.TypeOf(q.SameSiteNoneMode), constant.MakeInt64(int64(q.SameSiteNoneMode))},
			"SameSiteStrictMode":  {reflect.TypeOf(q.SameSiteStrictMode), constant.MakeInt64(int64(q.SameSiteStrictMode))},
			"StateActive":         {reflect.TypeOf(q.StateActive), constant.MakeInt64(int64(q.StateActive))},
			"StateClosed":         {reflect.TypeOf(q.StateClosed), constant.MakeInt64(int64(q.StateClosed))},
			"StateHijacked":       {reflect.TypeOf(q.StateHijacked), constant.MakeInt64(int64(q.StateHijacked))},
			"StateIdle":           {reflect.TypeOf(q.StateIdle), constant.MakeInt64(int64(q.StateIdle))},
			"StateNew":            {reflect.TypeOf(q.StateNew), constant.MakeInt64(int64(q.StateNew))},
		},
		UntypedConsts: map[string]gossa.UntypedConst{
			"DefaultMaxHeaderBytes":               {"untyped int", constant.MakeInt64(int64(q.DefaultMaxHeaderBytes))},
			"DefaultMaxIdleConnsPerHost":          {"untyped int", constant.MakeInt64(int64(q.DefaultMaxIdleConnsPerHost))},
			"MethodConnect":                       {"untyped string", constant.MakeString(string(q.MethodConnect))},
			"MethodDelete":                        {"untyped string", constant.MakeString(string(q.MethodDelete))},
			"MethodGet":                           {"untyped string", constant.MakeString(string(q.MethodGet))},
			"MethodHead":                          {"untyped string", constant.MakeString(string(q.MethodHead))},
			"MethodOptions":                       {"untyped string", constant.MakeString(string(q.MethodOptions))},
			"MethodPatch":                         {"untyped string", constant.MakeString(string(q.MethodPatch))},
			"MethodPost":                          {"untyped string", constant.MakeString(string(q.MethodPost))},
			"MethodPut":                           {"untyped string", constant.MakeString(string(q.MethodPut))},
			"MethodTrace":                         {"untyped string", constant.MakeString(string(q.MethodTrace))},
			"StatusAccepted":                      {"untyped int", constant.MakeInt64(int64(q.StatusAccepted))},
			"StatusAlreadyReported":               {"untyped int", constant.MakeInt64(int64(q.StatusAlreadyReported))},
			"StatusBadGateway":                    {"untyped int", constant.MakeInt64(int64(q.StatusBadGateway))},
			"StatusBadRequest":                    {"untyped int", constant.MakeInt64(int64(q.StatusBadRequest))},
			"StatusConflict":                      {"untyped int", constant.MakeInt64(int64(q.StatusConflict))},
			"StatusContinue":                      {"untyped int", constant.MakeInt64(int64(q.StatusContinue))},
			"StatusCreated":                       {"untyped int", constant.MakeInt64(int64(q.StatusCreated))},
			"StatusEarlyHints":                    {"untyped int", constant.MakeInt64(int64(q.StatusEarlyHints))},
			"StatusExpectationFailed":             {"untyped int", constant.MakeInt64(int64(q.StatusExpectationFailed))},
			"StatusFailedDependency":              {"untyped int", constant.MakeInt64(int64(q.StatusFailedDependency))},
			"StatusForbidden":                     {"untyped int", constant.MakeInt64(int64(q.StatusForbidden))},
			"StatusFound":                         {"untyped int", constant.MakeInt64(int64(q.StatusFound))},
			"StatusGatewayTimeout":                {"untyped int", constant.MakeInt64(int64(q.StatusGatewayTimeout))},
			"StatusGone":                          {"untyped int", constant.MakeInt64(int64(q.StatusGone))},
			"StatusHTTPVersionNotSupported":       {"untyped int", constant.MakeInt64(int64(q.StatusHTTPVersionNotSupported))},
			"StatusIMUsed":                        {"untyped int", constant.MakeInt64(int64(q.StatusIMUsed))},
			"StatusInsufficientStorage":           {"untyped int", constant.MakeInt64(int64(q.StatusInsufficientStorage))},
			"StatusInternalServerError":           {"untyped int", constant.MakeInt64(int64(q.StatusInternalServerError))},
			"StatusLengthRequired":                {"untyped int", constant.MakeInt64(int64(q.StatusLengthRequired))},
			"StatusLocked":                        {"untyped int", constant.MakeInt64(int64(q.StatusLocked))},
			"StatusLoopDetected":                  {"untyped int", constant.MakeInt64(int64(q.StatusLoopDetected))},
			"StatusMethodNotAllowed":              {"untyped int", constant.MakeInt64(int64(q.StatusMethodNotAllowed))},
			"StatusMisdirectedRequest":            {"untyped int", constant.MakeInt64(int64(q.StatusMisdirectedRequest))},
			"StatusMovedPermanently":              {"untyped int", constant.MakeInt64(int64(q.StatusMovedPermanently))},
			"StatusMultiStatus":                   {"untyped int", constant.MakeInt64(int64(q.StatusMultiStatus))},
			"StatusMultipleChoices":               {"untyped int", constant.MakeInt64(int64(q.StatusMultipleChoices))},
			"StatusNetworkAuthenticationRequired": {"untyped int", constant.MakeInt64(int64(q.StatusNetworkAuthenticationRequired))},
			"StatusNoContent":                     {"untyped int", constant.MakeInt64(int64(q.StatusNoContent))},
			"StatusNonAuthoritativeInfo":          {"untyped int", constant.MakeInt64(int64(q.StatusNonAuthoritativeInfo))},
			"StatusNotAcceptable":                 {"untyped int", constant.MakeInt64(int64(q.StatusNotAcceptable))},
			"StatusNotExtended":                   {"untyped int", constant.MakeInt64(int64(q.StatusNotExtended))},
			"StatusNotFound":                      {"untyped int", constant.MakeInt64(int64(q.StatusNotFound))},
			"StatusNotImplemented":                {"untyped int", constant.MakeInt64(int64(q.StatusNotImplemented))},
			"StatusNotModified":                   {"untyped int", constant.MakeInt64(int64(q.StatusNotModified))},
			"StatusOK":                            {"untyped int", constant.MakeInt64(int64(q.StatusOK))},
			"StatusPartialContent":                {"untyped int", constant.MakeInt64(int64(q.StatusPartialContent))},
			"StatusPaymentRequired":               {"untyped int", constant.MakeInt64(int64(q.StatusPaymentRequired))},
			"StatusPermanentRedirect":             {"untyped int", constant.MakeInt64(int64(q.StatusPermanentRedirect))},
			"StatusPreconditionFailed":            {"untyped int", constant.MakeInt64(int64(q.StatusPreconditionFailed))},
			"StatusPreconditionRequired":          {"untyped int", constant.MakeInt64(int64(q.StatusPreconditionRequired))},
			"StatusProcessing":                    {"untyped int", constant.MakeInt64(int64(q.StatusProcessing))},
			"StatusProxyAuthRequired":             {"untyped int", constant.MakeInt64(int64(q.StatusProxyAuthRequired))},
			"StatusRequestEntityTooLarge":         {"untyped int", constant.MakeInt64(int64(q.StatusRequestEntityTooLarge))},
			"StatusRequestHeaderFieldsTooLarge":   {"untyped int", constant.MakeInt64(int64(q.StatusRequestHeaderFieldsTooLarge))},
			"StatusRequestTimeout":                {"untyped int", constant.MakeInt64(int64(q.StatusRequestTimeout))},
			"StatusRequestURITooLong":             {"untyped int", constant.MakeInt64(int64(q.StatusRequestURITooLong))},
			"StatusRequestedRangeNotSatisfiable":  {"untyped int", constant.MakeInt64(int64(q.StatusRequestedRangeNotSatisfiable))},
			"StatusResetContent":                  {"untyped int", constant.MakeInt64(int64(q.StatusResetContent))},
			"StatusSeeOther":                      {"untyped int", constant.MakeInt64(int64(q.StatusSeeOther))},
			"StatusServiceUnavailable":            {"untyped int", constant.MakeInt64(int64(q.StatusServiceUnavailable))},
			"StatusSwitchingProtocols":            {"untyped int", constant.MakeInt64(int64(q.StatusSwitchingProtocols))},
			"StatusTeapot":                        {"untyped int", constant.MakeInt64(int64(q.StatusTeapot))},
			"StatusTemporaryRedirect":             {"untyped int", constant.MakeInt64(int64(q.StatusTemporaryRedirect))},
			"StatusTooEarly":                      {"untyped int", constant.MakeInt64(int64(q.StatusTooEarly))},
			"StatusTooManyRequests":               {"untyped int", constant.MakeInt64(int64(q.StatusTooManyRequests))},
			"StatusUnauthorized":                  {"untyped int", constant.MakeInt64(int64(q.StatusUnauthorized))},
			"StatusUnavailableForLegalReasons":    {"untyped int", constant.MakeInt64(int64(q.StatusUnavailableForLegalReasons))},
			"StatusUnprocessableEntity":           {"untyped int", constant.MakeInt64(int64(q.StatusUnprocessableEntity))},
			"StatusUnsupportedMediaType":          {"untyped int", constant.MakeInt64(int64(q.StatusUnsupportedMediaType))},
			"StatusUpgradeRequired":               {"untyped int", constant.MakeInt64(int64(q.StatusUpgradeRequired))},
			"StatusUseProxy":                      {"untyped int", constant.MakeInt64(int64(q.StatusUseProxy))},
			"StatusVariantAlsoNegotiates":         {"untyped int", constant.MakeInt64(int64(q.StatusVariantAlsoNegotiates))},
			"TimeFormat":                          {"untyped string", constant.MakeString(string(q.TimeFormat))},
			"TrailerPrefix":                       {"untyped string", constant.MakeString(string(q.TrailerPrefix))},
		},
	})
}
