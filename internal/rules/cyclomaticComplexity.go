package rules

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/rajiknows/gsa/internal/engine"
)

const maxComplexity = 10

type CyclomaticComplexityRule struct{}

func (r CyclomaticComplexityRule) Name() string { return "cyclomatic-complexity" }

func (r CyclomaticComplexityRule) Apply(filename string, src []byte) ([]engine.Issue, error) {
	var issues []engine.Issue
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, src, 0)
	if err != nil {
		return nil, err
	}

	ast.Inspect(f, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok {
			complexity := 1
			ast.Inspect(fn.Body, func(n ast.Node) bool {
				switch n.(type) {
				case *ast.IfStmt, *ast.ForStmt, *ast.RangeStmt, *ast.CaseClause, *ast.CommClause:
					complexity++
				}
				return true
			})

			if complexity > maxComplexity {
				pos := fset.Position(fn.Pos())
				issues = append(issues, engine.Issue{
					Rule:     r.Name(),
					File:     filename,
					Line:     pos.Line,
					Severity: "warning",
					Message:  fmt.Sprintf("Function %s has a cyclomatic complexity of %d, which is higher than the recommended maximum of %d", fn.Name.Name, complexity, maxComplexity),
				})
			}
		}
		return true
	})

	return issues, nil
}
