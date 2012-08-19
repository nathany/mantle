package gl3

// #include <OpenGL/gl3.h>
import "C"

import (
  "log"
)

type ErrorCode Enum

func GetError() ErrorCode {
  return ErrorCode(C.glGetError())
}

// "glGetError should always be called in a loop, until it returns GL_NO_ERROR"
func LogErrors() {
  for code := GetError(); code != NoError; code = GetError() {
    log.Printf("gl error %s\n", code)
  }
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

var errorText = map[ErrorCode]string{
  InvalidEnum:                 "Invalid Enum",
  InvalidValue:                "Invalid Value",
  InvalidOperation:            "Invalid Operation",
  OutOfMemory:                 "Out of Memory",
  InvalidFramebufferOperation: "Invalid Framebuffer Operation",
}

func (e ErrorCode) Error() string {
  return errorText[e]
}
