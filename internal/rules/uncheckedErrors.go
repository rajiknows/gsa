package rules

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"

	"github.com/rajiknows/gsa/internal/engine"
)

type UncheckedErrorRule struct{}

func (s UncheckedErrorRule) Name() string { return "unchecked error" }

func (c UncheckedErrorRule) Apply(filename string, src []byte) ([]engine.Issue, error) {
	var issues []engine.Issue
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, src, 0)
	if err != nil {
		return nil, err
	}

	conf := types.Config{Importer: nil}
	info := types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}
	_, _ = conf.Check("main", fset, []*ast.File{f}, &info)

	ast.Inspect(f, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		tv := info.Types[call]
		pos := fset.Position(call.Pos())
		switch t := tv.Type.(type) {
		case *types.Tuple:
			for i := 0; i < t.Len(); i++ {
				if isErrorType(t.At(i).Type()) {
					issues = append(issues, engine.Issue{
						Rule:     "unchecked Error",
						File:     filename,
						Severity: "high",
						Line:     pos.Line,
						Message:  "Call return an error please handle it ",
					})
				}
			}
		default:
			if isErrorType(t) {
				issues = append(issues, engine.Issue{
					Rule:     "unchecked Error",
					File:     filename,
					Severity: "high",
					Line:     pos.Line,
					Message:  "Call return an error please handle it ",
				})

			}
		}
		return true
	})
	return issues, nil
}

func isErrorType(t types.Type) bool {
	n, ok := t.(*types.Named)
	if !ok {
		return false
	}
	return n.Obj().Pkg() == nil && n.Obj().Name() == "error"
}
