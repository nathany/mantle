package gl3

// #include <OpenGL/gl3.h>
import "C"

import (
  "fmt"
  "log"
  "runtime"
)

type ErrorCode Enum

func GetError() ErrorCode {
  return ErrorCode(C.glGetError())
}

// "glGetError should always be called in a loop, until it returns GL_NO_ERROR"
func LogErrors() {
  func_name, file_line := callerInfo()
  for code := GetError(); code != NoError; code = GetError() {
    log.Printf("GL error %s in function %s\n  %s\n", code, func_name, file_line)
  }
}

func callerInfo() (func_name, file_line string) {
  if pc, file, line, ok := runtime.Caller(2); ok == true {
    if fun := runtime.FuncForPC(pc); fun != nil {
      func_name = fun.Name()
    }
    file_line = fmt.Sprintf("%s:%d", file, line)
  }
  return func_name, file_line
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
