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

func InitializeStockShaders() bool {
  //  uiStockShaders[GLT_SHADER_IDENTITY]     = gltLoadShaderPairSrcWithAttributes(szIdentityShaderVP, szIdentityShaderFP, 1, GLT_ATTRIBUTE_VERTEX, "vVertex");

  vs, err := gl.NewShader(gl.VertexShader)
  if err != nil {
    log.Println(err)
  }
  log.Printf("%s Shader\n", vs.GetType())
  vs.SetSource(IdentityShaderVP)
  log.Println(vs.GetSource())
  vs.Compile()
  log.Printf("Compile successful? %v\n", vs.GetCompileStatus())
  log.Println(vs.GetInfoLog())

  fs, _ := gl.NewShader(gl.FragmentShader)
  fs.SetSource(IdentityShaderFP)
  log.Println(fs.GetSource())
  fs.Compile()
  log.Printf("Compile successful? %v\n", fs.GetCompileStatus())
  log.Println(fs.GetInfoLog())
  log.Println(fs.IsDeleted())

  vs.Delete()
  fs.Delete()
  log.Println(fs.GetDeletionFlag())
  log.Println(fs.IsDeleted())

  return true
}

// UseStockShader(GLT_STOCK_SHADER nShaderID, ...)
