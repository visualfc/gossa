// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package syscall

import (
	q "syscall"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "syscall",
		Path: "syscall",
		Deps: map[string]string{
			"errors":           "errors",
			"internal/itoa":    "itoa",
			"internal/oserror": "oserror",
			"runtime":          "runtime",
			"sync":             "sync",
			"unsafe":           "unsafe",
		},
		Interfaces: map[string]reflect.Type{
			"Conn":    reflect.TypeOf((*q.Conn)(nil)).Elem(),
			"RawConn": reflect.TypeOf((*q.RawConn)(nil)).Elem(),
		},
		NamedTypes: map[string]igop.NamedType{
			"Dir":         {reflect.TypeOf((*q.Dir)(nil)).Elem(), "", "Marshal,Null"},
			"ErrorString": {reflect.TypeOf((*q.ErrorString)(nil)).Elem(), "Error,Is,Temporary,Timeout", ""},
			"Note":        {reflect.TypeOf((*q.Note)(nil)).Elem(), "Signal,String", ""},
			"ProcAttr":    {reflect.TypeOf((*q.ProcAttr)(nil)).Elem(), "", ""},
			"Qid":         {reflect.TypeOf((*q.Qid)(nil)).Elem(), "", ""},
			"SysProcAttr": {reflect.TypeOf((*q.SysProcAttr)(nil)).Elem(), "", ""},
			"Timespec":    {reflect.TypeOf((*q.Timespec)(nil)).Elem(), "", "Nano,Unix"},
			"Timeval":     {reflect.TypeOf((*q.Timeval)(nil)).Elem(), "", "Nano,Unix"},
			"Waitmsg":     {reflect.TypeOf((*q.Waitmsg)(nil)).Elem(), "ExitStatus,Exited,Signaled", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"EACCES":            reflect.ValueOf(&q.EACCES),
			"EAFNOSUPPORT":      reflect.ValueOf(&q.EAFNOSUPPORT),
			"EBUSY":             reflect.ValueOf(&q.EBUSY),
			"EEXIST":            reflect.ValueOf(&q.EEXIST),
			"EINTR":             reflect.ValueOf(&q.EINTR),
			"EINVAL":            reflect.ValueOf(&q.EINVAL),
			"EIO":               reflect.ValueOf(&q.EIO),
			"EISDIR":            reflect.ValueOf(&q.EISDIR),
			"EMFILE":            reflect.ValueOf(&q.EMFILE),
			"ENAMETOOLONG":      reflect.ValueOf(&q.ENAMETOOLONG),
			"ENOENT":            reflect.ValueOf(&q.ENOENT),
			"ENOTDIR":           reflect.ValueOf(&q.ENOTDIR),
			"EPERM":             reflect.ValueOf(&q.EPERM),
			"EPLAN9":            reflect.ValueOf(&q.EPLAN9),
			"ESPIPE":            reflect.ValueOf(&q.ESPIPE),
			"ETIMEDOUT":         reflect.ValueOf(&q.ETIMEDOUT),
			"ErrBadName":        reflect.ValueOf(&q.ErrBadName),
			"ErrBadStat":        reflect.ValueOf(&q.ErrBadStat),
			"ErrShortStat":      reflect.ValueOf(&q.ErrShortStat),
			"ForkLock":          reflect.ValueOf(&q.ForkLock),
			"SocketDisableIPv6": reflect.ValueOf(&q.SocketDisableIPv6),
			"Stderr":            reflect.ValueOf(&q.Stderr),
			"Stdin":             reflect.ValueOf(&q.Stdin),
			"Stdout":            reflect.ValueOf(&q.Stdout),
		},
		Funcs: map[string]reflect.Value{
			"Await":               reflect.ValueOf(q.Await),
			"Bind":                reflect.ValueOf(q.Bind),
			"BytePtrFromString":   reflect.ValueOf(q.BytePtrFromString),
			"ByteSliceFromString": reflect.ValueOf(q.ByteSliceFromString),
			"Chdir":               reflect.ValueOf(q.Chdir),
			"Clearenv":            reflect.ValueOf(q.Clearenv),
			"Close":               reflect.ValueOf(q.Close),
			"Create":              reflect.ValueOf(q.Create),
			"Dup":                 reflect.ValueOf(q.Dup),
			"Environ":             reflect.ValueOf(q.Environ),
			"Exec":                reflect.ValueOf(q.Exec),
			"Exit":                reflect.ValueOf(q.Exit),
			"Fchdir":              reflect.ValueOf(q.Fchdir),
			"Fd2path":             reflect.ValueOf(q.Fd2path),
			"Fixwd":               reflect.ValueOf(q.Fixwd),
			"ForkExec":            reflect.ValueOf(q.ForkExec),
			"Fstat":               reflect.ValueOf(q.Fstat),
			"Fwstat":              reflect.ValueOf(q.Fwstat),
			"Getegid":             reflect.ValueOf(q.Getegid),
			"Getenv":              reflect.ValueOf(q.Getenv),
			"Geteuid":             reflect.ValueOf(q.Geteuid),
			"Getgid":              reflect.ValueOf(q.Getgid),
			"Getgroups":           reflect.ValueOf(q.Getgroups),
			"Getpagesize":         reflect.ValueOf(q.Getpagesize),
			"Getpid":              reflect.ValueOf(q.Getpid),
			"Getppid":             reflect.ValueOf(q.Getppid),
			"Gettimeofday":        reflect.ValueOf(q.Gettimeofday),
			"Getuid":              reflect.ValueOf(q.Getuid),
			"Getwd":               reflect.ValueOf(q.Getwd),
			"Mkdir":               reflect.ValueOf(q.Mkdir),
			"Mount":               reflect.ValueOf(q.Mount),
			"NewError":            reflect.ValueOf(q.NewError),
			"NsecToTimeval":       reflect.ValueOf(q.NsecToTimeval),
			"Open":                reflect.ValueOf(q.Open),
			"Pipe":                reflect.ValueOf(q.Pipe),
			"Pread":               reflect.ValueOf(q.Pread),
			"Pwrite":              reflect.ValueOf(q.Pwrite),
			"RawSyscall":          reflect.ValueOf(q.RawSyscall),
			"RawSyscall6":         reflect.ValueOf(q.RawSyscall6),
			"Read":                reflect.ValueOf(q.Read),
			"Remove":              reflect.ValueOf(q.Remove),
			"Seek":                reflect.ValueOf(q.Seek),
			"Setenv":              reflect.ValueOf(q.Setenv),
			"SlicePtrFromStrings": reflect.ValueOf(q.SlicePtrFromStrings),
			"StartProcess":        reflect.ValueOf(q.StartProcess),
			"Stat":                reflect.ValueOf(q.Stat),
			"StringBytePtr":       reflect.ValueOf(q.StringBytePtr),
			"StringByteSlice":     reflect.ValueOf(q.StringByteSlice),
			"StringSlicePtr":      reflect.ValueOf(q.StringSlicePtr),
			"Syscall":             reflect.ValueOf(q.Syscall),
			"Syscall6":            reflect.ValueOf(q.Syscall6),
			"UnmarshalDir":        reflect.ValueOf(q.UnmarshalDir),
			"Unmount":             reflect.ValueOf(q.Unmount),
			"Unsetenv":            reflect.ValueOf(q.Unsetenv),
			"WaitProcess":         reflect.ValueOf(q.WaitProcess),
			"Write":               reflect.ValueOf(q.Write),
			"Wstat":               reflect.ValueOf(q.Wstat),
		},
		TypedConsts: map[string]igop.TypedConst{
			"SIGABRT": {reflect.TypeOf(q.SIGABRT), constant.MakeString(string(q.SIGABRT))},
			"SIGALRM": {reflect.TypeOf(q.SIGALRM), constant.MakeString(string(q.SIGALRM))},
			"SIGHUP":  {reflect.TypeOf(q.SIGHUP), constant.MakeString(string(q.SIGHUP))},
			"SIGINT":  {reflect.TypeOf(q.SIGINT), constant.MakeString(string(q.SIGINT))},
			"SIGKILL": {reflect.TypeOf(q.SIGKILL), constant.MakeString(string(q.SIGKILL))},
			"SIGTERM": {reflect.TypeOf(q.SIGTERM), constant.MakeString(string(q.SIGTERM))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"DMAPPEND":        {"untyped int", constant.MakeInt64(int64(q.DMAPPEND))},
			"DMAUTH":          {"untyped int", constant.MakeInt64(int64(q.DMAUTH))},
			"DMDIR":           {"untyped int", constant.MakeInt64(int64(q.DMDIR))},
			"DMEXCL":          {"untyped int", constant.MakeInt64(int64(q.DMEXCL))},
			"DMEXEC":          {"untyped int", constant.MakeInt64(int64(q.DMEXEC))},
			"DMMOUNT":         {"untyped int", constant.MakeInt64(int64(q.DMMOUNT))},
			"DMREAD":          {"untyped int", constant.MakeInt64(int64(q.DMREAD))},
			"DMTMP":           {"untyped int", constant.MakeInt64(int64(q.DMTMP))},
			"DMWRITE":         {"untyped int", constant.MakeInt64(int64(q.DMWRITE))},
			"ERRMAX":          {"untyped int", constant.MakeInt64(int64(q.ERRMAX))},
			"ImplementsGetwd": {"untyped bool", constant.MakeBool(bool(q.ImplementsGetwd))},
			"MAFTER":          {"untyped int", constant.MakeInt64(int64(q.MAFTER))},
			"MBEFORE":         {"untyped int", constant.MakeInt64(int64(q.MBEFORE))},
			"MCACHE":          {"untyped int", constant.MakeInt64(int64(q.MCACHE))},
			"MCREATE":         {"untyped int", constant.MakeInt64(int64(q.MCREATE))},
			"MMASK":           {"untyped int", constant.MakeInt64(int64(q.MMASK))},
			"MORDER":          {"untyped int", constant.MakeInt64(int64(q.MORDER))},
			"MREPL":           {"untyped int", constant.MakeInt64(int64(q.MREPL))},
			"O_APPEND":        {"untyped int", constant.MakeInt64(int64(q.O_APPEND))},
			"O_ASYNC":         {"untyped int", constant.MakeInt64(int64(q.O_ASYNC))},
			"O_CLOEXEC":       {"untyped int", constant.MakeInt64(int64(q.O_CLOEXEC))},
			"O_CREAT":         {"untyped int", constant.MakeInt64(int64(q.O_CREAT))},
			"O_EXCL":          {"untyped int", constant.MakeInt64(int64(q.O_EXCL))},
			"O_NOCTTY":        {"untyped int", constant.MakeInt64(int64(q.O_NOCTTY))},
			"O_NONBLOCK":      {"untyped int", constant.MakeInt64(int64(q.O_NONBLOCK))},
			"O_RDONLY":        {"untyped int", constant.MakeInt64(int64(q.O_RDONLY))},
			"O_RDWR":          {"untyped int", constant.MakeInt64(int64(q.O_RDWR))},
			"O_SYNC":          {"untyped int", constant.MakeInt64(int64(q.O_SYNC))},
			"O_TRUNC":         {"untyped int", constant.MakeInt64(int64(q.O_TRUNC))},
			"O_WRONLY":        {"untyped int", constant.MakeInt64(int64(q.O_WRONLY))},
			"QTAPPEND":        {"untyped int", constant.MakeInt64(int64(q.QTAPPEND))},
			"QTAUTH":          {"untyped int", constant.MakeInt64(int64(q.QTAUTH))},
			"QTDIR":           {"untyped int", constant.MakeInt64(int64(q.QTDIR))},
			"QTEXCL":          {"untyped int", constant.MakeInt64(int64(q.QTEXCL))},
			"QTFILE":          {"untyped int", constant.MakeInt64(int64(q.QTFILE))},
			"QTMOUNT":         {"untyped int", constant.MakeInt64(int64(q.QTMOUNT))},
			"QTTMP":           {"untyped int", constant.MakeInt64(int64(q.QTTMP))},
			"RFCENVG":         {"untyped int", constant.MakeInt64(int64(q.RFCENVG))},
			"RFCFDG":          {"untyped int", constant.MakeInt64(int64(q.RFCFDG))},
			"RFCNAMEG":        {"untyped int", constant.MakeInt64(int64(q.RFCNAMEG))},
			"RFENVG":          {"untyped int", constant.MakeInt64(int64(q.RFENVG))},
			"RFFDG":           {"untyped int", constant.MakeInt64(int64(q.RFFDG))},
			"RFMEM":           {"untyped int", constant.MakeInt64(int64(q.RFMEM))},
			"RFNAMEG":         {"untyped int", constant.MakeInt64(int64(q.RFNAMEG))},
			"RFNOMNT":         {"untyped int", constant.MakeInt64(int64(q.RFNOMNT))},
			"RFNOTEG":         {"untyped int", constant.MakeInt64(int64(q.RFNOTEG))},
			"RFNOWAIT":        {"untyped int", constant.MakeInt64(int64(q.RFNOWAIT))},
			"RFPROC":          {"untyped int", constant.MakeInt64(int64(q.RFPROC))},
			"RFREND":          {"untyped int", constant.MakeInt64(int64(q.RFREND))},
			"STATFIXLEN":      {"untyped int", constant.MakeInt64(int64(q.STATFIXLEN))},
			"STATMAX":         {"untyped int", constant.MakeInt64(int64(q.STATMAX))},
			"SYS_ALARM":       {"untyped int", constant.MakeInt64(int64(q.SYS_ALARM))},
			"SYS_AWAIT":       {"untyped int", constant.MakeInt64(int64(q.SYS_AWAIT))},
			"SYS_BIND":        {"untyped int", constant.MakeInt64(int64(q.SYS_BIND))},
			"SYS_BRK_":        {"untyped int", constant.MakeInt64(int64(q.SYS_BRK_))},
			"SYS_CHDIR":       {"untyped int", constant.MakeInt64(int64(q.SYS_CHDIR))},
			"SYS_CLOSE":       {"untyped int", constant.MakeInt64(int64(q.SYS_CLOSE))},
			"SYS_CREATE":      {"untyped int", constant.MakeInt64(int64(q.SYS_CREATE))},
			"SYS_DUP":         {"untyped int", constant.MakeInt64(int64(q.SYS_DUP))},
			"SYS_ERRSTR":      {"untyped int", constant.MakeInt64(int64(q.SYS_ERRSTR))},
			"SYS_EXEC":        {"untyped int", constant.MakeInt64(int64(q.SYS_EXEC))},
			"SYS_EXITS":       {"untyped int", constant.MakeInt64(int64(q.SYS_EXITS))},
			"SYS_FAUTH":       {"untyped int", constant.MakeInt64(int64(q.SYS_FAUTH))},
			"SYS_FD2PATH":     {"untyped int", constant.MakeInt64(int64(q.SYS_FD2PATH))},
			"SYS_FSTAT":       {"untyped int", constant.MakeInt64(int64(q.SYS_FSTAT))},
			"SYS_FVERSION":    {"untyped int", constant.MakeInt64(int64(q.SYS_FVERSION))},
			"SYS_FWSTAT":      {"untyped int", constant.MakeInt64(int64(q.SYS_FWSTAT))},
			"SYS_MOUNT":       {"untyped int", constant.MakeInt64(int64(q.SYS_MOUNT))},
			"SYS_NOTED":       {"untyped int", constant.MakeInt64(int64(q.SYS_NOTED))},
			"SYS_NOTIFY":      {"untyped int", constant.MakeInt64(int64(q.SYS_NOTIFY))},
			"SYS_NSEC":        {"untyped int", constant.MakeInt64(int64(q.SYS_NSEC))},
			"SYS_OPEN":        {"untyped int", constant.MakeInt64(int64(q.SYS_OPEN))},
			"SYS_OSEEK":       {"untyped int", constant.MakeInt64(int64(q.SYS_OSEEK))},
			"SYS_PIPE":        {"untyped int", constant.MakeInt64(int64(q.SYS_PIPE))},
			"SYS_PREAD":       {"untyped int", constant.MakeInt64(int64(q.SYS_PREAD))},
			"SYS_PWRITE":      {"untyped int", constant.MakeInt64(int64(q.SYS_PWRITE))},
			"SYS_REMOVE":      {"untyped int", constant.MakeInt64(int64(q.SYS_REMOVE))},
			"SYS_RENDEZVOUS":  {"untyped int", constant.MakeInt64(int64(q.SYS_RENDEZVOUS))},
			"SYS_RFORK":       {"untyped int", constant.MakeInt64(int64(q.SYS_RFORK))},
			"SYS_SEEK":        {"untyped int", constant.MakeInt64(int64(q.SYS_SEEK))},
			"SYS_SEGATTACH":   {"untyped int", constant.MakeInt64(int64(q.SYS_SEGATTACH))},
			"SYS_SEGBRK":      {"untyped int", constant.MakeInt64(int64(q.SYS_SEGBRK))},
			"SYS_SEGDETACH":   {"untyped int", constant.MakeInt64(int64(q.SYS_SEGDETACH))},
			"SYS_SEGFLUSH":    {"untyped int", constant.MakeInt64(int64(q.SYS_SEGFLUSH))},
			"SYS_SEGFREE":     {"untyped int", constant.MakeInt64(int64(q.SYS_SEGFREE))},
			"SYS_SEMACQUIRE":  {"untyped int", constant.MakeInt64(int64(q.SYS_SEMACQUIRE))},
			"SYS_SEMRELEASE":  {"untyped int", constant.MakeInt64(int64(q.SYS_SEMRELEASE))},
			"SYS_SLEEP":       {"untyped int", constant.MakeInt64(int64(q.SYS_SLEEP))},
			"SYS_STAT":        {"untyped int", constant.MakeInt64(int64(q.SYS_STAT))},
			"SYS_SYSR1":       {"untyped int", constant.MakeInt64(int64(q.SYS_SYSR1))},
			"SYS_TSEMACQUIRE": {"untyped int", constant.MakeInt64(int64(q.SYS_TSEMACQUIRE))},
			"SYS_UNMOUNT":     {"untyped int", constant.MakeInt64(int64(q.SYS_UNMOUNT))},
			"SYS_WSTAT":       {"untyped int", constant.MakeInt64(int64(q.SYS_WSTAT))},
			"S_IFBLK":         {"untyped int", constant.MakeInt64(int64(q.S_IFBLK))},
			"S_IFCHR":         {"untyped int", constant.MakeInt64(int64(q.S_IFCHR))},
			"S_IFDIR":         {"untyped int", constant.MakeInt64(int64(q.S_IFDIR))},
			"S_IFIFO":         {"untyped int", constant.MakeInt64(int64(q.S_IFIFO))},
			"S_IFLNK":         {"untyped int", constant.MakeInt64(int64(q.S_IFLNK))},
			"S_IFMT":          {"untyped int", constant.MakeInt64(int64(q.S_IFMT))},
			"S_IFREG":         {"untyped int", constant.MakeInt64(int64(q.S_IFREG))},
			"S_IFSOCK":        {"untyped int", constant.MakeInt64(int64(q.S_IFSOCK))},
		},
	})
}
