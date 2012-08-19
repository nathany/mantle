/*

*/
package error

import (
  "fmt"
  "log"
  "runtime"
  "strings"
)

// Error (could be multiple)
type Error map[ErrorCode]bool

// Log the error
func (e Error) Log(skip int) {
  func_name, file_line := callerInfo(skip)
  log.Printf("GL error %s in function %s\n  %s\n", e, func_name, file_line)
}

// Error interface
func (e Error) Error() string {
  var error_strings []string
  for code, _ := range e {
    error_strings = append(error_strings, code.Error())
  }
  return strings.Join(error_strings, "|")
}

// Get Error
func GetError() Error {
  err := Error{}
  // "glGetError should always be called in a loop, until it returns GL_NO_ERROR"
  for code := getError(); code != NoError; code = getError() {
    err[code] = true
  }
  if len(err) > 0 {
    return err
  }
  return nil
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
