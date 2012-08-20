package error

/* #cgo darwin LDFLAGS: -framework OpenGL

   #include <OpenGL/gl3.h>
*/
import "C"

type ErrorCode C.GLenum

func getError() ErrorCode {
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
