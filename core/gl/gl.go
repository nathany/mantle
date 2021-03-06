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

type Context struct {
    ClearColor Color

    errorHandler errorHandlerFunc
}

// create New context and setup error handler
func NewContext(handler errorHandlerFunc) *Context {
    return &Context{errorHandler: handler}
}
