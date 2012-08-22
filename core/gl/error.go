package gl

// #include <OpenGL/gl3.h>
import "C"

import "fmt"

type Errorable interface {
  GetError() ErrorCode
}

type errorHandlerFunc func(Errorable)

/*
  Error will query OpenGL for an error code resulting from past operations.
  More than one error may have been flagged since Error() was last called,
  so Error() should be called multiple times until NoError is returned.
*/
func (_ Context) GetError() ErrorCode {
  return ErrorCode(C.glGetError())
}

// ErrorCode as returned directly from OpenGL.
type ErrorCode C.GLenum

// In case of an error, the requested operation will be ignored, with the exception of OutOfMemory.
const (
  InvalidEnum      ErrorCode = 0x0500 + iota // enum argument out of range
  InvalidValue                               // numeric argument out of range
  InvalidOperation                           // operation illegal in current state
  StackOverflow                              // defined in a later verison (glcoreab)
  StackUnderflow                             // defined in a later verison (glcoreab)
  OutOfMemory                                // GL memory exhausted, result is undefined
  InvalidFramebufferOperation

  NoError ErrorCode = 0
)

var errorText = map[ErrorCode]string{
  InvalidEnum:                 "Invalid Enum",
  InvalidValue:                "Invalid Value",
  InvalidOperation:            "Invalid Operation",
  StackOverflow:               "Stack Overflow",
  StackUnderflow:              "Stack Underflow",
  OutOfMemory:                 "Out of Memory",
  InvalidFramebufferOperation: "Invalid Framebuffer Operation",
}

// String returns the English text for an error message.
func (e ErrorCode) String() string {
  return errorText[e]
}

// Error returns an English error string with the error codes hexidecimal value.
func (e ErrorCode) Error() string {
  return fmt.Sprintf("%s (%#x)", errorText[e], int(e))
}
