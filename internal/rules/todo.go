package rules

import (
	"bufio"
	"bytes"

	"github.com/rajiknows/gsa/internal/engine"
)

type TodoRule struct{}

func (t TodoRule) Name() string { return "todo" }

func (t TodoRule) Apply(filename string, src []byte) ([]engine.Issue, error) {
	var issues []engine.Issue
	scanner := bufio.NewScanner(bytes.NewReader(src))
	line := 1
	for scanner.Scan() {
		txt := scanner.Text()
		if bytes.Contains([]byte(txt), []byte("TODO")) {
			issues = append(issues, engine.Issue{
				Rule: "todo", File: filename, Line: line,
				Severity: "info", Message: txt,
			})
		}
		line++
	}
	return issues, nil
}
