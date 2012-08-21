package gfx

import "github.com/nathany/mantle/core/gl"

const (
  Opaque gl.Alpha = 1.0
)

var (
  Black     = gl.Color{0.0, 0.0, 0.0, Opaque}
  Red       = gl.Color{1.0, 0.0, 0.0, Opaque}
  Green     = gl.Color{0.0, 1.0, 0.0, Opaque}
  Yellow    = gl.Color{1.0, 1.0, 0.0, Opaque}
  Blue      = gl.Color{0.0, 0.0, 1.0, Opaque}
  Magenta   = gl.Color{1.0, 0.0, 1.0, Opaque}
  Cyan      = gl.Color{0.0, 1.0, 1.0, Opaque}
  DarkGray  = gl.Color{0.25, 0.25, 0.25, Opaque}
  LightGray = gl.Color{0.75, 0.75, 0.75, Opaque}
  Brown     = gl.Color{0.60, 0.40, 0.12, Opaque}
  Orange    = gl.Color{0.98, 0.625, 0.12, Opaque}
  Pink      = gl.Color{0.98, 0.04, 0.7, Opaque}
  Purple    = gl.Color{0.60, 0.40, 0.70, Opaque}
  White     = gl.Color{1.0, 1.0, 1.0, Opaque}
)
