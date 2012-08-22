/*
   Package gl is a hand-written binding to OpenGL 3.2 Core Profile.

   Some code is conditionally compiled for release or development builds.

   All GL functions call checkError() to catch errors, but in release builds
   it does nothing, and should be optimized away by the compiler.

   To use this optimization, specify the following build tag:

      go build -tags 'release'
*/
package gl

// #cgo darwin LDFLAGS: -framework OpenGL
import "C"

type afterFunc func(Errorable)

type Errorable interface {
  GetError() ErrorCode
}

type Context struct {
  ClearColor Color

  callAfter afterFunc
}

// create New context and setup error handler
func NewContext(after afterFunc) *Context {
  return &Context{callAfter: after}
}
