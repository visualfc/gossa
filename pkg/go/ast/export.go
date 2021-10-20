// export by github.com/goplus/gossa/cmd/qexp

package ast

import (
	"go/ast"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("go/ast", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*go/ast.ArrayType).End":       (*ast.ArrayType).End,
	"(*go/ast.ArrayType).Pos":       (*ast.ArrayType).Pos,
	"(*go/ast.AssignStmt).End":      (*ast.AssignStmt).End,
	"(*go/ast.AssignStmt).Pos":      (*ast.AssignStmt).Pos,
	"(*go/ast.BadDecl).End":         (*ast.BadDecl).End,
	"(*go/ast.BadDecl).Pos":         (*ast.BadDecl).Pos,
	"(*go/ast.BadExpr).End":         (*ast.BadExpr).End,
	"(*go/ast.BadExpr).Pos":         (*ast.BadExpr).Pos,
	"(*go/ast.BadStmt).End":         (*ast.BadStmt).End,
	"(*go/ast.BadStmt).Pos":         (*ast.BadStmt).Pos,
	"(*go/ast.BasicLit).End":        (*ast.BasicLit).End,
	"(*go/ast.BasicLit).Pos":        (*ast.BasicLit).Pos,
	"(*go/ast.BinaryExpr).End":      (*ast.BinaryExpr).End,
	"(*go/ast.BinaryExpr).Pos":      (*ast.BinaryExpr).Pos,
	"(*go/ast.BlockStmt).End":       (*ast.BlockStmt).End,
	"(*go/ast.BlockStmt).Pos":       (*ast.BlockStmt).Pos,
	"(*go/ast.BranchStmt).End":      (*ast.BranchStmt).End,
	"(*go/ast.BranchStmt).Pos":      (*ast.BranchStmt).Pos,
	"(*go/ast.CallExpr).End":        (*ast.CallExpr).End,
	"(*go/ast.CallExpr).Pos":        (*ast.CallExpr).Pos,
	"(*go/ast.CaseClause).End":      (*ast.CaseClause).End,
	"(*go/ast.CaseClause).Pos":      (*ast.CaseClause).Pos,
	"(*go/ast.ChanType).End":        (*ast.ChanType).End,
	"(*go/ast.ChanType).Pos":        (*ast.ChanType).Pos,
	"(*go/ast.CommClause).End":      (*ast.CommClause).End,
	"(*go/ast.CommClause).Pos":      (*ast.CommClause).Pos,
	"(*go/ast.Comment).End":         (*ast.Comment).End,
	"(*go/ast.Comment).Pos":         (*ast.Comment).Pos,
	"(*go/ast.CommentGroup).End":    (*ast.CommentGroup).End,
	"(*go/ast.CommentGroup).Pos":    (*ast.CommentGroup).Pos,
	"(*go/ast.CommentGroup).Text":   (*ast.CommentGroup).Text,
	"(*go/ast.CompositeLit).End":    (*ast.CompositeLit).End,
	"(*go/ast.CompositeLit).Pos":    (*ast.CompositeLit).Pos,
	"(*go/ast.DeclStmt).End":        (*ast.DeclStmt).End,
	"(*go/ast.DeclStmt).Pos":        (*ast.DeclStmt).Pos,
	"(*go/ast.DeferStmt).End":       (*ast.DeferStmt).End,
	"(*go/ast.DeferStmt).Pos":       (*ast.DeferStmt).Pos,
	"(*go/ast.Ellipsis).End":        (*ast.Ellipsis).End,
	"(*go/ast.Ellipsis).Pos":        (*ast.Ellipsis).Pos,
	"(*go/ast.EmptyStmt).End":       (*ast.EmptyStmt).End,
	"(*go/ast.EmptyStmt).Pos":       (*ast.EmptyStmt).Pos,
	"(*go/ast.ExprStmt).End":        (*ast.ExprStmt).End,
	"(*go/ast.ExprStmt).Pos":        (*ast.ExprStmt).Pos,
	"(*go/ast.Field).End":           (*ast.Field).End,
	"(*go/ast.Field).Pos":           (*ast.Field).Pos,
	"(*go/ast.FieldList).End":       (*ast.FieldList).End,
	"(*go/ast.FieldList).NumFields": (*ast.FieldList).NumFields,
	"(*go/ast.FieldList).Pos":       (*ast.FieldList).Pos,
	"(*go/ast.File).End":            (*ast.File).End,
	"(*go/ast.File).Pos":            (*ast.File).Pos,
	"(*go/ast.ForStmt).End":         (*ast.ForStmt).End,
	"(*go/ast.ForStmt).Pos":         (*ast.ForStmt).Pos,
	"(*go/ast.FuncDecl).End":        (*ast.FuncDecl).End,
	"(*go/ast.FuncDecl).Pos":        (*ast.FuncDecl).Pos,
	"(*go/ast.FuncLit).End":         (*ast.FuncLit).End,
	"(*go/ast.FuncLit).Pos":         (*ast.FuncLit).Pos,
	"(*go/ast.FuncType).End":        (*ast.FuncType).End,
	"(*go/ast.FuncType).Pos":        (*ast.FuncType).Pos,
	"(*go/ast.GenDecl).End":         (*ast.GenDecl).End,
	"(*go/ast.GenDecl).Pos":         (*ast.GenDecl).Pos,
	"(*go/ast.GoStmt).End":          (*ast.GoStmt).End,
	"(*go/ast.GoStmt).Pos":          (*ast.GoStmt).Pos,
	"(*go/ast.Ident).End":           (*ast.Ident).End,
	"(*go/ast.Ident).IsExported":    (*ast.Ident).IsExported,
	"(*go/ast.Ident).Pos":           (*ast.Ident).Pos,
	"(*go/ast.Ident).String":        (*ast.Ident).String,
	"(*go/ast.IfStmt).End":          (*ast.IfStmt).End,
	"(*go/ast.IfStmt).Pos":          (*ast.IfStmt).Pos,
	"(*go/ast.ImportSpec).End":      (*ast.ImportSpec).End,
	"(*go/ast.ImportSpec).Pos":      (*ast.ImportSpec).Pos,
	"(*go/ast.IncDecStmt).End":      (*ast.IncDecStmt).End,
	"(*go/ast.IncDecStmt).Pos":      (*ast.IncDecStmt).Pos,
	"(*go/ast.IndexExpr).End":       (*ast.IndexExpr).End,
	"(*go/ast.IndexExpr).Pos":       (*ast.IndexExpr).Pos,
	"(*go/ast.InterfaceType).End":   (*ast.InterfaceType).End,
	"(*go/ast.InterfaceType).Pos":   (*ast.InterfaceType).Pos,
	"(*go/ast.KeyValueExpr).End":    (*ast.KeyValueExpr).End,
	"(*go/ast.KeyValueExpr).Pos":    (*ast.KeyValueExpr).Pos,
	"(*go/ast.LabeledStmt).End":     (*ast.LabeledStmt).End,
	"(*go/ast.LabeledStmt).Pos":     (*ast.LabeledStmt).Pos,
	"(*go/ast.MapType).End":         (*ast.MapType).End,
	"(*go/ast.MapType).Pos":         (*ast.MapType).Pos,
	"(*go/ast.Object).Pos":          (*ast.Object).Pos,
	"(*go/ast.Package).End":         (*ast.Package).End,
	"(*go/ast.Package).Pos":         (*ast.Package).Pos,
	"(*go/ast.ParenExpr).End":       (*ast.ParenExpr).End,
	"(*go/ast.ParenExpr).Pos":       (*ast.ParenExpr).Pos,
	"(*go/ast.RangeStmt).End":       (*ast.RangeStmt).End,
	"(*go/ast.RangeStmt).Pos":       (*ast.RangeStmt).Pos,
	"(*go/ast.ReturnStmt).End":      (*ast.ReturnStmt).End,
	"(*go/ast.ReturnStmt).Pos":      (*ast.ReturnStmt).Pos,
	"(*go/ast.Scope).Insert":        (*ast.Scope).Insert,
	"(*go/ast.Scope).Lookup":        (*ast.Scope).Lookup,
	"(*go/ast.Scope).String":        (*ast.Scope).String,
	"(*go/ast.SelectStmt).End":      (*ast.SelectStmt).End,
	"(*go/ast.SelectStmt).Pos":      (*ast.SelectStmt).Pos,
	"(*go/ast.SelectorExpr).End":    (*ast.SelectorExpr).End,
	"(*go/ast.SelectorExpr).Pos":    (*ast.SelectorExpr).Pos,
	"(*go/ast.SendStmt).End":        (*ast.SendStmt).End,
	"(*go/ast.SendStmt).Pos":        (*ast.SendStmt).Pos,
	"(*go/ast.SliceExpr).End":       (*ast.SliceExpr).End,
	"(*go/ast.SliceExpr).Pos":       (*ast.SliceExpr).Pos,
	"(*go/ast.StarExpr).End":        (*ast.StarExpr).End,
	"(*go/ast.StarExpr).Pos":        (*ast.StarExpr).Pos,
	"(*go/ast.StructType).End":      (*ast.StructType).End,
	"(*go/ast.StructType).Pos":      (*ast.StructType).Pos,
	"(*go/ast.SwitchStmt).End":      (*ast.SwitchStmt).End,
	"(*go/ast.SwitchStmt).Pos":      (*ast.SwitchStmt).Pos,
	"(*go/ast.TypeAssertExpr).End":  (*ast.TypeAssertExpr).End,
	"(*go/ast.TypeAssertExpr).Pos":  (*ast.TypeAssertExpr).Pos,
	"(*go/ast.TypeSpec).End":        (*ast.TypeSpec).End,
	"(*go/ast.TypeSpec).Pos":        (*ast.TypeSpec).Pos,
	"(*go/ast.TypeSwitchStmt).End":  (*ast.TypeSwitchStmt).End,
	"(*go/ast.TypeSwitchStmt).Pos":  (*ast.TypeSwitchStmt).Pos,
	"(*go/ast.UnaryExpr).End":       (*ast.UnaryExpr).End,
	"(*go/ast.UnaryExpr).Pos":       (*ast.UnaryExpr).Pos,
	"(*go/ast.ValueSpec).End":       (*ast.ValueSpec).End,
	"(*go/ast.ValueSpec).Pos":       (*ast.ValueSpec).Pos,
	"(go/ast.CommentMap).Comments":  (ast.CommentMap).Comments,
	"(go/ast.CommentMap).Filter":    (ast.CommentMap).Filter,
	"(go/ast.CommentMap).String":    (ast.CommentMap).String,
	"(go/ast.CommentMap).Update":    (ast.CommentMap).Update,
	"(go/ast.Decl).End":             (ast.Decl).End,
	"(go/ast.Decl).Pos":             (ast.Decl).Pos,
	"(go/ast.Expr).End":             (ast.Expr).End,
	"(go/ast.Expr).Pos":             (ast.Expr).Pos,
	"(go/ast.Node).End":             (ast.Node).End,
	"(go/ast.Node).Pos":             (ast.Node).Pos,
	"(go/ast.ObjKind).String":       (ast.ObjKind).String,
	"(go/ast.Spec).End":             (ast.Spec).End,
	"(go/ast.Spec).Pos":             (ast.Spec).Pos,
	"(go/ast.Stmt).End":             (ast.Stmt).End,
	"(go/ast.Stmt).Pos":             (ast.Stmt).Pos,
	"(go/ast.Visitor).Visit":        (ast.Visitor).Visit,
	"go/ast.FileExports":            ast.FileExports,
	"go/ast.FilterDecl":             ast.FilterDecl,
	"go/ast.FilterFile":             ast.FilterFile,
	"go/ast.FilterPackage":          ast.FilterPackage,
	"go/ast.Fprint":                 ast.Fprint,
	"go/ast.Inspect":                ast.Inspect,
	"go/ast.IsExported":             ast.IsExported,
	"go/ast.MergePackageFiles":      ast.MergePackageFiles,
	"go/ast.NewCommentMap":          ast.NewCommentMap,
	"go/ast.NewIdent":               ast.NewIdent,
	"go/ast.NewObj":                 ast.NewObj,
	"go/ast.NewPackage":             ast.NewPackage,
	"go/ast.NewScope":               ast.NewScope,
	"go/ast.NotNilFilter":           ast.NotNilFilter,
	"go/ast.PackageExports":         ast.PackageExports,
	"go/ast.Print":                  ast.Print,
	"go/ast.SortImports":            ast.SortImports,
	"go/ast.Walk":                   ast.Walk,
}

