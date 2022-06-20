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

type OutputManager interface {
	Std() *MOut
	Err() *MOut
	SetParser(parser OutParser) *MOut
	GetParser() *OutParser
	Named(key string) *MOut
	SetNamedWriter(key string, io io.ReadWriter) *MOut
	ToString(i ...interface{}) string
	Out(i ...interface{}) (n int, err error)
	OutLn(i ...interface{}) (n int, err error)
}
