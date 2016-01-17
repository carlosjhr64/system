// Currently just a wrapper on exec.Command(c, a...).Output()
package system

import (
  "os"
  "fmt"
  "os/exec"
  "strings"
)

const VERSION string = "0.0.0"

type Command struct {
  command string
  arity int
}

func Oops(err error) {
  fmt.Fprintln(os.Stderr, err)
  os.Exit(69)
}

func New(command string) *Command {
  n := strings.Count(command, "%") // arity
  return &Command{command, n}
}

func (cmd *Command) Output(a ...interface{}) string {
  if len(a) != cmd.arity { panic("Wrong number of arguments.") }
  c := strings.Fields(fmt.Sprintf(cmd.command, a...))
  out, err := exec.Command(c[0], c[1:]...).Output()
  if err != nil { Oops(err) }
  return string(out)
}
