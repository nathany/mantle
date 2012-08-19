/*

*/
package error

import (
  "fmt"
  "strings"
)

// Error (could be multiple)
type Error map[ErrorCode]bool

// Error interface
func (e Error) Error() string {
  var error_strings []string
  for code, _ := range e {
    error_strings = append(error_strings, code.Error())
  }
  return strings.Join(error_strings, "|")
}

// Get Error
func GetError() Error {
  err := Error{}
  // "glGetError should always be called in a loop, until it returns GL_NO_ERROR"
  for code := getError(); code != NoError; code = getError() {
    err[code] = true
  }
  if len(err) > 0 {
    return err
  }
  return nil
}

var errorText = map[ErrorCode]string{
  InvalidEnum:                 "Invalid Enum",
  InvalidValue:                "Invalid Value",
  InvalidOperation:            "Invalid Operation",
  OutOfMemory:                 "Out of Memory",
  InvalidFramebufferOperation: "Invalid Framebuffer Operation",
}

func (e ErrorCode) Error() string {
  return fmt.Sprintf("%s (%#x)", errorText[e], int(e))
}
