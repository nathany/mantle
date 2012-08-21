package gl

// #include <OpenGL/gl3.h>
import "C"

import "fmt"

type ErrorCode C.GLenum

func Error() ErrorCode {
  return ErrorCode(C.glGetError())
}

const (
  InvalidEnum ErrorCode = 0x0500 + iota
  InvalidValue
  InvalidOperation
  StackOverflow  // defined in a later verison (glcoreab)
  StackUnderflow // defined in a later verison (glcoreab)
  OutOfMemory
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

func (e ErrorCode) String() string {
  return errorText[e]
}

func (e ErrorCode) Error() string {
  return fmt.Sprintf("%s (%#x)", errorText[e], int(e))
}
