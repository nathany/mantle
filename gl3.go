package gl3

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

// clear buffers to preset values
// glClear sets the bitplane area of the window to values previously selected
// by glClearColor, glClearDepth, and glClearStencil
// http://www.opengl.org/sdk/docs/man3/xhtml/glClear.xml
func Clear(mask AttribMask) {
  C.glClear(C.GLbitfield(mask))
}

type AttribMask Bitfield

const (
  DepthBufferBit   AttribMask = 0x00000100
  StencilBufferBit AttribMask = 0x00000400
  ColorBufferBit   AttribMask = 0x00004000
)
