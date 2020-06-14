package delete

import (
	"github.com/akin-ozer/containerless/library/shell"
)

func Delete(deployment string) {
	cmdString := "kubectl delete service.serving.knative.dev " + deployment
	shell.Piped(cmdString)
}
