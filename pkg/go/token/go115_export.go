// export by github.com/goplus/igop/cmd/qexp

//+build go1.15,!go1.16

package token

import (
	q "go/token"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "token",
		Path: "go/token",
		Deps: map[string]string{
			"fmt":          "fmt",
			"sort":         "sort",
			"strconv":      "strconv",
			"sync":         "sync",
			"unicode":      "unicode",
			"unicode/utf8": "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"File":     {reflect.TypeOf((*q.File)(nil)).Elem(), "", "AddLine,AddLineColumnInfo,AddLineInfo,Base,Line,LineCount,LineStart,MergeLine,Name,Offset,Pos,Position,PositionFor,SetLines,SetLinesForContent,Size,position,unpack"},
			"FileSet":  {reflect.TypeOf((*q.FileSet)(nil)).Elem(), "", "AddFile,Base,File,Iterate,Position,PositionFor,Read,Write,file"},
			"Pos":      {reflect.TypeOf((*q.Pos)(nil)).Elem(), "IsValid", ""},
			"Position": {reflect.TypeOf((*q.Position)(nil)).Elem(), "String", "IsValid"},
			"Token":    {reflect.TypeOf((*q.Token)(nil)).Elem(), "IsKeyword,IsLiteral,IsOperator,Precedence,String", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"IsExported":   reflect.ValueOf(q.IsExported),
			"IsIdentifier": reflect.ValueOf(q.IsIdentifier),
			"IsKeyword":    reflect.ValueOf(q.IsKeyword),
			"Lookup":       reflect.ValueOf(q.Lookup),
			"NewFileSet":   reflect.ValueOf(q.NewFileSet),
		},
		TypedConsts: map[string]igop.TypedConst{
			"ADD":            {reflect.TypeOf(q.ADD), constant.MakeInt64(int64(q.ADD))},
			"ADD_ASSIGN":     {reflect.TypeOf(q.ADD_ASSIGN), constant.MakeInt64(int64(q.ADD_ASSIGN))},
			"AND":            {reflect.TypeOf(q.AND), constant.MakeInt64(int64(q.AND))},
			"AND_ASSIGN":     {reflect.TypeOf(q.AND_ASSIGN), constant.MakeInt64(int64(q.AND_ASSIGN))},
			"AND_NOT":        {reflect.TypeOf(q.AND_NOT), constant.MakeInt64(int64(q.AND_NOT))},
			"AND_NOT_ASSIGN": {reflect.TypeOf(q.AND_NOT_ASSIGN), constant.MakeInt64(int64(q.AND_NOT_ASSIGN))},
			"ARROW":          {reflect.TypeOf(q.ARROW), constant.MakeInt64(int64(q.ARROW))},
			"ASSIGN":         {reflect.TypeOf(q.ASSIGN), constant.MakeInt64(int64(q.ASSIGN))},
			"BREAK":          {reflect.TypeOf(q.BREAK), constant.MakeInt64(int64(q.BREAK))},
			"CASE":           {reflect.TypeOf(q.CASE), constant.MakeInt64(int64(q.CASE))},
			"CHAN":           {reflect.TypeOf(q.CHAN), constant.MakeInt64(int64(q.CHAN))},
			"CHAR":           {reflect.TypeOf(q.CHAR), constant.MakeInt64(int64(q.CHAR))},
			"COLON":          {reflect.TypeOf(q.COLON), constant.MakeInt64(int64(q.COLON))},
			"COMMA":          {reflect.TypeOf(q.COMMA), constant.MakeInt64(int64(q.COMMA))},
			"COMMENT":        {reflect.TypeOf(q.COMMENT), constant.MakeInt64(int64(q.COMMENT))},
			"CONST":          {reflect.TypeOf(q.CONST), constant.MakeInt64(int64(q.CONST))},
			"CONTINUE":       {reflect.TypeOf(q.CONTINUE), constant.MakeInt64(int64(q.CONTINUE))},
			"DEC":            {reflect.TypeOf(q.DEC), constant.MakeInt64(int64(q.DEC))},
			"DEFAULT":        {reflect.TypeOf(q.DEFAULT), constant.MakeInt64(int64(q.DEFAULT))},
			"DEFER":          {reflect.TypeOf(q.DEFER), constant.MakeInt64(int64(q.DEFER))},
			"DEFINE":         {reflect.TypeOf(q.DEFINE), constant.MakeInt64(int64(q.DEFINE))},
			"ELLIPSIS":       {reflect.TypeOf(q.ELLIPSIS), constant.MakeInt64(int64(q.ELLIPSIS))},
			"ELSE":           {reflect.TypeOf(q.ELSE), constant.MakeInt64(int64(q.ELSE))},
			"EOF":            {reflect.TypeOf(q.EOF), constant.MakeInt64(int64(q.EOF))},
			"EQL":            {reflect.TypeOf(q.EQL), constant.MakeInt64(int64(q.EQL))},
			"FALLTHROUGH":    {reflect.TypeOf(q.FALLTHROUGH), constant.MakeInt64(int64(q.FALLTHROUGH))},
			"FLOAT":          {reflect.TypeOf(q.FLOAT), constant.MakeInt64(int64(q.FLOAT))},
			"FOR":            {reflect.TypeOf(q.FOR), constant.MakeInt64(int64(q.FOR))},
			"FUNC":           {reflect.TypeOf(q.FUNC), constant.MakeInt64(int64(q.FUNC))},
			"GEQ":            {reflect.TypeOf(q.GEQ), constant.MakeInt64(int64(q.GEQ))},
			"GO":             {reflect.TypeOf(q.GO), constant.MakeInt64(int64(q.GO))},
			"GOTO":           {reflect.TypeOf(q.GOTO), constant.MakeInt64(int64(q.GOTO))},
			"GTR":            {reflect.TypeOf(q.GTR), constant.MakeInt64(int64(q.GTR))},
			"IDENT":          {reflect.TypeOf(q.IDENT), constant.MakeInt64(int64(q.IDENT))},
			"IF":             {reflect.TypeOf(q.IF), constant.MakeInt64(int64(q.IF))},
			"ILLEGAL":        {reflect.TypeOf(q.ILLEGAL), constant.MakeInt64(int64(q.ILLEGAL))},
			"IMAG":           {reflect.TypeOf(q.IMAG), constant.MakeInt64(int64(q.IMAG))},
			"IMPORT":         {reflect.TypeOf(q.IMPORT), constant.MakeInt64(int64(q.IMPORT))},
			"INC":            {reflect.TypeOf(q.INC), constant.MakeInt64(int64(q.INC))},
			"INT":            {reflect.TypeOf(q.INT), constant.MakeInt64(int64(q.INT))},
			"INTERFACE":      {reflect.TypeOf(q.INTERFACE), constant.MakeInt64(int64(q.INTERFACE))},
			"LAND":           {reflect.TypeOf(q.LAND), constant.MakeInt64(int64(q.LAND))},
			"LBRACE":         {reflect.TypeOf(q.LBRACE), constant.MakeInt64(int64(q.LBRACE))},
			"LBRACK":         {reflect.TypeOf(q.LBRACK), constant.MakeInt64(int64(q.LBRACK))},
			"LEQ":            {reflect.TypeOf(q.LEQ), constant.MakeInt64(int64(q.LEQ))},
			"LOR":            {reflect.TypeOf(q.LOR), constant.MakeInt64(int64(q.LOR))},
			"LPAREN":         {reflect.TypeOf(q.LPAREN), constant.MakeInt64(int64(q.LPAREN))},
			"LSS":            {reflect.TypeOf(q.LSS), constant.MakeInt64(int64(q.LSS))},
			"MAP":            {reflect.TypeOf(q.MAP), constant.MakeInt64(int64(q.MAP))},
			"MUL":            {reflect.TypeOf(q.MUL), constant.MakeInt64(int64(q.MUL))},
			"MUL_ASSIGN":     {reflect.TypeOf(q.MUL_ASSIGN), constant.MakeInt64(int64(q.MUL_ASSIGN))},
			"NEQ":            {reflect.TypeOf(q.NEQ), constant.MakeInt64(int64(q.NEQ))},
			"NOT":            {reflect.TypeOf(q.NOT), constant.MakeInt64(int64(q.NOT))},
			"NoPos":          {reflect.TypeOf(q.NoPos), constant.MakeInt64(int64(q.NoPos))},
			"OR":             {reflect.TypeOf(q.OR), constant.MakeInt64(int64(q.OR))},
			"OR_ASSIGN":      {reflect.TypeOf(q.OR_ASSIGN), constant.MakeInt64(int64(q.OR_ASSIGN))},
			"PACKAGE":        {reflect.TypeOf(q.PACKAGE), constant.MakeInt64(int64(q.PACKAGE))},
			"PERIOD":         {reflect.TypeOf(q.PERIOD), constant.MakeInt64(int64(q.PERIOD))},
			"QUO":            {reflect.TypeOf(q.QUO), constant.MakeInt64(int64(q.QUO))},
			"QUO_ASSIGN":     {reflect.TypeOf(q.QUO_ASSIGN), constant.MakeInt64(int64(q.QUO_ASSIGN))},
			"RANGE":          {reflect.TypeOf(q.RANGE), constant.MakeInt64(int64(q.RANGE))},
			"RBRACE":         {reflect.TypeOf(q.RBRACE), constant.MakeInt64(int64(q.RBRACE))},
			"RBRACK":         {reflect.TypeOf(q.RBRACK), constant.MakeInt64(int64(q.RBRACK))},
			"REM":            {reflect.TypeOf(q.REM), constant.MakeInt64(int64(q.REM))},
			"REM_ASSIGN":     {reflect.TypeOf(q.REM_ASSIGN), constant.MakeInt64(int64(q.REM_ASSIGN))},
			"RETURN":         {reflect.TypeOf(q.RETURN), constant.MakeInt64(int64(q.RETURN))},
			"RPAREN":         {reflect.TypeOf(q.RPAREN), constant.MakeInt64(int64(q.RPAREN))},
			"SELECT":         {reflect.TypeOf(q.SELECT), constant.MakeInt64(int64(q.SELECT))},
			"SEMICOLON":      {reflect.TypeOf(q.SEMICOLON), constant.MakeInt64(int64(q.SEMICOLON))},
			"SHL":            {reflect.TypeOf(q.SHL), constant.MakeInt64(int64(q.SHL))},
			"SHL_ASSIGN":     {reflect.TypeOf(q.SHL_ASSIGN), constant.MakeInt64(int64(q.SHL_ASSIGN))},
			"SHR":            {reflect.TypeOf(q.SHR), constant.MakeInt64(int64(q.SHR))},
			"SHR_ASSIGN":     {reflect.TypeOf(q.SHR_ASSIGN), constant.MakeInt64(int64(q.SHR_ASSIGN))},
			"STRING":         {reflect.TypeOf(q.STRING), constant.MakeInt64(int64(q.STRING))},
			"STRUCT":         {reflect.TypeOf(q.STRUCT), constant.MakeInt64(int64(q.STRUCT))},
			"SUB":            {reflect.TypeOf(q.SUB), constant.MakeInt64(int64(q.SUB))},
			"SUB_ASSIGN":     {reflect.TypeOf(q.SUB_ASSIGN), constant.MakeInt64(int64(q.SUB_ASSIGN))},
			"SWITCH":         {reflect.TypeOf(q.SWITCH), constant.MakeInt64(int64(q.SWITCH))},
			"TYPE":           {reflect.TypeOf(q.TYPE), constant.MakeInt64(int64(q.TYPE))},
			"VAR":            {reflect.TypeOf(q.VAR), constant.MakeInt64(int64(q.VAR))},
			"XOR":            {reflect.TypeOf(q.XOR), constant.MakeInt64(int64(q.XOR))},
			"XOR_ASSIGN":     {reflect.TypeOf(q.XOR_ASSIGN), constant.MakeInt64(int64(q.XOR_ASSIGN))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"HighestPrec": {"untyped int", constant.MakeInt64(int64(q.HighestPrec))},
			"LowestPrec":  {"untyped int", constant.MakeInt64(int64(q.LowestPrec))},
			"UnaryPrec":   {"untyped int", constant.MakeInt64(int64(q.UnaryPrec))},
		},
	})
}
