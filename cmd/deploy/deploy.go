package deploy

import (
	"bytes"
	"github.com/akin-ozer/containerless/library/access"
	"github.com/akin-ozer/containerless/library/shell"
	"text/template"
)

type Artifact struct {
	Image string
}

func Deploy(image string){
	s := access.ReadFile("template/service.yaml")
	t := template.New("knative deployment")
	t, _ = t.Parse(s)
	var tpl bytes.Buffer
	t.Execute(&tpl, Artifact{Image: image})

	result := tpl.String()
	//_ = t.Execute(os.Stdout, Artifact{Image: image})
	shell.PipedStdin("kubectl apply -f -", result)
}
