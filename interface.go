package outinject

import "io"

// OutParser is the interface
// that is used by parsers
type OutParser interface {
	Parse(i ...interface{}) string // parses all arguments and returns a composed string
	Enable(mo *MOut) bool          // hook to disable or enable some features
}

type MOut struct {
	Io          io.ReadWriter // currently used ReadWriter
	NamedWriter string        // if a named write is set, this will be the name
	Parser      OutParser     // used parser
	IsTerminal  bool          // if we cloud detect an terminal support
	Width       int           // the witdth of a terminal, if we have detect a terminal
	Height      int           // same for the height
}

type OutputManager interface {
	Std() *MOut                                        // Points to the stdout
	Err() *MOut                                        // Ponints to the Errout
	SetParser(parser OutParser) *MOut                  // sets the current parser
	GetParser() *OutParser                             // returns the surrent parser
	Named(key string) *MOut                            // Like Std() or Err() but for a assigned ReadWriter (SetNamedWriter)
	SetNamedWriter(key string, io io.ReadWriter) *MOut // adds an costume ReadWriter togehter with an name
	ToString(i ...interface{}) string                  // creates an String by using the current Parser
	Out(i ...interface{}) (n int, err error)           // Print the output by using the current Parser and current ReadWriter, without line ending
	OutLn(i ...interface{}) (n int, err error)         // sam as Out() but with line ending
}
