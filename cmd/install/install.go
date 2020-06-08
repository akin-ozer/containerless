package install

import (
	"github.com/akin-ozer/containerless/library/access"
	"github.com/akin-ozer/containerless/library/shell"
)

func Install() {
	s := access.ReadFile("static/scripts")
	shell.PipedStdin("sh", s)
}
