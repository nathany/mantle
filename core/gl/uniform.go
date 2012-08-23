package gl

// #include <OpenGL/gl3.h>
import "C"

/*
  SetUniform4f
*/
func (location UniformLocation) Set4f(v0, v1, v2, v3 float32) {
  // defer rc.handleErrors()
  C.glUniform4f(C.GLint(location), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2), C.GLfloat(v3))
}

func (location UniformLocation) SetColor(color Color) {
  location.Set4f(float32(color.Red), float32(color.Green), float32(color.Blue), float32(color.Alpha))
}
