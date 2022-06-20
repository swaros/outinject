package outinject

import "io"

// OutParser is the interface
// that is used by parsers
type OutParser interface {
	Parse(i ...interface{}) string // parses all arguments and returns a composed string
	Enable(mo *MOut) bool          // hook to disable or enable some features
}

type MOut struct {
	Io          io.ReadWriter
	NamedWriter string
	Parser      OutParser
	IsTerminal  bool
}
