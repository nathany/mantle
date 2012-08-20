package gl

// #include <OpenGL/gl3.h>
import "C"

import (
  "unsafe"
)

const (
  TRUE  = 1
  FALSE = 0
)

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
