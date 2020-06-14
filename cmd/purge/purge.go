package purge

import (
	"github.com/akin-ozer/containerless/library/shell"
)

func Purge() {
	shell.PipedStdin("sh", "kind delete cluster")
}
