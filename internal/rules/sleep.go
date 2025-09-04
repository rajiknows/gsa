package rules

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/rajiknows/gsa/internal/engine"
)

type SleepRule struct{}

func (s SleepRule) Name() string { return "time.sleep" }

func (s SleepRule) Apply(filename string, src []byte) ([]engine.Issue, error) {
	var issues []engine.Issue
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, src, 0)
	if err != nil {
		return nil, err
	}

	ast.Inspect(f, func(n ast.Node) bool {
		if call, ok := n.(*ast.CallExpr); ok {
			if fun, ok := call.Fun.(*ast.SelectorExpr); ok {
				if ident, ok := fun.X.(*ast.Ident); ok && ident.Name == "time" && fun.Sel.Name == "Sleep" {
					pos := fset.Position(call.Pos())
					issues = append(issues, engine.Issue{
						Rule:     "sleep",
						File:     filename,
						Line:     pos.Line,
						Severity: "warn",
						Message:  "Use of time.Sleep detected; Consider using sync primitives",
					})
				}
			}
		}
		return true
	})
	return issues, nil
}
