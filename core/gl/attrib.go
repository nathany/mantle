package gl

// #include <OpenGL/gl3.h>
import "C"

import "unsafe"

const FLOAT Enum = 0x1406

/*

  glEnableVertexAttribArray: http://www.opengl.org/sdk/docs/man3/xhtml/glEnableVertexAttribArray.xml
*/
func (rc *Context) EnableVertexAttribArray(index AttribLocation) {
    defer rc.handleErrors()
    C.glEnableVertexAttribArray(C.GLuint(index))
}

/*
  glVertexAttribPointer: http://www.opengl.org/sdk/docs/man3/xhtml/glVertexAttribPointer.xml
*/
func (rc *Context) VertexAttribPointer(index AttribLocation, size int, type_ Enum, normalized bool, stride int, pointer interface{}) {
    defer rc.handleErrors()
    var cNormalized C.GLboolean
    if normalized {
        cNormalized = TRUE
    } else {
        cNormalized = FALSE
    }
    C.glVertexAttribPointer(C.GLuint(index), C.GLint(size), C.GLenum(type_), cNormalized, C.GLsizei(stride), unsafe.Pointer(nil))
}
