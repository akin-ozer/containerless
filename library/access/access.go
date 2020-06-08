package access

import (
	"github.com/gobuffalo/packr/v2"

)

func ReadFile(fileName string) string {
	box := packr.New("myBox", "../../resources")
	s, _ := box.FindString(fileName)
	return s
}