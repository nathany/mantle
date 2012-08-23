package main

import (
  "fmt"
  "github.com/jteeuwen/glfw"
  "github.com/nathany/mantle/core/gl"
  "github.com/nathany/mantle/gfx"
  "log"
)

var (
  rc      *gl.Context // rendering context
  program *gl.Program // identity shader program
)

func init() {
  rc = gl.NewContext(gfx.ErrorHandler)
}

// GLBatch﻿  triangleBatch
// GLShaderManager﻿  shaderManager

// setup rendering context
func setupRC() {
  rc.SetClearColor(gfx.Blue)

  program = gfx.IdentityShader(rc)

  vertices := []float32{
    -0.5, 0.0, 0.0,
    0.5, 0.0, 0.0,
    0.0, 0.5, 0.0}

  fmt.Println(vertices)

  // triangleBatch.Begin(GL_TRIANGLES, 3)
  // triangleBatch.CopyVertexData3f(vertices)
  // triangleBatch.End()
}

func renderScene() {
  rc.Clear(gl.ColorBufferBit)
  if program != nil { // program is not yet defined first call from onResize
    gfx.UseIdentityShader(program, gfx.Red)
    // triangleBatch.Draw()
  }
  glfw.SwapBuffers()
}

func main() {
  if err := glfw.Init(); err != nil {
    log.Fatal(err)
  }
  defer glfw.Terminate()

  openWindow()
  defer glfw.CloseWindow()
  glfw.SetWindowTitle("Triangle")
  glfw.SetWindowSizeCallback(onResize)

  setupRC()
  defer program.Delete()

  for glfw.WindowParam(glfw.Opened) == 1 && glfw.Key(glfw.KeyEsc) == 0 {
    renderScene()
  }
}

func openWindow() {
  glfw.OpenWindowHint(glfw.OpenGLVersionMajor, 3)
  glfw.OpenWindowHint(glfw.OpenGLVersionMinor, 2)
  glfw.OpenWindowHint(glfw.OpenGLForwardCompat, gl.TRUE)
  glfw.OpenWindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

  // width, height, r, g, b, a, depth, stencil, mode
  if err := glfw.OpenWindow(640, 480, 8, 8, 8, 0, 0, 0, glfw.Windowed); err != nil {
    log.Fatal(err)
  }
}

func onResize(w, h int) {
  rc.SetViewPort(0, 0, w, h)
  renderScene()
}
