package gl3

// #include <OpenGL/gl3.h>
import "C"

type ErrorCode Enum

func GetError() ErrorCode {
  return ErrorCode(C.glGetError())
}

const (
  InvalidEnum ErrorCode = 0x0500 + iota
  InvalidValue
  InvalidOperation
  _
  _
  OutOfMemory
  InvalidFramebufferOperation
  NoError ErrorCode = 0
)
