package gl3

// #cgo darwin LDFLAGS: -framework OpenGL

// #if defined(__APPLE_CC__)
// #include <OpenGL/gl3.h>
// #endif
import "C"
import "unsafe"

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
