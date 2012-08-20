// +build !release

package gl

// callAfterHook is used to call a user-specified function after each GL call
// It's intended use is to handle OpenGL errors in development builds
func callAfterHook() {
  if afterHook != nil {
    afterHook()
  }
}
