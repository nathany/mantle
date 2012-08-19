// +build !release

package gl

import (
  "github.com/nathany/gl3/error"
)

// checkError is used internally to handle OpenGL errors in development builds
func checkError() {
  if err := error.GetError(); err != nil {
    err.Log(1)
  }
}
