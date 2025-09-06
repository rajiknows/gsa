package rules

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/rajiknows/gsa/internal/engine"
)

type ConcurrencyRule struct{}

func (s ConcurrencyRule) Name() string { return "time.sleep" }

func (c ConcurrencyRule) Apply(filename string, src []byte) ([]engine.Issue, error) {
	var issues []engine.Issue
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, src, 0)
	if err != nil {
		return nil, err
	}

	ast.Inspect(f, func(n ast.Node) bool {
		if forStmt, ok := n.(*ast.ForStmt); ok {
			// Check if the loop has a simple structure (for i := 0; i < N; i++).
			// This is a very basic check and can be improved.
			if forStmt.Init != nil && forStmt.Cond != nil && forStmt.Post != nil {
				pos := fset.Position(forStmt.Pos())
				issues = append(issues, engine.Issue{
					Rule:     c.Name(),
					File:     filename,
					Line:     pos.Line,
					Severity: "info",
					Message:  "This for loop could potentially be parallelized.",
				})
			}
		}
		return true
	})

	return issues, nil
}
