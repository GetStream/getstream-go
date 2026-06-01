package getstream

import (
	"fmt"
	"io"
	"runtime"
)

// stackErr wraps an error with a single captured stack frame at the
// wrap site. It is fully compatible with errors.Is / errors.As / errors.Unwrap.
//
// The %+v verb on fmt.Sprintf prints the full chain including the captured
// frame; %v / %s print only the formatted message of this layer.
type stackErr struct {
	msg   string
	cause error
	pc    [32]uintptr
	n     int
}

// stackWrap wraps err with the call site captured via runtime.Callers.
// The wrap site is rendered by Format("%+v") and accessible via
// StackTrace(). Use stackWrap instead of fmt.Errorf("...: %w", ...) in
// user-facing paths so the original wrap site survives the cause chain.
func stackWrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	se := &stackErr{msg: msg, cause: err}
	se.n = runtime.Callers(2, se.pc[:])
	return se
}

func (e *stackErr) Error() string {
	if e.msg == "" {
		return e.cause.Error()
	}
	return e.msg + ": " + e.cause.Error()
}

func (e *stackErr) Unwrap() error { return e.cause }

// StackTrace returns the captured frames as "func\n\tfile:line" entries.
// Returns nil if no frames were captured.
func (e *stackErr) StackTrace() []string {
	if e.n == 0 {
		return nil
	}
	frames := runtime.CallersFrames(e.pc[:e.n])
	var out []string
	for {
		f, more := frames.Next()
		if f.Function != "" || f.File != "" {
			out = append(out, fmt.Sprintf("%s\n\t%s:%d", f.Function, f.File, f.Line))
		}
		if !more {
			break
		}
	}
	return out
}

// Format implements fmt.Formatter. The %+v verb prints message + stack frames
// plus the wrapped cause (recursively if the cause also implements %+v).
// %v / %s print the flat error string.
func (e *stackErr) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = io.WriteString(s, e.Error())
			for _, frame := range e.StackTrace() {
				_, _ = io.WriteString(s, "\n")
				_, _ = io.WriteString(s, frame)
			}
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}
