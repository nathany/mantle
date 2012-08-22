package gl

// #include <OpenGL/gl3.h>
import "C"

type Shadeable interface {
  NewShader(t ShaderType) *Shader
}

type Shader struct {
  Id   Uint
  Type ShaderType

  id C.GLuint // GL shader id
  rc *Context // render context
}

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

/*
  NewShader creates a shader object

    ShaderType  One of VertexShader, GeometryShader or FragmentShader

    Shader      Returns a shader, or nil on error.

    Error()     InvalidEnum

  glCreateShader: http://www.opengl.org/sdk/docs/man3/xhtml/glCreateShader.xml
*/
func (rc *Context) NewShader(t ShaderType) *Shader {
  defer rc.handleErrors()
  if id := C.glCreateShader(C.GLenum(t)); id != 0 {
    return &Shader{id: id, rc: rc, Type: t, Id: Uint(id)}
  }
  return nil
}

/*
  SetSource

    Error()     Invalid Value, Invalid Operation

  glShaderSource: http://www.opengl.org/sdk/docs/man3/xhtml/glShaderSource.xml
*/
func (shader Shader) SetSource(source string) {
  defer shader.rc.handleErrors()
  csource, length := allocCString(source)
  defer freeCString(csource)

  C.glShaderSource(shader.id, 1, &csource, &length)
}

/*
  Compile a shader object

  glCompileShader: http://www.opengl.org/sdk/docs/man3/xhtml/glCompileShader.xml
*/
func (shader Shader) Compile() {
  defer shader.rc.handleErrors()
  C.glCompileShader(shader.id)
}

func (shader Shader) GetCompileStatus() bool {
  defer shader.rc.handleErrors()
  return shader.get(compileStatus) == TRUE
}

func (shader *Shader) GetType() ShaderType {
  defer shader.rc.handleErrors()
  shader.Type = ShaderType(shader.get(shaderType))
  return shader.Type
}

/*
  IsShader checks if GL considers this a shader
  If the shader has been deleted (not just flagged for deletion), this will return false.

  glIsShader: http://www.opengl.org/sdk/docs/man3/xhtml/glIsShader.xml
*/
func (shader Shader) IsShader() bool {
  defer shader.rc.handleErrors()
  return C.glIsShader(shader.id) == TRUE
}

/*
  GetDeletionStatus checks if a shader is flagged for deletion.

    Error()     InvalidParam if shader has been deleted
*/
func (shader Shader) GetDeletionStatus() bool {
  defer shader.rc.handleErrors()
  return shader.get(deleteStatus) == TRUE
}

/*

  glGetShaderInfoLog: http://www.opengl.org/sdk/docs/man3/xhtml/glGetShaderInfoLog.xml
*/
func (shader Shader) GetInfoLog() string {
  defer shader.rc.handleErrors()
  length := shader.get(infoLogLength)

  if length > 1 {
    return fillGoString(length, func(ptr *C.GLchar, size C.GLsizei) {
      C.glGetShaderInfoLog(shader.id, size, nil, ptr)
    })
  }
  return ""
}

/*

  glGetShaderSource: http://www.opengl.org/sdk/docs/man3/xhtml/glGetShaderSource.xml
*/
func (shader Shader) GetSource() string {
  defer shader.rc.handleErrors()
  length := shader.get(shaderSourceLength)

  if length > 1 {
    return fillGoString(length, func(ptr *C.GLchar, size C.GLsizei) {
      C.glGetShaderSource(shader.id, size, nil, ptr)
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
  defer shader.rc.handleErrors()
  C.glDeleteShader(shader.id)
}

/*
  get

  glGetShaderiv: http://www.opengl.org/sdk/docs/man3/xhtml/glGetShader.xml
*/
func (shader Shader) get(pname shaderPname) int {
  // GL errors for are handled by the caller for better error reporting
  var params C.GLint
  C.glGetShaderiv(shader.id, C.GLenum(pname), &params)
  return int(params)
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
