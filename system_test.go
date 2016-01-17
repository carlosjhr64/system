package system

import "testing"

func TestOutput(test *testing.T) {
  bad := test.Error
  echo := New("echo %s")
  hello := echo.Output("Hello!")
  if hello != "Hello!\n" { bad("Hello!") }
}
