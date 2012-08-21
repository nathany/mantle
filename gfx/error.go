package gfx

import (
  "fmt"
  "log"
  "path"
  "runtime"
  "strings"

  "github.com/nathany/mantle/core/gl"
)

func ErrorHandler() {
  if err := GetError(); err != nil {
    // Log the error
    gl_func_name, _ := callerInfo(2)
    caller_func_name, caller_file_line := callerInfo(3)
    log.Printf("GL Error %s calling %s() from %s at:\n  %s\n", err, path.Base(gl_func_name), caller_func_name, caller_file_line)
  }
}

// callerInfo retrieves a partial stack trace for logging purposes
func callerInfo(skip int) (func_name, file_line string) {
  // 19 Aug 2012 (Go 1.0.2) tried runtime.Callers instead, but wasn't returning anything
  if pc, file, line, ok := runtime.Caller(skip + 1); ok == true {
    if fun := runtime.FuncForPC(pc); fun != nil {
      func_name = fun.Name()
    }
    file_line = fmt.Sprintf("%s:%d", file, line)
  }
  return func_name, file_line
}

// Error (could be multiple)
type Error map[gl.ErrorCode]bool

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
  for code := gl.Error(); code != gl.NoError; code = gl.Error() {
    err[code] = true
  }
  if len(err) > 0 {
    return err
  }
  return nil
}
