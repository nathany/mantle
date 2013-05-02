package gl

// #include <OpenGL/gl3.h>
import "C"

/*
  SetViewPort defines the 2D area of the framebuffer that OpenGL draws to.

  glViewport: http://www.opengl.org/sdk/docs/man3/xhtml/glViewport.xml
*/
func (rc *Context) SetViewPort(x, y, width, height int) {
    defer rc.handleErrors()
    C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}
