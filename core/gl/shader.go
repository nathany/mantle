package gl

// #include <OpenGL/gl3.h>
import "C"

import (
  "errors"
)

type Shader Uint

type ShaderType Enum

const (
  VertexShader   ShaderType = 0x8B31
  GeometryShader ShaderType = 0x8DD9
  FragmentShader ShaderType = 0x8B30
)

var shaderTypeText = map[ShaderType]string{
  VertexShader:   "Vertex",
  GeometryShader: "Geometry",
  FragmentShader: "Fragment",
}

func (t ShaderType) String() string {
  return shaderTypeText[t]
}

// Parameter names for Shader.get()
type shaderPname Enum

const (
  shaderType         shaderPname = 0x8B4F
  deleteStatus       shaderPname = 0x8B80
  compileStatus      shaderPname = 0x8B81
  infoLogLength      shaderPname = 0x8B84
  shaderSourceLength shaderPname = 0x8B88
)

/*
  NewShader creates a shader object

    ShaderType  One of VertexShader, GeometryShader or FragmentShader

    Shader      Returns a shader.

    Error()     InvalidEnum

  glCreateShader: http://www.opengl.org/sdk/docs/man3/xhtml/glCreateShader.xml
*/
func NewShader(t ShaderType) (Shader, error) {
  s := Shader(C.glCreateShader(C.GLenum(t)))
  callAfterHook()

  if s == 0 {
    return 0, errors.New("Shader creation error")
  }
  return s, nil
}

/*
  SetSource

    Error()     Invalid Value, Invalid Operation

  glShaderSource: http://www.opengl.org/sdk/docs/man3/xhtml/glShaderSource.xml
*/
func (shader Shader) SetSource(source string) {
  csource, length := allocCString(source)
  defer freeCString(csource)

  C.glShaderSource(C.GLuint(shader), 1, &csource, &length)
}

// from https://github.com/banthar/gl
// func glString(s string) *C.GLchar { return (*C.GLchar)(C.CString(s)) }
// func freeString(ptr *C.GLchar)    { C.free(unsafe.Pointer(ptr)) }

/*
  Compile a shader object

  glCompileShader: http://www.opengl.org/sdk/docs/man3/xhtml/glCompileShader.xml
*/
func (shader Shader) Compile() bool {
  C.glCompileShader(C.GLuint(shader))
  return shader.compileStatus()
}

/*
  get

  glGetShaderiv: http://www.opengl.org/sdk/docs/man3/xhtml/glGetShader.xml
*/
func (shader Shader) get(pname shaderPname) int {
  var params C.GLint
  C.glGetShaderiv(C.GLuint(shader), C.GLenum(pname), &params)
  callAfterHook()
  return int(params)
}

/*
  glIsShader: http://www.opengl.org/sdk/docs/man3/xhtml/glIsShader.xml
*/
func (shader Shader) isShader() bool {
  return C.glIsShader(C.GLuint(shader)) == TRUE
}

func (shader Shader) compileStatus() bool {
  return shader.get(compileStatus) == TRUE
}

func (shader Shader) Type() ShaderType {
  return ShaderType(shader.get(shaderType))
}

func (shader Shader) IsDeleted() bool {
  return !shader.isShader() // deleted (no longer a shader)
}

func (shader Shader) DeletionFlag() bool {
  if shader.IsDeleted() {
    return false
  }
  return shader.get(deleteStatus) == TRUE // flagged
}

/*

  glGetShaderInfoLog: http://www.opengl.org/sdk/docs/man3/xhtml/glGetShaderInfoLog.xml
*/
func (shader Shader) Log() string {
  length := shader.get(infoLogLength)

  if length > 1 {
    return fillGoString(length, func(ptr *C.GLchar, size C.GLsizei) {
      C.glGetShaderInfoLog(C.GLuint(shader), size, nil, ptr)
    })
  }
  return ""
}

/*

  glGetShaderSource: http://www.opengl.org/sdk/docs/man3/xhtml/glGetShaderSource.xml
*/
func (shader Shader) GetSource() string {
  length := shader.get(shaderSourceLength)

  if length > 1 {
    return fillGoString(length, func(ptr *C.GLchar, size C.GLsizei) {
      C.glGetShaderSource(C.GLuint(shader), size, nil, ptr)
    })
  }
  return ""
}

/*
  Delete a shader object, or flag for deletion if attached to a program.
  The shader must be detached from programs before the shader memory is freed.

    Error()   Invalid Value

  glDeleteShader: http://www.opengl.org/sdk/docs/man3/xhtml/glDeleteShader.xml
*/
func (shader Shader) Delete() {
  C.glDeleteShader(C.GLuint(shader))
}
