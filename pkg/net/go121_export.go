// export by github.com/goplus/igop/cmd/qexp

//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package net

import (
	q "net"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "net",
		Path: "net",
		Deps: map[string]string{
			"context":                                "context",
			"errors":                                 "errors",
			"internal/bytealg":                       "bytealg",
			"internal/godebug":                       "godebug",
			"internal/itoa":                          "itoa",
			"internal/nettrace":                      "nettrace",
			"internal/poll":                          "poll",
			"internal/singleflight":                  "singleflight",
			"internal/syscall/unix":                  "unix",
			"io":                                     "io",
			"io/fs":                                  "fs",
			"net/netip":                              "netip",
			"os":                                     "os",
			"runtime":                                "runtime",
			"runtime/cgo":                            "cgo",
			"sort":                                   "sort",
			"sync":                                   "sync",
			"sync/atomic":                            "atomic",
			"syscall":                                "syscall",
			"time":                                   "time",
			"unsafe":                                 "unsafe",
			"vendor/golang.org/x/net/dns/dnsmessage": "dnsmessage",
			"vendor/golang.org/x/net/route":          "route",
		},
		Interfaces: map[string]reflect.Type{
			"Addr":       reflect.TypeOf((*q.Addr)(nil)).Elem(),
			"Conn":       reflect.TypeOf((*q.Conn)(nil)).Elem(),
			"Error":      reflect.TypeOf((*q.Error)(nil)).Elem(),
			"Listener":   reflect.TypeOf((*q.Listener)(nil)).Elem(),
			"PacketConn": reflect.TypeOf((*q.PacketConn)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"AddrError":           reflect.TypeOf((*q.AddrError)(nil)).Elem(),
			"Buffers":             reflect.TypeOf((*q.Buffers)(nil)).Elem(),
			"DNSConfigError":      reflect.TypeOf((*q.DNSConfigError)(nil)).Elem(),
			"DNSError":            reflect.TypeOf((*q.DNSError)(nil)).Elem(),
			"Dialer":              reflect.TypeOf((*q.Dialer)(nil)).Elem(),
			"Flags":               reflect.TypeOf((*q.Flags)(nil)).Elem(),
			"HardwareAddr":        reflect.TypeOf((*q.HardwareAddr)(nil)).Elem(),
			"IP":                  reflect.TypeOf((*q.IP)(nil)).Elem(),
			"IPAddr":              reflect.TypeOf((*q.IPAddr)(nil)).Elem(),
			"IPConn":              reflect.TypeOf((*q.IPConn)(nil)).Elem(),
			"IPMask":              reflect.TypeOf((*q.IPMask)(nil)).Elem(),
			"IPNet":               reflect.TypeOf((*q.IPNet)(nil)).Elem(),
			"Interface":           reflect.TypeOf((*q.Interface)(nil)).Elem(),
			"InvalidAddrError":    reflect.TypeOf((*q.InvalidAddrError)(nil)).Elem(),
			"ListenConfig":        reflect.TypeOf((*q.ListenConfig)(nil)).Elem(),
			"MX":                  reflect.TypeOf((*q.MX)(nil)).Elem(),
			"NS":                  reflect.TypeOf((*q.NS)(nil)).Elem(),
			"OpError":             reflect.TypeOf((*q.OpError)(nil)).Elem(),
			"ParseError":          reflect.TypeOf((*q.ParseError)(nil)).Elem(),
			"Resolver":            reflect.TypeOf((*q.Resolver)(nil)).Elem(),
			"SRV":                 reflect.TypeOf((*q.SRV)(nil)).Elem(),
			"TCPAddr":             reflect.TypeOf((*q.TCPAddr)(nil)).Elem(),
			"TCPConn":             reflect.TypeOf((*q.TCPConn)(nil)).Elem(),
			"TCPListener":         reflect.TypeOf((*q.TCPListener)(nil)).Elem(),
			"UDPAddr":             reflect.TypeOf((*q.UDPAddr)(nil)).Elem(),
			"UDPConn":             reflect.TypeOf((*q.UDPConn)(nil)).Elem(),
			"UnixAddr":            reflect.TypeOf((*q.UnixAddr)(nil)).Elem(),
			"UnixConn":            reflect.TypeOf((*q.UnixConn)(nil)).Elem(),
			"UnixListener":        reflect.TypeOf((*q.UnixListener)(nil)).Elem(),
			"UnknownNetworkError": reflect.TypeOf((*q.UnknownNetworkError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"DefaultResolver":            reflect.ValueOf(&q.DefaultResolver),
			"ErrClosed":                  reflect.ValueOf(&q.ErrClosed),
			"ErrWriteToConnected":        reflect.ValueOf(&q.ErrWriteToConnected),
			"IPv4allrouter":              reflect.ValueOf(&q.IPv4allrouter),
			"IPv4allsys":                 reflect.ValueOf(&q.IPv4allsys),
			"IPv4bcast":                  reflect.ValueOf(&q.IPv4bcast),
			"IPv4zero":                   reflect.ValueOf(&q.IPv4zero),
			"IPv6interfacelocalallnodes": reflect.ValueOf(&q.IPv6interfacelocalallnodes),
			"IPv6linklocalallnodes":      reflect.ValueOf(&q.IPv6linklocalallnodes),
			"IPv6linklocalallrouters":    reflect.ValueOf(&q.IPv6linklocalallrouters),
			"IPv6loopback":               reflect.ValueOf(&q.IPv6loopback),
			"IPv6unspecified":            reflect.ValueOf(&q.IPv6unspecified),
			"IPv6zero":                   reflect.ValueOf(&q.IPv6zero),
		},
		Funcs: map[string]reflect.Value{
			"CIDRMask":            reflect.ValueOf(q.CIDRMask),
			"Dial":                reflect.ValueOf(q.Dial),
			"DialIP":              reflect.ValueOf(q.DialIP),
			"DialTCP":             reflect.ValueOf(q.DialTCP),
			"DialTimeout":         reflect.ValueOf(q.DialTimeout),
			"DialUDP":             reflect.ValueOf(q.DialUDP),
			"DialUnix":            reflect.ValueOf(q.DialUnix),
			"FileConn":            reflect.ValueOf(q.FileConn),
			"FileListener":        reflect.ValueOf(q.FileListener),
			"FilePacketConn":      reflect.ValueOf(q.FilePacketConn),
			"IPv4":                reflect.ValueOf(q.IPv4),
			"IPv4Mask":            reflect.ValueOf(q.IPv4Mask),
			"InterfaceAddrs":      reflect.ValueOf(q.InterfaceAddrs),
			"InterfaceByIndex":    reflect.ValueOf(q.InterfaceByIndex),
			"InterfaceByName":     reflect.ValueOf(q.InterfaceByName),
			"Interfaces":          reflect.ValueOf(q.Interfaces),
			"JoinHostPort":        reflect.ValueOf(q.JoinHostPort),
			"Listen":              reflect.ValueOf(q.Listen),
			"ListenIP":            reflect.ValueOf(q.ListenIP),
			"ListenMulticastUDP":  reflect.ValueOf(q.ListenMulticastUDP),
			"ListenPacket":        reflect.ValueOf(q.ListenPacket),
			"ListenTCP":           reflect.ValueOf(q.ListenTCP),
			"ListenUDP":           reflect.ValueOf(q.ListenUDP),
			"ListenUnix":          reflect.ValueOf(q.ListenUnix),
			"ListenUnixgram":      reflect.ValueOf(q.ListenUnixgram),
			"LookupAddr":          reflect.ValueOf(q.LookupAddr),
			"LookupCNAME":         reflect.ValueOf(q.LookupCNAME),
			"LookupHost":          reflect.ValueOf(q.LookupHost),
			"LookupIP":            reflect.ValueOf(q.LookupIP),
			"LookupMX":            reflect.ValueOf(q.LookupMX),
			"LookupNS":            reflect.ValueOf(q.LookupNS),
			"LookupPort":          reflect.ValueOf(q.LookupPort),
			"LookupSRV":           reflect.ValueOf(q.LookupSRV),
			"LookupTXT":           reflect.ValueOf(q.LookupTXT),
			"ParseCIDR":           reflect.ValueOf(q.ParseCIDR),
			"ParseIP":             reflect.ValueOf(q.ParseIP),
			"ParseMAC":            reflect.ValueOf(q.ParseMAC),
			"Pipe":                reflect.ValueOf(q.Pipe),
			"ResolveIPAddr":       reflect.ValueOf(q.ResolveIPAddr),
			"ResolveTCPAddr":      reflect.ValueOf(q.ResolveTCPAddr),
			"ResolveUDPAddr":      reflect.ValueOf(q.ResolveUDPAddr),
			"ResolveUnixAddr":     reflect.ValueOf(q.ResolveUnixAddr),
			"SplitHostPort":       reflect.ValueOf(q.SplitHostPort),
			"TCPAddrFromAddrPort": reflect.ValueOf(q.TCPAddrFromAddrPort),
			"UDPAddrFromAddrPort": reflect.ValueOf(q.UDPAddrFromAddrPort),
		},
		TypedConsts: map[string]igop.TypedConst{
			"FlagBroadcast":    {reflect.TypeOf(q.FlagBroadcast), constant.MakeInt64(int64(q.FlagBroadcast))},
			"FlagLoopback":     {reflect.TypeOf(q.FlagLoopback), constant.MakeInt64(int64(q.FlagLoopback))},
			"FlagMulticast":    {reflect.TypeOf(q.FlagMulticast), constant.MakeInt64(int64(q.FlagMulticast))},
			"FlagPointToPoint": {reflect.TypeOf(q.FlagPointToPoint), constant.MakeInt64(int64(q.FlagPointToPoint))},
			"FlagRunning":      {reflect.TypeOf(q.FlagRunning), constant.MakeInt64(int64(q.FlagRunning))},
			"FlagUp":           {reflect.TypeOf(q.FlagUp), constant.MakeInt64(int64(q.FlagUp))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"IPv4len": {"untyped int", constant.MakeInt64(int64(q.IPv4len))},
			"IPv6len": {"untyped int", constant.MakeInt64(int64(q.IPv6len))},
		},
	})
}
