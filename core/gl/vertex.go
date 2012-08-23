package gl

// #include <OpenGL/gl3.h>
import "C"

type VertexArray Uint

/*
  glGenVertexArrays: http://www.opengl.org/sdk/docs/man3/xhtml/glGenVertexArrays.xml
*/
func (rc *Context) GenVertexArrays(number int) []VertexArray {
  defer rc.handleErrors()
  arrays := make([]VertexArray, number) // a slice
  C.glGenVertexArrays(C.GLsizei(number), (*C.GLuint)(&arrays[0]))
  return arrays
}

/*
  glBindVertexArray: http://www.opengl.org/sdk/docs/man3/xhtml/glBindVertexArray.xml
*/
func (array VertexArray) Bind() {
  // defer rc.handleErrors()
  C.glBindVertexArray(C.GLuint(array))
}

func (rc *Context) UnbindVertexArray() {
  defer rc.handleErrors()
  C.glBindVertexArray(0)
}

/*
  glDeleteVertexArrays:
*/
