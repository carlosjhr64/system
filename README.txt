package system // import "github.com/carlosjhr64/system"

Currently just a wrapper on exec.Command(c, a...).Output()

const VERSION string = "0.0.0"

func Oops(err error)
func New(command string) *Command
type Command struct { ... }
