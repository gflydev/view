package pongo

import (
	"fmt"
	"github.com/flosch/pongo2/v6"
	"github.com/gflydev/core"
	"github.com/gflydev/core/utils"
	"io"
	"log"
)

func New() core.IView {
	return &defaultView{
		ViewPath: utils.Getenv("VIEW_PATH", "resources/views"),
		ViewExt:  utils.Getenv("VIEW_EXT", "tpl"),
	}
}

type defaultView struct {
	ViewPath string
	ViewExt  string
}

func (v *defaultView) Parse(tpl string, data core.Data) string {
	file, err := loadFile(tpl)
	if err != nil {
		log.Fatal(err)
	}
	ctx := pongo2.Context(data)
	res, err := file.Execute(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return res
}

func (v *defaultView) Writer(tpl string, data core.Data, writer io.Writer) error {
	var templateFile = pongo2.Must(loadFile(tpl))
	ctx := pongo2.Context(data)

	return templateFile.ExecuteWriter(ctx, writer)
}

func loadFile(template string) (*pongo2.Template, error) {
	viewPath := utils.Getenv("VIEW_PATH", "resources/views")
	viewExt := utils.Getenv("VIEW_EXT", "tpl")

	return pongo2.FromFile(fmt.Sprintf("%s/%s.%s", viewPath, template, viewExt))
}
