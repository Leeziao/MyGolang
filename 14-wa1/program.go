package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
)

type Program struct {
	fs   map[string]string
	ast  map[string]*ast.File
	pkgs map[string]*types.Package
	infos map[string]*types.Info
	fset *token.FileSet
}

func NewProgram(fs map[string]string) *Program {
	return &Program{
		fs:   fs,
		ast:  make(map[string]*ast.File),
		pkgs: make(map[string]*types.Package),
		fset: token.NewFileSet(),
		infos: make(map[string]*types.Info),
	}
}

func (p *Program) LoadPackage(path string) (pkg *types.Package, f *ast.File, in *types.Info, err error) {
	if pkg, ok := p.pkgs[path]; ok {
		return pkg, p.ast[path], p.infos[path], nil
	}

	f, err = parser.ParseFile(p.fset, path, p.fs[path], parser.AllErrors)
	if err != nil {
		return nil, nil, nil, err
	}

	conf := types.Config{Importer: p}
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs: make(map[*ast.Ident]types.Object),
		Uses: make(map[*ast.Ident]types.Object),
		Implicits: make(map[ast.Node]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Scopes: make(map[ast.Node]*types.Scope),
	}

	pkg, err = conf.Check(path, p.fset, []*ast.File{f}, info) // Returns a Package
	if err != nil {
		return nil, nil, nil, err
	}

	p.ast[path] = f
	p.pkgs[path] = pkg
	p.infos[path] = info

	return pkg, f, info, nil
}

func (p *Program) Import(path string) (*types.Package, error) {
	if pkg, ok := p.pkgs[path]; ok {
		return pkg, nil
	}
	pkg, _, _, err := p.LoadPackage(path)
	return pkg, err
}