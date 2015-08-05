// Package minterm implements minimal terminal functions.
package minterm

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/keybase/gopass"
	//"golang.org/x/crypto/ssh/terminal"
)

// MinTerm is a minimal terminal interface.
type MinTerm struct {
	tty    *os.File
	in     *os.File
	out    *os.File
	width  int
	height int
}

var ErrPromptInterrupted = errors.New("prompt interrupted")

// New creates a new MinTerm and opens the terminal file.  Any
// errors that happen while opening or getting the terminal size
// are returned.
func New() (*MinTerm, error) {
	m := &MinTerm{}
	if err := m.open(); err != nil {
		return nil, err
	}
	return m, nil
}

// Shutdown closes the terminal.
func (m *MinTerm) Shutdown() error {
	if m.tty == nil {
		return nil
	}
	// this can hang waiting for newline, so do it in a goroutine.
	// application shutting down, so will get closed by os anyway...
	go m.tty.Close()
	return nil
}

// Size returns the width and height of the terminal.
func (m *MinTerm) Size() (int, int) {
	return m.width, m.height
}

// Write writes a string to the terminal.
func (m *MinTerm) Write(s string) error {
	_, err := fmt.Fprint(m.out, s)
	return err
}

// Prompt gets a line of input from the terminal.  It displays the text in
// the prompt parameter first.
func (m *MinTerm) Prompt(prompt string) (string, error) {
	m.Write(prompt)
	r := bufio.NewReader(m.in)
	p, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	// strip off the trailing newline
	if len(p) > 0 {
		p = p[:len(p)-1]
	}
	return p, nil
}

// PromptPassword gets a line of input from the terminal, but
// nothing is echoed to the terminal to hide the text.
func (m *MinTerm) PromptPassword(prompt string) (string, error) {
	m.Write(prompt)
	if !strings.HasSuffix(prompt, ": ") {
		m.Write(": ")
	}
	b, err := gopass.GetPasswd()
	if err != nil {
		if err == gopass.ErrInterrupted {
			err = ErrPromptInterrupted
		}
		return "", err
	}
	return string(b), nil
}
