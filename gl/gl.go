/*
   Package gl is a hand-written binding to OpenGL 3.2 Core Profile.

   Some code is conditionally compiled for release or development builds.

   All GL functions call checkError() to catch errors, but in release builds
   it does nothing, and should be optimized away by the compiler.

   To use this optimization, specify the following build tag:

      go build -tags 'release'
*/
package gl

/* #cgo darwin LDFLAGS: -framework OpenGL

   #include <OpenGL/gl3.h>
*/
import "C"

import (
  "unsafe"
)

const (
  TRUE  = 1
  FALSE = 0
)

// btou converts a bool to a 0/1 value.
func Btou(b bool) uint8 {
  if b {
    return 1
  }
  return 0
}

type (
  Enum     C.GLenum
  Boolean  C.GLboolean
  Bitfield C.GLbitfield
  Byte     C.GLbyte
  Short    C.GLshort
  Int      C.GLint
  Sizei    C.GLsizei
  Ubyte    C.GLubyte
  Ushort   C.GLushort
  Uint     C.GLuint
  Half     C.GLhalf
  Float    C.GLfloat
  Clampf   C.GLclampf
  Double   C.GLdouble
  Clampd   C.GLclampd
  Char     C.GLchar
  Pointer  unsafe.Pointer
  Sync     C.GLsync
  Int64    C.GLint64
  Uint64   C.GLuint64
  Intptr   C.GLintptr
  Sizeiptr C.GLsizeiptr
)

// Color contains red, green, blue and alpha values between 0.0 and 1.0
type Color struct {
  Red, Green, Blue, Alpha Clampf
}

var (
  White = Color{1, 1, 1, 0}
)
