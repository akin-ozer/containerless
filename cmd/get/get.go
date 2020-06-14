package get

import "github.com/akin-ozer/containerless/library/shell"

func Get() {
	shell.Piped("kubectl get service.serving.knative.dev")
}
