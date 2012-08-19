// +build !release

package gl

import (
  "fmt"
  "github.com/nathany/gl3/error"
  "log"
  "path"
  "runtime"
)

// checkError is used internally to handle OpenGL errors in development builds
func checkError() {
  if err := error.GetError(); err != nil {
    // Log the error
    gl_func_name, _ := callerInfo(1)
    caller_func_name, caller_file_line := callerInfo(2)
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
