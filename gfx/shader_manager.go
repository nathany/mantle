package gfx

// Identity Shader
// GLT_SHADER_IDENTITY = 0
// GLT_ATTRIBUTE_VERTEX = 0  // index for glBindAttribLocation
const (
  IdentityShaderVP = "attribute vec4 vVertex;" +
    "void main(void) " +
    "{ gl_Position = vVertex; " +
    "}"

  IdentityShaderFP = "uniform vec4 vColor;" +
    "void main(void) " +
    "{ gl_FragColor = vColor;" +
    "}"
)

func InitializeStockShaders() bool {
  //  uiStockShaders[GLT_SHADER_IDENTITY]     = gltLoadShaderPairSrcWithAttributes(szIdentityShaderVP, szIdentityShaderFP, 1, GLT_ATTRIBUTE_VERTEX, "vVertex");
  return true
}

// UseStockShader(GLT_STOCK_SHADER nShaderID, ...)
