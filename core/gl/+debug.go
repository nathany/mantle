// +build !release

package gl

// handleErrors is used to call a user-specified error handler after each GL call
func (rc *Context) handleErrors() {
  if rc.errorHandler != nil {
    rc.errorHandler(rc)
  }
}
