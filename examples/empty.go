package main

import (
  "fmt"
  "github.com/jteeuwen/glfw"
  "github.com/nathany/gl3"
  "log"
)

func main() {
  var x gl3.Int = 1
  fmt.Println("Hello World", x)

  if err := glfw.Init(); err != nil {
    log.Fatal(err)
  }
  defer glfw.Terminate()

  if err := glfw.OpenWindow(256, 256, 8, 8, 8, 0, 0, 0, glfw.Windowed); err != nil {
    log.Fatal(err)
  }
  defer glfw.CloseWindow()

  glfw.SetWindowTitle("Simple GLFW window")

  for glfw.WindowParam(glfw.Opened) == 1 {
    // do gl3 stuff
    glfw.SwapBuffers()
  }
}
