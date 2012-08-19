package main

import (
  "github.com/jteeuwen/glfw"
  "github.com/nathany/gl3/gl"
  "log"
)

func main() {
  if err := glfw.Init(); err != nil {
    log.Fatal(err)
  }
  defer glfw.Terminate()

  openWindow()
  defer glfw.CloseWindow()

  glfw.SetWindowTitle("Simple GLFW window")

  for glfw.WindowParam(glfw.Opened) == 1 {
    // do gl3 stuff
    gl.Clear(gl.ColorBufferBit)

    glfw.SwapBuffers()
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
