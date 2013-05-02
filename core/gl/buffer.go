package gl

// #include <OpenGL/gl3.h>
import "C"

import (
    "unsafe"
)

type Buffer Uint

type Target Enum

const (
    ArrayBuffer Target = 0x8892
)

type Usage Enum

const (
    DynamicDraw Usage = 0x88E8
)

/*
  glGenBuffers: http://www.opengl.org/sdk/docs/man3/xhtml/glGenBuffers.xml
*/
func (rc *Context) GenBuffers(number int) []Buffer {
    defer rc.handleErrors()
    buffers := make([]Buffer, number) // a slice
    C.glGenBuffers(C.GLsizei(len(buffers)), (*C.GLuint)(&buffers[0]))
    return buffers
}

/*
  Hm: If no buffer object with name buffer exists, one is created with that name
  When a buffer object is bound to a target, the previous binding for that target is automatically broken.

  glBindBuffer: http://www.opengl.org/sdk/docs/man3/xhtml/glBindBuffer.xml
*/
func (buffer Buffer) Bind(target Target) {
    // defer rc.handleErrors()
    C.glBindBuffer(C.GLenum(target), C.GLuint(buffer))
}

/*
  TODO data interface{}? I dislike the runtime reflection in banthar's implementation.

  glBufferData: http://www.opengl.org/sdk/docs/man3/xhtml/glBufferData.xml
*/
func (target Target) SetBufferData(data []float32, usage Usage) {
    // defer rc.handleErrors()
    size := int(unsafe.Sizeof(data[0])) * len(data)
    // println(size)
    C.glBufferData(C.GLenum(target), C.GLsizeiptr(size), unsafe.Pointer(&data[0]), C.GLenum(usage))
}

/*
  glBufferSubData
*/

/*
  glDeleteBuffers: http://www.opengl.org/sdk/docs/man3/xhtml/glDeleteBuffers.xml
*/

/*
  glUnmapBuffer
*/
