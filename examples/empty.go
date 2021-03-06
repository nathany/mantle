package main

import (
    "github.com/jteeuwen/glfw"
    "github.com/nathany/mantle/core/gl"
    "github.com/nathany/mantle/gfx"
    "log"
)

var (
    rc *gl.Context // rendering context
)

func init() {
    rc = gl.NewContext(gfx.ErrorHandler)
}

func renderScene() {
    rc.Clear(gl.ColorBufferBit)
    glfw.SwapBuffers()
}

func main() {
    if err := glfw.Init(); err != nil {
        log.Fatal(err)
    }
    defer glfw.Terminate()

    openWindow()
    defer glfw.CloseWindow()
    glfw.SetWindowTitle("Mantle")
    glfw.SetWindowSizeCallback(onResize)

    rc.SetClearColor(gfx.Blue)

    for glfw.WindowParam(glfw.Opened) == 1 {
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
