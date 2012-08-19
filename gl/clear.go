package gl

// #include <OpenGL/gl3.h>
import "C"

// Clear buffers to preset values
// glClear sets the bitplane area of the window to values previously selected
// by glClearColor, glClearDepth, and glClearStencil
// http://www.opengl.org/sdk/docs/man3/xhtml/glClear.xml
func Clear(mask AttribMask) {
  C.glClear(C.GLbitfield(mask))
  checkError()
}

type AttribMask Bitfield

const (
  DepthBufferBit   AttribMask = 0x00000100
  StencilBufferBit AttribMask = 0x00000400
  ColorBufferBit   AttribMask = 0x00004000
)
