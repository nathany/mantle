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
func (rc *Context) Clear(mask AttribMask) {
    defer rc.handleErrors()
    C.glClear(C.GLbitfield(mask))
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

    Color   {Red, Green, Blue, Alpha}

  glClearColor: http://www.opengl.org/sdk/docs/man3/xhtml/glClearColor.xml
*/
func (rc *Context) SetClearColor(c Color) {
    defer rc.handleErrors()
    C.glClearColor(C.GLfloat(c.Red), C.GLfloat(c.Green), C.GLfloat(c.Blue), C.GLfloat(c.Alpha))
    rc.ClearColor = c
}
