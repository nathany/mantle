package gfx

import (
  "github.com/nathany/mantle/core/gl"
  "log"
)

// Identity Shader
// GLT_SHADER_IDENTITY = 0
// GLT_ATTRIBUTE_VERTEX = 0  // index for glBindAttribLocation
const (
  IdentityShaderVP = "#version 150\n" +
    "in vec4 vVertex;\n" +
    "void main(void) {\n" +
    "  gl_Position = vVertex;\n" +
    "}"

  IdentityShaderFP = "#version 150\n" +
    "uniform vec4 vColor;\n" +
    "out vec4 FragColor;\n" + // gl_FragColor is deprecated, must explicitly define output
    "void main(void) {\n" +
    "  FragColor = vColor;\n" +
    "}"
)

type Shadeable interface {
  NewShader(t gl.ShaderType) *gl.Shader
  NewProgram() *gl.Program
}

// func InitializeStockShaders()

func IdentityShader(rc Shadeable) *gl.Program {
  //  uiStockShaders[GLT_SHADER_IDENTITY]     = gltLoadShaderPairSrcWithAttributes(szIdentityShaderVP, szIdentityShaderFP, 1, GLT_ATTRIBUTE_VERTEX, "vVertex");

  vs := rc.NewShader(gl.VertexShader)
  if vs == nil {
    log.Println("Error creating Vertex shader")
  }
  // log.Printf("%s Shader %d\n", vs.Type, vs.Id)
  vs.SetSource(IdentityShaderVP)
  // log.Println(vs.GetSource())
  if ok := vs.Compile(); !ok {
    log.Println("Compile failed:\n", vs.GetInfoLog())
  }

  fs := rc.NewShader(gl.FragmentShader)
  fs.SetSource(IdentityShaderFP)
  // log.Println(fs.GetSource())
  if ok := fs.Compile(); !ok {
    log.Println("Compile failed:\n", fs.GetInfoLog())
  }
  // log.Println("Deleted?", !fs.IsShader())

  p := rc.NewProgram()
  p.AttachShader(vs)
  p.AttachShader(fs)
  if ok := p.Link(); !ok {
    log.Println("Link failed:\n", p.GetInfoLog())
  }

  // Vertex attribute
  p.BindAttribLocation(0, "vVertex")

  // what do I need to do for FragColor?

  vs.Delete()
  fs.Delete()
  // log.Println("Deleted?", !fs.IsShader())
  // log.Println("Flagged for Deletion?", fs.GetDeletionStatus())

  return p
}

// UseStockShader(GLT_STOCK_SHADER nShaderID, ...)
func UseIdentityShader(program *gl.Program, color gl.Color) {
  program.Use()
  colorLocation := program.UniformLocation("vColor")
  colorLocation.SetColor(color)
}
