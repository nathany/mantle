package gl

// #include <OpenGL/gl3.h>
import "C"

type Program struct {
  Id Uint

  id C.GLuint // GL program id
  rc *Context // render context
}

type AttribLocation uint
type UniformLocation int

/*
  NewProgram

  glCreateProgram: http://www.opengl.org/sdk/docs/man3/xhtml/glCreateProgram.xml
*/
func (rc *Context) NewProgram() *Program {
  defer rc.handleErrors()
  if id := C.glCreateProgram(); id != 0 {
    return &Program{id: id, rc: rc, Id: Uint(id)}
  }
  return nil
}

/*
  AttachShader

  glAttachShader: http://www.opengl.org/sdk/docs/man3/xhtml/glAttachShader.xml
*/
func (program Program) AttachShader(shader *Shader) {
  defer program.rc.handleErrors()
  C.glAttachShader(program.id, shader.id)
}

/*
  LinkProgram

  glLinkProgram: http://www.opengl.org/sdk/docs/man3/xhtml/glLinkProgram.xml
*/
func (program Program) Link() (ok bool) {
  defer program.rc.handleErrors()
  C.glLinkProgram(program.id)
  return program.get(linkStatus) == TRUE
}

/*
  GetInfoLog

  glGetProgramInfoLog: http://www.opengl.org/sdk/docs/man3/xhtml/glGetProgramInfoLog.xml
*/
func (program Program) GetInfoLog() string {
  defer program.rc.handleErrors()
  length := program.get(programInfoLogLength)

  if length > 1 {
    return fillGoString(length, func(ptr *C.GLchar, size C.GLsizei) {
      C.glGetProgramInfoLog(program.id, size, nil, ptr)
    })
  }
  return ""
}

/*
  Use

  glUseProgram: http://www.opengl.org/sdk/docs/man3/xhtml/glUseProgram.xml
*/
func (program Program) Use() {
  defer program.rc.handleErrors()
  C.glUseProgram(program.id)
}

/*
  BindAttribLocation

  glBindAttribLocation: http://www.opengl.org/sdk/docs/man3/xhtml/glBindAttribLocation.xml
*/
func (program Program) BindAttribLocation(index AttribLocation, name string) {
  defer program.rc.handleErrors()
  cName, _ := allocCString(name)
  defer freeCString(cName)
  C.glBindAttribLocation(program.id, C.GLuint(index), cName)
}

/*
  UniformLocation

  glGetUniformLocation: http://www.opengl.org/sdk/docs/man3/xhtml/glGetUniformLocation.xml
*/
func (program Program) UniformLocation(name string) UniformLocation {
  defer program.rc.handleErrors()
  cName, _ := allocCString(name)
  defer freeCString(cName)

  return UniformLocation(C.glGetUniformLocation(program.id, cName))
}

/*

  glBindFragDataLocation: http://www.opengl.org/sdk/docs/man3/xhtml/glBindFragDataLocation.xml
*/
func (program Program) BindFragDataLocation(colorNumber uint, name string) {
  defer program.rc.handleErrors()
  cName, _ := allocCString(name)
  defer freeCString(cName)

  C.glBindFragDataLocation(program.id, C.GLuint(colorNumber), cName)
}

/*
  Delete the program, or flag for deletion if part of the current rendering state.
  Shaders will be automtically detatched, but not deleted unless they were already flagged for deletion.

    GetError()    InvalidValue

  glDeleteProgram: http://www.opengl.org/sdk/docs/man3/xhtml/glDeleteProgram.xml
*/
func (program Program) Delete() {
  defer program.rc.handleErrors()
  C.glDeleteProgram(program.id)
}

/*
  get

  glGetProgramiv: http://www.opengl.org/sdk/docs/man3/xhtml/glGetProgram.xml
*/
func (program Program) get(pname programPname) int {
  // GL errors for are handled by the caller for better error reporting
  var params C.GLint
  C.glGetProgramiv(program.id, C.GLenum(pname), &params)
  return int(params)
}

// Parameter names for Program.get()
type programPname Enum

const (
  linkStatus           programPname = 0x8B82
  programInfoLogLength programPname = 0x8B84 // same as shaderInfoLogLength
)
