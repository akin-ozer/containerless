package install

import (
	"github.com/akin-ozer/containerless/library/access"
	"github.com/akin-ozer/containerless/library/shell"
)

func Install() {
	s := access.ReadFile("static/scripts/install.sh")
	shell.PipedStdin("sh", s)
}
