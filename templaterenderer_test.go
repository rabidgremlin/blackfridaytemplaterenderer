package blackfridaytemplaterenderer

import (
	"testing"
	"text/template"

	"github.com/russross/blackfriday"
)

func TestBasicRenderer(t *testing.T) {

	rendererTemplates := RendererTemplates{}
	rendererTemplates.NormalText, _ = template.New("NormalText").Parse("{{ .Text | html }}")
	rendererTemplates.ParagraphStart, _ = template.New("ParagraphStart").Parse("<p>")
	rendererTemplates.ParagraphEnd, _ = template.New("ParagraphEnd").Parse("</p>\n")
	rendererTemplates.Emphasis, _ = template.New("Emphasis").Parse("<em>{{ .Text }}</em>")

	renderer := TemplatedRenderer(0, rendererTemplates)

	input := []byte("This > is a *test*")
	output := blackfriday.Markdown(input, renderer, 0)

	outputAsString := string(output)
	expectedOutput := "<p>This &gt; is a <em>test</em></p>\n"

	if outputAsString != expectedOutput {
		t.Error("Expected ", expectedOutput, " got ", outputAsString)
	}
}

func TestLatexRenderer(t *testing.T) {

	rendererTemplates := RendererTemplates{}
	rendererTemplates.NormalText, _ = template.New("NormalText").Parse("{{ .Text }}")
	rendererTemplates.DocumentHeader, _ = template.New("DocumentHeader").Parse(`\documentclass{article}` + "\n" + `\begin{document}` + "\n")
	rendererTemplates.DocumentFooter, _ = template.New("DocumentFooter").Parse("\n" + `\end{document}`)

	rendererTemplates.ParagraphStart, _ = template.New("ParagraphStart").Parse("\n")
	rendererTemplates.ParagraphEnd, _ = template.New("ParagraphEnd").Parse("\n")
	rendererTemplates.Emphasis, _ = template.New("Emphasis").Parse(`\textit{ {{- .Text -}} }`)

	renderer := TemplatedRenderer(0, rendererTemplates)

	input := []byte("This is the first paragraph.\n\nThis is the second paragraph. I am the *emphasis*.")
	output := blackfriday.Markdown(input, renderer, 0)

	outputAsString := string(output)
	expectedOutput := `\documentclass{article}` + "\n" +
		`\begin{document}` + "\n\n" +
		`This is the first paragraph.` + "\n\n" +
		`This is the second paragraph. I am the \textit{emphasis}.` + "\n\n" +
		`\end{document}`

	if outputAsString != expectedOutput {
		t.Error("Expected ", expectedOutput, " got ", outputAsString)
	}
}
