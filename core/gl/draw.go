package gl

// #include <OpenGL/gl3.h>
import "C"

type PrimitiveMode Enum

const (
    Triangles PrimitiveMode = 0x0004
)

/*
  glDrawArrays: http://www.opengl.org/sdk/docs/man3/xhtml/glDrawArrays.xml
*/
func (rc *Context) DrawArrays(mode PrimitiveMode, first, count int) {
    defer rc.handleErrors()
    C.glDrawArrays(C.GLenum(mode), C.GLint(first), C.GLsizei(count))
}
