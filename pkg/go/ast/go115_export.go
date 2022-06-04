// export by github.com/goplus/igop/cmd/qexp

//+build go1.15,!go1.16

package ast

import (
	q "go/ast"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "ast",
		Path: "go/ast",
		Deps: map[string]string{
			"bytes":      "bytes",
			"fmt":        "fmt",
			"go/scanner": "scanner",
			"go/token":   "token",
			"io":         "io",
			"os":         "os",
			"reflect":    "reflect",
			"sort":       "sort",
			"strconv":    "strconv",
			"strings":    "strings",
		},
		Interfaces: map[string]reflect.Type{
			"Decl":    reflect.TypeOf((*q.Decl)(nil)).Elem(),
			"Expr":    reflect.TypeOf((*q.Expr)(nil)).Elem(),
			"Node":    reflect.TypeOf((*q.Node)(nil)).Elem(),
			"Spec":    reflect.TypeOf((*q.Spec)(nil)).Elem(),
			"Stmt":    reflect.TypeOf((*q.Stmt)(nil)).Elem(),
			"Visitor": reflect.TypeOf((*q.Visitor)(nil)).Elem(),
		},
		NamedTypes: map[string]igop.NamedType{
			"ArrayType":      {reflect.TypeOf((*q.ArrayType)(nil)).Elem(), "", "End,Pos,exprNode"},
			"AssignStmt":     {reflect.TypeOf((*q.AssignStmt)(nil)).Elem(), "", "End,Pos,stmtNode"},
			"BadDecl":        {reflect.TypeOf((*q.BadDecl)(nil)).Elem(), "", "End,Pos,declNode"},
			"BadExpr":        {reflect.TypeOf((*q.BadExpr)(nil)).Elem(), "", "End,Pos,exprNode"},
			"BadStmt":        {reflect.TypeOf((*q.BadStmt)(nil)).Elem(), "", "End,Pos,stmtNode"},
			"BasicLit":       {reflect.TypeOf((*q.BasicLit)(nil)).Elem(), "", "End,Pos,exprNode"},
			"BinaryExpr":     {reflect.TypeOf((*q.BinaryExpr)(nil)).Elem(), "", "End,Pos,exprNode"},
			"BlockStmt":      {reflect.TypeOf((*q.BlockStmt)(nil)).Elem(), "", "End,Pos,stmtNode"},
			"BranchStmt":     {reflect.TypeOf((*q.BranchStmt)(nil)).Elem(), "", "End,Pos,stmtNode"},
			"CallExpr":       {reflect.TypeOf((*q.CallExpr)(nil)).Elem(), "", "End,Pos,exprNode"},
			"CaseClause":     {reflect.TypeOf((*q.CaseClause)(nil)).Elem(), "", "End,Pos,stmtNode"},
			"ChanDir":        {reflect.TypeOf((*q.ChanDir)(nil)).Elem(), "", ""},
			"ChanType":       {reflect.TypeOf((*q.ChanType)(nil)).Elem(), "", "End,Pos,exprNode"},
			"CommClause":     {reflect.TypeOf((*q.CommClause)(nil)).Elem(), "", "End,Pos,stmtNode"},
			"Comment":        {reflect.TypeOf((*q.Comment)(nil)).Elem(), "", "End,Pos"},
			"CommentGroup":   {reflect.TypeOf((*q.CommentGroup)(nil)).Elem(), "", "End,Pos,Text"},
			"CommentMap":     {reflect.TypeOf((*q.CommentMap)(nil)).Elem(), "Comments,Filter,String,Update,addComment", ""},
			"CompositeLit":   {reflect.TypeOf((*q.CompositeLit)(nil)).Elem(), "", "End,Pos,exprNode"},
			"DeclStmt":       {reflect.TypeOf((*q.DeclStmt)(nil)).Elem(), "", "End,Pos,stmtNode"},
			"DeferStmt":      {reflect.TypeOf((*q.DeferStmt)(nil)).Elem(), "", "End,Pos,stmtNode"},
			"Ellipsis":       {reflect.TypeOf((*q.Ellipsis)(nil)).Elem(), "", "End,Pos,exprNode"},
			"EmptyStmt":      {reflect.TypeOf((*q.EmptyStmt)(nil)).Elem(), "", "End,Pos,stmtNode"},
			"ExprStmt":       {reflect.TypeOf((*q.ExprStmt)(nil)).Elem(), "", "End,Pos,stmtNode"},
			"Field":          {reflect.TypeOf((*q.Field)(nil)).Elem(), "", "End,Pos"},
			"FieldFilter":    {reflect.TypeOf((*q.FieldFilter)(nil)).Elem(), "", ""},
			"FieldList":      {reflect.TypeOf((*q.FieldList)(nil)).Elem(), "", "End,NumFields,Pos"},
			"File":           {reflect.TypeOf((*q.File)(nil)).Elem(), "", "End,Pos"},
			"Filter":         {reflect.TypeOf((*q.Filter)(nil)).Elem(), "", ""},
			"ForStmt":        {reflect.TypeOf((*q.ForStmt)(nil)).Elem(), "", "End,Pos,stmtNode"},
			"FuncDecl":       {reflect.TypeOf((*q.FuncDecl)(nil)).Elem(), "", "End,Pos,declNode"},
			"FuncLit":        {reflect.TypeOf((*q.FuncLit)(nil)).Elem(), "", "End,Pos,exprNode"},
			"FuncType":       {reflect.TypeOf((*q.FuncType)(nil)).Elem(), "", "End,Pos,exprNode"},
			"GenDecl":        {reflect.TypeOf((*q.GenDecl)(nil)).Elem(), "", "End,Pos,declNode"},
			"GoStmt":         {reflect.TypeOf((*q.GoStmt)(nil)).Elem(), "", "End,Pos,stmtNode"},
			"Ident":          {reflect.TypeOf((*q.Ident)(nil)).Elem(), "", "End,IsExported,Pos,String,exprNode"},
			"IfStmt":         {reflect.TypeOf((*q.IfStmt)(nil)).Elem(), "", "End,Pos,stmtNode"},
			"ImportSpec":     {reflect.TypeOf((*q.ImportSpec)(nil)).Elem(), "", "End,Pos,specNode"},
			"Importer":       {reflect.TypeOf((*q.Importer)(nil)).Elem(), "", ""},
			"IncDecStmt":     {reflect.TypeOf((*q.IncDecStmt)(nil)).Elem(), "", "End,Pos,stmtNode"},
			"IndexExpr":      {reflect.TypeOf((*q.IndexExpr)(nil)).Elem(), "", "End,Pos,exprNode"},
			"InterfaceType":  {reflect.TypeOf((*q.InterfaceType)(nil)).Elem(), "", "End,Pos,exprNode"},
			"KeyValueExpr":   {reflect.TypeOf((*q.KeyValueExpr)(nil)).Elem(), "", "End,Pos,exprNode"},
			"LabeledStmt":    {reflect.TypeOf((*q.LabeledStmt)(nil)).Elem(), "", "End,Pos,stmtNode"},
			"MapType":        {reflect.TypeOf((*q.MapType)(nil)).Elem(), "", "End,Pos,exprNode"},
			"MergeMode":      {reflect.TypeOf((*q.MergeMode)(nil)).Elem(), "", ""},
			"ObjKind":        {reflect.TypeOf((*q.ObjKind)(nil)).Elem(), "String", ""},
			"Object":         {reflect.TypeOf((*q.Object)(nil)).Elem(), "", "Pos"},
			"Package":        {reflect.TypeOf((*q.Package)(nil)).Elem(), "", "End,Pos"},
			"ParenExpr":      {reflect.TypeOf((*q.ParenExpr)(nil)).Elem(), "", "End,Pos,exprNode"},
			"RangeStmt":      {reflect.TypeOf((*q.RangeStmt)(nil)).Elem(), "", "End,Pos,stmtNode"},
			"ReturnStmt":     {reflect.TypeOf((*q.ReturnStmt)(nil)).Elem(), "", "End,Pos,stmtNode"},
			"Scope":          {reflect.TypeOf((*q.Scope)(nil)).Elem(), "", "Insert,Lookup,String"},
			"SelectStmt":     {reflect.TypeOf((*q.SelectStmt)(nil)).Elem(), "", "End,Pos,stmtNode"},
			"SelectorExpr":   {reflect.TypeOf((*q.SelectorExpr)(nil)).Elem(), "", "End,Pos,exprNode"},
			"SendStmt":       {reflect.TypeOf((*q.SendStmt)(nil)).Elem(), "", "End,Pos,stmtNode"},
			"SliceExpr":      {reflect.TypeOf((*q.SliceExpr)(nil)).Elem(), "", "End,Pos,exprNode"},
			"StarExpr":       {reflect.TypeOf((*q.StarExpr)(nil)).Elem(), "", "End,Pos,exprNode"},
			"StructType":     {reflect.TypeOf((*q.StructType)(nil)).Elem(), "", "End,Pos,exprNode"},
			"SwitchStmt":     {reflect.TypeOf((*q.SwitchStmt)(nil)).Elem(), "", "End,Pos,stmtNode"},
			"TypeAssertExpr": {reflect.TypeOf((*q.TypeAssertExpr)(nil)).Elem(), "", "End,Pos,exprNode"},
			"TypeSpec":       {reflect.TypeOf((*q.TypeSpec)(nil)).Elem(), "", "End,Pos,specNode"},
			"TypeSwitchStmt": {reflect.TypeOf((*q.TypeSwitchStmt)(nil)).Elem(), "", "End,Pos,stmtNode"},
			"UnaryExpr":      {reflect.TypeOf((*q.UnaryExpr)(nil)).Elem(), "", "End,Pos,exprNode"},
			"ValueSpec":      {reflect.TypeOf((*q.ValueSpec)(nil)).Elem(), "", "End,Pos,specNode"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"FileExports":       reflect.ValueOf(q.FileExports),
			"FilterDecl":        reflect.ValueOf(q.FilterDecl),
			"FilterFile":        reflect.ValueOf(q.FilterFile),
			"FilterPackage":     reflect.ValueOf(q.FilterPackage),
			"Fprint":            reflect.ValueOf(q.Fprint),
			"Inspect":           reflect.ValueOf(q.Inspect),
			"IsExported":        reflect.ValueOf(q.IsExported),
			"MergePackageFiles": reflect.ValueOf(q.MergePackageFiles),
			"NewCommentMap":     reflect.ValueOf(q.NewCommentMap),
			"NewIdent":          reflect.ValueOf(q.NewIdent),
			"NewObj":            reflect.ValueOf(q.NewObj),
			"NewPackage":        reflect.ValueOf(q.NewPackage),
			"NewScope":          reflect.ValueOf(q.NewScope),
			"NotNilFilter":      reflect.ValueOf(q.NotNilFilter),
			"PackageExports":    reflect.ValueOf(q.PackageExports),
			"Print":             reflect.ValueOf(q.Print),
			"SortImports":       reflect.ValueOf(q.SortImports),
			"Walk":              reflect.ValueOf(q.Walk),
		},
		TypedConsts: map[string]igop.TypedConst{
			"Bad":                        {reflect.TypeOf(q.Bad), constant.MakeInt64(int64(q.Bad))},
			"Con":                        {reflect.TypeOf(q.Con), constant.MakeInt64(int64(q.Con))},
			"FilterFuncDuplicates":       {reflect.TypeOf(q.FilterFuncDuplicates), constant.MakeInt64(int64(q.FilterFuncDuplicates))},
			"FilterImportDuplicates":     {reflect.TypeOf(q.FilterImportDuplicates), constant.MakeInt64(int64(q.FilterImportDuplicates))},
			"FilterUnassociatedComments": {reflect.TypeOf(q.FilterUnassociatedComments), constant.MakeInt64(int64(q.FilterUnassociatedComments))},
			"Fun":                        {reflect.TypeOf(q.Fun), constant.MakeInt64(int64(q.Fun))},
			"Lbl":                        {reflect.TypeOf(q.Lbl), constant.MakeInt64(int64(q.Lbl))},
			"Pkg":                        {reflect.TypeOf(q.Pkg), constant.MakeInt64(int64(q.Pkg))},
			"RECV":                       {reflect.TypeOf(q.RECV), constant.MakeInt64(int64(q.RECV))},
			"SEND":                       {reflect.TypeOf(q.SEND), constant.MakeInt64(int64(q.SEND))},
			"Typ":                        {reflect.TypeOf(q.Typ), constant.MakeInt64(int64(q.Typ))},
			"Var":                        {reflect.TypeOf(q.Var), constant.MakeInt64(int64(q.Var))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
