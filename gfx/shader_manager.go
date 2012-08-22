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

func InitializeStockShaders(rc gl.Shadeable) bool {
  //  uiStockShaders[GLT_SHADER_IDENTITY]     = gltLoadShaderPairSrcWithAttributes(szIdentityShaderVP, szIdentityShaderFP, 1, GLT_ATTRIBUTE_VERTEX, "vVertex");

  vs := rc.NewShader(gl.VertexShader)
  if vs == nil {
    log.Println("Error creating Vertex shader")
  }
  log.Printf("%s Shader %d\n", vs.Type, vs.Id)
  vs.SetSource(IdentityShaderVP)
  log.Println(vs.GetSource())
  log.Printf("Compile successful? %v\n", vs.Compile())
  log.Println(vs.GetInfoLog())

  fs := rc.NewShader(gl.FragmentShader)
  fs.SetSource(IdentityShaderFP)
  log.Println(fs.GetSource())
  log.Printf("Compile successful? %v\n", fs.Compile())
  log.Println(fs.GetInfoLog())
  // log.Println("Deleted?", !fs.IsShader())

  vs.Delete()
  fs.Delete()
  // log.Println("Deleted?", !fs.IsShader())
  // log.Println("Flagged for Deletion?", fs.GetDeletionStatus())

  return true
}

// UseStockShader(GLT_STOCK_SHADER nShaderID, ...)
