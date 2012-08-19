package gl3

// #include <OpenGL/gl3.h>
import "C"

import (
  "fmt"
  "log"
  "runtime"
  "strings"
)

func checkError() {
  errs := GetErrors()
  if len(errs) > 0 {
    LogErrors(errs, 1) // could register a global error delegation function
  }
}

type ErrorCode Enum

func getError() ErrorCode {
  return ErrorCode(C.glGetError())
}

// "glGetError should always be called in a loop, until it returns GL_NO_ERROR"
func GetErrors() (error_codes []ErrorCode) {
  for code := getError(); code != NoError; code = getError() {
    error_codes = append(error_codes, code)
  }
  return error_codes
}

func LogErrors(error_codes []ErrorCode, skip int) {
  if len(error_codes) > 0 {
    func_name, file_line := callerInfo(skip)

    var error_strings []string
    for _, code := range error_codes {
      error_strings = append(error_strings, code.Error())
    }

    errors := strings.Join(error_strings, "|")
    log.Printf("GL error %s in function %s\n  %s\n", errors, func_name, file_line)
  }
}

func callerInfo(skip int) (func_name, file_line string) {
  if pc, file, line, ok := runtime.Caller(skip + 2); ok == true {
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
