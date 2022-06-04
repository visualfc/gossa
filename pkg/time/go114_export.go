// export by github.com/goplus/igop/cmd/qexp

//+build go1.14,!go1.15

package time

import (
	q "time"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "time",
		Path: "time",
		Deps: map[string]string{
			"errors":  "errors",
			"runtime": "runtime",
			"sync":    "sync",
			"syscall": "syscall",
			"unsafe":  "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"Duration":   {reflect.TypeOf((*q.Duration)(nil)).Elem(), "Hours,Microseconds,Milliseconds,Minutes,Nanoseconds,Round,Seconds,String,Truncate", ""},
			"Location":   {reflect.TypeOf((*q.Location)(nil)).Elem(), "", "String,firstZoneUsed,get,lookup,lookupFirstZone,lookupName"},
			"Month":      {reflect.TypeOf((*q.Month)(nil)).Elem(), "String", ""},
			"ParseError": {reflect.TypeOf((*q.ParseError)(nil)).Elem(), "", "Error"},
			"Ticker":     {reflect.TypeOf((*q.Ticker)(nil)).Elem(), "", "Stop"},
			"Time":       {reflect.TypeOf((*q.Time)(nil)).Elem(), "Add,AddDate,After,AppendFormat,Before,Clock,Date,Day,Equal,Format,GobEncode,Hour,ISOWeek,In,IsZero,Local,Location,MarshalBinary,MarshalJSON,MarshalText,Minute,Month,Nanosecond,Round,Second,String,Sub,Truncate,UTC,Unix,UnixNano,Weekday,Year,YearDay,Zone,abs,date,locabs", "GobDecode,UnmarshalBinary,UnmarshalJSON,UnmarshalText,addSec,mono,nsec,sec,setLoc,setMono,stripMono,unixSec"},
			"Timer":      {reflect.TypeOf((*q.Timer)(nil)).Elem(), "", "Reset,Stop"},
			"Weekday":    {reflect.TypeOf((*q.Weekday)(nil)).Elem(), "String", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Local": reflect.ValueOf(&q.Local),
			"UTC":   reflect.ValueOf(&q.UTC),
		},
		Funcs: map[string]reflect.Value{
			"After":                  reflect.ValueOf(q.After),
			"AfterFunc":              reflect.ValueOf(q.AfterFunc),
			"Date":                   reflect.ValueOf(q.Date),
			"FixedZone":              reflect.ValueOf(q.FixedZone),
			"LoadLocation":           reflect.ValueOf(q.LoadLocation),
			"LoadLocationFromTZData": reflect.ValueOf(q.LoadLocationFromTZData),
			"NewTicker":              reflect.ValueOf(q.NewTicker),
			"NewTimer":               reflect.ValueOf(q.NewTimer),
			"Now":                    reflect.ValueOf(q.Now),
			"Parse":                  reflect.ValueOf(q.Parse),
			"ParseDuration":          reflect.ValueOf(q.ParseDuration),
			"ParseInLocation":        reflect.ValueOf(q.ParseInLocation),
			"Since":                  reflect.ValueOf(q.Since),
			"Sleep":                  reflect.ValueOf(q.Sleep),
			"Tick":                   reflect.ValueOf(q.Tick),
			"Unix":                   reflect.ValueOf(q.Unix),
			"Until":                  reflect.ValueOf(q.Until),
		},
		TypedConsts: map[string]igop.TypedConst{
			"April":       {reflect.TypeOf(q.April), constant.MakeInt64(int64(q.April))},
			"August":      {reflect.TypeOf(q.August), constant.MakeInt64(int64(q.August))},
			"December":    {reflect.TypeOf(q.December), constant.MakeInt64(int64(q.December))},
			"February":    {reflect.TypeOf(q.February), constant.MakeInt64(int64(q.February))},
			"Friday":      {reflect.TypeOf(q.Friday), constant.MakeInt64(int64(q.Friday))},
			"Hour":        {reflect.TypeOf(q.Hour), constant.MakeInt64(int64(q.Hour))},
			"January":     {reflect.TypeOf(q.January), constant.MakeInt64(int64(q.January))},
			"July":        {reflect.TypeOf(q.July), constant.MakeInt64(int64(q.July))},
			"June":        {reflect.TypeOf(q.June), constant.MakeInt64(int64(q.June))},
			"March":       {reflect.TypeOf(q.March), constant.MakeInt64(int64(q.March))},
			"May":         {reflect.TypeOf(q.May), constant.MakeInt64(int64(q.May))},
			"Microsecond": {reflect.TypeOf(q.Microsecond), constant.MakeInt64(int64(q.Microsecond))},
			"Millisecond": {reflect.TypeOf(q.Millisecond), constant.MakeInt64(int64(q.Millisecond))},
			"Minute":      {reflect.TypeOf(q.Minute), constant.MakeInt64(int64(q.Minute))},
			"Monday":      {reflect.TypeOf(q.Monday), constant.MakeInt64(int64(q.Monday))},
			"Nanosecond":  {reflect.TypeOf(q.Nanosecond), constant.MakeInt64(int64(q.Nanosecond))},
			"November":    {reflect.TypeOf(q.November), constant.MakeInt64(int64(q.November))},
			"October":     {reflect.TypeOf(q.October), constant.MakeInt64(int64(q.October))},
			"Saturday":    {reflect.TypeOf(q.Saturday), constant.MakeInt64(int64(q.Saturday))},
			"Second":      {reflect.TypeOf(q.Second), constant.MakeInt64(int64(q.Second))},
			"September":   {reflect.TypeOf(q.September), constant.MakeInt64(int64(q.September))},
			"Sunday":      {reflect.TypeOf(q.Sunday), constant.MakeInt64(int64(q.Sunday))},
			"Thursday":    {reflect.TypeOf(q.Thursday), constant.MakeInt64(int64(q.Thursday))},
			"Tuesday":     {reflect.TypeOf(q.Tuesday), constant.MakeInt64(int64(q.Tuesday))},
			"Wednesday":   {reflect.TypeOf(q.Wednesday), constant.MakeInt64(int64(q.Wednesday))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"ANSIC":       {"untyped string", constant.MakeString(string(q.ANSIC))},
			"Kitchen":     {"untyped string", constant.MakeString(string(q.Kitchen))},
			"RFC1123":     {"untyped string", constant.MakeString(string(q.RFC1123))},
			"RFC1123Z":    {"untyped string", constant.MakeString(string(q.RFC1123Z))},
			"RFC3339":     {"untyped string", constant.MakeString(string(q.RFC3339))},
			"RFC3339Nano": {"untyped string", constant.MakeString(string(q.RFC3339Nano))},
			"RFC822":      {"untyped string", constant.MakeString(string(q.RFC822))},
			"RFC822Z":     {"untyped string", constant.MakeString(string(q.RFC822Z))},
			"RFC850":      {"untyped string", constant.MakeString(string(q.RFC850))},
			"RubyDate":    {"untyped string", constant.MakeString(string(q.RubyDate))},
			"Stamp":       {"untyped string", constant.MakeString(string(q.Stamp))},
			"StampMicro":  {"untyped string", constant.MakeString(string(q.StampMicro))},
			"StampMilli":  {"untyped string", constant.MakeString(string(q.StampMilli))},
			"StampNano":   {"untyped string", constant.MakeString(string(q.StampNano))},
			"UnixDate":    {"untyped string", constant.MakeString(string(q.UnixDate))},
		},
	})
}
