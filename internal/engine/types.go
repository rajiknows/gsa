package engine

type Issue struct {
	Rule     string
	File     string
	Line     int
	Severity string
	Message  string
}

type Rule interface {
	Name() string
	Apply(filename string, src []byte) ([]Issue, error)
}
