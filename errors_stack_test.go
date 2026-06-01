package getstream

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestStackWrap_NilCause(t *testing.T) {
	if got := stackWrap(nil, "msg"); got != nil {
		t.Fatalf("stackWrap(nil) = %v, want nil", got)
	}
}

func TestStackWrap_PreservesUnwrapChain(t *testing.T) {
	sentinel := errors.New("root cause")
	wrapped := stackWrap(sentinel, "outer")

	if !errors.Is(wrapped, sentinel) {
		t.Fatalf("errors.Is should reach the wrapped sentinel; chain broken")
	}
	if got := errors.Unwrap(wrapped).Error(); got != "root cause" {
		t.Fatalf("Unwrap = %q, want %q", got, "root cause")
	}
}

func TestStackWrap_ErrorString(t *testing.T) {
	wrapped := stackWrap(errors.New("inner"), "outer")
	if got := wrapped.Error(); got != "outer: inner" {
		t.Fatalf("Error() = %q, want %q", got, "outer: inner")
	}
}

func TestStackWrap_PlusVFormatIncludesStack(t *testing.T) {
	wrapped := stackWrap(errors.New("inner"), "outer")
	formatted := fmt.Sprintf("%+v", wrapped)

	if !strings.Contains(formatted, "outer: inner") {
		t.Fatalf("%%+v output missing flat message: %q", formatted)
	}
	if !strings.Contains(formatted, "errors_stack_test.go") {
		t.Fatalf("%%+v output missing stack frame referencing the wrap site: %q", formatted)
	}
}

func TestStackWrap_VFormatNoStack(t *testing.T) {
	wrapped := stackWrap(errors.New("inner"), "outer")
	formatted := fmt.Sprintf("%v", wrapped)
	if formatted != "outer: inner" {
		t.Fatalf("%%v output = %q, want %q", formatted, "outer: inner")
	}
}

func TestStackWrap_StackTraceAccessor(t *testing.T) {
	wrapped := stackWrap(errors.New("inner"), "outer")
	var se *stackErr
	if !errors.As(wrapped, &se) {
		t.Fatalf("errors.As did not extract *stackErr")
	}
	frames := se.StackTrace()
	if len(frames) == 0 {
		t.Fatalf("StackTrace() empty, expected at least one frame")
	}
	// First frame should be in this test file.
	if !strings.Contains(frames[0], "errors_stack_test.go") {
		t.Fatalf("first frame = %q, want errors_stack_test.go", frames[0])
	}
}
