/*MIT License

Copyright (c) 2022 Thomas Ziegler, <thomas.zglr@googlemail.com>. All rights reserved.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package outinject

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/term"
)

// list of writers accessible by the keyname
var (
	writers map[string]io.ReadWriter = make(map[string]io.ReadWriter)
)

// interface for parsers

func NewStdout() *MOut {
	var m MOut
	return m.Std()
}

// NewStdColored creates an new stdout
// with color support

// set stdout as writer
func (m *MOut) Std() *MOut {
	m.Io = os.Stdout
	m.detectTerminal(int(os.Stdout.Fd()))

	return m
}

// detectTerminal if we get information about the terminal, we assigne them
func (m *MOut) detectTerminal(fd int) {
	m.IsTerminal = term.IsTerminal(fd)
	if w, h, err := term.GetSize(fd); err == nil {
		m.Height = h
		m.Width = w
	} else {
		m.Height = -1
		m.Width = -1
	}
}

// set stderr as writer
func (m *MOut) Err() *MOut {
	m.Io = os.Stderr
	m.detectTerminal(int(os.Stderr.Fd()))
	return m
}

// sets the parser they is responsible for formatting
func (m *MOut) SetParser(parser OutParser) *MOut {
	m.Parser = parser
	parser.Enable(m)
	return m
}

// GetParser returns the instance of the parser.
// can be nil
func (m *MOut) GetParser() *OutParser {
	return &m.Parser
}

// sets an named writer if exists.
// or ignores if not.
// writer have to be set with SetNamedWriter before.
func (m *MOut) Named(key string) *MOut {
	if io, exists := writers[key]; exists {
		m.Io = io
		m.NamedWriter = key
	}
	return m
}

// register or overidde a io.Writer by key-name.
// also it will be set as the current writer
func (m *MOut) SetNamedWriter(key string, io io.ReadWriter) *MOut {
	if key == "" {
		key = "default"
	}
	writers[key] = io
	m.NamedWriter = key
	m.Io = io
	return m
}

// ToString get the formated string, depending on the used formatter
func (m *MOut) ToString(i ...interface{}) string {
	if m.Parser == nil {
		var plain PlainParse
		m.Parser = plain
	}
	return m.Parser.Parse(i...)

}

// Out print the formatted content by using fmt.Fprint
// have the same return values.
func (m *MOut) Out(i ...interface{}) (n int, err error) {
	out := m.ToString(i...)
	if m.Io == nil {
		m.Std()
	}
	return fmt.Fprint(m.Io, out)
}

// OutLn print the formatted content by using fmt.Fprintln
// it have the same return values like fmt.Fprintln
func (m *MOut) OutLn(i ...interface{}) (n int, err error) {
	out := m.ToString(i...)
	if m.Io == nil {
		m.Std()
	}
	return fmt.Fprintln(m.Io, out)
}
