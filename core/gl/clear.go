package gl

// #include <OpenGL/gl3.h>
import "C"

/*
  Clear the specified window buffers to values previously selected by
  SetClearColor, SetClearDepth, SetClearStencil.

    mask    (DepthBufferBit|StencilBufferBit|ColorBufferBit)

    Error() InvalidValue

  glClear: http://www.opengl.org/sdk/docs/man3/xhtml/glClear.xml
*/
func Clear(mask AttribMask) {
  C.glClear(C.GLbitfield(mask))
  callAfterHook()
}

type AttribMask Bitfield

const (
  DepthBufferBit   AttribMask = 0x00000100
  StencilBufferBit AttribMask = 0x00000400
  ColorBufferBit   AttribMask = 0x00004000
)

/*
  SetClearColor to preselect a color for Clear().
  Values are clamped between 0.0 and 1.0.

    Color   {Red, Green, Blue}
    Alpha   fully transparent 0.0 to 1.0 fully opaque

  glClearColor: http://www.opengl.org/sdk/docs/man3/xhtml/glClearColor.xml
*/
func SetClearColor(c Color, a Alpha) {
  C.glClearColor(C.GLclampf(c.Red), C.GLclampf(c.Green), C.GLclampf(c.Blue), C.GLclampf(a))
  callAfterHook()
}
