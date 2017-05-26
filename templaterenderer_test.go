package blackfridaytemplaterenderer

import (
	"fmt"
	"testing"
	"text/template"

	"github.com/russross/blackfriday"
)

func TestBasicRenderer(t *testing.T) {

	rendererTemplates := RendererTemplates{}
	rendererTemplates.Entity, _ = template.New("Entity").Parse("{{ .Entity }}")
	rendererTemplates.NormalText, _ = template.New("NormalText").Parse("{{ .Text }}")
	rendererTemplates.ParagraphStart, _ = template.New("ParagraphStart").Parse("<p>")
	rendererTemplates.ParagraphEnd, _ = template.New("ParagraphEnd").Parse("</p>\n")
	rendererTemplates.Emphasis, _ = template.New("Emphasis").Parse("<em>{{ .Text }}</em>")

	renderer := TemplatedRenderer(0, rendererTemplates)

	input := []byte("This is a *test*")
	output := blackfriday.Markdown(input, renderer, 0)

	fmt.Printf("out:" + string(output) + "\n")
}