var typList = []interface{}{
	(*ast.ArrayType)(nil),
	(*ast.AssignStmt)(nil),
	(*ast.BadDecl)(nil),
	(*ast.BadExpr)(nil),
	(*ast.BadStmt)(nil),
	(*ast.BasicLit)(nil),
	(*ast.BinaryExpr)(nil),
	(*ast.BlockStmt)(nil),
	(*ast.BranchStmt)(nil),
	(*ast.CallExpr)(nil),
	(*ast.CaseClause)(nil),
	(*ast.ChanDir)(nil),
	(*ast.ChanType)(nil),
	(*ast.CommClause)(nil),
	(*ast.Comment)(nil),
	(*ast.CommentGroup)(nil),
	(*ast.CommentMap)(nil),
	(*ast.CompositeLit)(nil),
	(*ast.Decl)(nil),
	(*ast.DeclStmt)(nil),
	(*ast.DeferStmt)(nil),
	(*ast.Ellipsis)(nil),
	(*ast.EmptyStmt)(nil),
	(*ast.Expr)(nil),
	(*ast.ExprStmt)(nil),
	(*ast.Field)(nil),
	(*ast.FieldFilter)(nil),
	(*ast.FieldList)(nil),
	(*ast.File)(nil),
	(*ast.Filter)(nil),
	(*ast.ForStmt)(nil),
	(*ast.FuncDecl)(nil),
	(*ast.FuncLit)(nil),
	(*ast.FuncType)(nil),
	(*ast.GenDecl)(nil),
	(*ast.GoStmt)(nil),
	(*ast.Ident)(nil),
	(*ast.IfStmt)(nil),
	(*ast.ImportSpec)(nil),
	(*ast.Importer)(nil),
	(*ast.IncDecStmt)(nil),
	(*ast.IndexExpr)(nil),
	(*ast.InterfaceType)(nil),
	(*ast.KeyValueExpr)(nil),
	(*ast.LabeledStmt)(nil),
	(*ast.MapType)(nil),
	(*ast.MergeMode)(nil),
	(*ast.Node)(nil),
	(*ast.ObjKind)(nil),
	(*ast.Object)(nil),
	(*ast.Package)(nil),
	(*ast.ParenExpr)(nil),
	(*ast.RangeStmt)(nil),
	(*ast.ReturnStmt)(nil),
	(*ast.Scope)(nil),
	(*ast.SelectStmt)(nil),
	(*ast.SelectorExpr)(nil),
	(*ast.SendStmt)(nil),
	(*ast.SliceExpr)(nil),
	(*ast.Spec)(nil),
	(*ast.StarExpr)(nil),
	(*ast.Stmt)(nil),
	(*ast.StructType)(nil),
	(*ast.SwitchStmt)(nil),
	(*ast.TypeAssertExpr)(nil),
	(*ast.TypeSpec)(nil),
	(*ast.TypeSwitchStmt)(nil),
	(*ast.UnaryExpr)(nil),
	(*ast.ValueSpec)(nil),
	(*ast.Visitor)(nil),
}
