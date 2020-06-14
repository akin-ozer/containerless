package check

import (
	"fmt"
	"github.com/akin-ozer/containerless/library/shell"
	"os"
)

func Check() {
	returnCodeKind := shell.Execute("kind version")
	returnCodeKubectl := shell.Execute("kubectl version")

	if returnCodeKind != 0 || returnCodeKubectl != 0 {
		if returnCodeKind != 0 {
			fmt.Println("Kind can't be found")
		}
		if returnCodeKubectl != 0 {
			fmt.Println("Kubectl can't be found")
		}
		os.Exit(1)

	}

}
