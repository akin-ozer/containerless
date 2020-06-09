package deploy

import (
	"bytes"
	"fmt"
	"github.com/akin-ozer/containerless/library/access"
	"github.com/akin-ozer/containerless/library/shell"
	"text/template"
)

type Artifact struct {
	Image string
	Name  string
}

func Deploy(image string, name string) {
	s := access.ReadFile("template/service.yaml")
	t := template.New("knative deployment")
	t, _ = t.Parse(s)

	var tpl bytes.Buffer
	t.Execute(&tpl, Artifact{Image: image, Name: name})
	result := tpl.String()

	shell.PipedStdin("kubectl apply -f -", result)
	shell.Piped("nohup kubectl port-forward -n ambassador svc/ambassador 10000:80 &>/dev/null &")
	fmt.Println("-----------")
	fmt.Println("access from host: ")
	fmt.Println("curl -H \"Host: " + name + ".default.knative.example.com\" http://127.0.0.1:10000")
}
