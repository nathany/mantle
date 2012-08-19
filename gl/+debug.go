// +build !release

package gl

import (
  "github.com/nathany/gl3/error"
)

// checkError is used internally to handle OpenGL errors in development builds
func checkError() {
  errs := error.GetErrors()
  if len(errs) > 0 {
    error.LogErrors(errs, 1) // could register a global error delegation function
  }
}
