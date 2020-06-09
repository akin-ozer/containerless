package remove

import (
	"github.com/akin-ozer/containerless/library/shell"
)

func Remove() {
	shell.PipedStdin("sh", "kind delete cluster")
}
